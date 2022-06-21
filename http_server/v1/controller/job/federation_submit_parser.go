package jobcontroller

import (
  "fl/common/error"
  "fl/http_server/v1/form"
)


func PartyParse (
  role2PartyMap map[string]([]string),
  f form.JobSubmitForm,
) (map[string]form.JobCreateForm, *error.Error) {
  party2RoleMap := transferRole2PartToParty2RoleMap(role2PartyMap)
  party2Form := make(map[string]form.JobCreateForm)
  common := f.Parameter.Common
  for party, roles := range party2RoleMap {
    jcf := form.JobCreateForm {
      Name: f.Name,
      Parameter: form.RoleParameter{
        Common: common,
        RoleParameter: make(map[string]form.Kv),
      },
      RoleDag: make(map[string]form.Kv),
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
      jcf.Parameter.RoleParameter[r] = f.Parameter.RoleParameter[r]
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
