package form

import ()


type Kv map[string]interface{}


type RoleParameter struct {
  RoleParameter   map[string]Kv           `json:"role_parameter" binding:"required"`
  Common          string                  `json:"common" binding:"required"`
}


type JobCreateForm struct {
  Name          string                  `json:"name" binding:"required"`
  RoleDag       map[string]Kv           `json:"role_dag" binding:"required"`
  Parameter     RoleParameter           `json:"parameter" binding:"required"`
}
