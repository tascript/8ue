package main

import (
	"context"
	"flag"

	"github.com/tetratelabs/wazero"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)
}
