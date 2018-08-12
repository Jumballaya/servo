package server

import (
	"fmt"
	"log"
	"net/http"
)

type MiddlewareRoute func(http.ResponseWriter, *http.Request, RouteMethod)

func stringInSlice(a string, b []string) bool {
	for _, str := range b {
		if a == str {
			return true
		}
	}
	return false
}

func checkMethodMiddleware(method string) RouteMiddleware {
	fn := func(w http.ResponseWriter, r *http.Request, rm RouteMethod) {
		if r.Method != method {
			msg := fmt.Sprintf("Path %s has no method %s", r.URL.Path, r.Method)
			fmt.Fprintf(w, msg)
		} else {
			rm.ServeHTTP(w, r)
		}
	}
	return NewMiddleware(fn)
}

func checkNotFoundMiddleware(routeList []string, handler RouteMethod) RouteMiddleware {
	fn := func(w http.ResponseWriter, r *http.Request, rm RouteMethod) {
		if stringInSlice(r.URL.Path, routeList) {
			rm.ServeHTTP(w, r)
		} else {
			handler.ServeHTTP(w, r)
		}
	}
	return NewMiddleware(fn)
}

func NewMiddleware(route MiddlewareRoute) RouteMiddleware {
	return func(rm RouteMethod) RouteMethod {
		fn := func(w http.ResponseWriter, r *http.Request) {
			route(w, r, rm)
		}
		return http.HandlerFunc(fn)
	}
}

func LoggerMiddleware(logger *log.Logger) RouteMiddleware {
	fn := func(w http.ResponseWriter, r *http.Request, rm RouteMethod) {
		defer func() {
			log.Printf("%s - %s\t%s", r.Method, r.URL.Path, r.Header["User-Agent"])
		}()
		rm.ServeHTTP(w, r)
	}
	return NewMiddleware(fn)
}
