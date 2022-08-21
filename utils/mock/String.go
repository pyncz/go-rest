package mock

import (
	"math/rand"
	"time"
)

var symbols = []int32(
	"1234567890 abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
)

func MockString(length uint) string {
	rand.Seed(time.Now().UnixNano())

	res := make([]int32, length)
	for i := range res {
		res[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(res)
}
