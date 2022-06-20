build:
	go build

run: build
	# Use the current working directory as the local repo path
	REPO_LOCAL_PATH="." ./homepage

# can be more for multiplatform images
PLATFORMS := linux/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

release: $(PLATFORMS)

$(PLATFORMS):
	docker build --build-arg GOOS=$(os) --build-arg GOARCH=$(arch) -t homepage-$(os)-$(arch) .

	# Push image to public registry - hub.docker.com
	docker login
	docker tag homepage-$(os)-$(arch):latest reisinge/homepage-$(os)-$(arch):latest
	docker push reisinge/homepage-$(os)-$(arch):latest
