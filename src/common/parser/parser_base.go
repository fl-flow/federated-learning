package parser

import (
  "dag/common/parser/dag_parser"
  "dag/common/parser/parameter_parser"
)

type Conf struct {
  Dag         map[string](dagparser.DagTask)  `json:"dag"`
  Parameter   parameterparser.Parameter       `json:"parameter"`
}
