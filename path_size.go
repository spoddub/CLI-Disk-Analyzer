package pathsize

import "github.com/urfave/cli/v3"

func NewApp() *cli.Command {
	app := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory;",
		UsageText: "hexlet-path-size [global options]",
	}

	return app
}
