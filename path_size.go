package pathsize

import (
	"github.com/urfave/cli/v3"
	"os"
	"path/filepath"
)

func NewApp() *cli.Command {
	app := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory;",
		UsageText: "hexlet-path-size [global options]",
	}

	return app
}

func GetSize(path string) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		childInfo, err := os.Lstat(childPath)

		if err != nil {
			return 0, err
		}

		if childInfo.Mode().IsRegular() {
			size += childInfo.Size()
		}
	}
	return size, nil
}
