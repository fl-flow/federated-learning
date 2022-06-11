package dagParser


type DagTask struct {
  Input       []string  `json:"input"`
  Output      []string  `json:"output"`
  Cmd         string    `json:"cmd"`
}


type TaskDepandent struct {
  Up          []string
  Down        []string
}
