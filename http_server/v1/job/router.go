package job

import (
	"github.com/gin-gonic/gin"

  "fl/http_server/v1/job/view"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.POST("/submit/", view.JobSubmit)
  Router.POST("/", view.JobCreate)
	Router.GET("/", view.JobList)

  Router.POST("/notify/task/", view.NotifyTask)
	Router.POST("/notify/job/", view.NotifyJob)
	Router.POST("/notify/faderated/task/", view.FederatedNotifyTask)
}
