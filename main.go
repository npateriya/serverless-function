package main

import (
	"os"
        "fmt"
        "github.com/npateriya/serverless-agent/connectors"
        "github.com/npateriya/serverless-agent/utils"

)


func main() {
     filename, err := utils.DownloadFile(".test", "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/go/helloworld.go")
     fmt.Println(filename, err)
     _ = connectors.RunContainer(os.Args)
}
