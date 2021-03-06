package bili

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	}, true)

	assert.Nil(t, err)
	actualMD5, _ := calculateMD5("/tmp/test.flv")
	assert.Equal(t, "b748f8d0e3ec41cee2d016d91dcf5396", actualMD5)
	if err = os.Remove("/tmp/test.flv"); err != nil {
		assert.Fail(t, "cannot clean up the test file")
	}
}
