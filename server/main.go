package main

import (
	"fmt"
	"foo/server/proto"
	"foo/server/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = 50051
)

func main() {
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーの生成
	server := grpc.NewServer()

	// 実装したサービスの登録
	proto.RegisterUserManagerServer(server, &service.UserManagerServer{})

	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(server)

	// サーバーを起動
	server.Serve(listenPort)
}
