package main

import (
  "fl/dag_scheduler/parser"
  "fl/dag_scheduler/parser/dag_parser"
  "fl/dag_scheduler/parser/parameter_parser"
  "fmt"
)


func main() {
      // dag parser unit test
      _, ok1 := dagparser.TestParse(`{"a": 4}`)
      fmt.Println(ok1)
      _, ok2 := dagparser.TestParse(`
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
      _, ok3 := dagparser.TestParse(`
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
      _, ok4 := dagparser.TestParse(`
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
      _, ok5 := dagparser.TestParse(`
        {
          "a": {
            "input": ["input_lalala"],
            "output": ["output_lalalal"]
          }
        }
      `)
      fmt.Println(ok5)

      task, ok := dagparser.TestParse(`
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


      // parameter parser unit test
      _, parameterparserok1 := parameterparser.TestParse(`
        {
          "common": "CCCC",
          "tasks": ""
        }
      `)
      fmt.Println(parameterparserok1)

      parameterparser, parameterparserok := parameterparser.TestParse(`
        {
          "common": "CCCC",
          "tasks": {"a": "z", "f": "d"}
        }
      `)
      fmt.Println(parameterparserok)
      fmt.Println(parameterparser)


      // parser unit test
      tasks, parameter, e := parser.Parse(`
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
