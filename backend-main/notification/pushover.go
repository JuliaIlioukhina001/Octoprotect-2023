package notification

import (
	"backend/config"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PushoverNotify(user string, message string) error {
	body := gin.H{
		"token":   config.AppConfig.PushoverToken,
		"user":    user,
		"message": message,
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = http.Post("https://api.pushover.net/1/messages.json", "application/json", bytes.NewBuffer(bodyJson))
	if err != nil {
		return err
	}
	return nil
}
