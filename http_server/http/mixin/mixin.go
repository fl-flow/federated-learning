package mixin

import (
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/dag-scheduler/http_server/http/response"
)


func List(context *gin.Context, qs *gorm.DB)(
    *gorm.DB, int64, int, int,
) {
  var pagination PageNumberPagination
  var total int64
  qs.Count(&total)
  context.ShouldBindQuery(&pagination)
  size := pagination.Size
  if size <= 0 {
    size = DefaultSize
  }else if size > MaxSize {
    size = MaxSize
  }
  page := pagination.Page
  if page <= 0 {
    page = 1
  }
  qs = qs.Offset((page - 1) * size).Limit(size)
  return qs, total, page, size

}


func ListResponse(
    context *gin.Context,
    data interface{},
    total int64,
    page int,
    size int,
) {
  response.R(
    context,
    0,
    "success",
    map[string]interface{}{
      "count": total,
      "list": data,
      "page": page,
      "size": size,
    },
  )
}
