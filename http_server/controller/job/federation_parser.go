package jobcontroller

import (
  "fl/common/error"
  "fl/http_server/form"
)


func FederationParse(f form.JobCreateForm) (DagConf, *error.Error) {
  var dagConf DagConf

  roleDag := f.RoleDag
  roleParameter := f.Parameter

  dagRoles, e := getRoles(roleDag)
  if e != nil {
    return dagConf, e
  }
  parameterRoles, pe := getRoles(roleParameter)
  if pe != nil {
    return dagConf, pe
  }
  for _, pRole := range parameterRoles{
    has := inArray(pRole, dagRoles)
    if !has {
      return dagConf, &error.Error{
              Code: 101020,
              Hits: pRole,
          }
    }
  }
  return buildDagConf(f), nil
}


func inArray(item string, array []string) bool {
  var has bool = false
  for _, item_ := range array {
    if item == item_ {
      has = true
      break
    }
  }
  return has
}


func getRoles(roleValueMap map[string]interface{}) ([]string, *error.Error) {
  var roles []string
  for role, _ := range roleValueMap {
    has := inArray(role, RoleList)
    if has {
      roles = append(roles, role)
    }else {
      return roles, &error.Error{
              Code: 101010,
              Hits: role,
          }
    }
  }
  return roles, nil
}


func buildDagConf(f form.JobCreateForm) DagConf {
  return DagConf {
    Name: f.Name,
    Dag: f.RoleDag,
    Parameter: f.Parameter,
  }
}
