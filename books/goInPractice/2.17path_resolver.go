package main

import (
	"fmt"
	"net/http"
	"path"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc // e.g. [hello]hello or [goodbye]goodbye
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range p.handlers { // for key, value
		if matched, err := path.Match(pattern, check); matched && err == nil {
			handlerFunc(w, r)
			return
		} else if err != nil {
			_, _ = fmt.Fprint(w, err)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "William Noble"
	}
	fmt.Fprint
}

func main() {

}
