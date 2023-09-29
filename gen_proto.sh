protoc proto/user.proto \
  --go_out=server/proto \
  --go-grpc_out=server/proto \
  --go-grpc_opt=require_unimplemented_servers=false
