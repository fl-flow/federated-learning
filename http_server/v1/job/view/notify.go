package view

import (
  "fmt"
  "encoding/json"
  "github.com/gin-gonic/gin"

  "fl/etc"
  "fl/common/db"
  "fl/util/requests"
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
      etc.LocalParty,
    ).Updates(model.Task{
      Status: model.TaskStatusType(f.Status),
    })
    var j model.Job
    db.DataBase.Debug().Find(&j, model.Task{
      ID: f.Extra.JobID,
    })
    var jobForm form.JobForm
    json.Unmarshal(j.Conf, &jobForm)

    for _, parties := range jobForm.PartyMap { // duplicate send
      for _, partyIpPort := range parties {
          b, _ := json.Marshal(
            form.FederatedTaskNotify{
              Status: f.Status,
              JobID: f.Extra.JobID,
              Group: f.Extra.Group,
              Task: f.Extra.Task,
              Party: etc.LocalParty,
            },
          )
          go func () {
            requests.Post(
              fmt.Sprintf(
                "http://%v/api/v1/job/notify/faderated/task/",
                partyIpPort,
              ),
              b,
              map[string]string{},
            )
            }()
      }
    }
  }
}


func NotifyJob(context *gin.Context) {
  var f form.JobNotify
  if ok := mixin.CheckJSON(context, &f); !ok {
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


func FederatedNotifyTask(context *gin.Context) {
  var f form.FederatedTaskNotify
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  db.DataBase.Debug().Where(
    "job_id=? AND role=? AND name=? AND party=?",
    f.JobID,
    f.Group,
    f.Task,
    f.Party, // TODO:
  ).Updates(model.Task{
    Status: model.TaskStatusType(f.Status),
  })
}
