package mixin

import (
  "strings"
  "reflect"
)


func r(key string, ts reflect.Type, vs reflect.Value, m *(map[string]interface{})){
  for t := 0; t < ts.NumField(); t++ {
    n := ts.Field(t).Name
    if key == n {
      (*m)[key] = vs.Field(t).Interface()
      break
    } else if (n == "Model") {
      in := vs.Field(t).Interface()
      r(key, reflect.TypeOf(in), reflect.ValueOf(in), m)
    }
  }
}


func ParseStruct(rawData interface{}, fields string) map[string]interface{} {
  fieldsSplited := strings.Split(fields, "|")
  responseData := make(map[string]interface{})
  ts := reflect.TypeOf(rawData)
  vs := reflect.ValueOf(rawData)
  for _, f := range fieldsSplited {
    ss := strings.SplitN(f, ".", 2)
    s := ss[0]
    if len(ss) != 1 {
      for t := 0; t < ts.NumField(); t++ {
        if s == ts.Field(t).Name {
          d := vs.Field(t).Interface()
          if reflect.ValueOf(d).Kind() == reflect.Struct {
            responseData[s] = ParseStruct(d, ss[1])
          } else {
            responseData[s] = ParseSlice(d, ss[1])
          }
          break
        }
      }
    } else {
      r(s, ts, vs, &responseData)
      // for t := 0; t < ts.NumField(); t++ {
      //   fmt.Println(ts.Field(t).Name)
      //   n := ts.Field(t).Name
      //   if s == n {
      //     responseData[s] = vs.Field(t).Interface()
      //     break
      //   } else if (n == "Model") {
      //
      //   }
      // }
    }
  }
  return responseData
}


func ParseSlice(rawData interface{}, fields ...string) []map[string]interface{} {
  var responseDatas []map[string]interface{}
  dataSlice := reflect.ValueOf(rawData)

  for i := 0; i < dataSlice.Len(); i++ {
    responseData := ParseStruct(dataSlice.Index(i).Interface(), strings.Join(fields, "|"))
    responseDatas = append(responseDatas, responseData)
  }
  return responseDatas
}
