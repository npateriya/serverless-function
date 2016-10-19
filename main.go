package main

import (
	_ "os"

	"github.com/npateriya/serverless-agent/connectors"
	"github.com/npateriya/serverless-agent/models"
)

func main() {
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
	//	connectors.RunDexecContainer(functionArgs)

	functionArgsURL := &models.Function{
		Type: models.FUNCTION_TYPE_URL,
		//SourceURL: "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/go/helloworld.go",
		//SourceURL: "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/php/helloworld.php",
		SourceURL: "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/go/echochamber.go",
		//SourceFile: ".test/helloworld2.go",
		//SourceBlob: "",
		//SourceLang;"",
		//BaseImage:"",
		//BuildArgs:"",
		RunParams: []string{"hello world", "echo me"},
		//Version:"",
		CacheDir: ".test",
	}

	connectors.RunContainer(functionArgsURL)
}
