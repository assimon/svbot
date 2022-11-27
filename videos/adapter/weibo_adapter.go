package adapter

import (
	"errors"
	"fmt"
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"strings"
)

type WeiboAdapter struct{}

type WeiboAdapterResponse struct {
	Ok   int `json:"ok"`
	Data struct {
		ObjectID   string `json:"object_id"`
		ObjectType string `json:"object_type"`
		Object     struct {
			Summary string `json:"summary"`
			Author  struct {
				ID                int64  `json:"id"`
				ScreenName        string `json:"screen_name"`
				ProfileImageURL   string `json:"profile_image_url"`
				ProfileURL        string `json:"profile_url"`
				StatusesCount     int    `json:"statuses_count"`
				Verified          bool   `json:"verified"`
				VerifiedType      int    `json:"verified_type"`
				VerifiedTypeExt   int    `json:"verified_type_ext"`
				VerifiedReason    string `json:"verified_reason"`
				CloseBlueV        bool   `json:"close_blue_v"`
				Description       string `json:"description"`
				Gender            string `json:"gender"`
				Mbtype            int    `json:"mbtype"`
				Svip              int    `json:"svip"`
				Urank             int    `json:"urank"`
				Mbrank            int    `json:"mbrank"`
				FollowMe          bool   `json:"follow_me"`
				Following         bool   `json:"following"`
				FollowCount       int    `json:"follow_count"`
				FollowersCount    string `json:"followers_count"`
				FollowersCountStr string `json:"followers_count_str"`
				CoverImagePhone   string `json:"cover_image_phone"`
				AvatarHd          string `json:"avatar_hd"`
				Like              bool   `json:"like"`
				LikeMe            bool   `json:"like_me"`
			} `json:"author"`
			Stream struct {
				Duration float64 `json:"duration"`
				Format   string  `json:"format"`
				Width    int     `json:"width"`
				HdURL    string  `json:"hd_url"`
				URL      string  `json:"url"`
				Height   int     `json:"height"`
			} `json:"stream"`
			CreatedAt string `json:"created_at"`
			Image     struct {
				Width       int    `json:"width"`
				Pid         string `json:"pid"`
				Source      int    `json:"source"`
				IsSelfCover int    `json:"is_self_cover"`
				Type        int    `json:"type"`
				URL         string `json:"url"`
				Height      int    `json:"height"`
			} `json:"image"`
		} `json:"object"`
	} `json:"data"`
}

func (a WeiboAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Sugar.Error(r)
		}
	}()
	headerResp, err := restyHttp.GetMobileHttpRequest().Head(url)
	if err != nil {
		return nil, err
	}
	loc := headerResp.RawResponse.Request.Response.Header.Get("Location")
	pathList := strings.Split(loc, "/")
	if len(pathList) != 5 {
		return nil, errors.New("weibo id is err")
	}
	path := fmt.Sprintf("/s/video/object?object_id=%s&mid=%s", pathList[3], pathList[4])
	getResp, err := restyHttp.GetMobileHttpRequest().SetHeader("path", path).Get("https://video.h5.weibo.cn" + path)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var weiboResponse WeiboAdapterResponse
	err = json.Cjson.Unmarshal(getResp.Body(), &weiboResponse)
	if err != nil {
		return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             weiboResponse.Data.Object.Author.ScreenName,
		AuthorAvatar:           weiboResponse.Data.Object.Author.ProfileImageURL,
		Title:                  weiboResponse.Data.Object.Summary,
		Cover:                  "https:" + weiboResponse.Data.Object.Image.URL,
		NoWatermarkDownloadUrl: "https:" + weiboResponse.Data.Object.Stream.HdURL,
	}
	return shortVideoInfo, nil
}
