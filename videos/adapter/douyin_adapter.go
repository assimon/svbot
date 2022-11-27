package adapter

import (
	"errors"
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
	"strings"
)

type DouyinAdapter struct{}

type DouyinAdapterResponse struct {
	StatusCode int `json:"status_code"`
	ItemList   []struct {
		AwemeID    string `json:"aweme_id"`
		Desc       string `json:"desc"`
		CreateTime int    `json:"create_time"`
		Author     struct {
			UID          string `json:"uid"`
			ShortID      string `json:"short_id"`
			Nickname     string `json:"nickname"`
			Signature    string `json:"signature"`
			AvatarLarger struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_larger"`
			AvatarThumb struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			AvatarMedium struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_medium"`
			FollowStatus     int         `json:"follow_status"`
			UniqueID         string      `json:"unique_id"`
			FollowersDetail  interface{} `json:"followers_detail"`
			PlatformSyncInfo interface{} `json:"platform_sync_info"`
			Geofencing       interface{} `json:"geofencing"`
			PolicyVersion    interface{} `json:"policy_version"`
			TypeLabel        interface{} `json:"type_label"`
			CardEntries      interface{} `json:"card_entries"`
			MixInfo          interface{} `json:"mix_info"`
		} `json:"author"`
		Music struct {
			ID      int64  `json:"id"`
			Mid     string `json:"mid"`
			Title   string `json:"title"`
			Author  string `json:"author"`
			CoverHd struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_hd"`
			CoverLarge struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_large"`
			CoverMedium struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_medium"`
			CoverThumb struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_thumb"`
			PlayURL struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_url"`
			Duration int         `json:"duration"`
			Position interface{} `json:"position"`
			Status   int         `json:"status"`
		} `json:"music"`
		ChaList []struct {
			Cid          string      `json:"cid"`
			ChaName      string      `json:"cha_name"`
			Desc         string      `json:"desc"`
			UserCount    int         `json:"user_count"`
			ConnectMusic interface{} `json:"connect_music"`
			Type         int         `json:"type"`
			CoverItem    struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_item"`
			ViewCount      int    `json:"view_count"`
			HashTagProfile string `json:"hash_tag_profile"`
			IsCommerce     bool   `json:"is_commerce"`
		} `json:"cha_list"`
		Video struct {
			PlayAddr struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_addr"`
			Cover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover"`
			Height       int `json:"height"`
			Width        int `json:"width"`
			DynamicCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"dynamic_cover"`
			OriginCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"origin_cover"`
			Ratio        string      `json:"ratio"`
			HasWatermark bool        `json:"has_watermark"`
			BitRate      interface{} `json:"bit_rate"`
			Duration     int         `json:"duration"`
			Vid          string      `json:"vid"`
		} `json:"video"`
		ShareURL   string `json:"share_url"`
		Statistics struct {
			AwemeID      string `json:"aweme_id"`
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
		} `json:"statistics"`
		TextExtra []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
			HashtagID   int64  `json:"hashtag_id"`
		} `json:"text_extra"`
		ShareInfo struct {
			ShareWeiboDesc string `json:"share_weibo_desc"`
			ShareDesc      string `json:"share_desc"`
			ShareTitle     string `json:"share_title"`
		} `json:"share_info"`
		VideoLabels interface{} `json:"video_labels"`
		Duration    int         `json:"duration"`
		AwemeType   int         `json:"aweme_type"`
		ImageInfos  interface{} `json:"image_infos"`
		RiskInfos   struct {
			Warn             bool   `json:"warn"`
			Type             int    `json:"type"`
			Content          string `json:"content"`
			ReflowUnplayable int    `json:"reflow_unplayable"`
		} `json:"risk_infos"`
		CommentList  interface{} `json:"comment_list"`
		AuthorUserID int64       `json:"author_user_id"`
		Geofencing   interface{} `json:"geofencing"`
		VideoText    interface{} `json:"video_text"`
		LabelTopText interface{} `json:"label_top_text"`
		Promotions   interface{} `json:"promotions"`
		LongVideo    interface{} `json:"long_video"`
		IsPreview    int         `json:"is_preview"`
		GroupID      int64       `json:"group_id"`
		IsLiveReplay bool        `json:"is_live_replay"`
		ForwardID    string      `json:"forward_id"`
		Images       interface{} `json:"images"`
		GroupIDStr   string      `json:"group_id_str"`
	} `json:"item_list"`
	FilterList []interface{} `json:"filter_list"`
	Extra      struct {
		Logid string `json:"logid"`
		Now   int64  `json:"now"`
	} `json:"extra"`
}

func (a DouyinAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
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
	re := regexp.MustCompile(`/video\/(.*)\?`)
	match := re.FindStringSubmatch(loc)
	if len(match) != 2 {
		return nil, errors.New("匹配视频id失败")
	}
	getResp, err := restyHttp.GetMobileHttpRequest().Get("https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=" + match[1])
	if err != nil {
		return nil, err
	}
	var douyinResponse DouyinAdapterResponse
	err = json.Cjson.Unmarshal(getResp.Body(), &douyinResponse)
	if err != nil {
		return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             douyinResponse.ItemList[0].Author.Nickname,
		AuthorAvatar:           douyinResponse.ItemList[0].Author.AvatarLarger.URLList[0],
		Title:                  douyinResponse.ItemList[0].ShareInfo.ShareTitle,
		Cover:                  douyinResponse.ItemList[0].Video.OriginCover.URLList[0],
		CreatedAt:              "",
		NoWatermarkDownloadUrl: strings.Replace(douyinResponse.ItemList[0].Video.PlayAddr.URLList[0], "playwm", "play", -1),
	}
	return shortVideoInfo, nil
}
