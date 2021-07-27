package confwatcher

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const TEST_TMPDIR = "TEST_TMPDIR"

// TestTmpDir returns the path the Bazel test temp directory.
// If TEST_TMPDIR is not defined, it returns the OS default temp dir.
func testTmpDir() string {
	if tmp, ok := os.LookupEnv(TEST_TMPDIR); ok {
		return tmp
	}
	return os.TempDir()
}

func writeTempFile(byts []byte) (string, error) {
	tmpf, err := ioutil.TempFile(testTmpDir(), "confwatcher-")
	if err != nil {
		return "", err
	}
	defer tmpf.Close()

	_, err = tmpf.Write(byts)
	if err != nil {
		return "", err
	}

	return tmpf.Name(), nil
}

func TestNoFile(t *testing.T) {
	_, err := New("/nonexistent")
	require.Error(t, err)
}

func TestWrite(t *testing.T) {
	fpath, err := writeTempFile([]byte("{}"))
	require.NoError(t, err)

	w, err := New(fpath)
	require.NoError(t, err)
	defer w.Close()

	func() {
		f, err := os.Create(fpath)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write([]byte("{}"))
		require.NoError(t, err)
	}()

	select {
	case <-w.Watch():
	case <-time.After(500 * time.Millisecond):
		t.Errorf("timed out")
		return
	}
}

func TestWriteMultipleTimes(t *testing.T) {
	fpath, err := writeTempFile([]byte("{}"))
	require.NoError(t, err)

	w, err := New(fpath)
	require.NoError(t, err)
	defer w.Close()

	func() {
		f, err := os.Create(fpath)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write([]byte("{}"))
		require.NoError(t, err)
	}()

	time.Sleep(10 * time.Millisecond)

	func() {
		f, err := os.Create(fpath)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write([]byte("{}"))
		require.NoError(t, err)
	}()

	select {
	case <-w.Watch():
	case <-time.After(500 * time.Millisecond):
		t.Errorf("timed out")
		return
	}

	select {
	case <-time.After(500 * time.Millisecond):
	case <-w.Watch():
		t.Errorf("should not happen")
		return
	}
}

func TestDeleteCreate(t *testing.T) {
	fpath, err := writeTempFile([]byte("{}"))
	require.NoError(t, err)

	w, err := New(fpath)
	require.NoError(t, err)
	defer w.Close()

	os.Remove(fpath)
	time.Sleep(10 * time.Millisecond)

	func() {
		f, err := os.Create(fpath)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write([]byte("{}"))
		require.NoError(t, err)
	}()

	select {
	case <-w.Watch():
	case <-time.After(500 * time.Millisecond):
		t.Errorf("timed out")
		return
	}
}

func TestSymlinkDeleteCreate(t *testing.T) {
	fpath, err := writeTempFile([]byte("{}"))
	require.NoError(t, err)

	err = os.Symlink(fpath, fpath+"-sym")
	require.NoError(t, err)

	w, err := New(fpath + "-sym")
	require.NoError(t, err)
	defer w.Close()

	os.Remove(fpath)

	func() {
		f, err := os.Create(fpath)
		require.NoError(t, err)
		defer f.Close()

		_, err = f.Write([]byte("{}"))
		require.NoError(t, err)
	}()

	select {
	case <-w.Watch():
	case <-time.After(500 * time.Millisecond):
		t.Errorf("timed out")
		return
	}
}
