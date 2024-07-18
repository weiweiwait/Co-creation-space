package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	srv "my_project/project-common"
	"my_project/project-project/config"
	"my_project/project-project/router"
	"my_project/project-project/tracing"
)

func main() {
	r := gin.Default()
	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	//路由
	router.InitRouter(r)
	//初始化rpc调用
	router.InitUserRpc()
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()
	//初始化kafka
	c := config.InitKafkaWriter()
	stop := func() {
		gc.Stop()
		c()
	}

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
