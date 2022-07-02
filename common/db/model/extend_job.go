package model

import (
	"database/sql/driver"
)


type JobStatusType int

const (
	JobInit			JobStatusType = 1
)

func (c *JobStatusType) Scan(value interface{}) error {
	*c = JobStatusType(value.(int64))
  return nil
}

func (c JobStatusType) Value() (driver.Value, error) {
  return int64(c), nil
}
