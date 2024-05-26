package controller

import (
	"backend/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type UserConn struct {
	BaseConn
	UserID     uint
	CachedUser *model.User
	log        log.FieldLogger
}

func CreateUserConn(conn *websocket.Conn, userID uint) *UserConn {
	c := UserConn{}
	c.BaseConn = CreateConn(conn)
	c.closeHandler = c.CloseHandler
	c.UserID = userID
	c.log = c.BaseConn.log.WithField("userID", userID)
	if err := c.ReloadUser(); err != nil {
		c.ReportError("internal error")
		c.log.Error(err)
		return nil
	}
	UserConnMap[userID] = &c

	c.RegisterCommand("pair", c.Pair)
	c.RegisterCommand("unpair", c.Unpair)
	c.RegisterCommand("fetch-device-list", c.FetchDeviceList)
	c.RegisterCommand("request-state", c.ForwardRequestToNexus)
	c.RegisterCommand("start-stream", c.ForwardRequestToNexus)
	c.RegisterCommand("stop-stream", c.ForwardRequestToNexus)
	c.RegisterCommand("arm", c.ForwardRequestToNexus)
	c.RegisterCommand("disarm", c.ForwardRequestToNexus)
	c.RegisterCommand("update-config", c.UpdateConfig)

	return &c
}

func (c *UserConn) CloseHandler() {
	c.BaseConn.Conn.Close()
	if UserConnMap[c.UserID] == c {
		UserConnMap[c.UserID] = nil
		close(c.Send)
	}
}

func (c *UserConn) ReloadUser() error {
	var newUser model.User
	result := model.DB.Model(&model.User{}).Preload("Nexuses").First(&newUser, "id = ?", c.UserID)
	if result.Error != nil {
		return result.Error
	}
	c.CachedUser = &newUser
	return nil
}

type PairRequest struct {
	NexusMac   string `json:"nexusMac"`
	PairSecret string `json:"pairSecret"`
	NickName   string `json:"nickName"`
}

func (c *UserConn) Pair(jsonStr string) error {
	var request PairRequest
	err := json.Unmarshal([]byte(jsonStr), &request)
	if err != nil {
		c.ReportError("invalid request")
		c.log.Error(err)
		return nil
	}
	var nexus model.Nexus
	result := model.DB.First(&nexus, "mac_address = ?", request.NexusMac)
	if result.Error != nil {
		c.ReportError("target nexus is not registered")
		return nil
	}
	if nexus.PairSecret != request.PairSecret {
		c.ReportError("wrong pair secret")
		return nil
	}
	if request.NickName != "" {
		nexus.NickName = request.NickName
	}
	nexus.Users = append(nexus.Users, c.CachedUser)
	result = model.DB.Save(nexus)
	if result.Error != nil {
		return result.Error
	}
	if err := c.ReloadUser(); err != nil {
		return err
	}
	if NexusConnMap[nexus.ID] != nil {
		if err := NexusConnMap[nexus.ID].ReloadNexus(); err != nil {
			return err
		}
	}
	c.Send <- gin.H{
		"type": "pair-success",
	}
	return nil
}

type UnpairRequest struct {
	NexusID uint `json:"nexusID"`
}

func (c *UserConn) Unpair(jsonStr string) error {
	var request UnpairRequest
	err := json.Unmarshal([]byte(jsonStr), &request)
	if err != nil {
		c.ReportError("invalid request")
		c.log.Error(err)
		return nil
	}
	if request.NexusID == 0 {
		c.ReportError("no nexusID presented")
		return nil
	}
	if !c.checkNexusPermission(request.NexusID) {
		c.ReportError("permission denied")
		c.log.Error("permission denied for unpairing the nexus", request.NexusID)
		return nil
	}
	var nexus model.Nexus
	result := model.DB.First(&nexus, "id = ?", request.NexusID)
	if result.Error != nil {
		return result.Error
	}

	// remove user from nexus users
	i := 0
	for idx, n := range nexus.Users {
		if c.CachedUser.ID == n.ID {
			i = idx
		}
	}
	nexus.Users[i] = nexus.Users[len(nexus.Users)-1]
	nexus.Users[len(nexus.Users)-1] = nil
	nexus.Users = nexus.Users[:len(nexus.Users)-1]

	result = model.DB.Save(nexus)
	if result.Error != nil {
		return result.Error
	}
	err = c.ReloadUser()
	if err != nil {
		return err
	}
	if NexusConnMap[nexus.ID] != nil {
		err = NexusConnMap[nexus.ID].ReloadNexus()
		if err != nil {
			return err
		}
	}

	c.Send <- gin.H{
		"type":    "unpair-success",
		"nexusID": nexus.ID,
	}
	return nil
}

