package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
	"github.com/npateriya/serverless-function/models"
	"github.com/npateriya/serverless-function/server/connectors"
	"github.com/npateriya/serverless-function/server/datastore"
	"github.com/npateriya/serverless-function/server/datastore/sqlitedatastore"
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
	log.Printf("%+v", function)
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

func UpdateFunction(w http.ResponseWriter, r *http.Request) {

	// This should come from context
	var fds datastore.FunctionDataStore
	fds = sqlitedatastore.NewsqliteFunctionStore()

	decoder := json.NewDecoder(r.Body)
	var functionnew models.Function
	err := decoder.Decode(&functionnew)
	if err != nil {
		http.Error(w, JsonErr(err), http.StatusInternalServerError)
		return
	}
	functionorig := fds.GetFunctionByName(functionnew.Name)
	functionorig = mergeFunction(functionorig, &functionnew)
	fds.UpdateFunction(functionorig)
	ServeJsonResponse(w, &functionorig)
}
func mergeFunction(forig *models.Function, fnew *models.Function) *models.Function {
	if len(fnew.SourceBlob) > 0 {
		forig.SourceBlob = fnew.SourceBlob
	}
	if len(fnew.RunParams) > 0 {
		forig.RunParams = fnew.RunParams
	}
	if len(fnew.BuildArgs) > 0 {
		forig.BuildArgs = forig.BuildArgs
	}
	if len(fnew.IncludeDir) > 0 {
		forig.IncludeDir = fnew.IncludeDir
	}
	if len(fnew.SourceFile) > 0 {
		forig.SourceFile = fnew.SourceFile
	}
	if len(fnew.SourceLang) > 0 {
		forig.SourceLang = fnew.SourceLang
	}
	if len(fnew.Version) > 0 {
		forig.Version = fnew.Version
	}
	return forig
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

func ListFunction(w http.ResponseWriter, r *http.Request) {

	log.Println("ListFunction: start")

	// This should come from context
	var fds datastore.FunctionDataStore
	fds = sqlitedatastore.NewsqliteFunctionStore()

	// TODO add name space logic
	funcdata := fds.ListFunction("")
	ServeJsonResponse(w, &funcdata)
	log.Println("ListFunction: end")

}

func DeleteFunction(w http.ResponseWriter, r *http.Request) {

	log.Println("DeleteFunction: start")

	// This should come from context
	var fds datastore.FunctionDataStore
	fds = sqlitedatastore.NewsqliteFunctionStore()

	vars := mux.Vars(r)
	fname := vars["funcname"]
	log.Println(fname)
	fds.DeleteFunctionByName(fname, "default")

	var funcresp models.FunctionResponse
	funcresp.Message = fmt.Sprintf("Function delete default/%s", fname)
	ServeJsonResponse(w, &funcresp)
	log.Println("DeleteFunction: end")

}
func DeleteFunctionNS(w http.ResponseWriter, r *http.Request) {

	log.Println("DeleteFunction: start")

	// This should come from context
	var fds datastore.FunctionDataStore
	fds = sqlitedatastore.NewsqliteFunctionStore()

	vars := mux.Vars(r)
	fname := vars["funcname"]
	ns := vars["namespace"]

	fds.DeleteFunctionByName(fname, ns)
	var funcresp models.FunctionResponse
	funcresp.Message = fmt.Sprintf("Function delete %s/%s", ns, fname)
	ServeJsonResponse(w, &funcresp)
	log.Println("DeleteFunction: end")

}
