package dagparser


type DagTask struct {
  Input       []string  `json:"input"`
  Output      []string  `json:"output"`
  Cmd         string    `json:"cmd"`
  ValidateCmd string    `json:"validate_cmd"`
}


type TaskInput struct {
  UpTask        string
  Tag           string
}


type TaskDepandent struct {
  Up          []TaskInput
  Down        []string
}


type TaskParsered struct {
  Name          string
  Depandent     TaskDepandent
  Output        []string
  Cmd           string
  ValidateCmd   string
}
