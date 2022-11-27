package adapter

import (
	"errors"
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
)

type HuoshanAdapter struct{}

type HuoshanAdapterResponse struct {
	StatusCode int `json:"status_code"`
	Data       struct {
		UserInfo struct {
			EncryptedID string `json:"encrypted_id"`
		} `json:"user_info"`
		ItemInfo struct {
			ItemID         string `json:"item_id"`
			Cover          string `json:"cover"`
			URL            string `json:"url"`
			ReflowPlayable int    `json:"ReflowPlayable"`
		} `json:"item_info"`
	} `json:"data"`
}

func (a HuoshanAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Sugar.Error(r)
		}
	}()
	headerResp, err := restyHttp.GetHttpClient().R().Head(url)
	if err != nil {
		return nil, err
	}
	loc := headerResp.RawResponse.Request.Response.Header.Get("Location")
	re := regexp.MustCompile(`item_id=(.*)&tag`)
	match := re.FindStringSubmatch(loc)
	if len(match) != 2 {
		return nil, errors.New("匹配视频id失败")
	}
	getResp, err := restyHttp.GetMobileHttpRequest().Get("https://share.huoshan.com/api/item/info?item_id=" + match[1])
	if err != nil {
		return nil, err
	}
	var huoshanResponse HuoshanAdapterResponse
	err = json.Cjson.Unmarshal(getResp.Body(), &huoshanResponse)
	if err != nil {
		return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
		Cover:                  huoshanResponse.Data.ItemInfo.Cover,
		NoWatermarkDownloadUrl: huoshanResponse.Data.ItemInfo.URL,
	}
	return shortVideoInfo, nil
}
