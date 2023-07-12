Reviewed: 2023-03-21

* defines (repeatable) steps to build your program
* used since 1976 to build programs on Unix

```
install: test
	go install

test:
	go test ./...
```

```
.DEFAULT_GOAL := build

test:
	if [ ! -z "${DB_USER}" ]; then\
		go clean -testcache ./...;\
		go test ./... -cover -tags manual;\
	else\
		go test ./... -cover;\
	fi
.PHONY:test

build: test
	CGO_ENABLED=0 go build -o foo

run: build
	./foo

install: test
	go install
```

* target - each possible operation (like `build`)
* [goals](https://www.gnu.org/software/make/manual/html_node/Goals.html) - targets that make should strive ultimately to update
* `PHONY` - keeps make from getting confused if there's a directory with the same name as a target
