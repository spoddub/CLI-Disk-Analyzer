package main

import (
	"code"
	"context"
	"log"
	"os"
)

func main() {
	app := code.NewApp()

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
