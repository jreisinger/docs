```
FROM golang:1.13 AS build

# Set the current working directory inside container.
WORKDIR /app

# Download all dependencies.
COPY go.* ./
RUN go mod download

# Build the app.
COPY . ./
RUN make build

# Create a single layer image.
FROM alpine:latest
COPY --from=build /app/foo /bin/foo

ENTRYPOINT ["/bin/foo"]
```
