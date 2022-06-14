package jobcontroller

import (
  "fl/common/error"
  "fl/http_server/form"
)


func FederationParse(f form.JobCreateForm) ([]DagConf, *error.Error) {
  var dagConfs []DagConf

  roleDag := f.RoleDag
  roleParameter := f.Parameter.RoleParameter

  dagRoles, e := getRoles(roleDag)
  if e != nil {
    return dagConfs, e
  }
  parameterRoles, pe := getRoles(roleParameter)
  if pe != nil {
    return dagConfs, pe
  }
  for _, pRole := range parameterRoles{
    has := inArray(pRole, dagRoles)
    if !has {
      return dagConfs, &error.Error{
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


func buildDagConf(f form.JobCreateForm) []DagConf {
  var dagConfs []DagConf
  name := f.Name
  roleDag := f.RoleDag
  commonParameter := f.Parameter.Common
  roleParameter := f.Parameter.RoleParameter
  for role, dag := range roleDag {
    parameter := Parameter {
      Common: commonParameter,
      Tasks: roleParameter[role],
    }
    dagConf := DagConf {
      Name: name + role,
      Dag: dag,
      Parameter: parameter,
    }
    dagConfs = append(dagConfs, dagConf)
  }
  return dagConfs
}
