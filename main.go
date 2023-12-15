package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/dyncoch/orders-api/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Printf("failed to start application: %s", err)
	}

}
