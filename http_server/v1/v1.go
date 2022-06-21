package v1

import (
  "net/http"
	"github.com/gin-gonic/gin"

  "fl/http_server/v1/app"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.GET("/version", func(c *gin.Context) {
    c.String(http.StatusOK, "1.0.0")
  })

  Router.POST("/job/submit/", app.JobSubmit)
  Router.POST("/job/", app.JobCreate)
	Router.GET("/job/", app.JobList)
}
