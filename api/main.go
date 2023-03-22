package main

import (
	_ "github.com/blog-api/api"
	"github.com/blog-api/config"
	"github.com/blog-api/midd"
	"github.com/blog-api/router"
	srv "github.com/blog-common"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {
	r := gin.Default()
	//tp, tpErr := tracing.JaegerTraceProvider(config.C.JaegerConfig.Endpoints)
	//if tpErr != nil {
	//	log.Fatal(tpErr)
	//}
	//otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	r.Use(midd.RequestLog())
	r.Use(otelgin.Middleware("project-api"))

	//r.StaticFS("/upload", http.Dir("upload"))
	//路由
	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
