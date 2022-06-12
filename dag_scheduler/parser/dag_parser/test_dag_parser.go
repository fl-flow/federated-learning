package dagparser

import(
  "encoding/json"
  "fl/dag_scheduler/dag_error"
)


func TestParse(dag string) ([]TaskParsered, error) {
  var dagMap map[string]DagTask
  ok := json.Unmarshal([]byte(dag), &dagMap)
  var tasksParsed []TaskParsered
  if ok != nil {
    return tasksParsed, &dagError.DagError{Code: 11000}
  }
  tasksParsed, e := Parse(dagMap)
  if e != nil {
    return tasksParsed, e
  }
  return tasksParsed, e
}
