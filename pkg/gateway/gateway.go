package gateway

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/takafk9/go-grpc-sample/pkg/api"

)

type Config struct {
	Server string
	Port   string
}

func RunServer() error {

	cfg := getConfig()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterBookServiceHandlerFromEndpoint(ctx, mux, cfg.Server, opts)
	if err != nil {
		return err
	}
	log.Println("starting proxy server...")

	return http.ListenAndServe(":"+cfg.Port, mux)
}

func getConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.Server, "server", "", "grpc server and port")
	flag.StringVar(&cfg.Port, "port", "", "gateway server port")
	flag.Parse()

	return cfg
}
