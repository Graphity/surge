package app

import "github.com/gin-gonic/gin"

func RegisterAllHandlers(r *gin.Engine) {
	r.GET("/", Index)
	r.GET("/ping", Ping)
}
