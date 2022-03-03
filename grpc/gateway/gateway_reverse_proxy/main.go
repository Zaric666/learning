package main

import (
	"context"
	"flag"
	gw "github.com/Zaric666/learning/grpc/gateway/proto"
	lrf "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"os"
)

func injectHeadersIntoMetadata(ctx context.Context, req *http.Request) metadata.MD {
	//https://aspenmesh.io/2018/04/tracing-grpc-with-istio/
	var (
		otHeaders = []string{
			"x-request-id",
			"x-b3-traceid",
			"x-b3-spanid",
			"x-b3-parentspanid",
			"x-b3-sampled",
			"x-b3-flags",
			"x-ot-span-context"}
	)
	var pairs []string

	for _, h := range otHeaders {
		if v := req.Header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	return metadata.Pairs(pairs...)
}

type annotator func(context.Context, *http.Request) metadata.MD

func chainGrpcAnnotators(annotators ...annotator) annotator {
	return func(c context.Context, r *http.Request) metadata.MD {
		var mds []metadata.MD
		for _, a := range annotators {
			mds = append(mds, a(c, r))
		}
		return metadata.Join(mds...)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	annotators := []annotator{injectHeadersIntoMetadata}

	mux := runtime.NewServeMux(
		runtime.WithMetadata(chainGrpcAnnotators(annotators...)),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	formatter := lrf.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(getEnv("LOG_LEVEL", "info"))
	if err != nil {
		log.Error(err)
	}
	log.SetLevel(level)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
