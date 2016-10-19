build:
	go build
test:
	go run .test/helloworld.go 
glide-install:
	glide install --strip-vcs --strip-vendor
glide-dedup:
	glide install --strip-vcs --strip-vendor
glide-update:
	glide update
