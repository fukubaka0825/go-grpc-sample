package main

import (
	"fmt"
	"github.com/takafk9/go-grpc-sample/pkg/gateway"
	"os"


)

func main() {
	if err := gateway.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
