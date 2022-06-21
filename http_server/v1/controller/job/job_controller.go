package jobcontroller

import (
  "fmt"
  "github.com/fl-flow/dag-scheduler/common/db/model"
  "github.com/fl-flow/dag-scheduler/dag_scheduler_client"

  "fl/common/error"
  "fl/http_server/v1/form"
)


func JobCreate(f form.JobCreateForm) (model.Job, *error.Error) {
  jobConf, e := FederationParse(f)
  if e != nil {
    return model.Job{}, e
  }
  job, clientE := dagschedulerclient.CreateJob(jobConf)
  if clientE != nil {
    errorMessage := clientE.Message()
    return job, &error.Error{
      Code: clientE.Code,
      Msg: errorMessage["message"].(string),
      Hits: errorMessage["hits"].(string),
    }
  }
  return job, nil
}


func JobSubmit(f form.JobSubmitForm) (form.JobSubmitForm, *error.Error) {
  _, e := FederationParse(form.JobCreateForm{
    Name: f.Name,
    RoleDag: f.RoleDag,
    Parameter: f.Parameter,
  })
  if e != nil {
    return f, e
  }
  a, er := PartyParse(f.PartyMap, f)
  if er != nil {
    return f, er
  }
  fmt.Println(len(a), a, er)
  for party_id, createForm := range a {
    TransferJob(party_id, createForm)
  }
  return f, nil
}
