package jobcontroller


var RoleList = []string {
  "HOST",
  "GUEST",
  "ARBITER",
}


type DagConf struct {
  Name          string              `json:"name"`
  Dag           interface {}        `json:"dag"`
  Parameter     interface {}        `json:"parameter"`
}
