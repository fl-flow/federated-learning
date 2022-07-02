package v1

import (
  "net/http"
	"github.com/gin-gonic/gin"

  "fl/http_server/v1/job"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.GET("/version", func(c *gin.Context) {
    c.String(http.StatusOK, "1.0.0")
  })

  job.RegisterRouter(Router.Group("/job"))
}
