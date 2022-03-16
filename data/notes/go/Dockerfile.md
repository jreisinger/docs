```Dockerfile
# Use official Go image containing environment for building Go programs.
FROM golang:1.17-alpine AS build

# Create a directory inside the image. Docker will use this directory as the
# default destination for all subsequent commands. You can use relative paths.
WORKDIR /app

# Download all modules necessary to compile the app, i.e. download dependencies.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Compile (build the source code of) the app.
COPY *.go ./
RUN go build -o app
#RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Create a single layer image.
FROM alpine:latest
COPY --from=build /app/app /bin/app

# Doesn't actually publish the port but functions as documentation. To actually
# publish the port use `-p` on `docker run` or `-P` to publish all exposed ports
# and map them to high-order ports.
EXPOSE 8080

# Tell Docker to execute this command on a "docker run".
ENTRYPOINT ["/bin/app"]
```

```
docker build . -t app
docker run app
```

See also:

* https://docs.docker.com/language/golang/build-images/
* https://github.com/jreisinger/kvstore/blob/master/Dockerfile
* https://github.com/jreisinger/restful-api-example/blob/main/Dockerfile
