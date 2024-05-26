package notification

import (
	"backend/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TelegramNotify(user string, message string) error {
	obj := gin.H{
		"chat_id": user,
		"text":    message,
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(obj)
	if err != nil {
		return err
	}
	_, err = http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.AppConfig.TelegramToken), "application/json", buf)
	return err
}
