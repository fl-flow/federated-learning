package dagparser

import(
  "encoding/json"
  "dag/common/dag_error"
  "testing"
  "fmt"
)


func testParse(dag string) ([]TaskParsered, *dagerror.DagError) {
  var dagMap map[string]DagTask
  ok := json.Unmarshal([]byte(dag), &dagMap)
  var tasksParsed []TaskParsered
  if ok != nil {
    return tasksParsed, &dagerror.DagError{Code: 11000}
  }
  tasksParsed, e := Parse(dagMap)
  if e != nil {
    return tasksParsed, e
  }
  return tasksParsed, e
}


func TestParse(t *testing.T) {
  _, ok1 := testParse(`{"a": 4}`)
  fmt.Println(ok1)
  _, ok2 := testParse(`
    {
      "a": {
        "input": ["input_lalala"],
        "output": ["output_lalalal"],
        "cmd": "cmd"
      },
      "b": {
        "input": ["input_a"],
        "output": ["output_b"],
        "cmd": "cmd"
      }
    }
  `)
  fmt.Println(ok2)
  _, ok3 := testParse(`
    {
      "a": {
        "input": ["input_lalala.c"],
        "output": ["output_lalalal.d"],
        "cmd": "cmd"
      },
      "b": {
        "input": ["input_a.e"],
        "output": ["output_b.f"],
        "cmd": "cmd"
      }
    }
  `)
  fmt.Println(ok3)
  _, ok4 := testParse(`
    {
      "a": {
        "input": ["c.f"],
        "output": ["d"],
        "cmd": "cmd"
      },
      "b": {
        "input": ["a.d"],
        "output": ["e"],
        "cmd": "cmd"
      },
      "c": {
        "input": ["b.e"],
        "output": ["f"],
        "cmd": "cmd"
      },
      "d": {
        "input": [],
        "output": ["f"],
        "cmd": "cmd"
      }
    }
  `)
  fmt.Println(ok4)
  _, ok5 := testParse(`
    {
      "a": {
        "input": ["input_lalala"],
        "output": ["output_lalalal"]
      }
    }
  `)
  fmt.Println(ok5)

  task, ok := testParse(`
    {
      "a": {
        "input": ["d.f"],
        "output": ["d"],
        "cmd": "cmd"
      },
      "b": {
        "input": ["d.f"],
        "output": ["e"],
        "cmd": "cmd"
      },
      "c": {
        "input": ["b.e"],
        "output": ["f"],
        "cmd": "cmd"
      },
      "d": {
        "input": [],
        "output": ["f"],
        "cmd": "cmd"
      }
    }
  `)
  fmt.Println(ok, "ok")
  fmt.Println(task, "ok")
}
