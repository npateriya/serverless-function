package sqlitedatastore

import (
	//	"database/sql"
	//	"fmt"
	"log"
	//	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/npateriya/serverless-agent/models"
	"github.com/npateriya/serverless-agent/server/datastore"
	"github.com/npateriya/serverless-agent/utils"
)

type sqliteFunctionStore struct {
	SQLDS *sqliteStore
}

func NewsqliteFunctionStore() datastore.FunctionDataStore {
	sqlfs := sqliteFunctionStore{SQLDS: sqliteStoreConnect("")}

	sqlStmt := `
	create table  if not exists function (name type string not null primary key, 
	functype text, sourcefile text, sourceurl text, sourceblob text, sourcelang text,
	baseimage text, buildargs text, runparams text, includedir text, 
	cachedir text,namespace string, version string );
	`
	_, err := sqlfs.SQLDS.CDB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	var fds datastore.FunctionDataStore = &sqlfs

	return fds
}

func (ref *sqliteFunctionStore) SaveFunction(funcdata *models.Function) {

	var insertSqlStmt = `INSERT INTO function (name, functype, sourcefile, 
	sourceurl,sourceblob, sourcelang, 
	baseimage, buildargs, runparams, 
	includedir, cachedir, namespace , 
	version)
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ;`

	argsstr := utils.EncodeSlice(funcdata.BuildArgs)
	paramsstr := utils.EncodeSlice(funcdata.RunParams)
	inludedir := utils.EncodeSlice(funcdata.IncludeDir)
	if len(funcdata.Namespace) == 0 {
		funcdata.Namespace = "default"
	}

	stmt, err := ref.SQLDS.CDB.Prepare(insertSqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, insertSqlStmt)
	}
	defer stmt.Close()
	_, err = stmt.Exec(&funcdata.Name, &funcdata.Type, &funcdata.SourceFile,
		&funcdata.SourceURL, &funcdata.SourceBlob, &funcdata.SourceLang,
		&funcdata.BaseImage, &argsstr, &paramsstr,
		&inludedir, &funcdata.CacheDir, &funcdata.Namespace,
		&funcdata.Version)
	if err != nil {
		log.Printf("%q: %s\n", err, insertSqlStmt)
	}
	return
}

func (ref *sqliteFunctionStore) GetFunctionByName(name string) *models.Function {
	var funcdata models.Function
	var selectSqlStmt = `SELECT name, functype, sourcefile, 
	sourceurl,sourceblob, sourcelang, 
	baseimage, buildargs, runparams, 
	includedir, cachedir, namespace , 
	version 
	FROM function WHERE name=?`
	var argsstr, paramsstr, includedir string

	stmt, err := ref.SQLDS.CDB.Prepare(selectSqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, selectSqlStmt)
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(
		&funcdata.Name, &funcdata.Type, &funcdata.SourceFile,
		&funcdata.SourceURL, &funcdata.SourceBlob, &funcdata.SourceLang,
		&funcdata.BaseImage, &argsstr, &paramsstr,
		&includedir, &funcdata.CacheDir, &funcdata.Namespace,
		&funcdata.Version)
	if err != nil {
		log.Printf("%q: %s\n", err, selectSqlStmt)
		return nil
	}
	funcdata.BuildArgs = utils.DecodeToMap(argsstr)
	funcdata.RunParams = utils.DecodeToMap(paramsstr)
	funcdata.IncludeDir = utils.DecodeToMap(includedir)
	return &funcdata
}

func (ref *sqliteFunctionStore) ListFunction(namespace string) *[]models.Function {
	var funclist []models.Function
	var selectSqlStmt = `SELECT name, functype, sourcefile, 
	sourceurl,sourceblob, sourcelang, 
	baseimage, buildargs, runparams, 
	includedir, cachedir, namespace , 
	version 
	FROM function WHERE namespace=?`
	var argsstr, paramsstr, includedir string
	if len(namespace) == 0 {
		namespace = "default"
	}

	rows, err := ref.SQLDS.CDB.Query(selectSqlStmt, namespace)
	if err != nil {
		log.Printf("%q: %s\n", err, selectSqlStmt)
	}
	defer rows.Close()

	var funcdata models.Function
	for rows.Next() {
		err = rows.Scan(
			&funcdata.Name, &funcdata.Type, &funcdata.SourceFile,
			&funcdata.SourceURL, &funcdata.SourceBlob, &funcdata.SourceLang,
			&funcdata.BaseImage, &argsstr, &paramsstr,
			&includedir, &funcdata.CacheDir, &funcdata.Namespace,
			&funcdata.Version)
		if err != nil {
			log.Printf("%q: %s\n", err, selectSqlStmt)
			return nil
		}
		funcdata.BuildArgs = utils.DecodeToMap(argsstr)
		funcdata.RunParams = utils.DecodeToMap(paramsstr)
		funcdata.IncludeDir = utils.DecodeToMap(includedir)
		funclist = append(funclist, funcdata)
	}
	return &funclist
}

//func (ref *sqliteFunctionStore) GetFunctionList() []models.Function {

//}
//func (ref *sqliteFunctionStore) GetFunctionMap() map[string]models.Function {

//}

//func (ref *sqliteFunctionStore) UpdateFunction(funcdata *models.Function) {

//}

//func (ref *sqliteFunctionStore) DeleteFunctionByName(funcname string) {

//}
