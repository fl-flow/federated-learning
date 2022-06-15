package form

import ()


type JobCreateForm struct {
  Name          string                  `json:"name" binding:"required"`
  RoleDag       map[string]interface{}  `json:"role_dag" binding:"required"`
  Parameter     map[string]interface{}   `json:"role_parameter"`
}
