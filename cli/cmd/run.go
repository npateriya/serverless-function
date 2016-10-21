// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/npateriya/serverless-agent/models"
	"github.com/npateriya/serverless-agent/utils/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	server       string
	file         string
	url          string
	funcparam    []string
	buildimage   string
	argsbuild    string
	targetdir    string
	namespace    string
	functiontype string
)

func makeRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "run serverless function",
		Long: `
User can run serverless function, function can be dowloaded from a public url or uploaded from local filesytem.
It supports function written in most of language, including node, python, go,php, java, scala, perl, c, c++, bash etc. 

Example:
./cli  run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.php?token=AFQsZXU2KxxgReBY5MOoGyimCEn8H58Rks5YEkaTwA%3D%3D
./cli  run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.py?token=AFQsZRl3aBnfjhRfw3lmxBB-bas0LtQyks5YEkaswA%3D%3D
./cli  run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.go?token=AFQsZfRwyoQqlcMcKZhwjlNvTqR62MRSks5YEjPewA%3D%3D
./cli  run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.js?token=AFQsZdzuufjXWtMuZZPpDrZ7Ae8Xn8jUks5YEkZtwA%3D%3D
./cli  run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.c?token=AFQsZaBEQJLO0ivNjWQx7uMUdb-afH33ks5YEkbMwA%3D%3D

./cli run --file testsource/helloworld.go
./cli run --file testsource/helloworld.py
./cli run --file testsource/helloworld.js
./cli run --file testsource/helloworld.c

./cli run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/toupper.go?token=AFQsZfMzSVPF4cvpV8doC05x9vydaIOsks5YEkbzwA%3D%3D --funcparam lower

Example Cisco Functions:
./cli run  -f testsource/spark.py  -p npateriya@gmail.com
./cli run  -f testsource/tropo.py -p 1669777xxxx

Example, running custom python based function:
>>tee helloworld.py 
print("hello world")

>>./cli run -f hello.py
Response from Function execuition:
StdOut   : hello world
StdErr   :  
ExitCode : 0 `,
		Run: func(cmd *cobra.Command, args []string) {
			path := "/function"
			funreq := models.Function{}
			funreq.CacheDir = ".cache" // TODO remove this
			funresp := models.FunctionResponse{}
			restClient := rest.New(rest.Config{Server: server})

			if len(funcparam) > 0 {
				fmt.Println(funcparam)
				funreq.RunParams = funcparam
			}
			if len(url) > 0 {
				funreq.SourceURL = url
				funreq.Type = models.FUNCTION_TYPE_URL
				err := restClient.Post(path, nil, &funreq, &funresp)
				if err != nil {
					fmt.Errorf("%s\n", err)
				}
				printResponse(&funresp)
			} else if len(file) > 0 {
				funreq.Type = models.FUNCTION_TYPE_BLOB
				srcBlob, err := ioutil.ReadFile(file)
				if err != nil || len(srcBlob) == 0 {
					fmt.Errorf("Unable to read file % err %s\n", file, err)
				}
				funreq.SourceBlob = srcBlob
				funreq.SourceFile = file
				err = restClient.Post(path, nil, &funreq, &funresp)
				if err != nil {
					fmt.Errorf("%s\n", err)
				}
				printResponse(&funresp)
			} else {
				fmt.Println("Either --url or --file is required.")
			}

		},
	}

	server = viper.GetString("SERVER")
	file = viper.GetString("FILE")
	url = viper.GetString("URL")
	//funcparam = viper.GetString("FUNCPARAM")

	runCmd.Flags().StringVarP(&server, "server", "s", "http://localhost:8888", "Agent API server endpoint")
	runCmd.Flags().StringVarP(&file, "file", "f", "", "Function local file: E.g. .test/helloworld.php")
	runCmd.Flags().StringVarP(&url, "url", "u", "", "URL to dowload function file E.g https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.c?token=AFQsZaBEQJLO0ivNjWQx7uMUdb-afH33ks5YEkbMwA%3D%3D")
	runCmd.Flags().StringSliceVarP(&funcparam, "funcparam", "p", []string{}, "Run time function parameters")

	viper.BindPFlag("server", runCmd.Flags().Lookup("server"))
	viper.BindPFlag("file", runCmd.Flags().Lookup("file"))
	viper.BindPFlag("url", runCmd.Flags().Lookup("url"))
	//viper.BindPFlag("funcparam", runCmd.Flags().Lookup("funcparam"))

	return runCmd
}

func printResponse(funresp *models.FunctionResponse) {
	if funresp != nil {
		fmt.Printf("\nResponse from Function execuition:\n")
		fmt.Printf("StdOut   : %s", funresp.StdOut)
		fmt.Printf("StdErr   : %s \n", funresp.StdErr)
		fmt.Printf("ExitCode : %d \n", funresp.ExitCode)
		if funresp.Error != nil {
			fmt.Printf("Error    :%s \n", funresp.Error)
		}
	}
}

func init() {
}
