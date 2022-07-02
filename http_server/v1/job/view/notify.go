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


func Notify(context *gin.Context) {
  var f form.Notify
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  if f.Type == "task" {
    taskNotifyKV := f.Extra.(map[string]interface{})
    taskNotify := form.TaskNotify {
      JobID: uint(taskNotifyKV["job_id"].(float64)),
      Group: taskNotifyKV["group"].(string),
      Task: taskNotifyKV["task"].(string),
    }
    db.DataBase.Debug().Where(
      "job_id=? AND role=? AND name=? AND party=?",
      taskNotify.JobID,
      taskNotify.Group,
      taskNotify.Task,
      etc.LocalParty, // TODO:
    ).Updates(model.Task{
      Status: model.TaskStatusType(f.Status),
    })

    fmt.Println("notify task data // TODO: ", taskNotify)
  }
  fmt.Println("notify data // TODO: ", f)
}
