package utils

import (
	"math/rand"
	"time"
)

type RandReader = func(p []byte) (n int, err error)

func GetRandReader() RandReader {
	rand.Seed(time.Now().UnixNano())
	return rand.Read
}
