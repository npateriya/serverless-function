package datastore

import (
	"github.com/npateriya/serverless-function/models"
)

type FunctionDataStore interface {
	GetFunctionByName(name string) *models.Function
	ListFunction(namespace string) *[]models.Function
	//	GetFunctionMap() map[string]models.Function
	SaveFunction(funcdata *models.Function)
	UpdateFunction(funcdata *models.Function)
	//	DeleteFunctionByName(funcname string)
}
