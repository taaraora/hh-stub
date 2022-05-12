MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(patsubst %/,%,$(dir $(MAKEFILE_PATH)))


.PHONY: genproto
genproto:
	docker run --user $(shell id -u):$(shell id -g) --rm -v $(CURRENT_DIR):/defs namely/protoc-all:1.46_0 -l go -f pkg/hh/hh.proto -o . --go-source-relative

.PHONY: run
run:
	go run ./cmd/hh-stub