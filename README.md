# serverless
Serverless agent that can execute functions written in multiple languages. 
Serverless cli for triggering function execution.

[![asciicast](https://asciinema.org/a/7uozu03v844frxn3xpnvxm53u.png)](https://asciinema.org/a/7uozu03v844frxn3xpnvxm53u)

## Start Agent
- Pull source code 
```
    go get github.com/npateriya/serverless-agent
```
- build 
```
make build
```

- start serverless agent
  ### Note: before starting serverless-agent ensure 'docker ps' is working on shell ###
```
./serverless-agent
```

## Build CLI
- cd cli
- build 
```
make build
```

## Test
```
/cli function 
User can save, list, get and run serverless function, function can implemented as code dowloaded from a public url or uploaded from local filesytem.
It supports function written in most of language, including node, python, go,php, java, scala, perl, c, c++, bash etc.

Usage:
  cli function [command]

Available Commands:
  get         Get serverless function
  run         Run serverless function
  save        Save serverless function

Use "cli function [command] --help" for more information about a command.
```

- Save functions

```
NeeleshP-ServelessDemo>>./cli function save -n hello-go-local --file testsource/helloworld.go
Function details:
Name   : hello-go-local
Type   : FUNCTION_TYPE_BLOB 
URL    :  
Namespace: default

NeeleshP-ServelessDemo>>./cli  function save -n hello-php-url -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.php?token=AFQsZXU2KxxgReBY5MOoGyimCEn8H58Rks5YEkaTwA%3D%3D
Function details:
Name   : hello-php-url
Type   : FUNCTION_TYPE_URL 
URL    : https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.php?token=AFQsZXU2KxxgReBY5MOoGyimCEn8H58Rks5YEkaTwA%3D%3D 
Namespace: default

```


- Run functions

```
NeeleshP-ServelessDemo>>./cli function run -n hello-go-local
Response from Function execuition:
StdOut   : hello world cli
StdErr   :  
ExitCode : 0 

NeeleshP-ServelessDemo>>./cli function run -n hello-php-url
Response from Function execuition:
StdOut   : hello world
StdErr   :  
ExitCode : 0 


```

- Get Function 

```
NeeleshP-ServelessDemo>>./cli function get -n hello-php-url
Name   : hello-php-url
Type   : FUNCTION_TYPE_URL 
URL    : https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.php?token=AFQsZXU2KxxgReBY5MOoGyimCEn8H58Rks5YEkaTwA%3D%3D 
Namespace: default

NeeleshP-ServelessDemo>>./cli function get -n hello-go-local
Name   : hello-go-local
Type   : FUNCTION_TYPE_BLOB 
URL    :  
Namespace: default

```