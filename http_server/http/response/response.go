package response

import (
  "net/http"
  "github.com/gin-gonic/gin"
)


func R(c *gin.Context, code int, msg interface{}, data interface{})  {
  var httpStatus int
  if code == 0 {
    httpStatus = http.StatusOK
  }else {
    httpStatus = http.StatusBadRequest
  }
	c.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}


type Ret struct {
  Code    int             `json:"code"`
  Data    interface{}     `json:"data"`
  Msg     interface{}     `json:"msg"`
}
