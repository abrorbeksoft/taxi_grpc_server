package main

import (
	"context"
	"fmt"
	"github.com/abrorbeksoft/taxi_grpc_server/genproto/user_service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	
}

func (s Server) Login(ctx context.Context, user *user_service.User) (*user_service.SuccessMessage, error) {

	fmt.Println(user)

	return &user_service.SuccessMessage{
		Responce: "hello world",
	},nil
}

func main() {
	lis,err:=net.Listen("tcp",":9000")
	if err!=nil {
		fmt.Println(err)
	}
	s:=Server{}
	grpcServer:=grpc.NewServer();

	user_service.RegisterRegistrationSystemServer(grpcServer,&s)

	if err:=grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s",err)
	}
}
