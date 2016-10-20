package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonErr(err interface{}) string {
	jsonBytes, jsonErr := json.Marshal(map[string]string{"error": TypeCastToString(err)})
	jsonStr := string(jsonBytes[:])
	if jsonErr != nil {
		panic(jsonErr)
	}
	return jsonStr
}

func ServeJsonResponse(w http.ResponseWriter, responseBodyObj interface{}) {
	ServeJsonResponseWithCode(w, responseBodyObj, http.StatusOK)
}

func ServeJsonResponseWithCode(w http.ResponseWriter, responseBodyObj interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff") // https://www.owasp.org/index.php/List_of_useful_HTTP_headers

	entityContent, err := json.Marshal(responseBodyObj)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(code)
	if _, err := w.Write(entityContent); err != nil {
		panic(err)
	}
}

func TypeCastToString(err interface{}) string {
	if errStr, ok := err.(fmt.Stringer); ok {
		return errStr.String()
	} else if errE, ok := err.(error); ok {
		return errE.Error()
	} else if errStr, ok := err.(string); ok {
		return errStr
	} else {
		return fmt.Sprintf("%s", err)
	}
}
