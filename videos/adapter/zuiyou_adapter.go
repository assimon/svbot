package adapter

import (
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
)

type ZuiyouAdapter struct{}

func (a ZuiyouAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Sugar.Error(r)
		}
	}()
	getResp, err := restyHttp.GetMobileHttpRequest().Get(url)
	if err != nil {
		return nil, err
	}
	videoUrl := regexp.MustCompile(`fullscreen=\"false\" src=\"(.*?)\"`).FindStringSubmatch(getResp.String())
	videoTitle := regexp.MustCompile(`:<\/span><h1>(.*?)<\/h1><\/div><div class=`).FindStringSubmatch(getResp.String())
	videoCover := regexp.MustCompile(`poster=\"(.*?)\">`).FindStringSubmatch(getResp.String())
	videoAuthor := regexp.MustCompile(`<span class=\"SharePostCard__name\">(.*?)<\/span>`).FindStringSubmatch(getResp.String())
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             videoAuthor[1],
		Title:                  videoTitle[1],
		Cover:                  videoCover[1],
		NoWatermarkDownloadUrl: videoUrl[1],
	}
	return shortVideoInfo, nil
}
