APP = urls-random-generator
GOOS = linux
GOARCH = amd64

all: clean build

clean:
	go clean

build: 
	export GOOS=${GOOS}
	export GOARCH=${GOARCH}
	go build 

