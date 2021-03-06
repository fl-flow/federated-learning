package form

import ()


type Dag struct {
  Input   struct {
    Data    []string  `json:"data" binding:"required"`
    Tensor  []string  `json:"tensor" binding:"required"`
    Model   []string  `json:"model" binding:"required"`
  } `json:"input" binding:"required"`
  Output   struct {
    Data    []string  `json:"data" binding:"required"`
    Tensor  []string  `json:"tensor" binding:"required"`
    Model   []string  `json:"model" binding:"required"`
  } `json:"output" binding:"required"`
  Cmd         []string  `json:"cmd" binding:"required"`
  ValidateCmd string    `json:"validate_cmd" binding:"required"`
}


type Task2Dag map[string]Dag


type CommonParameter struct {
  Communication struct {
    Engine string `json:"engine" binding:"required"`
  } `json:"communication" binding:"required"`
  // Storage struct {
  //   Engine string `json:"engine" binding:"required"`
  // } `json:"storage" binding:"required"`
  // Computing struct {
  //   Engine string `json:"engine" binding:"required"`
  // } `json:"computing" binding:"required"`
  PartyMap map[string]([]string) `json:"party_map,omitempty" binding:"-"`
}


type RoleParameter struct {
  RoleParameter   map[string]interface{}  `json:"role_parameter" binding:"required"`
  Common          CommonParameter         `json:"common" binding:"required"`
}


type JobCreateRawConf struct {
  Name          string                  `json:"name" binding:"required"`
  RoleDag       map[string]Task2Dag     `json:"role_dag" binding:"required"`
  Parameter     RoleParameter           `json:"parameter" binding:"required"`
}


type JobForm struct {
  JobCreateRawConf
  PartyMap        map[string]([]string) `json:"party_map" binding:"required"`
  ID              uint                  `json:"id"`
}


type TaskNotify struct {
  Status  int             `json:"status"`
  Type    string          `json:"type"`
  ID      uint            `json:"id"`
  Extra   struct {
    JobID     uint      `json:"job_id"`
    Group     string    `json:"group"`
    Task      string    `json:"task"`
  } `json:"extra"`
}


type JobNotify struct {
  Status  int             `json:"status"`
  Type    string          `json:"type"`
  ID      uint            `json:"id"`
}


type FederatedTaskNotify struct {
  Status    int       `json:"status"`
  JobID     uint      `json:"job_id"`
  Group     string    `json:"group"`
  Task      string    `json:"task"`
  Party     string    `json:"party"`
}
