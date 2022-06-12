package dagparser

import(
  "fmt"
  "strings"
  "fl/common/dag_error"
)


func Parse(dagMap map[string]DagTask) ([]TaskParsered, error) {
  var tasksParsed []TaskParsered
  tasksDepandentMap, inDegreeMap, error := findTasksDepandent(dagMap)
  if error != nil {
    return tasksParsed, error
  }
  orderedTasks, loopE := checkLoop(inDegreeMap, tasksDepandentMap)
  if loopE != nil {
    return tasksParsed, loopE
  }
  tasksParsed = buildTaskParsered(
      orderedTasks,
      tasksDepandentMap,
      dagMap,
  )
  return tasksParsed, nil
}


func parseTaskDepandent(value string) (string, string, error) {
  // task.tag
  rets := strings.Split(value, ".")
  if (len(rets) != 2){
    return "", "", &dagError.DagError{
        Code: 11010,
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
    if taskInfo.Cmd == "" {
      return tasksDepandentMap, inDegreeMap, &dagError.DagError{
          Code: 11040,
          Msg: fmt.Sprintf(
              "parser error( task %v's cmd is required )",
              taskName,
          ),
      }
    }
    input := taskInfo.Input
    for _, inputItem := range input {
      upTaskName, upTag, error := parseTaskDepandent(inputItem)
      if error != nil {
        return tasksDepandentMap, inDegreeMap, error
      }
      _, inputOk := tasksDepandentMap[upTaskName]
      if !inputOk {
        return tasksDepandentMap, inDegreeMap, &dagError.DagError{
            Code: 11020,
            Msg: fmt.Sprintf(
                "parser error( %v; task %v not exits )",
                inputItem,
                upTaskName,
            ),
        }
      }
      tasksDepandentMap[taskName].Up = append(
        tasksDepandentMap[taskName].Up,
        TaskInput {
          UpTask: upTaskName,
          Tag: upTag,
        })
      tasksDepandentMap[upTaskName].Down = append(tasksDepandentMap[upTaskName].Down, taskName)
    }
  }
  for taskName, taskInfo := range tasksDepandentMap {
    inDegreeMap[taskName] = len(taskInfo.Up)
  }
  return tasksDepandentMap, inDegreeMap, nil
}


func checkLoop(inDegreeMap map[string]int, tasksDepandentMap map[string]*TaskDepandent) ([]string, error) {
  var queue, orderedTasks []string
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
    orderedTasks = append(orderedTasks, taskName)
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
    return orderedTasks, &dagError.DagError{
        Code: 11030,
    }
  }
  return orderedTasks, nil
}


func buildTaskParsered(
    orderedTasks []string,
    tasksDepandentMap map[string]*TaskDepandent,
    dagTaskMap map[string]DagTask,
) []TaskParsered {
    var tasksParsed []TaskParsered
    for _, taskName := range orderedTasks {
      tasksParsed = append(
        tasksParsed,
        TaskParsered {
          Name: taskName,
          Depandent: tasksDepandentMap[taskName],
          Output: dagTaskMap[taskName].Output,
          Cmd: dagTaskMap[taskName].Cmd,
        },
      )
    }
    return tasksParsed
}
