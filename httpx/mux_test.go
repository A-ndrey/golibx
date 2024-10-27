package httpx

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	testMux := http.NewServeMux()
	Route(testMux,
		Group("g11",
			Group("g21",
				GET("get", newTestHandler("g11-g21-get"))),
			Group("g22",
				GET("get", newTestHandler("g11-g22-get"))),
		),
		GET("get", newTestHandler("get")),
		POST("post", newTestHandler("post")),
		PUT("put", newTestHandler("put")),
		PATCH("patch", newTestHandler("patch")),
		DELETE("delete", newTestHandler("delete")),
		HEAD("head", newTestHandler("head")),
		TRACE("trace", newTestHandler("trace")),
		CONNECT("connect", newTestHandler("connect")),
		OPTIONS("options", newTestHandler("options")),
		Any("any", newTestHandler("any")),
		Any("/", newTestHandler("root")),
	)

	testRequest := func(method string, target string, expected string) {
		req := httptest.NewRequest(method, target, nil)
		h, _ := testMux.Handler(req)
		wr := httptest.NewRecorder()
		h.ServeHTTP(wr, req)
		response := wr.Result()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			t.Error(err)
		}
		if string(responseBody) != expected {
			t.Errorf("expecting %s, got %s", expected, string(responseBody))
		}
	}

	testRequest(http.MethodGet, "/g11/g21/get", "g11-g21-get")
	testRequest(http.MethodGet, "/g11/g22/get", "g11-g22-get")
	testRequest(http.MethodGet, "/get", "get")
	testRequest(http.MethodPost, "/post", "post")
	testRequest(http.MethodPut, "/put", "put")
	testRequest(http.MethodPatch, "/patch", "patch")
	testRequest(http.MethodDelete, "/delete", "delete")
	testRequest(http.MethodHead, "/head", "head")
	testRequest(http.MethodConnect, "/connect", "connect")
	testRequest(http.MethodOptions, "/options", "options")
	testRequest(http.MethodTrace, "/trace", "trace")
	testRequest(http.MethodGet, "/any", "any")
	testRequest(http.MethodPost, "/any", "any")
	testRequest(http.MethodGet, "/", "root")
	testRequest(http.MethodGet, "/unknown", "root")
}

func newTestHandler(response string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(response))
	})
}
