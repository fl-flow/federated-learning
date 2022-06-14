package form

import ()


type Parameter struct {
  Common          string                  `json:"common"`
  RoleParameter   map[string]interface{}  `json:"role_parameter"`
}


type JobCreateForm struct {
  Name          string                  `json:"name" binding:"required"`
  RoleDag       map[string]interface{}  `json:"role_dag" binding:"required"`
  Parameter     Parameter               `json:"parameter" binding:"required"`
}
