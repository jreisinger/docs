# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS build

# Set the current working directory inside container.
WORKDIR /go/src/homepage

# Install tools required for building the app.
RUN apk add --no-cache git

# Download all dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Build the app.
COPY . ./
RUN go build -o /bin/homepage

# Create a single layer image.
#FROM scratch # -> this doesn't work
FROM alpine:latest
WORKDIR /app/homepage
COPY --from=build /bin/homepage /app/homepage/homepage
COPY --from=build /go/src/homepage/tmpl /app/homepage/tmpl
RUN apk add --no-cache git

EXPOSE 5001
ENTRYPOINT ["/app/homepage/homepage"]
