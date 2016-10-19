package main

import (
	"io/ioutil"
	"log"
	_ "os"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/npateriya/serverless-agent/connectors"
	"github.com/npateriya/serverless-agent/models"
)

func main() {

	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	//	functionArgs := &models.Function{
	//		Type: models.FUNCTION_TYPE_FILE,
	//		//SourceURL:  ".test/helloworld2.go",
	//		//SourceFile: ".test/helloworld2.go",
	//		SourceFile: ".test/helloworld.go",
	//		//SourceBlob: "",
	//		//SourceLang;"",
	//		//BaseImage:"",
	//		//BuildArgs:"",
	//		RunParams: []string{"hello world"},
	//		//Version:"",
	//	}
	//	connectors.RunDexecContainer(functionArgs,client)

	//	functionArgsURL := &models.Function{
	//		Type: models.FUNCTION_TYPE_URL,
	//		//SourceURL: "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/go/helloworld.go",
	//		//SourceURL: "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/php/helloworld.php",
	//		SourceURL: "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/go/echochamber.go",
	//		//SourceFile: ".test/helloworld2.go",
	//		//SourceBlob: "",
	//		//SourceLang;"",
	//		//BaseImage:"",
	//		//BuildArgs:"",
	//		RunParams: []string{"hello world", "echo me"},
	//		//Version:"",
	//		CacheDir: ".test",
	//	}

	//	connectors.RunContainer(functionArgsURL,client)

	srcBlob, err := ioutil.ReadFile(".test/helloworld2.go")
	if err != nil {
		log.Fatal(err)
	}
	functionArgsBlob := &models.Function{
		Type:       models.FUNCTION_TYPE_BLOB,
		SourceFile: "helloworldblob.go",
		SourceBlob: srcBlob,
		//SourceLang;"",
		//BaseImage:"",
		//BuildArgs:"",
		RunParams: []string{"hello world", "echo me"},
		//Version:"",
		CacheDir: ".test",
	}

	connectors.RunContainer(functionArgsBlob, client)
}
