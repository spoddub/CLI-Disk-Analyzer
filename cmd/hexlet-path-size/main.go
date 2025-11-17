package main

import (
	"code"
	"context"
	"log"
	"os"
)

func main() {
	app := pathsize.NewApp()

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
