package middleware

import (
  "fmt"
  "github.com/gin-gonic/gin"
)


func AuthMiddleware(c *gin.Context) {
	fmt.Println("// TODO: auth middleware start")
	c.Next()
  fmt.Println("// TODO: auth middleware end")
}
