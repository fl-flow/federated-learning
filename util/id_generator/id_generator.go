package idgenerator

import (
  "time"
)


func NewID() uint {
  timeUnixNano := time.Now().UnixNano()
  return uint(timeUnixNano)
}
