package jsonapi

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc[T any] func(*http.Request) (T, error)

func WrapFunc[T any](handlerFunc HandlerFunc[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)

		resp, err := handlerFunc(r)
		if err != nil {
			respErr := errorResponse{}

			switch e := err.(type) {
			case Error:
				respErr.Error = e
			default:
				respErr.Error = defaultError
			}

			respErr.write(w)

			return
		}

		bytes, err := json.Marshal(resp)
		if err != nil {
			errorResponse{defaultError}.write(w)
			return
		}

		setStatus(w, resp)
		w.Write(bytes)
	}
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func setStatus(w http.ResponseWriter, resp any) {
	if r, ok := resp.(Response); ok {
		w.WriteHeader(r.Status())
	}
}
