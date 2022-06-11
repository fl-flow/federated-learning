package dagError


var Conf = map[int]string {
  0: "success",

  10000: "parser error", // base
  11000: "parser error( task.tag )", // base
  11010: "parser error( task.tag; task not exits)",
  11020: "parser error( loop found)",
}
