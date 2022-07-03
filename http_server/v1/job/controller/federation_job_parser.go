package jobcontroller

import (
  "fmt"
  "strings"
  "reflect"
  "encoding/json"
  "fl/common/error"
  "fl/http_server/v1/job/form"
)


func FederationParse(f form.JobCreateRawConf, partyMap map[string]([]string)) (DagConf, *error.Error) {
  var dagConf DagConf

  roleDag := f.RoleDag
  roleParameter := f.Parameter.RoleParameter

  dagRoles, e := getRoles(reflect.ValueOf(roleDag).MapKeys())
  if e != nil {
    return dagConf, e
  }
  parameterRoles, pe := getRoles(reflect.ValueOf(roleParameter).MapKeys())
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
  df, e := buildDagConf(f, partyMap)
  if e != nil {
    return dagConf, e
  }
  return df, nil
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


func getRoles(roless []reflect.Value) ([]string, *error.Error) {
  var roles []string
  for _, role_ := range roless {
    role := string(role_.String())
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


func buildDagConf(f form.JobCreateRawConf, partyMap map[string]([]string)) (DagConf, *error.Error) {
  var common form.CommonParameter = f.Parameter.Common
  common.PartyMap = partyMap
  commonByte, _ := json.Marshal(common)
  RoleParameter := make(map[string](map[string]interface{}))
  for role, v := range f.Parameter.RoleParameter {
    RoleParameter[role] = map[string]interface{}{}
    RoleParameter[role]["tasks"] = v
    RoleParameter[role]["common"] = string(commonByte)
  }
  d, e := transferRoleDag(f.RoleDag)
  if e != nil {
    return DagConf{}, e
  }
  return DagConf {
    Name: f.Name,
    Dag: d,
    Parameter: RoleParameter,
    WaitCmdToRun: true,
  }, nil
}


func buildInputTag(taskAndTag string, task2Dag form.Task2Dag, type_ string) (string, *error.Error) {
  rets := strings.Split(taskAndTag, ".")
  if (len(rets) != 2){
    return "", &error.Error{
        Code: 101030, // TODO:
        Hits: taskAndTag,
    }
  }
  task, tag := rets[0], rets[1]
  var output []string
  if type_ == "data" {
    output = task2Dag[task].Output.Data
  } else if type_ == "model" {
    output = task2Dag[task].Output.Model
  } else {
    return "", &error.Error{
        Code: 101030, // TODO:
        Hits: taskAndTag,
    }
  }
  index := -1
  for i, outputTag := range output {
    if outputTag == tag {
      index = i
    }
  }
  if index == -1 {
    return "", &error.Error{
        Code: 101030, // TODO:
        Hits: taskAndTag,
    }
  }
  return fmt.Sprintf("%v.%v:%v", task, type_, index), nil
}


func transferRoleDag(role2task2Dag map[string]form.Task2Dag) (Role2Task2TaskConf, *error.Error) {
  fmt.Println(role2task2Dag, "zzzzz")
  role2Task2TaskConf := make(Role2Task2TaskConf)
  for role, task2Dag := range role2task2Dag {
    role2Task2TaskConf[role] = make(Task2TaskConf)
    for t, taskConf := range task2Dag {
      var inputs []string
      for _, inputData := range taskConf.Input.Data {
        in, e := buildInputTag(inputData, task2Dag, "data")
        if e != nil {
          return role2Task2TaskConf, e
        }
        inputs = append(
          inputs,
          in,
        )
      }
      for _, inputData := range taskConf.Input.Model {
        in_, err := buildInputTag(inputData, task2Dag, "model")
        if err != nil {
          return role2Task2TaskConf, err
        }
        inputs = append(
          inputs,
          in_,
        )
      }
      role2Task2TaskConf[role][t] = TaskConf {
        Input: inputs,
        Output: []string {
          "data",
          "model",
        },
        Cmd: taskConf.Cmd,
        ValidateCmd: taskConf.ValidateCmd,
      }
    }
  }
  return role2Task2TaskConf, nil
}
