package datastore

import (
	"github.com/npateriya/serverless-agent/models"
)

type FunctionDataStore interface {
	GetFunctionByName(name string) *models.Function
	//	GetFunctionList() []models.Function
	//	GetFunctionMap() map[string]models.Function
	SaveFunction(funcdata *models.Function)
	//	UpdateFunction(funcdata *models.Function)
	//	DeleteFunctionByName(funcname string)
}
