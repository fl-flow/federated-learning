package error

import (
  "fmt"
)


type Error struct {
  Code  int
  Msg   string
  Hits  string
}


func (e *Error) Error() string {
  var msg string
  if(e.Msg==""){
    msg = Conf[e.Code]
  } else {
    msg = e.Msg
  }
  return fmt.Sprintf(
      `{"code": %d, "msg": %v}`,
      e.Code,
      msg,
  )
}


func (e *Error) Message() interface{} {
  return map[string]interface{} {
    "message": Conf[e.Code],
    "hits": e.Hits,
  }
}
