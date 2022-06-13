package model


func init()  {
  JobStatusReverseMap = getReverseMap(JobStatusMap)
  TaskStatusReverseMap = getReverseMap(TaskStatusMap)
}
