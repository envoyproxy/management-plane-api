# Management Plane API
The Management Plane API (MPA) Working Group will allow the Envoy community to collaborate on the development of higher-level features, which can be defined as additional functionality not pertaining to normal control and data plane operations.

## Community

### Slack
Join Envoy Slack [per the instructions](https://github.com/envoyproxy/envoy#contact) in the Envoy repo, and join channel `#management-plane-api`.

### Meetings
A summary is available on the [Wiki](https://github.com/envoyproxy/management-plane-api/wiki) of past and current meetings and work items.

# API Definitions
MPA definitions will be in `proto3` format. Since the APIs may be accessed from clients without native support for gRPC (e.g. browsers), we will use [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to generate JSON endpoints that can be used in addition to gRPC. protoc-gen-swagger from the grpc-gateway project will also be used to generate Swagger files, which provide code-gen and documentation for clients that wish to use JSON endpoints.

For an example of all of the functionality that `grpc-gateway` offers in combination with `protoc-gen-swagger` OpenAPI annotations, see [a_bit_of_everything.proto](https://github.com/grpc-ecosystem/grpc-gateway/blob/67246c6b822b91a64e574b1773dbfd6f174b3827/examples/internal/proto/examplepb/a_bit_of_everything.proto) and the resulting Swagger [a_bit_of_everything.swagger.json](https://github.com/grpc-ecosystem/grpc-gateway/blob/67246c6b822b91a64e574b1773dbfd6f174b3827/examples/internal/proto/examplepb/a_bit_of_everything.swagger.json)

# Usage
- Define protos in the `api/` folder.
- `make` to validate the definition and generate the Swagger specification.
- `make view` to see documentation in the browser for the Swagger.

# Dependencies
This project requires the following dependencies.
- realpath
- go

## MacOS
To install the required dependencies, run:
`brew install coreutils`
`brew install go`

