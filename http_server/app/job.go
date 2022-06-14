package app

import (
  "fmt"

  "github.com/gin-gonic/gin"

  "fl/http_server/form"
  "fl/http_server/controller/job"
  "fl/http_server/http/response"
)


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
  error := jobcontroller.JobCreate(f)
  if error != nil {
    response.R(
      context,
      error.Code,
      error.Message(),
      error.Message(),
    )
    return
  }
  response.R(context, 0, "success", "success")
}
