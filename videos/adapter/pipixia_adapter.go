package adapter

import (
	"errors"
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
)

type PipixiaAdapter struct{}

type PipixiaAdapterResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Prompt     string `json:"prompt"`
	Time       int64  `json:"time"`
	Data       struct {
		Data struct {
			CellType    int    `json:"cell_type"`
			CellID      int64  `json:"cell_id"`
			CellIDStr   string `json:"cell_id_str"`
			DisplayTime int    `json:"display_time"`
			Item        struct {
				ItemID     int64         `json:"item_id"`
				ItemIDStr  string        `json:"item_id_str"`
				ItemType   int           `json:"item_type"`
				CreateTime int           `json:"create_time"`
				Status     int           `json:"status"`
				Comments   []interface{} `json:"comments"`
				Content    string        `json:"content"`
				Cover      struct {
					Width   int    `json:"width"`
					Height  int    `json:"height"`
					URI     string `json:"uri"`
					URLList []struct {
						URL string `json:"url"`
					} `json:"url_list"`
					IsGif        bool `json:"is_gif"`
					DownloadList []struct {
						URL string `json:"url"`
					} `json:"download_list"`
				} `json:"cover"`
				Share struct {
					ShareURL        string `json:"share_url"`
					Title           string `json:"title"`
					Content         string `json:"content"`
					ImageURL        string `json:"image_url"`
					CompoundPageURL string `json:"compound_page_url"`
					Ios             struct {
						Channel         int `json:"channel"`
						WeixinStrategy  int `json:"weixin_strategy"`
						MomentsStrategy int `json:"moments_strategy"`
						QqStrategy      int `json:"qq_strategy"`
						QzoneStrategy   int `json:"qzone_strategy"`
					} `json:"ios"`
					Android struct {
						Channel         int `json:"channel"`
						WeixinStrategy  int `json:"weixin_strategy"`
						MomentsStrategy int `json:"moments_strategy"`
						QqStrategy      int `json:"qq_strategy"`
						QzoneStrategy   int `json:"qzone_strategy"`
					} `json:"android"`
					LinkText        string `json:"link_text"`
					ShareText       string `json:"share_text"`
					WeixinStrategy  int    `json:"weixin_strategy"`
					MomentsStrategy int    `json:"moments_strategy"`
					QqStrategy      int    `json:"qq_strategy"`
					QzoneStrategy   int    `json:"qzone_strategy"`
					WechatURL       string `json:"wechat_url"`
					MomentsURL      string `json:"moments_url"`
					QqURL           string `json:"qq_url"`
					QzoneURL        string `json:"qzone_url"`
					DouyinURL       string `json:"douyin_url"`
					Schema          string `json:"schema"`
					LargeImageURL   string `json:"large_image_url"`
				} `json:"share"`
				Author struct {
					ID     int64  `json:"id"`
					IDStr  string `json:"id_str"`
					Name   string `json:"name"`
					Avatar struct {
						Width   int    `json:"width"`
						Height  int    `json:"height"`
						URI     string `json:"uri"`
						URLList []struct {
							URL string `json:"url"`
						} `json:"url_list"`
						IsGif        bool `json:"is_gif"`
						DownloadList []struct {
							URL string `json:"url"`
						} `json:"download_list"`
					} `json:"avatar"`
					IsFollowing            interface{}   `json:"is_following"`
					FollowersCount         int           `json:"followers_count"`
					FollowingsCount        int           `json:"followings_count"`
					LikeCount              int           `json:"like_count"`
					Level                  int           `json:"level"`
					Description            string        `json:"description"`
					ProfileSchema          string        `json:"profile_schema"`
					IsFollowed             interface{}   `json:"is_followed"`
					GodCommentCount        int           `json:"god_comment_count"`
					VoteCount              int           `json:"vote_count"`
					CertifyInfo            interface{}   `json:"certify_info"`
					Achievements           []interface{} `json:"achievements"`
					Status                 int           `json:"status"`
					AuthorInfo             interface{}   `json:"author_info"`
					BroadcastInfo          interface{}   `json:"broadcast_info"`
					LiveAuth               bool          `json:"live_auth"`
					Region                 string        `json:"region"`
					Language               string        `json:"language"`
					Punishments            interface{}   `json:"punishments"`
					RecommendTag           interface{}   `json:"recommend_tag"`
					RecommendReason        interface{}   `json:"recommend_reason"`
					CommercePermissionList interface{}   `json:"commerce_permission_list"`
					DecorationList         interface{}   `json:"decoration_list"`
					InteractionLimitation  interface{}   `json:"interaction_limitation"`
					CreativeLevelInfo      interface{}   `json:"creative_level_info"`
					Gender                 int           `json:"gender"`
					Age                    string        `json:"age"`
					Horoscope              string        `json:"horoscope"`
					HideAge                interface{}   `json:"hide_age"`
				} `json:"author"`
				Stats struct {
					GoDetailCount   int `json:"go_detail_count"`
					ImpressionCount int `json:"impression_count"`
					CommentCount    int `json:"comment_count"`
					LikeCount       int `json:"like_count"`
					ShareCount      int `json:"share_count"`
					PlayCount       int `json:"play_count"`
					DubbingCount    int `json:"dubbing_count"`
					BuryCount       int `json:"bury_count"`
					BulletCount     int `json:"bullet_count"`
					DiggCounts      []struct {
						DiggType  int `json:"digg_type"`
						DiggCount int `json:"digg_count"`
					} `json:"digg_counts"`
					BuryCounts []struct {
						BuryType  int `json:"bury_type"`
						BuryCount int `json:"bury_count"`
					} `json:"bury_counts"`
					ViewCount int `json:"view_count"`
				} `json:"stats"`
				ItemRelation struct {
					IsLike     bool `json:"is_like"`
					IsFavorite bool `json:"is_favorite"`
					IsBury     bool `json:"is_bury"`
					DiggType   int  `json:"digg_type"`
					BuryType   int  `json:"bury_type"`
				} `json:"item_relation"`
				DefaultSchema string      `json:"default_schema"`
				VitalComments interface{} `json:"vital_comments"`
				Note          interface{} `json:"note"`
				Video         struct {
					VideoID       string        `json:"video_id"`
					Text          string        `json:"text"`
					HashtagSchema []interface{} `json:"hashtag_schema"`
					CoverImage    struct {
						Width   int    `json:"width"`
						Height  int    `json:"height"`
						URI     string `json:"uri"`
						URLList []struct {
							URL string `json:"url"`
						} `json:"url_list"`
						IsGif        bool `json:"is_gif"`
						DownloadList []struct {
							URL string `json:"url"`
						} `json:"download_list"`
					} `json:"cover_image"`
					Duration      float64     `json:"duration"`
					VideoWidth    int         `json:"video_width"`
					VideoHeight   int         `json:"video_height"`
					Title         interface{} `json:"title"`
					VideoDownload struct {
						URI     string `json:"uri"`
						Width   int    `json:"width"`
						Height  int    `json:"height"`
						URLList []struct {
							URL     string `json:"url"`
							Expires int    `json:"expires"`
						} `json:"url_list"`
						CoverImage struct {
							Width   int    `json:"width"`
							Height  int    `json:"height"`
							URI     string `json:"uri"`
							URLList []struct {
								URL string `json:"url"`
							} `json:"url_list"`
							IsGif        bool `json:"is_gif"`
							DownloadList []struct {
								URL string `json:"url"`
							} `json:"download_list"`
						} `json:"cover_image"`
						Duration           float64     `json:"duration"`
						AlarmText          interface{} `json:"alarm_text"`
						CodecType          int         `json:"codec_type"`
						Definition         int         `json:"definition"`
						VideoModel         string      `json:"video_model"`
						P2PType            interface{} `json:"p2p_type"`
						FileHash           interface{} `json:"file_hash"`
						AnimatedCoverImage interface{} `json:"animated_cover_image"`
					} `json:"video_download"`
					VideoLow  interface{} `json:"video_low"`
					VideoMid  interface{} `json:"video_mid"`
					VideoHigh struct {
						URI     string `json:"uri"`
						Width   int    `json:"width"`
						Height  int    `json:"height"`
						URLList []struct {
							URL     string `json:"url"`
							Expires int    `json:"expires"`
						} `json:"url_list"`
						CoverImage struct {
							Width   int    `json:"width"`
							Height  int    `json:"height"`
							URI     string `json:"uri"`
							URLList []struct {
								URL string `json:"url"`
							} `json:"url_list"`
							IsGif        bool `json:"is_gif"`
							DownloadList []struct {
								URL string `json:"url"`
							} `json:"download_list"`
						} `json:"cover_image"`
						Duration           float64     `json:"duration"`
						AlarmText          interface{} `json:"alarm_text"`
						CodecType          int         `json:"codec_type"`
						Definition         int         `json:"definition"`
						VideoModel         string      `json:"video_model"`
						P2PType            int         `json:"p2p_type"`
						FileHash           interface{} `json:"file_hash"`
						AnimatedCoverImage interface{} `json:"animated_cover_image"`
					} `json:"video_high"`
					VideoFallback struct {
						URI        string        `json:"uri"`
						Width      int           `json:"width"`
						Height     int           `json:"height"`
						URLList    []interface{} `json:"url_list"`
						CoverImage struct {
							Width   int    `json:"width"`
							Height  int    `json:"height"`
							URI     string `json:"uri"`
							URLList []struct {
								URL string `json:"url"`
							} `json:"url_list"`
							IsGif        bool        `json:"is_gif"`
							DownloadList interface{} `json:"download_list"`
						} `json:"cover_image"`
						Duration           float64     `json:"duration"`
						AlarmText          interface{} `json:"alarm_text"`
						CodecType          interface{} `json:"codec_type"`
						Definition         interface{} `json:"definition"`
						VideoModel         string      `json:"video_model"`
						P2PType            interface{} `json:"p2p_type"`
						FileHash           interface{} `json:"file_hash"`
						AnimatedCoverImage interface{} `json:"animated_cover_image"`
					} `json:"video_fallback"`
					VideoGodCommentUrls interface{} `json:"video_god_comment_urls"`
					Animate             interface{} `json:"animate"`
					Transport           bool        `json:"transport"`
					TransportText       string      `json:"transport_text"`
					SupportLivePhoto    bool        `json:"support_live_photo"`
					TailAdPassthrough   string      `json:"tail_ad_passthrough"`
				} `json:"video"`
				Link                interface{}   `json:"link"`
				TextSchema          []interface{} `json:"text_schema"`
				Level               interface{}   `json:"level"`
				Source              interface{}   `json:"source"`
				DebugInfo           interface{}   `json:"debug_info"`
				Stage               interface{}   `json:"stage"`
				UserStatus          interface{}   `json:"user_status"`
				PostSource          int           `json:"post_source"`
				ExpireTime          interface{}   `json:"expire_time"`
				ReviewStatus        interface{}   `json:"review_status"`
				FrozenToast         string        `json:"frozen_toast"`
				Duration            float64       `json:"duration"`
				VideoType           int           `json:"video_type"`
				RelatedID           interface{}   `json:"related_id"`
				AncestorID          string        `json:"ancestor_id"`
				OriginVideoID       string        `json:"origin_video_id"`
				RelatedCommentID    int           `json:"related_comment_id"`
				AncestorCommentID   interface{}   `json:"ancestor_comment_id"`
				OriginVideoDownload struct {
					URI     string `json:"uri"`
					Width   int    `json:"width"`
					Height  int    `json:"height"`
					URLList []struct {
						URL     string `json:"url"`
						Expires int    `json:"expires"`
					} `json:"url_list"`
					CoverImage struct {
						Width   int    `json:"width"`
						Height  int    `json:"height"`
						URI     string `json:"uri"`
						URLList []struct {
							URL string `json:"url"`
						} `json:"url_list"`
						IsGif        bool `json:"is_gif"`
						DownloadList []struct {
							URL string `json:"url"`
						} `json:"download_list"`
					} `json:"cover_image"`
					Duration           float64     `json:"duration"`
					AlarmText          interface{} `json:"alarm_text"`
					CodecType          int         `json:"codec_type"`
					Definition         int         `json:"definition"`
					VideoModel         string      `json:"video_model"`
					P2PType            interface{} `json:"p2p_type"`
					FileHash           interface{} `json:"file_hash"`
					AnimatedCoverImage interface{} `json:"animated_cover_image"`
				} `json:"origin_video_download"`
				CommentPostItem      bool          `json:"comment_post_item"`
				HotInfo              interface{}   `json:"hot_info"`
				AuthorDisplayTags    interface{}   `json:"author_display_tags"`
				Rating               int           `json:"rating"`
				AncestorSchema       interface{}   `json:"ancestor_schema"`
				JumpLink             interface{}   `json:"jump_link"`
				MicroAppID           interface{}   `json:"micro_app_id"`
				RecreateMetaInfoList []interface{} `json:"recreate_meta_info_list"`
				NeihanStyle          int           `json:"neihan_style"`
				MappingGids          interface{}   `json:"mapping_gids"`
				HumanTags            interface{}   `json:"human_tags"`
				AdminDebug           interface{}   `json:"admin_debug"`
				AiTag                interface{}   `json:"ai_tag"`
				WardInfo             struct {
					IsWard        bool `json:"is_ward"`
					WardCommentID int  `json:"ward_comment_id"`
					WardReplyID   int  `json:"ward_reply_id"`
					WardUsers     []struct {
						ID     int64  `json:"id"`
						IDStr  string `json:"id_str"`
						Name   string `json:"name"`
						Avatar struct {
							Width   int    `json:"width"`
							Height  int    `json:"height"`
							URI     string `json:"uri"`
							URLList []struct {
								URL string `json:"url"`
							} `json:"url_list"`
							IsGif        bool `json:"is_gif"`
							DownloadList []struct {
								URL string `json:"url"`
							} `json:"download_list"`
						} `json:"avatar"`
						IsFollowing            interface{} `json:"is_following"`
						FollowersCount         interface{} `json:"followers_count"`
						FollowingsCount        interface{} `json:"followings_count"`
						LikeCount              interface{} `json:"like_count"`
						Level                  int         `json:"level"`
						Description            string      `json:"description"`
						ProfileSchema          string      `json:"profile_schema"`
						IsFollowed             interface{} `json:"is_followed"`
						GodCommentCount        interface{} `json:"god_comment_count"`
						VoteCount              int         `json:"vote_count"`
						CertifyInfo            interface{} `json:"certify_info"`
						Achievements           interface{} `json:"achievements"`
						Status                 interface{} `json:"status"`
						AuthorInfo             interface{} `json:"author_info"`
						BroadcastInfo          interface{} `json:"broadcast_info"`
						LiveAuth               bool        `json:"live_auth"`
						Region                 string      `json:"region"`
						Language               string      `json:"language"`
						Punishments            interface{} `json:"punishments"`
						RecommendTag           interface{} `json:"recommend_tag"`
						RecommendReason        interface{} `json:"recommend_reason"`
						CommercePermissionList interface{} `json:"commerce_permission_list"`
						DecorationList         interface{} `json:"decoration_list"`
						InteractionLimitation  interface{} `json:"interaction_limitation"`
						CreativeLevelInfo      interface{} `json:"creative_level_info"`
						Gender                 interface{} `json:"gender"`
						Age                    interface{} `json:"age"`
						Horoscope              interface{} `json:"horoscope"`
						HideAge                interface{} `json:"hide_age"`
					} `json:"ward_users"`
					WardCount     int           `json:"ward_count"`
					HasNew        bool          `json:"has_new"`
					TrendMessages []interface{} `json:"trend_messages"`
					IsVisible     bool          `json:"is_visible"`
				} `json:"ward_info"`
				RelatedCommentIDStr string `json:"related_comment_id_str"`
				InteractEgg         struct {
					LikeStyle string `json:"like_style"`
				} `json:"interact_egg"`
				ClubInfo     interface{} `json:"club_info"`
				AlbumIntro   interface{} `json:"album_intro"`
				EpisodeInfo  interface{} `json:"episode_info"`
				DrainageInfo interface{} `json:"drainage_info"`
				ItemCellType int         `json:"item_cell_type"`
				CanDownload  bool        `json:"can_download"`
				ClubRelation interface{} `json:"club_relation"`
				GameCardInfo interface{} `json:"game_card_info"`
				Privilege    struct {
					BulletPlay bool `json:"bullet_play"`
					BulletPost bool `json:"bullet_post"`
				} `json:"privilege"`
				Activity        interface{} `json:"activity"`
				ShowFeatureDigg bool        `json:"show_feature_digg"`
				ShowFeatureBury bool        `json:"show_feature_bury"`
				PromotionInfo   interface{} `json:"promotion_info"`
				AhaImage        []struct {
					Width   int    `json:"width"`
					Height  int    `json:"height"`
					URI     string `json:"uri"`
					URLList []struct {
						URL string `json:"url"`
					} `json:"url_list"`
					IsGif        bool `json:"is_gif"`
					DownloadList []struct {
						URL string `json:"url"`
					} `json:"download_list"`
				} `json:"aha_image"`
				QualityLevel      interface{} `json:"quality_level"`
				IsOrigin          interface{} `json:"is_origin"`
				Animations        interface{} `json:"animations"`
				NewAppVisible     interface{} `json:"new_app_visible"`
				DupItemSchema     interface{} `json:"dup_item_schema"`
				Position          interface{} `json:"position"`
				AuditInfo         interface{} `json:"audit_info"`
				LastCommentTime   interface{} `json:"last_comment_time"`
				RecommendTags     interface{} `json:"recommend_tags"`
				InteractionStatus interface{} `json:"interaction_status"`
				LogItemExtra      string      `json:"log_item_extra"`
				AuthorTags        string      `json:"author_tags"`
				CellUICtrl        struct {
					ShowAuthorTags bool        `json:"show_author_tags"`
					ImageStyle     interface{} `json:"image_style"`
					DisplayStyle   interface{} `json:"display_style"`
				} `json:"cell_ui_ctrl"`
			} `json:"item"`
			UserInfo         interface{} `json:"user_info"`
			Hashtag          interface{} `json:"hashtag"`
			UsersHashtag     interface{} `json:"users_hashtag"`
			Comment          interface{} `json:"comment"`
			CommentInfo      interface{} `json:"comment_info"`
			ReplyInfo        interface{} `json:"reply_info"`
			BannerInfo       interface{} `json:"banner_info"`
			Button           interface{} `json:"button"`
			DislikeOptions   interface{} `json:"dislike_options"`
			AdInfo           interface{} `json:"ad_info"`
			PassThrough      string      `json:"pass_through"`
			Stickup          bool        `json:"stickup"`
			FollowOptions    interface{} `json:"follow_options"`
			CommentShowAd    bool        `json:"comment_show_ad"`
			LastViewTime     int         `json:"last_view_time"`
			ClubInfo         interface{} `json:"club_info"`
			AlbumInfo        interface{} `json:"album_info"`
			BlockInfo        interface{} `json:"block_info"`
			UserHotItemList  interface{} `json:"user_hot_item_list"`
			LiveInfo         interface{} `json:"live_info"`
			CollectionInfo   interface{} `json:"collection_info"`
			UserItems        interface{} `json:"user_items"`
			SearchAttachInfo interface{} `json:"search_attach_info"`
			SaasLiveInfo     interface{} `json:"saas_live_info"`
			BoardInfo        interface{} `json:"board_info"`
			CardItem         interface{} `json:"card_item"`
		} `json:"data"`
	} `json:"data"`
}

func (a PipixiaAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
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
	re := regexp.MustCompile(`/item\/(.*)\?`)
	match := re.FindStringSubmatch(loc)
	if len(match) != 2 {
		return nil, errors.New("匹配视频id失败")
	}
	getResp, err := restyHttp.GetMobileHttpRequest().Get("https://is.snssdk.com/bds/cell/detail/?cell_type=1&aid=1319&app_name=super&cell_id=" + match[1])
	if err != nil {
		return nil, err
	}
	var pipixiaResponse PipixiaAdapterResponse
	err = json.Cjson.Unmarshal(getResp.Body(), &pipixiaResponse)
	if err != nil {
		return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             pipixiaResponse.Data.Data.Item.Author.Name,
		AuthorAvatar:           pipixiaResponse.Data.Data.Item.Author.Avatar.DownloadList[0].URL,
		Title:                  pipixiaResponse.Data.Data.Item.Content,
		Cover:                  pipixiaResponse.Data.Data.Item.Cover.URLList[0].URL,
		NoWatermarkDownloadUrl: pipixiaResponse.Data.Data.Item.OriginVideoDownload.URLList[0].URL,
	}
	return shortVideoInfo, nil
}
