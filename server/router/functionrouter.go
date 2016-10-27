package router

import (
	"github.com/gorilla/mux"
	"github.com/npateriya/serverless-function/server/controllers"
)

const (
	// This end point is running on server
	// Only function name is needed, function is populated by querying function
	// model data from db based on name.
	FunctionRun = "/function/{funcname}/run"
	// This endpoint is running on agent nodes, note in body
	// need to pass complete populated function model
	FunctionAgentRun = "/agent/function/{funcname}/run"
	FunctionSave     = "/function"
	FunctionGet      = "/function/{funcname}"
	FunctionList     = "/function"
)

func Function(muxRouter *mux.Router) {
	muxRouter.HandleFunc(FunctionRun, controllers.RunFunction).Methods("POST")
	muxRouter.HandleFunc(FunctionAgentRun, controllers.RunAgentFunction).Methods("POST")
	muxRouter.HandleFunc(FunctionSave, controllers.SaveFunction).Methods("POST")
	muxRouter.HandleFunc(FunctionGet, controllers.GetFunction).Methods("GET")
	muxRouter.HandleFunc(FunctionList, controllers.ListFunction).Methods("GET")
}
