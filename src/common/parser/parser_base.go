package parser

import (
  "fl/common/parser/dag_parser"
  "fl/common/parser/parameter_parser"
)

type Conf struct {
  Dag         map[string](dagparser.DagTask)  `json:"dag"`
  Parameter   parameterparser.Parameter       `json:"parameter"`
}
