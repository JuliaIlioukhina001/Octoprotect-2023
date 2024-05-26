package router

import (
	"backend/config"
	"backend/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServeHome(c *gin.Context) {
	c.String(http.StatusTeapot, "I'm a teapot.")
}

func ListenHTTP() {
	r := gin.Default()
	r.GET("/", ServeHome)
	r.GET("/ws/user", UserAuth, ServeUserWS)
	r.GET("/ws/nexus", NexusAuth, ServeNexusWS)
	r.POST("/nexus", AdminAuth, controller.ProvisionNexus)
	err := r.Run(config.AppConfig.HTTPListenAddr)
	if err != nil {
		panic(err)
	}
}
