package router

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeUserWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"remoteAddr": c.Request.RemoteAddr,
		}).Warn("Failed to upgrade to ws, err:", err)
		return
	}
	userConn := controller.CreateUserConn(conn, c.GetUint("userID"))
	go userConn.Writer()
	go userConn.CommandHandler()
}

func ServeNexusWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"remoteAddr": c.Request.RemoteAddr,
		}).Warn("Failed to upgrade to ws, err:", err)
		return
	}
	nexusConn := controller.CreateNexusConn(conn, c.GetUint("nexusID"))
	go nexusConn.Writer()
	go nexusConn.CommandHandler()
	nexusConn.SendInitializeConfig()
}
