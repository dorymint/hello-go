# example for ldflags

VERSION := $(shell git describe --tags)
ifeq ($(VERSION),)
	VERSION := dev
endif
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git rev-parse --verify --short HEAD)
DATE := $(shell date -u +%Y-%m-%d)

LDFLAGS := -ldflags " \
	-X 'main.Version=$(VERSION)' \
	-X 'main.Branch=$(BRANCH)' \
	-X 'main.Commit=$(COMMIT)' \
	-X 'main.Date=$(DATE)' \
	"

all :
	go run $(LDFLAGS) main.go
