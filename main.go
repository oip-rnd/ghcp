package main

import (
	"context"
	"os"

	"github.com/int128/ghcp/di"
)

var version = "HEAD"

func main() {
	os.Exit(di.NewCmd().Run(context.Background(), os.Args, version))
}
