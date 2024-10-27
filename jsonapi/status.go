package jsonapi

import "net/http"

type Response interface {
	Status() int
}

type OKResponse struct{}

func (OKResponse) Status() int {
	return http.StatusOK
}

type AcceptedResponse struct{}

func (AcceptedResponse) Status() int {
	return http.StatusAccepted
}
