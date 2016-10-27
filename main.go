package main

import (
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/npateriya/serverless-function/server"
)

func main() {

	_, err := docker.NewClientFromEnv()
	if err != nil {
		log.Println("Ensure docker env variables are set so that 'docker ps' works")
		log.Fatal(err)
	}

	srv := server.New()
	os.Exit(srv.Run())

}
