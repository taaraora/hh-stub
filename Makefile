MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(patsubst %/,%,$(dir $(MAKEFILE_PATH)))


.PHONY: genproto
genproto:
	docker run --user $(shell id -u):$(shell id -g) --rm -v $(CURRENT_DIR):/defs namely/protoc-all:1.46_0 -l go -f pkg/hh/sessionforwarder.proto -o . --go-source-relative

.PHONY: run-serv
run-serv:
	go run ./cmd/hh-stub


.PHONY: run-client
run-client:
	go run ./cmd/sf-stub