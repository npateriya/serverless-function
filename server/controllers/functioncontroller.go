package controllers

import (
	"encoding/json"
	"net/http"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/npateriya/serverless-agent/connectors"
	"github.com/npateriya/serverless-agent/models"
)

type Function []models.Function

type FunctionController struct{}

func New() *FunctionController {
	bpd := &FunctionController{}
	return bpd
}

func RunContainer(w http.ResponseWriter, r *http.Request) {

	client, err := docker.NewClientFromEnv()
	if err != nil {
		http.Error(w, JsonErr(err), http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var function models.Function
	err = decoder.Decode(&function)
	if err != nil {
		http.Error(w, JsonErr(err), http.StatusInternalServerError)
		return
	}
	resp := connectors.RunContainer(&function, client)
	ServeJsonResponse(w, resp)

}
