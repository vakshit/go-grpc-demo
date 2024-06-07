generate_protos:
	protoc demo.proto \
	--go_out=demo \
	--go_opt=paths=source_relative \
	--go-grpc_out=demo \
	--go-grpc_opt=paths=source_relative