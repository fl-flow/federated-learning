package rpctask

import (
  "fl/common/parser"
)


func (task *Task) Create(conf parser.Conf, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello world"
  // for _, v := range kv {
  //   fmt.Println(v)
  // }
	return nil
}
