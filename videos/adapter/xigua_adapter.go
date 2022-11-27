package adapter

import (
	"encoding/base64"
	"github.com/assimon/svbot/internal/config"
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
	"strings"
)

type XiGuaAdapter struct{}

type XiGuaAdapterResponse struct {
	RecommendFeed   interface{}   `json:"recommendFeed"`
	AttentionFeed   interface{}   `json:"attentionFeed"`
	NbaFeed         interface{}   `json:"nbaFeed"`
	LivingFeed      interface{}   `json:"livingFeed"`
	ChannelFeed     []interface{} `json:"channelFeed"`
	HomeFeed        interface{}   `json:"homeFeed"`
	AdBanner        []interface{} `json:"adBanner"`
	ChannelInfo     interface{}   `json:"channelInfo"`
	ChannelFeedList []interface{} `json:"ChannelFeedList"`
	UserDetail      struct {
		EnableTabs    []interface{} `json:"enableTabs"`
		HotPersonList []interface{} `json:"hotPersonList"`
		UserInfo      struct {
			Name           string `json:"name"`
			Description    string `json:"description"`
			Avatar         string `json:"avatar"`
			FollowersCount int    `json:"followersCount"`
			FollowingCount int    `json:"followingCount"`
			UserID         string `json:"user_id"`
			Follow         bool   `json:"follow"`
		} `json:"userInfo"`
		VideoData struct {
			VideoList []interface{} `json:"videoList"`
			Loading   bool          `json:"loading"`
		} `json:"videoData"`
		HotsoonData struct {
			HotsoonList []interface{} `json:"hotsoonList"`
		} `json:"hotsoonData"`
		PreviewSeries []interface{} `json:"preview_series"`
		SeriesData    struct {
			SeriesList []interface{} `json:"series_list"`
			HasMore    bool          `json:"hasMore"`
			NextCursor string        `json:"nextCursor"`
		} `json:"seriesData"`
	} `json:"UserDetail"`
	FooterLinks           []interface{} `json:"FooterLinks"`
	LvideoChannel         []interface{} `json:"LvideoChannel"`
	LvideoChannelOnTcc    []interface{} `json:"LvideoChannelOnTcc"`
	LvideoCategory        []interface{} `json:"LvideoCategory"`
	AlbumInCategory       []interface{} `json:"AlbumInCategory"`
	ChannelFeedV2         []interface{} `json:"ChannelFeedV2"`
	ChannelLevelOneConfig []interface{} `json:"ChannelLevelOneConfig"`
	ChannelLevelTwoConfig []interface{} `json:"ChannelLevelTwoConfig"`
	HighQualityFeed       []interface{} `json:"HighQualityFeed"`
	ChannelBannerConfig   []interface{} `json:"ChannelBannerConfig"`
	Teleplay              interface{}   `json:"Teleplay"`
	Projection            struct {
		Video struct {
		} `json:"video"`
		Series struct {
		} `json:"series"`
		PSeries struct {
		} `json:"pSeries"`
		Playlist struct {
			ItemNum int `json:"item_num"`
		} `json:"playlist"`
		ShouldReturn404 bool        `json:"shouldReturn404"`
		ItemID          string      `json:"item_id"`
		Key             interface{} `json:"key"`
	} `json:"Projection"`
	CinemaChannelFeed   []interface{} `json:"CinemaChannelFeed"`
	CinemaFeedRebojiemu []interface{} `json:"CinemaFeedRebojiemu"`
	CinemaFeedFromRedis []interface{} `json:"CinemaFeedFromRedis"`
	MyWatchHistory      []struct {
		Type      string        `json:"type"`
		VideoFeed []interface{} `json:"videoFeed"`
		HasMore   bool          `json:"hasMore"`
	} `json:"MyWatchHistory"`
	MyFavorite []struct {
		Type      string        `json:"type"`
		VideoFeed []interface{} `json:"videoFeed"`
		HasMore   bool          `json:"hasMore"`
	} `json:"MyFavorite"`
	AuthorDetailInfo struct {
	} `json:"AuthorDetailInfo"`
	AuthorEnableTabs    []interface{} `json:"AuthorEnableTabs"`
	AuthorHomeVideoList struct {
		Inited    bool          `json:"inited"`
		LiveVideo interface{}   `json:"liveVideo"`
		Videos    []interface{} `json:"videos"`
	} `json:"AuthorHomeVideoList"`
	AuthorVideoList struct {
		VideoList []interface{} `json:"videoList"`
		HasMore   bool          `json:"hasMore"`
		MaxTime   int           `json:"maxTime"`
		Offset    int           `json:"offset"`
	} `json:"AuthorVideoList"`
	AuthorVideoSearchList struct {
		VideoList []interface{} `json:"videoList"`
		HasMore   bool          `json:"hasMore"`
		MaxTime   int           `json:"maxTime"`
		Offset    int           `json:"offset"`
	} `json:"AuthorVideoSearchList"`
	AuthorLVideoList struct {
		VideoList []interface{} `json:"videoList"`
		HasMore   bool          `json:"hasMore"`
		Offset    int           `json:"offset"`
	} `json:"AuthorLVideoList"`
	AuthorHotsoonList struct {
		VideoList []interface{} `json:"videoList"`
		HasMore   bool          `json:"hasMore"`
		MaxTime   int           `json:"maxTime"`
		Offset    int           `json:"offset"`
	} `json:"AuthorHotsoonList"`
	AuthorPreviewSeries struct {
		List    []interface{} `json:"list"`
		HasMore bool          `json:"hasMore"`
		Offset  int           `json:"offset"`
	} `json:"AuthorPreviewSeries"`
	AuthorSeriesDataList struct {
		List    []interface{} `json:"list"`
		HasMore bool          `json:"hasMore"`
		Offset  int           `json:"offset"`
	} `json:"AuthorSeriesDataList"`
	AuthorSeriesData struct {
		SeriesList []interface{} `json:"series_list"`
		HasMore    bool          `json:"hasMore"`
		NextCursor string        `json:"nextCursor"`
	} `json:"AuthorSeriesData"`
	AuthorFansData struct {
		UserList []interface{} `json:"userList"`
		HasMore  bool          `json:"hasMore"`
		Cursor   int           `json:"cursor"`
	} `json:"AuthorFansData"`
	AuthorFollowData struct {
		UserList []interface{} `json:"userList"`
		HasMore  bool          `json:"hasMore"`
		Cursor   int           `json:"cursor"`
	} `json:"AuthorFollowData"`
	AuthorFollowAscData struct {
		UserList []interface{} `json:"userList"`
		HasMore  bool          `json:"hasMore"`
		Cursor   int           `json:"cursor"`
	} `json:"AuthorFollowAscData"`
	AuthorFollowSearchData struct {
		UserList []interface{} `json:"userList"`
		HasMore  bool          `json:"hasMore"`
		Cursor   int           `json:"cursor"`
	} `json:"AuthorFollowSearchData"`
	AuthorTabsCount struct {
	} `json:"AuthorTabsCount"`
	AuthorTabsShow struct {
	} `json:"AuthorTabsShow"`
	AuthorPlayedVideoList struct {
		VideoList []interface{} `json:"videoList"`
		HasMore   bool          `json:"hasMore"`
		Offset    int           `json:"offset"`
	} `json:"AuthorPlayedVideoList"`
	AuthorPlayedHotsoonList struct {
		VideoList []interface{} `json:"videoList"`
		HasMore   bool          `json:"hasMore"`
		Offset    int           `json:"offset"`
	} `json:"AuthorPlayedHotsoonList"`
	ClusterFeeds struct {
		List    []interface{} `json:"list"`
		HasMore bool          `json:"hasMore"`
	} `json:"ClusterFeeds"`
	HotWords      []interface{} `json:"HotWords"`
	Home          interface{}   `json:"Home"`
	Channel       interface{}   `json:"Channel"`
	SEOLinks      []interface{} `json:"SEOLinks"`
	SEOUpdateTime string        `json:"SEOUpdateTime"`
	Layout        struct {
		Default struct {
			Side        string `json:"side"`
			HeaderTheme string `json:"headerTheme"`
			ClassName   string `json:"className"`
		} `json:"default"`
	} `json:"layout"`
	Exception404 struct {
		Value bool `json:"value"`
	} `json:"exception404"`
	SearchPresetWord interface{} `json:"searchPresetWord"`
	AnyVideo         struct {
		GidInformation struct {
			Gid         string `json:"gid"`
			GroupSource int    `json:"groupSource"`
			PackerData  struct {
				Video struct {
					PseriesID       interface{} `json:"pseries_id"`
					VideoDuration   int         `json:"video_duration"`
					VideoWatchCount int         `json:"video_watch_count"`
					UserInfo        struct {
						AvatarURL       string `json:"avatar_url"`
						Follow          bool   `json:"follow"`
						FollowerCount   int    `json:"follower_count"`
						Name            string `json:"name"`
						UserID          int64  `json:"user_id"`
						VideoTotalCount int    `json:"video_total_count"`
						UserAuthInfo    string `json:"user_auth_info"`
					} `json:"user_info"`
					VideoUserLike    int    `json:"video_user_like"`
					VideoPublishTime string `json:"video_publish_time"`
					VideoLikeCount   int    `json:"video_like_count"`
					VideoAbstract    string `json:"video_abstract"`
					Vid              string `json:"vid"`
					UserDigg         int    `json:"user_digg"`
					UserSuperDigg    int    `json:"user_super_digg"`
					SuperDiggControl string `json:"super_digg_control"`
					Title            string `json:"title"`
					Tag              string `json:"tag"`
					GroupID          string `json:"group_id"`
					ItemID           string `json:"item_id"`
					IsOriginal       bool   `json:"is_original"`
					LogPb            struct {
						ImprID      string `json:"impr_id"`
						IsFollowing string `json:"is_following"`
						GroupID     string `json:"group_id"`
						GroupSource string `json:"group_source"`
						AuthorID    string `json:"author_id"`
					} `json:"log_pb"`
					GroupSource         int         `json:"group_source"`
					BanAction           interface{} `json:"ban_action"`
					PosterURL           string      `json:"poster_url"`
					BanDanmaku          int         `json:"ban_danmaku"`
					BanDanmakuSend      int         `json:"ban_danmaku_send"`
					DanmakuCount        int         `json:"danmaku_count"`
					Rank                int         `json:"rank"`
					IsContractArguement bool        `json:"is_contract_arguement"`
					VideoLogo           struct {
						LogoType  int `json:"logo_type"`
						LogoImage struct {
							URL     string `json:"url"`
							Width   int    `json:"width"`
							URLList []struct {
								URL string `json:"url"`
							} `json:"url_list"`
							URI    string `json:"uri"`
							Height int    `json:"height"`
						} `json:"logo_image"`
						LogoText     string `json:"logo_text"`
						LogoPosition int    `json:"logo_position"`
					} `json:"video_logo"`
					DanmakuMask     int         `json:"danmaku_mask"`
					HistoryDuration int         `json:"history_duration"`
					RiskWarning     interface{} `json:"risk_warning"`
					IsBanRobot      bool        `json:"IsBanRobot"`
					StickerList     interface{} `json:"sticker_list"`
					ShareIFrameInfo struct {
						InAllowList bool `json:"InAllowList"`
						InBanList   bool `json:"InBanList"`
						BanByVL     bool `json:"BanByVL"`
					} `json:"ShareIFrameInfo"`
					IsExclusive     bool `json:"IsExclusive"`
					IsCopyright     bool `json:"IsCopyright"`
					XgVideoRichText struct {
					} `json:"xg_video_rich_text"`
					Address       string `json:"address"`
					VideoResource struct {
						Vid    string      `json:"vid"`
						Dash   interface{} `json:"dash"`
						Normal struct {
							Status         int     `json:"status"`
							Message        string  `json:"message"`
							EnableSsl      bool    `json:"enable_ssl"`
							AutoDefinition string  `json:"auto_definition"`
							EnableAdaptive bool    `json:"enable_adaptive"`
							VideoID        string  `json:"video_id"`
							VideoDuration  float64 `json:"video_duration"`
							MediaType      string  `json:"media_type"`
							BigThumbs      []struct {
								ImgNum   int      `json:"img_num"`
								URI      string   `json:"uri"`
								ImgURL   string   `json:"img_url"`
								ImgUrls  []string `json:"img_urls"`
								ImgXSize int      `json:"img_x_size"`
								ImgYSize int      `json:"img_y_size"`
								ImgXLen  int      `json:"img_x_len"`
								ImgYLen  int      `json:"img_y_len"`
								Duration float64  `json:"duration"`
								Interval int      `json:"interval"`
								Fext     string   `json:"fext"`
							} `json:"big_thumbs"`
							VideoList struct {
								Video1 struct {
									Definition  string `json:"definition"`
									Quality     string `json:"quality"`
									Vtype       string `json:"vtype"`
									Vwidth      int    `json:"vwidth"`
									Vheight     int    `json:"vheight"`
									Bitrate     int    `json:"bitrate"`
									RealBitrate int    `json:"real_bitrate"`
									Fps         int    `json:"fps"`
									CodecType   string `json:"codec_type"`
									Size        int    `json:"size"`
									MainURL     string `json:"main_url"`
									BackupURL1  string `json:"backup_url_1"`
									FileID      string `json:"file_id"`
									Volume      struct {
										Loudness float64 `json:"loudness"`
										Peak     float64 `json:"peak"`
									} `json:"volume"`
									QualityType      int    `json:"quality_type"`
									EncryptionMethod string `json:"encryption_method"`
									URLExpire        int    `json:"url_expire"`
									PreloadSize      int    `json:"preload_size"`
									PreloadInterval  int    `json:"preload_interval"`
									PreloadMinStep   int    `json:"preload_min_step"`
									PreloadMaxStep   int    `json:"preload_max_step"`
									FileHash         string `json:"file_hash"`
									P2PVerifyURL     string `json:"p2p_verify_url"`
								} `json:"video_1"`
								Video2 struct {
									Definition  string `json:"definition"`
									Quality     string `json:"quality"`
									Vtype       string `json:"vtype"`
									Vwidth      int    `json:"vwidth"`
									Vheight     int    `json:"vheight"`
									Bitrate     int    `json:"bitrate"`
									RealBitrate int    `json:"real_bitrate"`
									Fps         int    `json:"fps"`
									CodecType   string `json:"codec_type"`
									Size        int    `json:"size"`
									MainURL     string `json:"main_url"`
									BackupURL1  string `json:"backup_url_1"`
									FileID      string `json:"file_id"`
									Volume      struct {
										Loudness float64 `json:"loudness"`
										Peak     float64 `json:"peak"`
									} `json:"volume"`
									QualityType      int    `json:"quality_type"`
									EncryptionMethod string `json:"encryption_method"`
									URLExpire        int    `json:"url_expire"`
									PreloadSize      int    `json:"preload_size"`
									PreloadInterval  int    `json:"preload_interval"`
									PreloadMinStep   int    `json:"preload_min_step"`
									PreloadMaxStep   int    `json:"preload_max_step"`
									FileHash         string `json:"file_hash"`
									P2PVerifyURL     string `json:"p2p_verify_url"`
								} `json:"video_2"`
								Video3 struct {
									Definition  string `json:"definition"`
									Quality     string `json:"quality"`
									Vtype       string `json:"vtype"`
									Vwidth      int    `json:"vwidth"`
									Vheight     int    `json:"vheight"`
									Bitrate     int    `json:"bitrate"`
									RealBitrate int    `json:"real_bitrate"`
									Fps         int    `json:"fps"`
									CodecType   string `json:"codec_type"`
									Size        int    `json:"size"`
									MainURL     string `json:"main_url"`
									BackupURL1  string `json:"backup_url_1"`
									FileID      string `json:"file_id"`
									Volume      struct {
										Loudness float64 `json:"loudness"`
										Peak     float64 `json:"peak"`
									} `json:"volume"`
									QualityType      int    `json:"quality_type"`
									EncryptionMethod string `json:"encryption_method"`
									URLExpire        int    `json:"url_expire"`
									PreloadSize      int    `json:"preload_size"`
									PreloadInterval  int    `json:"preload_interval"`
									PreloadMinStep   int    `json:"preload_min_step"`
									PreloadMaxStep   int    `json:"preload_max_step"`
									FileHash         string `json:"file_hash"`
									P2PVerifyURL     string `json:"p2p_verify_url"`
								} `json:"video_3"`
								Video4 struct {
									Definition  string `json:"definition"`
									Quality     string `json:"quality"`
									Vtype       string `json:"vtype"`
									Vwidth      int    `json:"vwidth"`
									Vheight     int    `json:"vheight"`
									Bitrate     int    `json:"bitrate"`
									RealBitrate int    `json:"real_bitrate"`
									Fps         int    `json:"fps"`
									CodecType   string `json:"codec_type"`
									Size        int    `json:"size"`
									MainURL     string `json:"main_url"`
									BackupURL1  string `json:"backup_url_1"`
									FileID      string `json:"file_id"`
									Volume      struct {
										Loudness float64 `json:"loudness"`
										Peak     float64 `json:"peak"`
									} `json:"volume"`
									QualityType      int    `json:"quality_type"`
									EncryptionMethod string `json:"encryption_method"`
									URLExpire        int    `json:"url_expire"`
									PreloadSize      int    `json:"preload_size"`
									PreloadInterval  int    `json:"preload_interval"`
									PreloadMinStep   int    `json:"preload_min_step"`
									PreloadMaxStep   int    `json:"preload_max_step"`
									FileHash         string `json:"file_hash"`
									P2PVerifyURL     string `json:"p2p_verify_url"`
								} `json:"video_4"`
							} `json:"video_list"`
							BarrageMaskURL string `json:"barrage_mask_url"`
							Volume         struct {
								Loudness float64 `json:"loudness"`
								Peak     int     `json:"peak"`
							} `json:"volume"`
							PopularityLevel     int  `json:"popularity_level"`
							HasEmbeddedSubtitle bool `json:"has_embedded_subtitle"`
							BarrageMaskInfo     struct {
								Version        string `json:"version"`
								BarrageMaskURL string `json:"barrage_mask_url"`
								FileID         string `json:"file_id"`
								FileHash       string `json:"file_hash"`
								FileSize       int    `json:"file_size"`
								UpdatedAt      int    `json:"updated_at"`
								Bitrate        int    `json:"bitrate"`
								HeadLen        int    `json:"head_len"`
							} `json:"barrage_mask_info"`
							PosterURL  string `json:"poster_url"`
							ExtraInfos struct {
								Status                 string      `json:"Status"`
								Message                string      `json:"Message"`
								LogoType               string      `json:"LogoType"`
								VideoModelVersion      int         `json:"VideoModelVersion"`
								HelpInfoURL            string      `json:"HelpInfoURL"`
								LengthOfVideoList      string      `json:"LengthOfVideoList"`
								IsDynamicVideo         bool        `json:"IsDynamicVideo"`
								UserAction             string      `json:"UserAction"`
								AccountName            string      `json:"AccountName"`
								DeniedVideoModelV1JSON string      `json:"DeniedVideoModelV1JSON"`
								ResTag                 string      `json:"ResTag"`
								EncodeUserTag          string      `json:"EncodeUserTag"`
								VideoMeta              interface{} `json:"VideoMeta"`
								RemovedVideoMeta       interface{} `json:"RemovedVideoMeta"`
							} `json:"extraInfos"`
							RefreshToken  string `json:"refreshToken"`
							InterfaceInfo struct {
								Code       int    `json:"code"`
								Message    string `json:"message"`
								Logid      string `json:"logid"`
								APIStr     string `json:"api_str"`
								Timestamep int64  `json:"timestamep"`
							} `json:"interfaceInfo"`
						} `json:"normal"`
						Dash120Fps interface{} `json:"dash_120fps"`
						Normal6Min interface{} `json:"normal_6min"`
					} `json:"videoResource"`
					VideoReprintedInfo interface{} `json:"video_reprinted_info"`
				} `json:"video"`
				Key string `json:"key"`
			} `json:"packerData"`
		} `json:"gidInformation"`
		QueryKey            string        `json:"queryKey"`
		SpiderVideo         []interface{} `json:"spiderVideo"`
		CommentForSSR       interface{}   `json:"commentForSSR"`
		SubtitleForScrawler interface{}   `json:"subtitleForScrawler"`
		KvideoData          interface{}   `json:"kvideoData"`
	} `json:"anyVideo"`
	SearchBanner  interface{}   `json:"searchBanner"`
	ComplexSearch interface{}   `json:"complexSearch"`
	LiveSearch    interface{}   `json:"liveSearch"`
	LvideoSearch  interface{}   `json:"lvideoSearch"`
	UserSearch    interface{}   `json:"userSearch"`
	SearchHotList []interface{} `json:"SearchHotList"`
}

