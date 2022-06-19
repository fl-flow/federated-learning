package mixin


const MaxSize int = 100
const DefaultSize int = 20


type PageNumberPagination struct {
  Page  int   `form:"page"`
  Size  int   `form:"size"`
}
