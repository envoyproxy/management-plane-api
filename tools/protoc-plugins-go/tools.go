// +build tools

// This package tracks build dependencies so they are not removed when `go mod tidy` is run.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-check-lint"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/go-swagger/go-swagger"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
)