func (a XiGuaAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Sugar.Error(r)
		}
	}()
	if strings.Contains(url, "v.ixigua.com") {
		headerResp, err := restyHttp.GetHttpClient().R().Head(url)
		if err != nil {
			return nil, err
		}
		loc := headerResp.RawResponse.Request.Response.Header.Get("Location")
		id := regexp.MustCompile(`video\/(.*)\/`).FindStringSubmatch(loc)
		url = "https://www.ixigua.com/" + id[1]
	}
	getResp, err := restyHttp.GetHttpClient().R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36 ").
		SetHeader("cookie", config.GetCookieInstance().XiguaCookie).
		Get(url)
	if err != nil {
		return nil, err
	}
	jsonData := regexp.MustCompile(`<script id=\"SSR_HYDRATED_DATA\">window._SSR_HYDRATED_DATA=(.*?)<\/script>`).FindStringSubmatch(getResp.String())
	checkJson := strings.ReplaceAll(jsonData[1], "undefined", "null")
	var xiguaResponse XiGuaAdapterResponse
	err = json.Cjson.UnmarshalFromString(checkJson, &xiguaResponse)
	if err != nil {
		return nil, err
	}
	videoUrl, err := base64.StdEncoding.DecodeString(xiguaResponse.AnyVideo.GidInformation.PackerData.Video.VideoResource.Normal.VideoList.Video3.MainURL)
	if err != nil {
		return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             xiguaResponse.UserDetail.UserInfo.Name,
		AuthorAvatar:           xiguaResponse.UserDetail.UserInfo.Avatar,
		Title:                  xiguaResponse.AnyVideo.GidInformation.PackerData.Video.Title,
		Cover:                  xiguaResponse.AnyVideo.GidInformation.PackerData.Video.PosterURL,
		NoWatermarkDownloadUrl: string(videoUrl),
	}
	return shortVideoInfo, nil
}
