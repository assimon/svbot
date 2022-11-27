package adapter

import (
	"fmt"
	"github.com/assimon/svbot/internal/config"
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
)

type KuaiShouAdapter struct{}

type KuaishouAdapterResponse struct {
	Result    int `json:"result"`
	ShareInfo struct {
		GroupName     string `json:"groupName"`
		ShareTitle    string `json:"shareTitle"`
		ShareSubTitle string `json:"shareSubTitle"`
		ShareID       string `json:"shareId"`
		ShareType     string `json:"shareType"`
		DocID         string `json:"docId"`
	} `json:"shareInfo"`
	Comments []struct {
		Content           string        `json:"content"`
		Time              string        `json:"time"`
		Type              int           `json:"type"`
		Timestamp         int64         `json:"timestamp"`
		AuthorVerified    bool          `json:"authorVerified"`
		CommentAuthorTags []interface{} `json:"commentAuthorTags"`
		CommentBottomTags []interface{} `json:"commentBottomTags"`
		CashTags          struct {
		} `json:"cashTags"`
		Headurl     string `json:"headurl"`
		PhotoID     int64  `json:"photo_id"`
		UserID      int    `json:"user_id"`
		AuthorID    int    `json:"author_id"`
		AuthorLiked bool   `json:"author_liked"`
		ReplyTo     int    `json:"reply_to"`
		UserSex     string `json:"user_sex"`
		Headurls    []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"headurls"`
		AuthorName      string `json:"author_name"`
		CommentID       int64  `json:"comment_id"`
		ReplyToUserName string `json:"replyToUserName,omitempty"`
	} `json:"comments"`
	Counts struct {
		FanCount        int `json:"fanCount"`
		FollowCount     int `json:"followCount"`
		CollectionCount int `json:"collectionCount"`
		PhotoCount      int `json:"photoCount"`
	} `json:"counts"`
	Photo struct {
		SinglePicture bool   `json:"singlePicture"`
		Type          int    `json:"type"`
		Duration      int    `json:"duration"`
		Timestamp     int64  `json:"timestamp"`
		HeadURL       string `json:"headUrl"`
		KwaiID        string `json:"kwaiId"`
		UserName      string `json:"userName"`
		Verified      bool   `json:"verified"`
		UserID        int    `json:"userId"`
		WebpCoverUrls []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"webpCoverUrls"`
		ForwardCount int `json:"forwardCount"`
		Music        struct {
			Duration           int     `json:"duration"`
			Image              string  `json:"image"`
			VocalStartTime     int     `json:"vocalStartTime"`
			Name               string  `json:"name"`
			ID                 int64   `json:"id"`
			Type               int     `json:"type"`
			UsageCount         int     `json:"usageCount"`
			URL                string  `json:"url"`
			PhotoCount         int     `json:"photoCount"`
			AudioType          int     `json:"audioType"`
			Loudness           float64 `json:"loudness"`
			CopyrightTimeLimit int     `json:"copyrightTimeLimit"`
			AudioUrls          []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"audioUrls"`
			AvatarURL      string `json:"avatarUrl"`
			GenreID        int    `json:"genreId"`
			AnalysisResult struct {
				Beats []struct {
					Algorithm string `json:"algorithm"`
					BeatsUrls []struct {
						Cdn string `json:"cdn"`
						URL string `json:"url"`
					} `json:"beatsUrls"`
				} `json:"beats"`
			} `json:"analysisResult"`
			Artist    string `json:"artist"`
			ImageUrls []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"imageUrls"`
			AvatarUrls []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"avatarUrls"`
			IsOffline bool `json:"isOffline"`
		} `json:"music"`
		TagShow struct {
			BizID      string `json:"bizId"`
			BannerType int    `json:"bannerType"`
			UsedCount  string `json:"usedCount"`
			Type       int    `json:"type"`
		} `json:"tagShow"`
		CoverUrls []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"coverUrls"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		ViewCount    int    `json:"viewCount"`
		LikeCount    int    `json:"likeCount"`
		CommentCount int    `json:"commentCount"`
		UserSex      string `json:"userSex"`
		HeadUrls     []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"headUrls"`
		MainMvUrls []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"mainMvUrls"`
		Caption   string `json:"caption"`
		SameFrame struct {
			Allow          bool `json:"allow"`
			AvailableDepth int  `json:"availableDepth"`
		} `json:"sameFrame"`
		UsD       int    `json:"us_d"`
		ShareInfo string `json:"share_info"`
		UserEid   string `json:"userEid"`
		PhotoID   string `json:"photoId"`
		ExtParams struct {
			Mtype    int    `json:"mtype"`
			Color    string `json:"color"`
			W        int    `json:"w"`
			Sound    int    `json:"sound"`
			H        int    `json:"h"`
			Interval int    `json:"interval"`
			Video    int    `json:"video"`
		} `json:"ext_params"`
		ExpTag       string        `json:"exp_tag"`
		ServerExpTag string        `json:"serverExpTag"`
		AdminTags    []interface{} `json:"adminTags"`
	} `json:"photo"`
	HostName string `json:"host-name"`
	Mp4URL   string `json:"mp4Url"`
	Photos   []struct {
		SinglePicture bool   `json:"singlePicture"`
		Type          int    `json:"type"`
		Duration      int    `json:"duration"`
		Timestamp     int64  `json:"timestamp"`
		HeadURL       string `json:"headUrl"`
		KwaiID        string `json:"kwaiId"`
		UserName      string `json:"userName"`
		Verified      bool   `json:"verified"`
		UserID        int    `json:"userId"`
		WebpCoverUrls []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"webpCoverUrls"`
		ForwardCount int `json:"forwardCount"`
		SoundTrack   struct {
			Name    string `json:"name"`
			PhotoID int64  `json:"photoId"`
			User    struct {
				KwaiID            string `json:"kwaiId"`
				Following         bool   `json:"following"`
				Headurl           string `json:"headurl"`
				UserName          string `json:"user_name"`
				UserID            int    `json:"user_id"`
				Eid               string `json:"eid"`
				VisitorBeFollowed bool   `json:"visitorBeFollowed"`
				UserSex           string `json:"user_sex"`
				Headurls          []struct {
					Cdn string `json:"cdn"`
					URL string `json:"url"`
				} `json:"headurls"`
			} `json:"user"`
			AudioUrls []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"audioUrls"`
			Artist    string `json:"artist"`
			ImageUrls []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"imageUrls"`
			AvatarUrls []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"avatarUrls"`
			ID                   int64   `json:"id"`
			Type                 int     `json:"type"`
			AudioType            int     `json:"audioType"`
			Loudness             float64 `json:"loudness"`
			GenreID              int     `json:"genreId"`
			IsOffline            bool    `json:"isOffline"`
			DisableEnhancedEntry bool    `json:"disableEnhancedEntry"`
		} `json:"soundTrack,omitempty"`
		TagShow struct {
			BizID      string `json:"bizId"`
			BannerType int    `json:"bannerType"`
			UsedCount  string `json:"usedCount"`
			Type       int    `json:"type"`
		} `json:"tagShow"`
		CoverUrls []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"coverUrls"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		ViewCount    int    `json:"viewCount"`
		LikeCount    int    `json:"likeCount"`
		CommentCount int    `json:"commentCount"`
		UserSex      string `json:"userSex"`
		Poi          struct {
			ID        int     `json:"id"`
			Address   string  `json:"address"`
			Title     string  `json:"title"`
			City      string  `json:"city"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Category  int     `json:"category"`
		} `json:"poi,omitempty"`
		HeadUrls []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"headUrls"`
		MainMvUrls []struct {
			Cdn string `json:"cdn"`
			URL string `json:"url"`
		} `json:"mainMvUrls"`
		Caption   string `json:"caption"`
		SameFrame struct {
			Allow          bool `json:"allow"`
			AvailableDepth int  `json:"availableDepth"`
		} `json:"sameFrame"`
		UsD       int    `json:"us_d"`
		ShareInfo string `json:"share_info"`
		UserEid   string `json:"userEid"`
		PhotoID   string `json:"photoId"`
		ExtParams struct {
			Mtype    int    `json:"mtype"`
			Color    string `json:"color"`
			W        int    `json:"w"`
			Sound    int    `json:"sound"`
			H        int    `json:"h"`
			Interval int    `json:"interval"`
			Video    int    `json:"video"`
		} `json:"ext_params"`
		ExpTag       string        `json:"exp_tag"`
		ServerExpTag string        `json:"serverExpTag"`
		AdminTags    []interface{} `json:"adminTags"`
		Music        struct {
			Duration           int     `json:"duration"`
			Image              string  `json:"image"`
			VocalStartTime     int     `json:"vocalStartTime"`
			Name               string  `json:"name"`
			ID                 int64   `json:"id"`
			Type               int     `json:"type"`
			URL                string  `json:"url"`
			AudioType          int     `json:"audioType"`
			Loudness           float64 `json:"loudness"`
			CopyrightTimeLimit int     `json:"copyrightTimeLimit"`
			AudioUrls          []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"audioUrls"`
			AvatarURL string `json:"avatarUrl"`
			GenreID   int    `json:"genreId"`
			Artist    string `json:"artist"`
			ImageUrls []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"imageUrls"`
			AvatarUrls []struct {
				Cdn string `json:"cdn"`
				URL string `json:"url"`
			} `json:"avatarUrls"`
			IsOffline bool `json:"isOffline"`
		} `json:"music,omitempty"`
	} `json:"photos"`
	SerialInfo struct {
		Valid      bool        `json:"valid"`
		Title      interface{} `json:"title"`
		Msg        interface{} `json:"msg"`
		SerialID   interface{} `json:"serialId"`
		SerialType interface{} `json:"serialType"`
		Show       bool        `json:"show"`
	} `json:"serialInfo"`
}

func (a KuaiShouAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
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
	photoId := regexp.MustCompile(`photoId=(.*?)\&`).FindStringSubmatch(loc)
	postData := fmt.Sprintf(`{"photoId": "%s","isLongVideo": false}`, photoId[1])
	getResp, err := restyHttp.GetHttpClient().R().
		SetHeader("cookie", config.GetCookieInstance().KuaishouCookie).
		SetHeader("Content-Type", "application/json").
		SetHeader("Referer", loc).
		SetBody(postData).
		Post("https://v.m.chenzhongtech.com/rest/wd/photo/info")
	if err != nil {
		return nil, err
	}
	var kuaishouResponse KuaishouAdapterResponse
	err = json.Cjson.Unmarshal(getResp.Body(), &kuaishouResponse)
	if err != nil {
		return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             kuaishouResponse.Photo.UserName,
		AuthorAvatar:           kuaishouResponse.Photo.HeadURL,
		Title:                  kuaishouResponse.Photo.Caption,
		Cover:                  kuaishouResponse.Photo.CoverUrls[0].URL,
		NoWatermarkDownloadUrl: kuaishouResponse.Mp4URL,
	}
	return shortVideoInfo, nil
}
