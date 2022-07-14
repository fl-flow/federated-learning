package jobcontroller

import (
  "fmt"
  "sync"
  "encoding/json"
  // dagModel "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/dag_scheduler_client"
  //
  //
  // "fl/etc"
  "fl/common/db"
  "fl/common/error"
  "fl/common/db/model"
  // "fl/util/id_generator"
  "fl/http_server/v1/job/form"
)


func JobFederationCancelController(f form.JobCancelForm) *error.Error {
  var j model.Job
  ret := db.DataBase.Find(&j, model.Job{ID: f.JobID})
  if ret.RowsAffected == 0 {
    return &error.Error{Code: 103010, Hits: fmt.Sprintf("%v", f.JobID)}
  }
  var c form.JobForm
  json.Unmarshal(j.Conf, &c)

  var ids []string
  for _, partyIds := range c.PartyMap {
    new := true
    for _, partyId := range partyIds {
      for _, i := range ids {
        if i == partyId {
          new = false
          break
        }
      }
      if new {
        ids = append(ids, partyId)
      }
    }
  }

  wait := &sync.WaitGroup{}
  wait.Add(len(ids))
  for _, partyId := range ids {
    fmt.Println(partyId)
    go TransferJob(
      partyId,
      "/api/v1/job/cancel/",
      map[string]uint {
        "job_id": f.JobID,
      },
      wait,
    )
  }
  wait.Wait()
  return nil
}


func JobCancelController(f form.JobCancelForm) *error.Error {
  ret := db.DataBase.Find(&model.Job{}, model.Job{ID: f.JobID})
  if ret.RowsAffected == 0 {
    return &error.Error{Code: 103010, Hits: fmt.Sprintf("%v", f.JobID)}
  }
  fmt.Println(f.JobID, "f.JobIDf.JobID")
  dagschedulerclient.Cancel(f.JobID)
  return nil
}
