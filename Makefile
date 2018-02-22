.PHONY: lint-install lint

lint-install:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

lint:
	@echo "lint"
	@gometalinter -s vendor --deadline=60s ./...

build:
	go build -race

install-deps:
	dep ensure -v

update-deps:
	dep ensure -v -update

all: lint-install lint install-deps build
