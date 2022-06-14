package jobcontroller


var RoleList = []string {
  "HOST",
  "GUEST",
  "ARBITER",
}


type Parameter struct {
  Common        string          `json:"common"`
  Tasks         interface {}    `json:"tasks"`
}


type DagConf struct {
  Name          string          `json:"name"`
  Dag           interface {}    `json:"dag"`
  Parameter     interface {}    `json:"parameter"`
}
