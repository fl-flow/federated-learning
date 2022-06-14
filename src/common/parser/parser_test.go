package parser

import (
  "fmt"
  "testing"
  "encoding/json"
  "dag/common/dag_error"
  "dag/common/parser/dag_parser"
  "dag/common/parser/parameter_parser"
)


func testParse(rawConf string) ([]dagparser.TaskParsered, parameterparser.Parameter, *dagerror.DagError){
  var conf Conf
  var tasks []dagparser.TaskParsered
  var parameters parameterparser.Parameter
  ok := json.Unmarshal([]byte(rawConf), &conf)
  if ok != nil {
    return tasks, parameters, &dagerror.DagError{Code: 10000}
  }
  return Parse(conf)
}


func TestParse(t *testing.T) {
  tasks, parameter, e := testParse(`
    {
      "dag": {
        "a": {
          "input": [],
          "output": ["d"],
          "cmd": "cmd"
        }
      },
      "parameter": {
        "common": "CCCC",
        "tasks": {"a": "z"}
      }
    }
  `)
  fmt.Println(e)
  fmt.Println(tasks)
  fmt.Println(parameter)
}
