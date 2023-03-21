Reviewed: 2023-03-21

* defines (repeatable) steps to build your program
* used since 1976 to build programs on Unix
* each possible operation (like `build`) is called a *target*
* `PHONY` keeps `make` from getting confused if there's a directory with the same name as a target

```
.DEFAULT_GOAL := build

test:
	if [ ! -z "${DB_USER}" ]; then\
		go clean -testcache ./...;\
		go test ./... -tags manual;\
	else\
		go test ./...;\
	fi
.PHONY:test

build: test
	CGO_ENABLED=0 go build -o foo

run: build
	./foo
```
