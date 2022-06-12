package parser

import (
  "fl/dag_scheduler/parser/dag_parser"
  "fl/dag_scheduler/parser/parameter_parser"
)

type Conf struct {
  Dag         map[string](dagparser.DagTask)  `json:"dag"`
  Parameter   parameterparser.Parameter       `json:"parameter"`
}
