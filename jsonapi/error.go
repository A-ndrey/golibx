package jsonapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var defaultError = Error{
	Status: http.StatusInternalServerError,
	Detail: http.StatusText(http.StatusInternalServerError),
}

type Error struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.Status, http.StatusText(e.Status), e.Detail)
}

type errorResponse struct {
	Error `json:"error"`
}

func (e errorResponse) write(w http.ResponseWriter) {
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(e)
}
