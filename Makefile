TAGS ?= "sqlite"

test:
	./test.sh
	go mod tidy -v

build:
	go build -v .
	go mod tidy -v

update:
	go get -u -tags ${TAGS}
	make test
