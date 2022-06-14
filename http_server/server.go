package httpserver

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "fl/docs"
	"fl/http_server/app"
	"fl/http_server/http/middleware"
)

func Run() {
	ginApp := gin.Default()
	ginApp.Use(middleware.AuthMiddleware)
	ginApp.GET("/version", func(c *gin.Context) {
    c.String(http.StatusOK, "1.0.0")
  })
	ginApp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginApp.POST("/v1/job/", app.JobCreate)


	log.Println("Server is running....")
	ginApp.Run(":8443")
}
