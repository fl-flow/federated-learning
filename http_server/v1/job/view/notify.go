package view

import (
  "fmt"
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/fl-flow/dag-scheduler/dag_scheduler_client"

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

    var partyList []string
    for _, parties := range jobForm.PartyMap {
      for _, partyIpPort := range parties{
        insert := true
        for _, i := range partyList {
          if i == partyIpPort {
            insert = false
            break
          }
        }
        if insert {
          partyList = append(partyList, partyIpPort)
        }
      }
    }
    // for _, parties := range jobForm.PartyMap { // duplicate send
    for _, partyIpPort := range partyList {
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
    // }
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

  // handle ready event, to run task
  if model.TaskStatusType(f.Status) != model.TaskReady {
    return
  }
  fmt.Println(f, "fffffff")
  if f.Party != etc.LocalParty {
    return
  }
  var tasks []model.Task
  db.DataBase.Find(&tasks, model.Task{
    JobID: f.JobID,
    Role: f.Group,
    Name: f.Task,
    Party: f.Party,
  })
  run := true
  for _, task := range tasks {
    if task.Status != model.TaskReady {
      run = false
      break
    }
  }
  fmt.Println(run, f, "fffffff")
  if run {
    dagschedulerclient.ToRunTask(map[string]interface{}{
      "job_id": f.JobID,
      "group": f.Group,
      "task": f.Task,
    })
  }
}
