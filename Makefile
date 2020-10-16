.PHONY: proto
proto:
	protoc -I api/proto ./api/proto/* --go_out=plugins=grpc:api/proto