package router

import (
	"github.com/gorilla/mux"
	"github.com/npateriya/serverless-function/server/controllers"
)

const (
	HealthStatus = "/health"
)

func HouseKeeping(muxRouter *mux.Router) {
	muxRouter.HandleFunc(HealthStatus, controllers.HealthStatus).Methods("GET")
}
