package main

import (
	"fmt"
	"os"
	"github.com/takafk9/go-grpc-sample/pkg/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

