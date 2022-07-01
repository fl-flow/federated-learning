package jobcontroller

import (
  "sync"
  "encoding/json"
  "gorm.io/datatypes"
  dagModel "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/dag_scheduler_client"

  "fl/common/db"
  "fl/common/error"
  "fl/common/db/model"
  "fl/http_server/v1/form"
)

const LocalParty = "127.0.0.1:8443"


func JobCreate(f form.JobForm) (dagModel.Job, *error.Error) {
  party2CreateForm, er := PartyParse(f.PartyMap, f)
  if er != nil {
    return dagModel.Job{}, er
  }
  dagConfMap := make(map[string]DagConf)
  for party_id, partConf := range party2CreateForm {
    dagConf, e := FederationParse(partConf, f.PartyMap)
    if e != nil {
      return dagModel.Job{}, e
    }
    dagConfMap[party_id] = dagConf
  }
  // TODO: LocalParty in dagConfMap
  dagConf := dagConfMap[LocalParty]

  job, clientE := dagschedulerclient.CreateJob(dagConf)
  if clientE != nil {
    errorMessage := clientE.Message()
    return job, &error.Error{
      Code: clientE.Code,
      Msg: errorMessage["message"].(string),
      Hits: errorMessage["hits"].(string),
    }
  }

  tx := db.DataBase.Begin()
  jByte, _ := json.Marshal(f)
  j := model.Job{
    Name: f.Name,
    Conf: datatypes.JSON(jByte),
  }
  db.DataBase.Create(&j)
  var insertingTasks []model.Task
  for party_id, dagConf := range dagConfMap {
    for role, tasks := range dagConf.Dag {
      for t, _ := range tasks {
        insertingTasks = append(insertingTasks, model.Task{
          Job: j,
          Party: party_id,
          Role: role,
          Name: t,
        })
      }
    }
  }
  db.DataBase.Debug().Create(&insertingTasks)
  tx.Commit()
  return job, nil
}


func JobSubmit(f form.JobForm) (form.JobForm, *error.Error) {
  party2CreateForm, er := PartyParse(f.PartyMap, f)
  if er != nil {
    return f, er
  }
  for _, partConf := range party2CreateForm {
    _, e := FederationParse(partConf, f.PartyMap)
    if e != nil {
      return f, e
    }
  }
  wait := &sync.WaitGroup{}
  wait.Add(len(party2CreateForm))
  for party_id, _ := range party2CreateForm {
    go TransferJob(party_id, f, wait)
  }
  wait.Wait()
  return f, nil
}
