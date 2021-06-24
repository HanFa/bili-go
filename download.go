package bili

import (
	"errors"
	"fmt"
)

// ProgressWriter is useful for progress report
type ProgressWriter interface {
	Write(p []byte) (n int, err error)
	SetContentLength(contentLength int) (err error) // set the total length of the video
}

// DownloadOptionCommon is the common part of the request payload for the video download
type DownloadOptionCommon struct {
	Page       int
	Resolution StreamResolutionMode
	Mode       StreamMode
	Allow4K    bool
	OutPath    string
}

// DownloadOptionAid represents a DownloadOptionCommon specified by video Aid
type DownloadOptionAid struct {
	Aid int
	DownloadOptionCommon
}

// DownloadOptionBvid represents a DownloadOptionCommon specified by video Bvid
type DownloadOptionBvid struct {
	Bvid string
	DownloadOptionCommon
}

// download downloads the video specified by the returned streaming url
// the progress bar will presented in the standard output if showProgress is True
// progressWriter provides a hook whenever the progress is updated
func (c *Client) download(streamUrlResponse StreamUrlResponse, option DownloadOptionCommon, showProgress bool, progressWriter ProgressWriter) error {
	if option.Mode == StreamFlv || option.Mode == StreamLowResMp4 {
		parts := len(streamUrlResponse.Data.Durl)
		if parts == 0 {
			return errors.New("no flv/mp4 url found")
		}
		if parts > 1 { // contains multi parts
			for _, durl := range streamUrlResponse.Data.Durl {
				if err := HttpGetAsFile(
					c.client, durl.Url, fmt.Sprintf("%s-%d", option.OutPath, durl.Order), showProgress, progressWriter); err != nil {
					return err
				}
			}
		} else {
			durl := streamUrlResponse.Data.Durl[0]
			if err := HttpGetAsFile(c.client, durl.Url, option.OutPath, showProgress, progressWriter); err != nil {
				return err
			}
		}
	} else if option.Mode == StreamDash {
		return errors.New("dash mode is not yet supported")
	} else {
		return errors.New("invalid streaming node")
	}

	return nil
}

// DownloadByAid download the video specified by Aid
func (c *Client) DownloadByAid(option DownloadOptionAid, showProgress bool, progressWriter ProgressWriter) error {
	urlResponse, err := c.GetStreamUrlAvid(option.Aid, option.Page, option.Resolution, option.Mode, option.Allow4K)
	if err != nil {
		return err
	}
	return c.download(urlResponse, option.DownloadOptionCommon, showProgress, progressWriter)
}

// DownloadByBvid download the video specified by Bvid
func (c *Client) DownloadByBvid(option DownloadOptionBvid, showProgress bool, progressWriter ProgressWriter) error {
	urlResponse, err := c.GetStreamUrlBvid(option.Bvid, option.Page, option.Resolution, option.Mode, option.Allow4K)
	if err != nil {
		return err
	}
	return c.download(urlResponse, option.DownloadOptionCommon, showProgress, progressWriter)
}
