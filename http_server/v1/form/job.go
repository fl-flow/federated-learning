package form

import ()


type Kv map[string]interface{}


type CommonParameter struct {
  Communication struct {
    Engine string `json:"engine" binding:"required"`
  } `json:"communication" binding:"required"`
  Storage struct {
    Engine string `json:"engine" binding:"required"`
  } `json:"storage" binding:"required"`
  Computing struct {
    Engine string `json:"engine" binding:"required"`
  } `json:"computing" binding:"required"`
  PartyMap map[string]([]string) `json:"party_map,omitempty" binding:"-"`
}


type RoleParameter struct {
  RoleParameter   map[string]interface{}  `json:"role_parameter" binding:"required"`
  Common          CommonParameter         `json:"common" binding:"required"`
}


type JobCreateRawConf struct {
  Name          string                  `json:"name" binding:"required"`
  RoleDag       map[string]Kv           `json:"role_dag" binding:"required"`
  Parameter     RoleParameter           `json:"parameter" binding:"required"`
}


type JobForm struct {
  JobCreateRawConf
  PartyMap        map[string]([]string) `json:"party_map" binding:"required"`
}
