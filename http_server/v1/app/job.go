package app

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/fl-flow/dag-scheduler/dag_scheduler_client"

  "fl/http_server/v1/form"
  "fl/http_server/http/response"
  "fl/http_server/http/mixin"
  "fl/http_server/v1/controller/job"
  "fl/http_server/http/reverse_proxy"
)


func JobSubmit(context *gin.Context) {
  fmt.Println("asdasd")
  var f form.JobSubmitForm
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  token, er := jobcontroller.JobSubmit(f)
  mixin.CommonResponse(context, token, er)
}


func JobCreate(context *gin.Context) {
  f := form.JobCreateForm{}
	if e := context.ShouldBindJSON(&f); e != nil {
    response.R(
      context,
      100,
      fmt.Sprintf("%v", e),
      fmt.Sprintf("%v", e),
    )
    return
	}
  job, error := jobcontroller.JobCreate(f)
  if error != nil {
    response.R(
      context,
      error.Code,
      error.Message(),
      error.Message(),
    )
    return
  }
  response.R(context, 0, "success", job)
}


func JobList(context *gin.Context) {
  // TODO: 别删 可能 还有用
  // var pagination PageNumberPagination
  // context.ShouldBindQuery(&pagination)
  // jobs, _ := dagschedulerclient.ListJob(pagination.Page, pagination.Size)
  // response.R(context, 0, "success", jobs)
  reverseproxy.ReverseProxy(context, dagschedulerclient.IPPort, dagschedulerclient.Protocol)
}
