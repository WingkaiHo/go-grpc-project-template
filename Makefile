.PHONY: proto
proto:
	protoc -I api/proto/demo ./api/proto/demo/*.proto --go_out=plugins=grpc:api/proto/demo


.PHONY: build
build:
	GOARCH="amd64" go build ./