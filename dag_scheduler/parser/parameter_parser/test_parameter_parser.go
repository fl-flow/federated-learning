package parameterparser

import (
  "encoding/json"
  "fl/dag_scheduler/dag_error"
)


func TestParse(rawParameter string) (Parameter, error){
  var parameter Parameter
  ok := json.Unmarshal([]byte(rawParameter), &parameter)
  if ok != nil {
    return parameter, &dagError.DagError{Code: 12000}
  }
  parameter1, ok1 := Parse(parameter)
  if ok1 != nil {
    return parameter, ok1
  }
  return parameter1, nil
}
