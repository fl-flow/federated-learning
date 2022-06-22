package model

import (
  "time"
  "gorm.io/datatypes"
)


type BaseModel struct {
  ID        uint        `gorm:"primarykey"`
  CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}


type Job struct {
  BaseModel
  Name            string
  Conf            datatypes.JSON      `gorm:"type:json"`
  Tasks           []Task
}


type JobStatus struct {
  BaseModel
  Job             Job
  Status          int
}


type Task struct {
  BaseModel
  ID              uint          `gorm:"primarykey"`
  JobID           uint
  Job             Job
  Party           string
  Role            string
  Name            string
}


type TaskStatus struct {
  BaseModel
  Task            Task
  Status          TaskStatusType    `gorm:"type:int"`
}
