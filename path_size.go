package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !all && strings.HasPrefix(info.Name(), ".") {
		return 0, nil
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var total int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		childPath := filepath.Join(path, entry.Name())

		if entry.IsDir() && !recursive {
			continue
		}

		childSize, err := GetSize(childPath, recursive, all)
		if err != nil {
			return 0, err
		}

		total += childSize
	}

	return total, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return FormatSize(size, human), nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)
	unitIndex := 0

	for value >= 1024 && unitIndex < len(units)-1 {
		value /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[unitIndex])
}
