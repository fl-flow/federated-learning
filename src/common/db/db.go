package db

import (
  "log"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"

  "dag/common/db/model"
)


func init()  {
  db, error := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  DataBase = db
  if error != nil {
    log.Fatalln("error db connect")
  }
  db.AutoMigrate(
    &model.Job{},
    &model.Task{},
    &model.TaskResult{},
  )
}
