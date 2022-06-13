package main

import (
  _ "dag/common/db"
  "dag/http_server"
)


func main() {
  httpserver.Run()
}