type NexusConfig struct {
	NexusID     uint                 `json:"nexusID"`
	Sensitivity float64              `json:"sensitivity"`
	TitanW      []model.TitanWConfig `json:"titanW"`
}

func (c *UserConn) UpdateConfig(jsonStr string) error {
	var request NexusConfig
	err := json.Unmarshal([]byte(jsonStr), &request)
	if err != nil {
		c.ReportError("invalid request")
		c.log.Error(err)
		return nil
	}
	if request.NexusID == 0 {
		c.ReportError("no nexusID presented")
		return nil
	}
	if !c.checkNexusPermission(request.NexusID) {
		c.ReportError("permission denied")
		c.log.Error("permission denied for updating config of nexus", request.NexusID)
		return nil
	}
	var nexus model.Nexus
	result := model.DB.First(&nexus, "id = ?", request.NexusID)
	if result.Error != nil {
		return result.Error
	}
	nexus.Config = model.NexusConfig{
		Sensitivity: request.Sensitivity,
		TitanW:      request.TitanW,
	}
	result = model.DB.Save(nexus)
	if result.Error != nil {
		return result.Error
	}
	if NexusConnMap[request.NexusID] != nil {
		NexusConnMap[request.NexusID].ReloadNexus()
		NexusConnMap[request.NexusID].SendInitializeConfig()
	}
	c.Send <- gin.H{
		"type": "update-config-success",
	}
	return nil
}

type DeviceListInfo struct {
	ID         uint              `json:"id"`
	MacAddress string            `json:"macAddress"`
	NickName   string            `json:"nickName"`
	Config     model.NexusConfig `json:"config"`
	Online     bool              `json:"online"`
}

func (c *UserConn) FetchDeviceList(_ string) error {
	if err := c.ReloadUser(); err != nil {
		return err
	}
	result := make([]DeviceListInfo, 0)
	for _, nexus := range c.CachedUser.Nexuses {
		result = append(result, DeviceListInfo{
			ID:         nexus.ID,
			MacAddress: nexus.MacAddress,
			NickName:   nexus.NickName,
			Config:     nexus.Config,
			Online:     NexusConnMap[nexus.ID] != nil,
		})
	}
	c.Send <- gin.H{
		"type": "device-list",
		"data": result,
	}
	return nil
}

func (c *UserConn) checkNexusPermission(nexusID uint) bool {
	_ = c.ReloadUser()
	for _, nexus := range c.CachedUser.Nexuses {
		if nexus.ID == nexusID {
			return true
		}
	}
	return false
}

func (c *UserConn) ForwardRequestToNexus(jsonStr string) error {
	var request map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &request)
	if err != nil {
		c.ReportError("invalid request")
		c.log.Error(err)
		return nil
	}
	if request["nexusID"] == nil {
		c.ReportError("no nexusID presented")
		return nil
	}
	nexusID := uint(request["nexusID"].(float64))
	if !c.checkNexusPermission(nexusID) {
		c.ReportError("permission denied")
		c.log.Error("trying to perform", request)
		return nil
	}
	if NexusConnMap[nexusID] == nil {
		c.ReportError("nexus is offline")
		return nil
	}
	NexusConnMap[nexusID].Send <- request
	return nil
}
