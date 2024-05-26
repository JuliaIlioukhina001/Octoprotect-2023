package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = pongWait / 2
	maxMessageSize = 1024
)

var UserConnMap = map[uint]*UserConn{}
var NexusConnMap = map[uint]*NexusConn{}

type BaseConn struct {
	Conn               *websocket.Conn
	Send               chan interface{}
	registeredCommands map[string]func(jsonStr string) error
	log                log.FieldLogger
	closeHandler       func()
}

func CreateConn(conn *websocket.Conn) BaseConn {
	return BaseConn{
		Conn:               conn,
		Send:               make(chan interface{}),
		log:                log.WithField("remoteAddr", conn.RemoteAddr()),
		registeredCommands: map[string]func(jsonStr string) error{},
	}
}

func (c *BaseConn) ReportError(error string) {
	c.Send <- gin.H{
		"type":    "error",
		"message": error,
	}
}

func (c *BaseConn) RegisterCommand(command string, fn func(jsonStr string) error) {
	c.registeredCommands[command] = fn
}

type BaseCommand struct {
	Type string `json:"type"`
}

func (c *BaseConn) CommandHandler() {
	c.Conn.SetReadLimit(maxMessageSize)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		kind, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				c.closeHandler()
			}
			break
		}
		if kind == websocket.TextMessage {
			var command BaseCommand
			if err := json.Unmarshal(message, &command); err != nil {
				c.ReportError("cannot deserialize command")
				c.log.Error(err)
				continue
			}
			if c.registeredCommands[command.Type] == nil {
				c.ReportError("invalid command")
				c.log.Error("invalid command:", command.Type, "original message:", message)
				continue
			}
			c.log.Debug("Received command", command.Type, message)
			err := c.registeredCommands[command.Type](string(message))
			if err != nil {
				c.ReportError("internal server error")
				c.log.Error(err)
				continue
			}
		}
	}
}

func (c *BaseConn) Writer() {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		c.closeHandler()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteJSON(message); err != nil {
				c.log.Error(err)
				return
			}
		case <-pingTicker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.log.Error(err)
				return
			}
		}
	}
}
