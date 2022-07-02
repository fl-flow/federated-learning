package etc

import (
  "log"
  "fmt"
  "flag"
)


func init() {
  // server
  ip := flag.String("ip", IP, "ip")
	port := flag.Int("port", PORT, "port")

  localSchemaIP := flag.String("localip", LocalSchemaIP, "local ip")

  // server
	flag.Parse()

  IP = *ip
  PORT = *port
  LocalSchemaIP = *localSchemaIP

  log.Println(fmt.Sprintf(
      `
      // server
      IP: %v
      PORT: %v
      LocalSchemaIP: %v
      `,
      // server
      IP,
      PORT,
      LocalSchemaIP,
  ),)
}
