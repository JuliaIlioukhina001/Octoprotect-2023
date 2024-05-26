package router

import (
	"backend/config"
	"backend/model"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func NexusAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header presented",
		})
		return
	}
	token, found := strings.CutPrefix(authHeader, "Bearer ")
	if !found {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no bearer authorization header presented",
		})
		return
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecretKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid JWT",
		})
		return
	}
	if claims["nexusMac"] == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no nexusMac in JWT",
		})
		return
	}
	var nexus model.Nexus
	result := model.DB.First(&nexus, "mac_address = ?", claims["nexusMac"])
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "nexus not registered",
			})
			return
		} else {
			log.Error(result.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	c.Set("nexusID", nexus.ID)
	c.Next()
}

func registerUser(username string, password string) (id uint) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := model.User{
		Username: username,
		Password: string(hash),
	}
	result := model.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user.ID
}

func UserAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header presented",
		})
		return
	}
	cred, found := strings.CutPrefix(authHeader, "Basic ")
	if !found {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no basic authorization header presented",
		})
		return
	}
	credBytes, err := base64.StdEncoding.DecodeString(cred)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid basic authorization header",
		})
		return
	}
	credPair := strings.SplitN(string(credBytes), ":", 2)
	if len(credPair) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid basic authorization header",
		})
		return
	}
	var user model.User
	result := model.DB.First(&user, "username = ?", credPair[0])
	if result.Error != nil {
		// Create user if not exists
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			id := registerUser(credPair[0], credPair[1])
			c.Set("userID", id)
			c.Next()
			return
		} else {
			panic(result.Error)
		}
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credPair[1])) != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "wrong password",
		})
		return
	}
	c.Set("userID", user.ID)
	c.Next()
}

func AdminAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization header presented",
		})
		return
	}
	cred, found := strings.CutPrefix(authHeader, "Token ")
	if !found {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no token authorization header presented",
		})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(config.AppConfig.AdminTokenBcrypted), []byte(cred)) != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "wrong admin token",
		})
		return
	}
	c.Next()
}
