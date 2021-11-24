package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Graphity/surge/server/app"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../../../web/static/*")
	app.RegisterAllHandlers(router)
	router.Run()
}
