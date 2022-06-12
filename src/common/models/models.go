type BaseModel struct {
  ID        uint        `gorm:"primarykey"`
  CreatedAt time.Time
	UpdatedAt time.Time
}


type Job struct {
  BaseModel
  // TODO: Status
  Name            string
  Description     string
  Dag             []dagparser.TaskParsered      `gorm:"type:json"`
  Parameter       parameterparser.Parameter     `gorm:"type:json"`
}


type Task struct {
  BaseModel
  Job             Job           `gorm:"ForeignKey:JobId"`
  // TODO: Status
  Name            string
  Description     string
  Pid             int
  OrderInJob      int
  UpTasks         []string      `gorm:"type:json"`
  MemoryLimited   int
}


type TaskResult struct {
  BaseModel
  Task            Task    `gorm:"ForeignKey:TaskId"`
  Tag             string
  Ret             string
}
