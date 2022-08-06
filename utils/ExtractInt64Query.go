package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractInt64Query(ctx *gin.Context, key string, defaults int64) (int64, error) {
	query := ctx.Query(key)
	if query == "" {
		return defaults, nil
	} else {
		return strconv.ParseInt(query, 10, 64)
	}
}
