package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	srv "my_project/project-common"
	"my_project/project-common/logs"
	"my_project/project-user/config"
	"my_project/project-user/router"
	"my_project/project-user/tracing"
)

func main() {
	r := gin.Default()
	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	lc := &logs.LogConfig{
		DebugFileName: "/home/fjw/GolandProjects/my_project/logs/debug/project-debug.log",
		InfoFileName:  "/home/fjw/GolandProjects/my_project/logs/info/project-info.log",
		WarnFileName:  "/home/fjw/GolandProjects/my_project/logs/error/project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	//路由
	router.InitRouter(r)
	//注册grpc服务
	gc := router.RegisterGrpc()
	//grpc注册到etcd
	router.RegisterEtcdServer()
	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
