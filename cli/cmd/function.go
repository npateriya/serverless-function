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

	"github.com/npateriya/serverless-function/models"
	"github.com/npateriya/serverless-function/utils/rest"
	"github.com/ryanuber/columnize"
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
	funcname     string
)

func makeFuncCommand() *cobra.Command {

	funcCmd := &cobra.Command{
		Use:     "function [subcommand]",
		Short:   "Manage and Run Functions",
		Aliases: []string{"func", "f"},
		Long: `
		
User can save, list, get and run serverless function, function can implemented as code dowloaded from a public url or uploaded from local filesytem.
It supports function written in most of language, including node, python, go,php, java, scala, perl, c, c++, bash etc.`,
	}
	return funcCmd
}
func makeFuncSaveCommand() *cobra.Command {
	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "Save serverless function",
		Long: `
User can run serverless function, function can be dowloaded from a public url or uploaded from local filesytem.
It supports function written in most of language, including node, python, go,php, java, scala, perl, c, c++, bash etc. 

Example:

./cli function save -n hello-go-local --file testsource/helloworld.go
./cli function save -n hello-py-local --file testsource/helloworld.py
./cli function save -n hello-js-local --file testsource/helloworld.js
./cli function save -n hello-c-local  --file testsource/helloworld.c

./cli func save -n spark -f  testsource/spark.py
./cli func save -n toupper -f  testsource/toupper.go
./cli func save -n tropo -f testsource/tropo.py

./cli  function save -n hello-php-url -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.php?token=AFQsZXU2KxxgReBY5MOoGyimCEn8H58Rks5YEkaTwA%3D%3D
./cli  function save -n hello-py-url  -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.py?token=AFQsZRl3aBnfjhRfw3lmxBB-bas0LtQyks5YEkaswA%3D%3D
./cli  function save -n hello-go-url  -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.go?token=AFQsZfRwyoQqlcMcKZhwjlNvTqR62MRSks5YEjPewA%3D%3D
./cli  function save -n hello-js-url  -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.js?token=AFQsZdzuufjXWtMuZZPpDrZ7Ae8Xn8jUks5YEkZtwA%3D%3D
./cli  function save -n hello-c-url   -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.c?token=AFQsZaBEQJLO0ivNjWQx7uMUdb-afH33ks5YEkbMwA%3D%3D

 `,
		Run: func(cmd *cobra.Command, args []string) {
			path := "/function"
			funreq := models.Function{}
			funreq.CacheDir = ".cache" // TODO remove this
			funreq.Namespace = namespace
			funreq.Name = funcname
			funresp := models.Function{}
			restClient := rest.New(rest.Config{Server: server})

			if len(funreq.Name) == 0 {
				fmt.Errorf("function name is required field")

			}
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
				printFunction(&funresp)
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
				printFunction(&funresp)
			} else {
				fmt.Println("Either --url or --file is required.")
			}

		},
	}

	server = viper.GetString("SERVER")

	saveCmd.Flags().StringVarP(&server, "server", "s", "http://localhost:8888", "Agent API server endpoint")
	saveCmd.Flags().StringVarP(&funcname, "funcname", "n", "", "Name of function")
	saveCmd.Flags().StringVarP(&namespace, "namespace", "x", "default", "Namespace( to add function")
	saveCmd.Flags().StringVarP(&file, "file", "f", "", "Function local file: E.g. .test/helloworld.php")
	saveCmd.Flags().StringVarP(&url, "url", "u", "", "URL to dowload function file E.g https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.c?token=AFQsZaBEQJLO0ivNjWQx7uMUdb-afH33ks5YEkbMwA%3D%3D")
	saveCmd.Flags().StringSliceVarP(&funcparam, "funcparam", "p", []string{}, "Run time function parameters")

	viper.BindPFlag("server", saveCmd.Flags().Lookup("server"))
	viper.BindPFlag("file", saveCmd.Flags().Lookup("file"))
	viper.BindPFlag("url", saveCmd.Flags().Lookup("url"))
	viper.BindPFlag("funcname", saveCmd.Flags().Lookup("funcname"))
	viper.BindPFlag("namespce", saveCmd.Flags().Lookup("namespace"))

	//viper.BindPFlag("funcparam", saveCmd.Flags().Lookup("funcparam"))

	return saveCmd
}

func makeFuncRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run serverless function",
		Long: `
User can run serverless function, functions executed need to be added first using 'function save' command.
It supports function written in most of language, including node, python, go,php, java, scala, perl, c, c++, bash etc. 

Example:
./cli function run  -n toupper-param-url -p hello

Example Cisco Functions:
./cli function run  -n spark-local -p npateriya@gmail.com
./cli function run  -n tropo-local -p 1669777xxxx



Example, running custom python based function:
>>tee helloworld.py 
print("hello world")

>>./cli function save -n myfirst  -f hello.py
>>./cli function run -n myfirst

Response from Function execuition:
StdOut   : hello world
StdErr   :  
ExitCode : 0 `,
		Run: func(cmd *cobra.Command, args []string) {
			path := "/function"
			funreq := models.Function{}
			funreq.CacheDir = ".cache" // TODO remove this
			funreq.Namespace = namespace
			funreq.Name = funcname
			funresp := models.FunctionResponse{}
			restClient := rest.New(rest.Config{Server: server})

			if len(funreq.Name) == 0 {
				fmt.Errorf("function name is required field")

			}
			path = fmt.Sprintf("%s/%s/run", path, funcname)
			if len(funcparam) > 0 {
				fmt.Println(funcparam)
				funreq.RunParams = funcparam
			}
			err := restClient.Post(path, nil, &funreq, &funresp)
			if err != nil {
				fmt.Errorf("%s\n", err)
			}
			printResponse(&funresp)
		},
	}

	server = viper.GetString("SERVER")
	file = viper.GetString("FILE")
	url = viper.GetString("URL")
	server = viper.GetString("SERVER")

	runCmd.Flags().StringVarP(&server, "server", "s", "http://localhost:8888", "Agent API server endpoint")
	runCmd.Flags().StringVarP(&funcname, "funcname", "n", "", "Name of function")
	runCmd.Flags().StringVarP(&namespace, "namespace", "x", "default", "Namespace( to add function")
	runCmd.Flags().StringSliceVarP(&funcparam, "funcparam", "p", []string{}, "Run time function parameters")

	viper.BindPFlag("server", runCmd.Flags().Lookup("server"))
	viper.BindPFlag("funcname", runCmd.Flags().Lookup("funcname"))
	viper.BindPFlag("namespce", runCmd.Flags().Lookup("namespace"))
	viper.BindPFlag("funcparam", runCmd.Flags().Lookup("funcparam"))
	viper.BindEnv("server", "SERVER")
	return runCmd
}

func makeFuncGetCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:     "get",
		Short:   "Get serverless function",
		Aliases: []string{"list"},
		Long: `
User can get serverless function, functions need to be added first using 'function save' command.

Example:
./cli function get  -n toupper-param-url
./cli function get  -n toupper-param-url --namespace default
`,
		Run: func(cmd *cobra.Command, args []string) {

			path := "/function"
			funreq := models.Function{}
			funreq.Namespace = namespace
			funreq.Name = funcname

			restClient := rest.New(rest.Config{Server: server})

			//			if len(funreq.Name) == 0 {
			//				fmt.Errorf("function name is required field")
			//			}
			if len(funcname) != 0 {
				funresp := models.Function{}
				path = fmt.Sprintf("%s/%s", path, funcname)
				err := restClient.Get(path, nil, &funreq, &funresp)
				if err != nil {
					fmt.Errorf("%s\n", err)
				}
				printFunction(&funresp)
			} else {
				funresplist := []models.Function{}
				err := restClient.Get(path, nil, &funreq, &funresplist)
				if err != nil {
					fmt.Errorf("%s\n", err)
				}
				printFunctionList(funresplist)
			}
		},
	}

	runCmd.Flags().StringVarP(&server, "server", "s", "http://localhost:8888", "Agent API server endpoint")
	runCmd.Flags().StringVarP(&funcname, "funcname", "n", "", "Name of function")
	runCmd.Flags().StringVarP(&namespace, "namespace", "x", "default", "Namespace( to add function")
	server = viper.GetString("SERVER")

	viper.BindPFlag("server", runCmd.Flags().Lookup("server"))
	viper.BindPFlag("funcname", runCmd.Flags().Lookup("funcname"))
	viper.BindPFlag("namespce", runCmd.Flags().Lookup("namespace"))

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

func printFunction(funcdata *models.Function) {
	funcarray := []string{"Name | NameSpace | Type | URL"}
	var funcstr string
	funcstr = fmt.Sprintf("%s| %s | %s | %s", funcdata.Name, funcdata.Namespace, funcdata.Type, funcdata.SourceURL)
	funcarray = append(funcarray, funcstr)
	result := columnize.SimpleFormat(funcarray)
	fmt.Println(result)
}

func printFunctionList(flist []models.Function) {
	funcarray := []string{"Name | NameSpace | Type | URL"}
	var funcstr string
	for _, funcdata := range flist {
		funcstr = fmt.Sprintf("%s| %s | %s | %s", funcdata.Name, funcdata.Namespace, funcdata.Type, funcdata.SourceURL)
		funcarray = append(funcarray, funcstr)
	}
	result := columnize.SimpleFormat(funcarray)
	fmt.Println(result)
}
func init() {
}
