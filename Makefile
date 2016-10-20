SRC=php
build:
	go build
glide-install:
	glide install --strip-vcs --strip-vendor
glide-dedup:
	glide install --strip-vcs --strip-vendor
glide-update:
	glide update

curl-test:
	echo "Source code running from below URL"
	curl -s "https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/php/helloworld.php"
	echo;echo "Running above source codei as function"
	curl -s  -X POST localhost:8888/function -d '{"type":"FUNCTION_TYPE_URL", "sourceurl":"https://raw.githubusercontent.com/docker-exec/dexec/master/.test/bats/fixtures/php/helloworld.php", "cachedir":".test"}' | python -m json.tool
