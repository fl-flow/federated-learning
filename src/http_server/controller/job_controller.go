package controller

import (
  "dag/common/db"
  "dag/common/parser"
  "dag/common/db/model"
  "dag/common/dag_error"
)


func JobCreate(name string, conf parser.Conf) *dagerror.DagError {
  // parse conf
  orderedTasks, _, error := parser.Parse(conf)
  if error != nil {
    return error
  }
  // Are there any tasks
  if len(orderedTasks) == 0 {
    return &dagerror.DagError{Code: 110010}
  }

  // TODO: cmd validate

  // insert to db (job and tasks)
  // insert job
  var job = model.Job {
    Name: name,
    Dag: orderedTasks,
    RawDagmap: conf.Dag,
    Parameter: model.JobParameter(conf.Parameter),
    Status: "init",
  }
  tx := db.DataBase.Begin()
  db.DataBase.Create(&job)
  // insert tasks
  var tasks []model.Task
  for index, t := range orderedTasks {
    var ups []string
    for _, up := range t.Depandent.Up {
      ups = append(ups, up.UpTask)
    }
    tasks = append(tasks, model.Task {
      Job: job,
      Name: t.Name,
      Description: "", // TODO:
      // Pid: nil,
      Status: "init",
      OrderInJob: index,
      UpTasks: ups,
      MemoryLimited: 0, // TODO:
      Cmd: t.Cmd,
      ValidateCmd: t.ValidateCmd,
    })
  }
  db.DataBase.Create(&tasks)
  tx.Commit()
  return nil
}
