package common

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func IDFromURL(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		return 0, errors.New("not found")
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, err
	}
	return id, nil
}
