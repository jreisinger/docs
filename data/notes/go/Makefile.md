```
test:
	if [ ! -z "${DB_USER}" ]; then\
		go clean -testcache ./...;\
		GO111MODULE=on go test ./... -tags manual;\
	else\
		GO111MODULE=on go test ./...;\
	fi

build: test
	GO111MODULE=on CGO_ENABLED=0 go build -o foo

run: build
	./foo
```
