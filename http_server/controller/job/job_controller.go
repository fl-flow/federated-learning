package jobcontroller

import (
  "fmt"

  "fl/http_server/form"
  "fl/common/error"
)


func JobCreate(f form.JobCreateForm) *error.Error {
  jobConfs, e := FederationParse(f)
  if e != nil {
    return e
  }
  fmt.Println(jobConfs, "jobConfs")
  return nil
}
