```
# Use an image containing environment for building Go programs.
FROM golang:1.17 AS build

# Set the current working directory inside container.
WORKDIR /src

# Download all dependencies.
COPY go.* ./
RUN go mod download

# Build the source code.
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o myprog

# Create a single layer image.
FROM alpine:latest
COPY --from=build /src/myprog /bin/myprog

# Tell Docker to execute this command on a "docker run".
ENTRYPOINT ["/bin/myprog"]
```

```
docker build . -t myprog
docker run myprog
```

See also kvstore's [Dockerfile](https://github.com/jreisinger/kvstore/blob/master/Dockerfile).
