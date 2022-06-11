package dagError

import (
  "fmt"
)


type DagError struct {
  Code  int
  Msg   string
}


func (e *DagError) Error() string {
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
