package error


var Conf = map[int]string {
  0: "success",

  100000: "http api error",

  101000: "http api error (job)", // base
  101010: "http api error; job.create role is not permitted",
  101020: "http api error; job.create role is not in dag",
}
