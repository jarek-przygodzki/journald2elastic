GOCMD=go

MODULE   = $(shell basename "$(PWD)")
BIN      = $(CURDIR)/bin
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo unknown)

LDFLAGS = -ldflags '-X github.com/jarek-przygodzki/journald2elastic/app.Version=$(VERSION) -X github.com/jarek-przygodzki/journald2elastic/app.BuildDate=$(DATE)'


install: build
	@echo "  >  Installing binary..."
	$(GOCMD) install -v
build: go-get
	@echo "  >  Building binary..."
	$(GOCMD) build \
		${LDFLAGS} \
		-o $(BIN)/$(basename $(MODULE)) main.go
go-get:
	@echo "  >  Checking if there are any missing dependencies..."
	$(GOCMD) get -t ./...

image:
	docker build -t github.com/jarek-przygodzki/journald2elastic .

test:
	$(GOCMD) test -v ./...