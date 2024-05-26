package controller

import (
	"backend/model"
	"backend/notification"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type NexusConn struct {
	BaseConn
	NexusID     uint
	CachedNexus *model.Nexus
}

func CreateNexusConn(conn *websocket.Conn, nexusID uint) *NexusConn {
	c := NexusConn{}
	c.BaseConn = CreateConn(conn)
	c.NexusID = nexusID
	c.closeHandler = c.CloseHandler
	c.log = c.BaseConn.log.WithField("nexusID", nexusID)
	if err := c.ReloadNexus(); err != nil {
		c.ReportError("internal error")
		c.log.Error(err)
		return nil
	}
	NexusConnMap[c.NexusID] = &c

	c.RegisterCommand("accel", c.ApplyNexusIDThenBroadcast)
	c.RegisterCommand("conn-state", c.ApplyNexusIDThenBroadcast)
	c.RegisterCommand("nexus-state", c.ApplyNexusIDThenBroadcast)
	c.RegisterCommand("movement-trigger", c.MovementTrigger)

	// broadcast new online status
	for _, user := range c.CachedNexus.Users {
		if UserConnMap[user.ID] != nil {
			_ = UserConnMap[user.ID].FetchDeviceList("")
		}
	}
	return &c
}

func (c *NexusConn) CloseHandler() {
	c.BaseConn.Conn.Close()
	if NexusConnMap[c.NexusID] == c {
		NexusConnMap[c.NexusID] = nil
		close(c.Send)
	}
}

func (c *NexusConn) ReloadNexus() error {
	var newNexus model.Nexus
	result := model.DB.Model(&model.Nexus{}).Preload("Users").First(&newNexus, "id = ?", c.NexusID)
	if result.Error != nil {
		return result.Error
	}
	c.CachedNexus = &newNexus
	return nil
}

func (c *NexusConn) SendInitializeConfig() {
	enabledTitanW := make([]string, 0)
	for _, titanW := range c.CachedNexus.Config.TitanW {
		if titanW.Enabled {
			enabledTitanW = append(enabledTitanW, titanW.UUID)
		}
	}
	c.Send <- gin.H{
		"type":        "initialize",
		"devices":     enabledTitanW,
		"sensitivity": c.CachedNexus.Config.Sensitivity,
	}
}

func (c *NexusConn) broadcastMessage(message interface{}) {
	for _, user := range c.CachedNexus.Users {
		if UserConnMap[user.ID] != nil {
			UserConnMap[user.ID].Send <- message
		}
	}
}

func (c *NexusConn) ApplyNexusIDThenBroadcast(jsonStr string) error {
	var request map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &request)
	if err != nil {
		c.ReportError("invalid request")
		c.log.Error(err)
		return nil
	}
	request["nexusID"] = c.NexusID
	c.broadcastMessage(request)
	return nil
}

type MovementTriggerPayload struct {
	Type      string  `json:"type"`
	NexusID   uint    `json:"nexusID"`
	TitanID   string  `json:"titanID"`
	Magnitude float64 `json:"magnitude"`
}

func (c *NexusConn) MovementTrigger(jsonStr string) error {
	var request MovementTriggerPayload
	err := json.Unmarshal([]byte(jsonStr), &request)
	if err != nil {
		c.ReportError("invalid request")
		c.log.Error(err)
		return nil
	}
	request.NexusID = c.NexusID
	c.broadcastMessage(request)
	for _, user := range c.CachedNexus.Users {
		notification.NotifyUser(user.ID, fmt.Sprintf("Detected movement! Nexus ID: %d, Titan ID: %s, Magnitude: %f", request.NexusID, request.TitanID, request.Magnitude))
	}
	return nil
}
