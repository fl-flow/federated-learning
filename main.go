package main

import (
  "flag"

  "fl/http_server"
)


func main() {
  ip := flag.String("ip", "127.0.0.1", "ip")
	port := flag.Int("port", 8443, "port")
	flag.Parse()

  httpserver.Run(*ip, *port)
}
