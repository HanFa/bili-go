package bili

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// calculateMD5 calculates the md5sum of the file pointed by the path
func calculateMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hasher := md5.New()
	if _, err = io.Copy(hasher, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func TestClient_DownloadByAid(t *testing.T) {
	err := client.DownloadByAid(DownloadOptionAid{
		Aid: 713884132,
		DownloadOptionCommon: DownloadOptionCommon{
			Page:       0,
			Resolution: Stream240P,
			Mode:       StreamFlv,
			Allow4K:    true,
			OutPath:    "/tmp/test.flv",
		},
	}, true, nil)

	assert.Nil(t, err)
	actualMD5, _ := calculateMD5("/tmp/test.flv")
	assert.Equal(t, "7bf29f794b6b700c846363a2839120d7", actualMD5)
	if err = os.Remove("/tmp/test.flv"); err != nil {
		assert.Fail(t, "cannot clean up the test file")
	}
}

type MyProgressWriter struct {
	curLength   int
	totalLength int
}

func (m *MyProgressWriter) Write(p []byte) (n int, err error) {
	m.curLength += len(p)
	return len(p), nil
}

func (m *MyProgressWriter) SetContentLength(contentLength int) (err error) {
	m.totalLength = contentLength
	return nil
}

func TestClient_DownloadByAid_withProgressWriter(t *testing.T) {
	pw := MyProgressWriter{}
	err := client.DownloadByAid(DownloadOptionAid{
		Aid: 713884132,
		DownloadOptionCommon: DownloadOptionCommon{
			Page:       0,
			Resolution: Stream240P,
			Mode:       StreamFlv,
			Allow4K:    true,
			OutPath:    "/tmp/test.flv",
		},
	}, true, &pw)

	assert.Nil(t, err)
	actualMD5, _ := calculateMD5("/tmp/test.flv")
	assert.Equal(t, "7bf29f794b6b700c846363a2839120d7", actualMD5)
	assert.Equal(t, 2514282, pw.curLength)
	assert.Equal(t, pw.totalLength, pw.curLength)
	if err = os.Remove("/tmp/test.flv"); err != nil {
		assert.Fail(t, "cannot clean up the test file")
	}
}
