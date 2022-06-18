package jobcontroller

import (
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