package model

import (
  "time"
  "encoding/json"
  // "gorm.io/datatypes"
)


type BaseModel struct {
  ID        uint        `gorm:"primarykey;type:bigint auto_increment"`
  CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}


type Job struct {
  BaseModel
  ID              uint                `gorm:"primarykey;type:bigint auto_increment"`
  Name            string
  Conf            json.RawMessage     `gorm:"type:json"`
  Tasks           []Task
  Status          JobStatusType       `gorm:"type:tinyint"`
}


type Task struct {
  BaseModel
  ID              uint            `gorm:"primarykey;type:bigint auto_increment"`
  JobID           uint            `gorm:"type:bigint"`
  Job             Job             `json:"-"`
  Status          TaskStatusType  `gorm:"type:tinyint"`
  Party           string
  Role            string
  Name            string
}
