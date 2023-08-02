package realpath_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gandarez/go-realpath"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealpath(t *testing.T) {
	tmpDir := t.TempDir()

	tmpFile, err := os.CreateTemp(tmpDir, "file")
	require.NoError(t, err)

	defer tmpFile.Close()

	path, err := realpath.Realpath(tmpFile.Name())
	require.NoError(t, err)

	assert.Equal(t, tmpFile.Name(), path)
}

func TestRealpath_ZeroLenght(t *testing.T) {
	_, err := realpath.Realpath("")

	assert.EqualError(t, err, "invalid argument")
}

func TestRealpath_NonFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test on windows platform")
	}

	wd, err := os.Getwd()
	require.NoError(t, err)

	_, err = realpath.Realpath("non-file")

	assert.EqualError(t, err, fmt.Sprintf("lstat %s: no such file or directory", filepath.Join(wd, "non-file")))
}

func TestRealpath_NonFile_Windows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skipping test on non-windows platform")
	}

	wd, err := os.Getwd()
	require.NoError(t, err)

	_, err = realpath.Realpath("non-file")

	assert.EqualError(
		t,
		err,
		fmt.Sprintf("CreateFile %s: The system cannot find the file specified.", filepath.Join(wd, "non-file")),
	)
}

func TestRealpath_RelativePath(t *testing.T) {
	path, err := realpath.Realpath("testdata/relative.go")
	require.NoError(t, err)

	wd, err := os.Getwd()
	require.NoError(t, err)

	assert.Equal(t, filepath.Join(wd, "testdata/relative.go"), path)
}

func TestRealpath_Symlink(t *testing.T) {
	tmpDir := t.TempDir()

	tmpFile, err := os.CreateTemp(tmpDir, "file")
	require.NoError(t, err)

	defer tmpFile.Close()

	tmpFileSymlink := filepath.Join(tmpDir, "file_symlink")

	err = os.Symlink(tmpFile.Name(), tmpFileSymlink)
	require.NoError(t, err)

	path, err := realpath.Realpath(tmpFileSymlink)
	require.NoError(t, err)

	assert.Equal(t, tmpFile.Name(), path)
}

func TestRealpath_TrailingSlash(t *testing.T) {
	tmpDir := t.TempDir()

	path, err := realpath.Realpath(tmpDir + string(os.PathSeparator))
	require.NoError(t, err)

	assert.Equal(t, tmpDir, path)
}
