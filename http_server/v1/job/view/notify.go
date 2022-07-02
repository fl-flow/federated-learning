package view

import (
  "fmt"
  "github.com/gin-gonic/gin"

  "fl/etc"
  "fl/common/db"
  "fl/common/db/model"
  "fl/http_server/http/mixin"
  "fl/http_server/v1/job/form"
)


func NotifyTask(context *gin.Context) {
  var f form.TaskNotify
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  if f.Type == "task" {
    db.DataBase.Debug().Where(
      "job_id=? AND role=? AND name=? AND party=?",
      f.Extra.JobID,
      f.Extra.Group,
      f.Extra.Task,
      etc.LocalParty, // TODO:
    ).Updates(model.Task{
      Status: model.TaskStatusType(f.Status),
    })
  }
}


func NotifyJob(context *gin.Context) {
  var f form.JobNotify
  if ok := mixin.CheckJSON(context, &f); !ok {
    fmt.Println(f, "asdasdasd")
    return
  }
  if f.Type == "job" {
    db.DataBase.Debug().Where(
      "id=?",
      f.ID,
    ).Updates(model.Job{
      Status: model.JobStatusType(f.Status),
    })
  }
}
