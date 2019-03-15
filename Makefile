GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOFMT=gofmt
BUILD=`git rev-parse HEAD`
TIMEBUILD=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
OUT_FILE=watcher

all: tests

init:
	dep ensure -v

build: clean
	$(GOBUILD) \
	-ldflags "-X main.buildstamp=$(TIMEBUILD) -X main.githash=$(BUILD)" \
	-o $(OUT_FILE) ./*.go

build_rpi: clean
	GOARCH=arm GOARM=7 \
	$(GOBUILD) \
	-ldflags "-X main.buildstamp=$(TIMEBUILD) -X main.githash=$(BUILD)" \
	-o $(OUT_FILE) ./*.go

clean:
	rm -f $(OUT_FILE)

run: build
	$(GORUN) ./*.go

run_test: get
	$(GORUN) ./*.go -t

tests:
	$(GOCMD) test -v ./...

format:
	$(GOFMT) -w .