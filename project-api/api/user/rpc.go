package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	loginServiceV1 "my_project/project-user/pkg/service/login.service.v1"
)

var LoginServiceClient loginServiceV1.LoginServiceClient

func InitRpcUserClient() {
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = loginServiceV1.NewLoginServiceClient(conn)
}
