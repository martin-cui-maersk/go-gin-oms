package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQueryInt(c *gin.Context, key string) (int, bool, error) {
	str, exists := c.GetQuery(key)
	if !exists {
		return 0, false, nil
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, true, fmt.Errorf("%s must be an integer", key)
	}
	return val, true, nil
}
