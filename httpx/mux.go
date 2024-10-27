package httpx

import (
	"net/http"
	"path"
)

type RouteFunc func(prefix string, mux *http.ServeMux)

func Route(mux *http.ServeMux, routes ...RouteFunc) {
	for _, route := range routes {
		route("/", mux)
	}
}

func Group(p string, routes ...RouteFunc) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		for _, route := range routes {
			route(path.Join(prefix, p), mux)
		}
	}
}

func GET(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("GET "+path.Join(prefix, pattern), handler)
	}
}

func HEAD(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("HEAD "+path.Join(prefix, pattern), handler)
	}
}

func OPTIONS(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("OPTIONS "+path.Join(prefix, pattern), handler)
	}
}

func TRACE(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("TRACE "+path.Join(prefix, pattern), handler)
	}
}

func PUT(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("PUT "+path.Join(prefix, pattern), handler)
	}
}

func DELETE(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("DELETE "+path.Join(prefix, pattern), handler)
	}
}

func POST(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("POST "+path.Join(prefix, pattern), handler)
	}
}

func PATCH(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("PATCH "+path.Join(prefix, pattern), handler)
	}
}

func CONNECT(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle("CONNECT "+path.Join(prefix, pattern), handler)
	}
}

func Any(pattern string, handler http.Handler) RouteFunc {
	return func(prefix string, mux *http.ServeMux) {
		mux.Handle(path.Join(prefix, pattern), handler)
	}
}
