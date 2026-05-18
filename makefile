
protoc:
	protoc --proto_path=contracts/proto \
       --go_out=contracts/pb --go_opt=paths=source_relative \
       --go-grpc_out=contracts/pb --go-grpc_opt=paths=source_relative \
       contracts/proto/**/*.proto