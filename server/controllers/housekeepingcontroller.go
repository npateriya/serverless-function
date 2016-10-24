package controllers

import (
	"net/http"
	"os"
)

type HouseKeepingController struct{}

func NewHouseKeepingController() *HouseKeepingController {
	hkc := &HouseKeepingController{}
	return hkc
}

func HealthStatus(w http.ResponseWriter, r *http.Request) {

	hname, _ := os.Hostname()
	resp := map[string]string{
		"Status":      "Healthy",
		"Description": "Serverless Function as a Service",
		"Hostname":    hname,
	}
	ServeJsonResponse(w, resp)

}
