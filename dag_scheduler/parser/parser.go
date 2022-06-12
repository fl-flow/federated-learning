package parser

import (
  "fmt"
  "encoding/json"
  "fl/dag_scheduler/dag_error"
  "fl/dag_scheduler/parser/dag_parser"
  "fl/dag_scheduler/parser/parameter_parser"
)


func Parse(rawConf string) ([]dagparser.TaskParsered, parameterparser.Parameter, error){
  var conf Conf
  var tasks []dagparser.TaskParsered
  var parameters parameterparser.Parameter
  ok := json.Unmarshal([]byte(rawConf), &conf)
  if ok != nil {
    return tasks, parameters, &dagError.DagError{Code: 10000}
  }

  tasks, dagError := dagparser.Parse(conf.Dag)
  if dagError != nil {
    return tasks, parameters, dagError
  }
  parameters, parameterError := parameterparser.Parse(conf.Parameter)
  if parameterError != nil {
    return tasks, parameters, parameterError
  }
  error := checkDagParameter(tasks, parameters)
  if error != nil {
    return tasks, parameters, error
  }
  return tasks, parameters, nil
}


func checkDagParameter(
  tasks []dagparser.TaskParsered,
  parameters parameterparser.Parameter) error {
  if len(tasks) != len(parameters.Tasks) {
    return &dagError.DagError{Code: 12010}
  }
  for _, task := range tasks {
    _, ok := parameters.Tasks[task.Name]
    if !ok {
      return &dagError.DagError{
        Code: 12020,
        Msg: fmt.Sprintf(
            "dag's task %v is not in parameters",
            task.Name,
        ),
      }
    }
  }
  return nil
}
