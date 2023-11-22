package client

import (
	"io"
	"net/http"
)

func ExecuteNodeEc() ([]byte, error) {
	response, err := http.Get("user-service:4500/ping")

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
