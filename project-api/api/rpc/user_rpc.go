package rpc

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"my_project/project-api/config"
	"my_project/project-common/discovery"
	"my_project/project-common/logs"
	"my_project/project-grpc/user/login"
)

var LoginServiceClient login.LoginServiceClient

//func InitRpcUserClient() {
//	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
//	resolver.Register(etcdRegister)
//
//	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	LoginServiceClient = login.NewLoginServiceClient(conn)
//}

func InitRpcUserClient() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.Dial(
		"etcd:///user",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
