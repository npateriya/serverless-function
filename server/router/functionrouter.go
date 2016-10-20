package router

import (
	"github.com/gorilla/mux"
	"github.com/npateriya/serverless-agent/server/controllers"
)

const (
	FunctionRun = "/function"
)

func Function(muxRouter *mux.Router) {
	muxRouter.HandleFunc(FunctionRun, controllers.RunContainer).Methods("POST")
}
