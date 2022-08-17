package utils

import (
	"strconv"
)

func ExtractInt64Query(value string, defaults int64) (int64, error) {
	if value == "" {
		return defaults, nil
	} else {
		return strconv.ParseInt(value, 10, 64)
	}
}
