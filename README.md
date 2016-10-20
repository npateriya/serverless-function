# serverless-agent
Serverless agent that will listen to queue and execute function.

## Local execution
- Pull source code 

	go get github.com/npateriya/serverless-agent
- build 

	make build
- start serverless agent
	./serverless-agent

### Adhoc testing

- make curl-test -e SRC=python -e EXT=py
- make curl-test -e SRC=go
- make curl-test -e SRC=php
- make curl-test -e SRC=node -e EXT=js

