package httpserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"fl/etc"
	_ "fl/docs"
	"fl/http_server/v1"
	"fl/http_server/http/middleware"
)

func Run() {
	ginApp := gin.Default()

	ginApp.Use(middleware.AuthMiddleware)

	ginApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1.RegisterRouter(ginApp.Group("/api/v1"))

	ginApp.Run(fmt.Sprintf("%v:%d", etc.IP, etc.PORT))
}
