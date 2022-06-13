package dagerror


var Conf = map[int]string {
  0: "success",

  10000: "parser error",

  11000: "dag parser error", // base
  11010: "dag parser error( task.tag )",
  11020: "dag parser error( task.tag; task not exits)",
  11030: "dag parser error( loop found)",
  11040: "dag parser error( cmd is required )",

  12000: "parameter parser error", // base
  12010: "num of parameter is not equal to num of dag",
  12020: "num of parameter is not equal to num of dag",

  110000: "job http api error", // base
  110010: "job http api error (no tasks)",

}
