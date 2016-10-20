package router

import (
	"net/http"

	_ "github.com/gorilla/mux"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

//func ContextWrap(handler ActionHandler) HttpHandler {
//	wrappedHandler := func(w http.ResponseWriter, r *http.Request) {
//		handler(util.Context(r), w, r)
//	}
//	return wrappedHandler
//}
