GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOGET=$(GOCMD) get
BUILD=`git rev-parse HEAD`
TIMEBUILD=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
OUT_FILE=watcher

all: build

build: clean
	$(GOBUILD) \
	-ldflags "-X main.buildstamp=$(TIMEBUILD) -X main.githash=$(BUILD)" \
	-o $(OUT_FILE) ./*.go

build_rpi: clean
	GOARCH=arm GOARM=7 \
	$(GOBUILD) \
	-ldflags "-X main.buildstamp=$(TIMEBUILD) -X main.githash=$(BUILD)" \
	-o $(OUT_FILE) main.go

clean:
	rm -f $(OUT_FILE)

get:
	$(GOGET)

run: get
	$(GORUN) ./*.go

run_test: get
	$(GORUN) ./*.go -t