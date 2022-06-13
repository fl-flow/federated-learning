package dagrpc

import (
  "log"
  "net"
  "net/rpc"
  "net/rpc/jsonrpc"

  "dag/rpc/methods/task"
)


func RunRpc() {
  rpc.RegisterName("Task", &rpctask.Task{})

  tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000");
  if err != nil {
    log.Fatal(err);
  }
  tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
  if err2 != nil {
    log.Fatal(err2);
  }
  for {
    conn, err3 := tcplisten.Accept();
    if err3 != nil {
      log.Println(err3);
      continue;
    }
    go jsonrpc.ServeConn(conn);
  }
}
