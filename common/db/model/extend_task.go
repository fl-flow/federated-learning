package model

import (
	"database/sql/driver"
)


type TaskStatusType	int

const (
	TaskReady			TaskStatusType = 2
)

func (c *TaskStatusType) Scan(value interface{}) error {
	*c = TaskStatusType(value.(int64))
  return nil
}

func (c TaskStatusType) Value() (driver.Value, error) {
  return int64(c), nil
}
