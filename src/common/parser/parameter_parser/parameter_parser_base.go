package parameterparser


type Parameter struct {
  Common      string              `json:"common"`
  Tasks       map[string]string   `json:"tasks"`
}
