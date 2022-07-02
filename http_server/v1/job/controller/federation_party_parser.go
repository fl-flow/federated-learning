package jobcontroller

import (
  "fmt"
  "fl/common/error"
  "fl/http_server/v1/job/form"
)


func PartyParse (
  role2PartyMap map[string]([]string),
  f form.JobForm,
) (map[string]form.JobCreateRawConf, *error.Error) {
  party2RoleMap := transferRole2PartToParty2RoleMap(role2PartyMap)
  party2Form := make(map[string]form.JobCreateRawConf)
  common := f.Parameter.Common
  for party, roles := range party2RoleMap {
    jcf := form.JobCreateRawConf {
      Name: f.Name,
      Parameter: form.RoleParameter{
        Common: common,
        RoleParameter: make(map[string]interface{}),
      },
      RoleDag: make(map[string]form.Task2Dag),
    }
    for _, r := range *roles {
      if f.RoleDag[r] == nil {
        return party2Form, &error.Error{
                Code: 102010,
                Hits: r,
            }
      }
      if f.Parameter.RoleParameter[r] == nil {
        return party2Form, &error.Error{
                Code: 102010,
                Hits: r,
            }
      }
      jcf.RoleDag[r] = f.RoleDag[r]
      p, e := processParameter(f.Parameter.RoleParameter[r], party)
      if e != nil {
        return party2Form, e
      }
      jcf.Parameter.RoleParameter[r] = p
    }
    party2Form[party] = jcf
  }
  return party2Form, nil
}


func transferRole2PartToParty2RoleMap(
  role2PartyMap map[string]([]string),
) map[string](*[]string) {
  party2RoleMap := map[string](*[]string){}
  for role, parties := range role2PartyMap {
    for _, p := range parties {
      // TODO: validate parties // ip ?
      if party2RoleMap[p] == nil {
        party2RoleMap[p] = &([]string{})
      }
      *(party2RoleMap[p]) = append(*(party2RoleMap[p]), role)
    }
  }
  return party2RoleMap
}


func processParameter(parameter interface{}, party string) (interface{}, *error.Error) {
  switch parameter.(type) {
    case map[string]interface{}:
      return parameter, nil
    case [](interface{}):
      tmp := parameter.([]interface{})
      for _, pParameter_ := range tmp {
        switch pParameter_.(type) {
          case map[string]interface{}:
            fmt.Println(parameter)
          default:
            return pParameter_, &error.Error{
              Code: 102021,
              Hits: fmt.Sprintf("%v", pParameter_),
            }
        }
        pParameter := pParameter_.(map[string]interface{})
        p := pParameter["party"]
        switch p.(type) {
          case string:
            fmt.Println(p)
          default:
            return pParameter, &error.Error{
              Code: 102020,
              Hits: fmt.Sprintf("%v", p),
            }
        }
        if p.(string) == party {
          if pParameter["data"] == nil {
            return pParameter, &error.Error{
              Code: 102025,
              Hits: fmt.Sprintf("%v", p),
            }
          }
          return pParameter["data"], nil
        }
      }
      return parameter, &error.Error{
        Code: 102030,
        Hits: party,
      }
    default:
      return parameter, &error.Error{
        Code: 102015,
        Hits: fmt.Sprintf("%v", parameter),
      }
  }
}
