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
User can run server les function, function can be dowloaded from a public url 
or uploaded from local file. Example:
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

serverless cli enable running multiple language function like go, php, python 
node etc.

Usage:
  cli run [flags]

Flags:
  -f, --file string        Function local file: E.g. .test/helloworld.php
  -p, --funcparam string   Run time function parameters
  -s, --server string      Agent API server endpoint (default "http://localhost:8888")
  -u, --url string         URL to dowload function file E.g https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/php/helloworld.php
```

- Run source function from remote url
```
./cli  run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/helloworld.js?token=AFQsZdzuufjXWtMuZZPpDrZ7Ae8Xn8jUks5YEkZtwA%3D%3D

Response from Function execuition:
StdOut   : hello world
StdErr   :  
ExitCode : 0 
```

- Run source function with parameters
```
./cli run -u https://raw.githubusercontent.com/npateriya/serverless-agent/master/.test/toupper.go?token=AFQsZfMzSVPF4cvpV8doC05x9vydaIOsks5YEkbzwA%3D%3D --funcparam lower

Response from Function execuition:
StdOut   : LOWER
StdErr   :  
ExitCode : 0 
```

- Run source from local directory
```
./cli run --file testsource/helloworld.go

Response from Function execuition:
StdOut   : hello world cli
StdErr   :  
ExitCode : 0 
```
