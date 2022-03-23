protoc -I protobuf/ protobuf/*.proto --go_out=api --go-grpc_out=api --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
