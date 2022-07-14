package view

import (
  "github.com/gin-gonic/gin"

  "fl/http_server/http/mixin"
  "fl/http_server/v1/job/form"
  "fl/http_server/v1/job/controller"
)


func JobFederationCancelView(context *gin.Context) {
  var f form.JobCancelForm
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  e := jobcontroller.JobFederationCancelController(f)
  mixin.CommonResponse(context, "success", e)
}


func JobCancelView(context *gin.Context) {
  var f form.JobCancelForm
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  e := jobcontroller.JobCancelController(f)
  mixin.CommonResponse(context, "success", e)
}
