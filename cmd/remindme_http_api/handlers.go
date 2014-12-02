
package main 

import (
  "net/http"
  "regexp"
)

type Route struct {
  Pattern *regexp.Regexp 
  Handler http.Handler
}

type RegexpHandlers struct {
  Handlers []*Route
}

func (rh *RegexpHandlers) Handler(pattern string, handler http.Handler) {
  var err error 
  var pattern_re *regexp.Regexp

  pattern_re, err = regexp.Compile(pattern)
  checkError(err)

  rh.Handlers = append(rh.Handlers, &Route{pattern_re, handler})
}

func (rh *RegexpHandlers) HandleFunc(pattern string, handler http.Handler) {
  var err error 
  var pattern_re *regexp.Regexp

  pattern_re, err = regexp.Compile(pattern)
  checkError(err)

  if rh.Handlers == nil {
    rh.Handlers = make([]*Route, 0)
  }

  rh.Handlers = append(rh.Handlers, &Route{pattern_re, handler})
}

func (rh *RegexpHandlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  for _, route := range rh.Handlers {
    if route.Pattern.MatchString(r.URL.Path) {
      route.Handler.ServeHTTP(w, r)
      return 
    }
  }

  http.NotFound(w, r)
}
