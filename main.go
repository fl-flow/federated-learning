package main

import (
  "fl/dagScheduler/parser/dagParser"
  "fmt"
)


func main() {
      // unit test
      ok1 := dagParser.Parse(`{"a": 4}`)
      fmt.Println(ok1)
      ok2 := dagParser.Parse(`{"a": {"input": ["input_lalala"], "output": ["output_lalalal"]}, "b": {"input": ["input_a"], "output": ["output_b"]}}`)
      fmt.Println(ok2)
      ok3 := dagParser.Parse(`{"a": {"input": ["input_lalala.c"], "output": ["output_lalalal.d"]}, "b": {"input": ["input_a.e"], "output": ["output_b.f"]}}`)
      fmt.Println(ok3)
      ok4 := dagParser.Parse(`{"a": {"input": ["c.f"], "output": ["d"]}, "b": {"input": ["a.d"], "output": ["e"]}, "c": {"input": ["b.e"], "output": ["f"]}, "d": {"input": [], "output": ["f"]}}`)
      fmt.Println(ok4)

      ok := dagParser.Parse(`{"a": {"input": [], "output": ["d"]}, "b": {"input": ["a.d"], "output": ["e"]}, "c": {"input": ["b.e"], "output": ["f"]}, "d": {"input": ["a.d"], "output": ["f"]}}`)
      fmt.Println(ok, "ok")
}
