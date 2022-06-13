package model

import (
	"errors"
	"encoding/json"
	"database/sql/driver"
	"dag/common/parser/dag_parser"
	"dag/common/parser/parameter_parser"
)


func getReverseMap(m map[int]string) map[string]int {
	var rm = map[string]int {}
	for v, d := range m {
		rm[d] = v
	}
	return rm
}


type JobDag 				[]dagparser.TaskParsered

func (c JobDag) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c JobDag) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}


type JobRawDagmap   map[string](dagparser.DagTask)

func (c JobRawDagmap) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c JobRawDagmap) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}


type JobParameter   parameterparser.Parameter

func (c JobParameter) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c JobParameter) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}


type TaskUpTasks 		[]string

func (c TaskUpTasks) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c TaskUpTasks) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}


type JobStatus	string
var JobStatusMap = map[int]string {
	0: "init",
	1: "running",
	2: "success",
	3: "failed",
	4: "cancelled",
}
var JobStatusReverseMap map[string]int

func (c *JobStatus) Scan(value int) error {
	d, ok := JobStatusMap[value]
	if !ok {
		for _, v := range JobStatusMap {
			*c = JobStatus(v)
			return nil
		}
		return errors.New("error enum scan")
	}
	*c = JobStatus(d)
  return nil
}

func (c JobStatus) Value() (driver.Value, error) {
	d, ok := JobStatusReverseMap[string(c)]
	if !ok {
		for _, v := range JobStatusReverseMap {
			return int64(v), nil
		}
		return d, errors.New("error enum value")
	}
  return int64(d), nil
}


type TaskStatus		string
var TaskStatusMap = map[int]string {
	0: "init",
	1: "waiting",
	2: "ready",
	3: "running",
	4: "success",
	5: "failed",
	6: "cancelled",
}
var TaskStatusReverseMap map[string]int

func (c *TaskStatus) Scan(value int) error {
	d, ok := TaskStatusMap[value]
	if !ok {
		for _, v := range TaskStatusMap {
			*c = TaskStatus(v)
			return nil
		}
		return errors.New("error enum scan")
	}
	*c = TaskStatus(d)
  return nil
}

func (c TaskStatus) Value() (driver.Value, error) {
	d, ok := TaskStatusReverseMap[string(c)]
	if !ok {
		for _, v := range TaskStatusReverseMap {
			return int64(v), nil
		}
		return d, errors.New("error enum value")
	}
  return int64(d), nil
}
