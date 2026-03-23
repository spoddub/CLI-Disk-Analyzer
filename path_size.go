package code

import (
	"fmt"
	"os"
)

func GetSize(path string) (int, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return int(info.Size()), nil
	}

	total := 0

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		if entryInfo.IsDir() {
			continue
		}

		total += int(entryInfo.Size())
	}

	return total, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%dB", size), nil
}
