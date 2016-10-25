package router

import (
	"github.com/gorilla/mux"
	"github.com/npateriya/serverless-agent/server/controllers"
)

const (
	FunctionRun  = "/function/{funcname}/run"
	FunctionSave = "/function"
	FunctionGet  = "/function/{funcname}"
)

func Function(muxRouter *mux.Router) {
	muxRouter.HandleFunc(FunctionRun, controllers.RunFunction).Methods("POST")
	muxRouter.HandleFunc(FunctionSave, controllers.SaveFunction).Methods("POST")
	muxRouter.HandleFunc(FunctionGet, controllers.GetFunction).Methods("GET")

}
