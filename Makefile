BINDATA_IGNORE = $(shell git ls-files -io --exclude-standard $< | sed 's/^/-ignore=/;s/[.]/[.]/g')
GO_IMAGE=golang:alpine

# Hardcoded based on vungle/golang:1.9.2 for now for convenience, for true value
# execute $(shell docker run --rm $(GO_IMAGE) /bin/bash -c 'echo $$GOPATH').
DOCKER_GOPATH=/go
DOCKER_WORKDIR=$(DOCKER_GOPATH)/src/Massmail
CMD=Massmail

DOCKER_SHELL=\
docker run --rm \
-w $(DOCKER_WORKDIR) \
-v `pwd`:$(DOCKER_WORKDIR) \
$(SHELL_OPTS) \
$(GO_IMAGE)

assets: static/
	go-bindata -o data/bindata.go -pkg data $(BINDATA_OPTS) $(BINDATA_IGNORE) -ignore=[.]gitignore -ignore=[.]gitkeep $<... views/...

front-end:
