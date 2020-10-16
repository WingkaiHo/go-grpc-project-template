.PHONY: proto
proto:
	protoc -I api/proto/demo ./api/proto/demo/* --go_out=plugins=grpc:api/proto/demo