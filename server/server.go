package server

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/npateriya/serverless-function/server/router"
	"github.com/rs/cors"
)

type Server struct {
}

func New() *Server {
	server := &Server{}
	return server
}

func (ref *Server) BuildServer() *negroni.Negroni {
	muxRouter := mux.NewRouter()
	router.Function(muxRouter)
	router.HouseKeeping(muxRouter)
	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "CONNECT"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
		cor,
	)
	n.UseHandler(muxRouter)
	return n
}

func (ref *Server) Run() int {
	fmt.Println("server: start")
	server := ref.BuildServer()
	server.Run("0.0.0.0:8888")
	fmt.Println("server: end")
	return 0
}
