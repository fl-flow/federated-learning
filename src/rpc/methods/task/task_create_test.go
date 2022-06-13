package rpctask

import (
  "fmt"
  "testing"
  "net/rpc/jsonrpc"
  "dag/common/parser"
)

func TestCreateTask(t *testing.T) {
  client, err := jsonrpc.Dial("tcp", "127.0.0.1:8000")
  if err != nil {
      t.Fatal("dialing:", err)
  }
  args := &parser.Conf{}
  var reply string
  err = client.Call("Task.Create", args, &reply)
  if err != nil {
      t.Fatal("error:", err)
  }
  fmt.Printf("response: %v \n", reply)
  // // Asynchronous call
  // quotient := new(Quotient)
  // divCall := client.Go("Arith.Divide", args, quotient, nil)
  // replyCall := <-divCall.Done // will be equal to divCall
  // if replyCall.Error != nil {
  //     log.Fatal("arith error:", replyCall.Error)
  // }
  // fmt.Printf("Arith: %d/%d=%d...%d", args.A, args.B, quotient.Quo, quotient.Rem)
  // // check errors, print, etc.
}
