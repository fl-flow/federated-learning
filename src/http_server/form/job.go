package form

import (
  "dag/common/parser"
)


type JobCreateForm struct {
  parser.Conf
  Name          string                `json:"name" binding:"required"`
}
