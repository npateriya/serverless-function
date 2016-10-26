package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
	"github.com/npateriya/serverless-agent/models"
	"github.com/npateriya/serverless-agent/server/connectors"
	"github.com/npateriya/serverless-agent/server/datastore"
	"github.com/npateriya/serverless-agent/server/datastore/sqlitedatastore"
)

type Function []models.Function

type FunctionController struct{}

func New() *FunctionController {
	bpd := &FunctionController{}
	return bpd
}

func RunFunction(w http.ResponseWriter, r *http.Request) {

	client, err := docker.NewClientFromEnv()
	if err != nil {
		http.Error(w, JsonErr(err), http.StatusInternalServerError)
		return
	}

	// This should come from context
	var fds datastore.FunctionDataStore
	fds = sqlitedatastore.NewsqliteFunctionStore()

	vars := mux.Vars(r)
	fname := vars["funcname"]
	log.Println(fname)

	funcdata := fds.GetFunctionByName(fname)
	if funcdata == nil {
		http.NotFound(w, r)
	}

	//Lets check if function data passed asbody and it have extra param that
	// need to be passed
	decoder := json.NewDecoder(r.Body)
	var function models.Function
	err = decoder.Decode(&function)
	if err == nil && len(function.RunParams) > 0 {
		funcdata.RunParams = function.RunParams
	}
	resp := connectors.RunContainer(funcdata, client)
	ServeJsonResponse(w, resp)

}

func RunAgentFunction(w http.ResponseWriter, r *http.Request) {

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

func SaveFunction(w http.ResponseWriter, r *http.Request) {

	// This should come from context
	var fds datastore.FunctionDataStore
	fds = sqlitedatastore.NewsqliteFunctionStore()

	decoder := json.NewDecoder(r.Body)
	var function models.Function
	err := decoder.Decode(&function)
	if err != nil {
		http.Error(w, JsonErr(err), http.StatusInternalServerError)
		return
	}
	fds.SaveFunction(&function)
	ServeJsonResponse(w, &function)
}

func GetFunction(w http.ResponseWriter, r *http.Request) {

	log.Println("GetFunction: start")

	// This should come from context
	var fds datastore.FunctionDataStore
	fds = sqlitedatastore.NewsqliteFunctionStore()

	vars := mux.Vars(r)
	fname := vars["funcname"]
	log.Println(fname)

	funcdata := fds.GetFunctionByName(fname)
	if funcdata == nil {
		http.NotFound(w, r)
	} else {
		ServeJsonResponse(w, &funcdata)
	}
	log.Println("GetFunction: end")

}
