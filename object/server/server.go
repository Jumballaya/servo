package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type App struct {
	Mux              *http.ServeMux
	Routers          []*Router
	GlobalMiddleware []RouteMiddleware
	genericRoutes    map[string]RouteMethod
	routeList        []string
	staticFolder     *staticFolder
}

type staticFolder struct {
	url  string
	path string
}

func NewApp() *App {
	return &App{
		Routers:          []*Router{},
		Mux:              http.NewServeMux(),
		GlobalMiddleware: []RouteMiddleware{},
		genericRoutes:    make(map[string]RouteMethod),
		routeList:        []string{},
	}
}

func (a *App) Use(root string, router *Router) {
	router.Root = root
	a.Routers = append(a.Routers, router)

	for _, route := range router.Routes {
		path := formatPath(root, route.Path)
		a.routeList = append(a.routeList, path)
	}
}

func (a *App) UseMiddleware(middleware ...RouteMiddleware) {
	a.GlobalMiddleware = append(a.GlobalMiddleware, middleware...)
}

func (a *App) NotFound(method RouteMethod) {
	a.genericRoutes["404"] = method
}

func (a *App) StaticFiles(folderPath, url string) {
	// Check if folder exists first
	a.staticFolder = &staticFolder{
		url:  url,
		path: folderPath,
	}
}

func (a *App) Run(port string) {
	// Static
	if a.staticFolder != nil {
		fs := http.StripPrefix(a.staticFolder.url+"/", http.FileServer(http.Dir(a.staticFolder.path)))
		a.Mux.Handle(a.staticFolder.url+"/", fs)
	}

	// 404 middleware
	notFoundHandler, ok := a.genericRoutes["404"]
	if !ok {
		notFoundHandler = http.HandlerFunc(http.NotFound)
	}
	notFound := checkNotFoundMiddleware(a.routeList, notFoundHandler)
	a.UseMiddleware(notFound)

	for _, router := range a.Routers {
		router.buildRoutes(a.Mux, a.GlobalMiddleware)
	}

	fmt.Println(fmt.Sprintf("Server listening on port %s", port))
	log.Fatal(http.ListenAndServe(port, a.Mux))
}

type Router struct {
	Root   string
	Routes []*Route
}

func NewRouter() *Router {
	return &Router{Routes: []*Route{}}
}

func formatPath(root, path string) string {
	if path == "/" && root == "/" {
		return path
	}
	if path == "/" {
		return root
	}
	if strings.HasPrefix(path, root) {
		index := strings.Index(path, root)
		return path[index:]
	}
	return fmt.Sprintf("%s%s", root, path)
}

func (r *Router) newRoute(method, path string, fn RouteMethod, middleware ...RouteMiddleware) {
	var mid RouteMiddleware
	if len(middleware) > 0 {
		if len(middleware) > 1 {
			mid = chainMiddleware(middleware[0], middleware[1:]...)
		} else {
			mid = chainMiddleware(middleware[0])
		}
	}

	route := &Route{
		Method:     method,
		Path:       path,
		Function:   fn,
		Middleware: mid,
	}

	r.Routes = append(r.Routes, route)
}

func (r *Router) buildRoutes(mux *http.ServeMux, globalMw []RouteMiddleware) {
	for _, route := range r.Routes {
		mw := []RouteMiddleware{}
		mw = append(mw, globalMw...)
		if route.Middleware != nil {
			mw = append(mw, route.Middleware)
		}

		routeHandler := handleRoute(route.Function, mw...)
		urlPath := formatPath(r.Root, route.Path)
		mux.Handle(urlPath, routeHandler)
	}
}

func (r *Router) Get(path string, fn RouteMethod, middleware ...RouteMiddleware) {
	middleware = append(middleware, checkMethodMiddleware(http.MethodGet))
	r.newRoute(http.MethodGet, path, fn, middleware...)
}

func (r *Router) Post(path string, fn RouteMethod, middleware ...RouteMiddleware) {
	middleware = append(middleware, checkMethodMiddleware(http.MethodPost))
	r.newRoute(http.MethodPost, path, fn, middleware...)
}

func (r *Router) Put(path string, fn RouteMethod, middleware ...RouteMiddleware) {
	middleware = append(middleware, checkMethodMiddleware(http.MethodPut))
	r.newRoute(http.MethodPut, path, fn, middleware...)
}

func (r *Router) Patch(path string, fn RouteMethod, middleware ...RouteMiddleware) {
	middleware = append(middleware, checkMethodMiddleware(http.MethodPatch))
	r.newRoute(http.MethodPatch, path, fn, middleware...)
}

func (r *Router) Delete(path string, fn RouteMethod, middleware ...RouteMiddleware) {
	middleware = append(middleware, checkMethodMiddleware(http.MethodDelete))
	r.newRoute(http.MethodDelete, path, fn, middleware...)
}

type Route struct {
	Method     string
	Path       string
	Function   RouteMethod
	Middleware RouteMiddleware
}

type RouteMiddleware func(RouteMethod) RouteMethod

type RouteMethod interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func handleRoute(main RouteMethod, middleware ...RouteMiddleware) RouteMethod {
	if len(middleware) > 0 {
		if len(middleware) > 1 {
			return chainMiddleware(middleware[0], middleware[1:]...)(main)
		}
		return chainMiddleware(middleware[0])(main)
	}
	return main
}

func chainMiddleware(outer RouteMiddleware, middleware ...RouteMiddleware) RouteMiddleware {
	return func(rm RouteMethod) RouteMethod {
		topIndex := len(middleware) - 1
		for i := range middleware {
			rm = middleware[topIndex-i](rm)
		}
		return outer(rm)
	}
}
