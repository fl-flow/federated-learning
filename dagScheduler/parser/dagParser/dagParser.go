package dagParser

import(
  "encoding/json";
  "fl/dagScheduler/dagError"
  "strings"
  "fmt"
)


func Parse(dag string) (error) {
  var dagMap map[string]DagTask
  ok := json.Unmarshal([]byte(dag), &dagMap)
  if ok != nil {
    return &dagError.DagError{Code: 10000}
  }

  tasksDepandentMap, inDegreeMap, error := findTasksDepandent(dagMap)
  if error != nil {
    return error
  }
  loopE := checkLoop(inDegreeMap, tasksDepandentMap)
  if loopE != nil {
    return loopE
  }
  return nil
}


func parseTaskDepandent(value string) (string, string, error) {
  // task.tag
  rets := strings.Split(value, ".")
  if (len(rets) != 2){
    return "", "", &dagError.DagError{
        Code: 11000,
        Msg: fmt.Sprintf(
            "error dagparser (%v ; required task.tag; task and tag can't contain '.')",
            value,
        ),
    }
  }
  return rets[0], rets[1], nil
}


func findTasksDepandent(
  dagTaskMap map[string]DagTask) (
  map[string]*TaskDepandent, map[string]int, error) {

  // get all tasks
  tasksDepandentMap := make(map[string]*TaskDepandent)
  inDegreeMap := make(map[string]int)
  for taskName, _ := range dagTaskMap {
    tasksDepandentMap[taskName] = &TaskDepandent{}
  }

  // build depandents
  for taskName, taskInfo := range dagTaskMap {
    input := taskInfo.Input
    for _, inputItem := range input {
      upTaskName, _, error := parseTaskDepandent(inputItem)
      if error != nil {
        return tasksDepandentMap, inDegreeMap, error
      }
      _, inputOk := tasksDepandentMap[upTaskName]
      if !inputOk {
        return tasksDepandentMap, inDegreeMap, &dagError.DagError{
            Code: 11010,
            Msg: fmt.Sprintf(
                "parser error( %v; task %v not exits )",
                inputItem,
                upTaskName,
            ),
        }
      }
      tasksDepandentMap[taskName].Up = append(tasksDepandentMap[taskName].Up, upTaskName)
      tasksDepandentMap[upTaskName].Down = append(tasksDepandentMap[upTaskName].Down, taskName)
    }
  }
  for taskName, taskInfo := range tasksDepandentMap {
    inDegreeMap[taskName] = len(taskInfo.Up)
  }
  return tasksDepandentMap, inDegreeMap, nil
}


func checkLoop(inDegreeMap map[string]int, tasksDepandentMap map[string]*TaskDepandent) error {
  queue := make([]string, 0)
  for taskName, inDegree := range inDegreeMap {
    if (inDegree == 0) {
      queue = append(queue, taskName)
    }
  }
  totals := 0
  qLength := len(queue)
  // TODO: order of task
  for (qLength > 0) {
    taskName := queue[0]
    queue = queue[1:]
    totals ++
    qLength --
    for _, downTaskName := range tasksDepandentMap[taskName].Down {
      inDegreeMap[downTaskName] --
      if inDegreeMap[downTaskName] == 0 {
        queue = append(queue, downTaskName)
        qLength ++
      }
    }
  }
  if (totals != len(tasksDepandentMap)){
    // TODO: find loop
    return &dagError.DagError{
        Code: 11020,
    }
  }
  return nil
}
