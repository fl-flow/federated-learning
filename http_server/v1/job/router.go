package job

import (
	"github.com/gin-gonic/gin"

  "fl/http_server/v1/job/view"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.POST("/submit/", view.JobSubmit)
  Router.POST("/", view.JobCreate)
	Router.GET("/", view.JobList)

  Router.POST("/notify/", view.Notify)
}
