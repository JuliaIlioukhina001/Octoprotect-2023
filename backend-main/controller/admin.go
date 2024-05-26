package controller

import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type ProvisionNexusRequest struct {
	NexusMac string `json:"nexusMac" binding:"required"`
}

func ProvisionNexus(c *gin.Context) {
	var request ProvisionNexusRequest
	e := c.ShouldBindJSON(&request)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request params",
		})
		return
	}
	pairSecret := uuid.New().String()
	nexus := model.Nexus{
		MacAddress: request.NexusMac,
		PairSecret: pairSecret,
		NickName:   "Unnamed Nexus",
		Config: model.NexusConfig{
			Sensitivity: 1.0,
		},
	}
	result := model.DB.Create(&nexus)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "the given nexusMac is already provisioned",
			})
		} else {
			log.WithFields(log.Fields{
				"remoteAddr": c.RemoteIP(),
				"nexusMac":   request.NexusMac,
			}).Error(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal error",
			})
		}
		return
	}
	c.JSON(http.StatusOK, nexus)
}
