package notification

import (
	"backend/model"
	log "github.com/sirupsen/logrus"
)

func NotifyUser(userID uint, message string) {
	var user model.User
	result := model.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		log.Error(result.Error)
		return
	}
	retries := 3
	for retries > 0 {
		if user.PushoverKey != "" {
			if PushoverNotify(user.PushoverKey, message) == nil {
				return
			}
		}
		if user.TelegramUserID != "" {
			if TelegramNotify(user.TelegramUserID, message) == nil {
				return
			}
		}
		retries--
	}
}
