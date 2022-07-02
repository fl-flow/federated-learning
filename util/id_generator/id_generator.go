package idgenerator

import (
  "time"
)


func NewID() uint {
  // timeUnixNano := time.Now().UnixNano()
  timeUnixNano := time.Now().Unix()
  return uint(timeUnixNano)
}
