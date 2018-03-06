GOCMD=go
GOBUILD=$(GOCMD) build
BUILD=`git rev-parse HEAD`
TIMEBUILD=`date -u '+%Y-%m-%d_%I:%M:%S%p'`

all: build

build:
	$(GOBUILD) \
	-ldflags "-X main.buildstamp=$(TIMEBUILD) -X main.githash=$(BUILD)" \
	-o watcher main.go

build_rpi:
	GOARCH=arm GOARM=7 \
	$(GOBUILD) \
	-ldflags "-X main.buildstamp=$(TIMEBUILD) -X main.githash=$(BUILD)" \
	-o watcher main.go