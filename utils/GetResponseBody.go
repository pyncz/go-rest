package utils

import (
	"io"
	"net/http"
)

func GetResponseBody(res *http.Response) string {
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyBytes)
}
