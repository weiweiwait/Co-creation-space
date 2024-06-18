package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"my_project/project-api/config"
	"my_project/project-common/discovery"
	"my_project/project-common/logs"
	loginServiceV1 "my_project/project-user/pkg/service/login.service.v1"
)

var LoginServiceClient loginServiceV1.LoginServiceClient

func InitRpcUserClient() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = loginServiceV1.NewLoginServiceClient(conn)
}
