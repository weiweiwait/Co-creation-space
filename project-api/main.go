package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	_ "my_project/project-api/api"
	"my_project/project-api/api/midd"
	"my_project/project-api/config"
	"my_project/project-api/router"
	"my_project/project-api/tracing"
	srv "my_project/project-common"
	"net/http"
)

func main() {
	r := gin.Default()
	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	r.Use(midd.RequestLog())
	r.Use(otelgin.Middleware("project-api"))
	r.StaticFS("/upload", http.Dir("upload"))
	//路由
	router.InitRouter(r)
	//开启pprof 默认的访问路径是/debug/pprof
	pprof.Register(r)

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
