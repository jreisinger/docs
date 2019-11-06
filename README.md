Code, Docker image and data for https://reisinge.net. If you edit stuff in `data` folder it will be reflected almost instantly in https://reisinge.net.

## Initialize the project

```
# Let's use Go modules for dependency management.
cd ~/github
mkdir homepage
go mod init github.com/jreisinger/homepage

# Create the simple app.
vim main.go
```

## Build, run and test the app

```
go build main.go  # will add modules to go.mod and go.sum if needed
./main
curl localhost:5001
```

## Docker stuff

```
# Build multiplatform images
export GOOS=linux  # darwin
export GOARCH=arm  # amd64
docker build --build-arg GOOS=$GOOS --build-arg GOARCH=$GOARCH -t homepage-$GOOS-$GOARCH .

# Push image to public registry - hub.docker.com
docker login
docker tag homepage-$GOOS-$GOARCH:latest reisinge/homepage-$GOOS-$GOARCH:latest
docker push reisinge/homepage-$GOOS-$GOARCH:latest
```

```
# Run the image.
docker run -d -p 5001:5001 homepage-$GOOS-$GOARCH

# Test the image.
curl localhost:5001
```

Docker image building is based on:

* https://www.callicoder.com/docker-golang-image-container-example/
* https://github.com/jreisinger/quotes
