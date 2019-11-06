FROM golang:1.13-alpine AS build

ARG GOOS=darwin
ARG GOARCH=amd64

# Set the current working directory inside container.
WORKDIR /go/src/homepage

# Install tools required for building the app.
RUN apk add git

# Download all dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Build the app.
COPY . ./
RUN go build -o /bin/homepage

# Create a single layer image.
#FROM scratch # -> this doesn't work
FROM alpine:latest
COPY --from=build /bin/homepage /bin/homepage
RUN apk update
RUN apk add git

EXPOSE 5001
ENTRYPOINT ["/bin/homepage"]
