package pathsize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSize_File(t *testing.T) {
	size1, err := GetSize("testdata/file1.txt")
	assert.NoError(t, err)
	size2, err := GetSize("testdata/file2.txt")
	assert.NoError(t, err)
	assert.Equal(t, size1, size2)
}

func TestGetSize_Dir(t *testing.T) {
	size1, err := GetSize("testdata/dir1/a.txt")
	assert.NoError(t, err)
	size2, err := GetSize("testdata/dir1/b.txt")
	assert.NoError(t, err)
	size3, err := GetSize("testdata/dir1")
	assert.NoError(t, err)
	assert.Equal(t, size1+size2, size3)
}

func TestGetSize_NonExistentDir(t *testing.T) {
	_, err := GetSize("testdata/nonexistent.txt")
	assert.Error(t, err)
}
