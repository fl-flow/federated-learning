package mixin

import (
  "fmt"
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
  "fl/common/error"
  "fl/http_server/http/response"
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


func CheckJSON(context *gin.Context, form any) bool {
	if e := context.ShouldBindJSON(form); e != nil {
    response.R(
      context,
      100,
      fmt.Sprintf("%v", e),
      fmt.Sprintf("%v", e),
    )
    return false
	}
	return true
}


func CommonResponse(context *gin.Context, res interface{}, er *error.Error) {
	if er != nil {
		response.R(
			context,
			er.Code,
			er.Message(),
			er.Message(),
		)
		return
	}
	response.R(
		context,
		0,
		"success",
		res,
	)
}
