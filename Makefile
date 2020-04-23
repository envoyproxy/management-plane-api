SHELL:=/usr/bin/env bash
.DEFAULT_GOAL:=all

MAKEFLAGS += --no-print-directory

.PHONY: all
all: proto

.PHONY: proto
proto:
	tools/compile-protos.sh

.PHONY: view
view:
	build/bin/swagger serve api/test/v1alpha1/pb.swagger.yaml

# TODO: lint, lint-fix
