package error


var Conf = map[int]string {
  0: "success",

  100000: "http api error",

  101000: "http api error (job)", // base
  101010: "http api error; job.create role is not permitted",
  101020: "http api error; job.create role is not in dag",

  102010: "http api error; job.submit role error",
  102015: "http api error; job.submit role error; required dict or list",
  102020: "http api error; job.submit parameter party error; party required string",
  102021: "http api error; job.submit parameter party error; required dict",
  102025: "http api error; job.submit parameter party error; data required",
  102030: "http api error; job.submit parameter party error; party is not existed",
  102050: "http api error; job.submit transfer remote error",
}
