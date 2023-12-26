package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/h2p2f/ntflw-probe/internal/probe/app"
)

func main() {

	ctx := context.Background()

	sigint := make(chan os.Signal, 1)
	connectionsClosed := make(chan struct{})

	signal.Notify(sigint, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	app.Run(ctx, sigint, connectionsClosed)

	<-connectionsClosed

	close(sigint)

}
