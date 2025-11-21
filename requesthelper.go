package requesthelper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetJSONData[T any](c *gin.Context) (T, error) {
	var result T

	return result, nil
}
func GetLimitAndOffset(c *gin.Context) (int, int) {
	limit, _ := strconv.Atoi(c.DefaultQuery("page-size", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// calculate offset form page and limit
	offset := (page - 1) * limit
	return limit, offset
}
