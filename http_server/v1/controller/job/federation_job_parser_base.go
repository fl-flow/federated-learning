package jobcontroller


var RoleList = []string {
  "HOST",
  "GUEST",
  "ARBITER",
}


type TaskConf struct {
  Input       []string  `json:"input" binding:"required"`
  Output      []string  `json:"output" binding:"required"`
  Cmd         []string  `json:"cmd" binding:"required"`
  ValidateCmd string    `json:"validate_cmd" binding:"required"`
}

type Task2TaskConf map[string]TaskConf


type Role2Task2TaskConf map[string]Task2TaskConf


type DagConf struct {
  Name          string              `json:"name"`
  Dag           Role2Task2TaskConf  `json:"dag"`
  Parameter     interface {}        `json:"parameter"`
}
