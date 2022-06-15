package jobcontroller

import (
  "github.com/fl-flow/dag-scheduler/dag_scheduler_client"

  "fl/http_server/form"
  "fl/common/error"
)


func JobCreate(f form.JobCreateForm) *error.Error {
  jobConf, e := FederationParse(f)
  if e != nil {
    return e
  }
  _, clientE := dagschedulerclient.CreateJob(jobConf)
  if clientE != nil {
    errorMessage := clientE.Message()
    return &error.Error{
      Code: clientE.Code,
      Msg: errorMessage["message"].(string),
      Hits: errorMessage["hits"].(string),
    }
  }
  return nil
}
