package adapter

import (
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
	"strings"
)

type LvZhouAdapter struct{}

func (a LvZhouAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Sugar.Error(r)
		}
	}()
	getResp, err := restyHttp.GetMobileHttpRequest().Get(url)
	if err != nil {
		return nil, err
	}
	videoCover := regexp.MustCompile(`<div style=\"background-image:url\((.*)\)`).FindStringSubmatch(getResp.String())
	videoUrl := regexp.MustCompile(`<video src=\"([^\"]*)\"`).FindStringSubmatch(getResp.String())
	videoAuthor := regexp.MustCompile(`<div class=\"nickname\">(.*)<\/div>`).FindStringSubmatch(getResp.String())
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             videoAuthor[1],
		Cover:                  videoCover[1],
		NoWatermarkDownloadUrl: strings.ReplaceAll(videoUrl[1], "amp;", ""),
	}
	return shortVideoInfo, nil
}
