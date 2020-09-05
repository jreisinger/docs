```
FROM golang:1.14 AS build

# Set the current working directory inside container.
WORKDIR /src

# Download all dependencies.
COPY go.* ./
RUN go mod download

# Build the source code.
COPY . ./
RUN make build

# Create a single layer image.
FROM alpine:latest # or search for distroless images
COPY --from=build /src/foo /bin/foo

ENTRYPOINT ["/bin/foo"]
```

See also kvstore's [Dockerfile](https://github.com/jreisinger/kvstore/blob/master/Dockerfile).
