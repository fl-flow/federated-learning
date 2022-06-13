package app

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"

  "dag/common/parser"
  "dag/http_server/form"
  "dag/http_server/controller"
)


func JobCreate(context *gin.Context) {
  f := form.JobCreateForm{}
	if e := context.ShouldBindJSON(&f); e != nil {
		// id = article.Insert()
    fmt.Println(e)
    context.JSON(http.StatusBadRequest, gin.H{})
    return // TODO:
	}
  error := controller.JobCreate(
    f.Name,
    parser.Conf {
      Dag: f.Dag,
      Parameter: f.Parameter,
    },
  )
  fmt.Println(error, "asdasd")
  context.JSON(http.StatusOK, gin.H{})
}
