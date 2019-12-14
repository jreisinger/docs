Code and data for my [home page](https://reisinge.net).

## Usage

Edit stuff in `data` folder. It will be reflected almost instantly in the [homepage](https://reisinge.net).

## Development

Initialize the project

```
# Let's use Go modules for dependency management.
cd ~/github
mkdir homepage
go mod init github.com/jreisinger/homepage

# Create the simple app.
vim main.go
```

Test locally:

```
make build
./main
curl localhost:5001
```

Release to dockerhub:

```
make release
```
