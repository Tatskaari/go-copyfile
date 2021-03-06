package copyfile

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	// Clean up any old test files before we begin
	paths, _ := filepath.Glob("test_data/test*.txt")
	for _, path := range paths {
		os.Remove(path)
	}
}

func TestCopy(t *testing.T) {
	b, err := ioutil.ReadFile("test_data/input.txt")
	assert.NoError(t, err)
	var c Copier
	err = c.Copy("test_data/input.txt", "test_data/test2.txt")
	assert.NoError(t, err)
	b2, err := ioutil.ReadFile("test_data/test2.txt")
	assert.NoError(t, err)
	assert.Equal(t, b, b2)
	// Write some stuff to the new file, it should not modify the original under
	// any circumstances.
	err = ioutil.WriteFile("test_data/test2.txt", []byte("testing"), 0644)
	assert.NoError(t, err)
	b, err = ioutil.ReadFile("test_data/input.txt")
	assert.NoError(t, err)
	assert.NotEqual(t, []byte("testing"), b)
}

func TestCopyGeneric(t *testing.T) {
	b, err := ioutil.ReadFile("test_data/input.txt")
	assert.NoError(t, err)
	c := Copier{Generic: true}
	err = c.Copy("test_data/input.txt", "test_data/test3.txt")
	assert.NoError(t, err)
	b2, err := ioutil.ReadFile("test_data/test3.txt")
	assert.NoError(t, err)
	assert.Equal(t, b, b2)
}

func TestCopyNonExistingFile(t *testing.T) {
	var c Copier
	err := c.Copy("test_data/doesnt_exist.txt", "test_data/test4.txt")
	assert.Error(t, err)
}

func TestCopyToNonWritableFile(t *testing.T) {
	// Ensure permissions are right, this seems to lose them after git clone?
	os.Chmod("test_data/readonly.txt", 0444)
	var c Copier
	err := c.Copy("test_data/input.txt", "test_data/readonly.txt")
	assert.Error(t, err)
}

func TestLink(t *testing.T) {
	var c Copier
	err := c.Link("test_data/input.txt", "test_data/test5.txt")
	assert.NoError(t, err)
}

func TestIsSameFile(t *testing.T) {
	var c Copier
	assert.True(t, c.IsSameFile("test_data/input.txt", "test_data/input.txt"))
	assert.False(t, c.IsSameFile("test_data/input.txt", "test_data"))
}
