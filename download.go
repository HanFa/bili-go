package bili

import (
	"errors"
	"fmt"
)

// ProgressWriter is useful for progress report
type ProgressWriter struct {
	contentLength int // the total length of the video
	curLength     int // received length
}

func (pw *ProgressWriter) Write(p []byte) (n int, err error) {
	pw.curLength += len(p)
	return len(p), nil
}

type DownloadOptionCommon struct {
	Page       int
	Resolution StreamResolutionMode
	Mode       StreamMode
	Allow4K    bool
	OutPath    string
}

type DownloadOptionAid struct {
	Aid int
	DownloadOptionCommon
}

type DownloadOptionBvid struct {
	Bvid string
	DownloadOptionCommon
}

func (c *Client) download(streamUrlResponse StreamUrlResponse, option DownloadOptionCommon, showProgress bool, progressWriter *ProgressWriter) error {
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

// DownloadByAid download video by Aid from BiliBili
func (c *Client) DownloadByAid(option DownloadOptionAid, showProgress bool, progressWriter *ProgressWriter) error {
	urlResponse, err := c.GetStreamUrlAvid(option.Aid, option.Page, option.Resolution, option.Mode, option.Allow4K)
	if err != nil {
		return err
	}
	return c.download(urlResponse, option.DownloadOptionCommon, showProgress, progressWriter)
}

// DownloadByBvid download video by Bvid from BiliBili
func (c *Client) DownloadByBvid(option DownloadOptionBvid, showProgress bool, progressWriter *ProgressWriter) error {
	urlResponse, err := c.GetStreamUrlBvid(option.Bvid, option.Page, option.Resolution, option.Mode, option.Allow4K)
	if err != nil {
		return err
	}
	return c.download(urlResponse, option.DownloadOptionCommon, showProgress, progressWriter)
}
