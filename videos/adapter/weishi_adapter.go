package adapter

import (
	"errors"
	"github.com/assimon/svbot/internal/json"
	"github.com/assimon/svbot/internal/log"
	"github.com/assimon/svbot/internal/restyHttp"
	"regexp"
)

type WeishiAdapter struct{}

type WeishiAdapterResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Feeds []struct {
			ID       string `json:"id"`
			Wording  string `json:"wording"`
			Type     int    `json:"type"`
			PosterID string `json:"poster_id"`
			Poster   struct {
				ID                string `json:"id"`
				Type              int    `json:"type"`
				UID               string `json:"uid"`
				Createtime        int    `json:"createtime"`
				Nick              string `json:"nick"`
				Avatar            string `json:"avatar"`
				Sex               int    `json:"sex"`
				FeedlistTimeID    string `json:"feedlist_time_id"`
				FeedlistHotID     string `json:"feedlist_hot_id"`
				RelatedFeedlistID string `json:"related_feedlist_id"`
				FollowerlistID    string `json:"followerlist_id"`
				InteresterlistID  string `json:"interesterlist_id"`
				ChatlistID        string `json:"chatlist_id"`
				RichFlag          int    `json:"rich_flag"`
				Age               int    `json:"age"`
				Address           string `json:"address"`
				Wealth            struct {
					FlowerNum int `json:"flower_num"`
					Score     int `json:"score"`
				} `json:"wealth"`
				Background        string `json:"background"`
				Status            string `json:"status"`
				FollowStatus      int    `json:"followStatus"`
				ChartScore        int    `json:"chartScore"`
				ChartRank         int    `json:"chartRank"`
				FeedGoldNum       int    `json:"feedGoldNum"`
				AvatarUpdatetime  int    `json:"avatar_updatetime"`
				DescFromOperator  string `json:"desc_from_operator"`
				SyncContent       int    `json:"sync_content"`
				FeedlistPraiseID  string `json:"feedlist_praise_id"`
				Settingmask       int    `json:"settingmask"`
				Originalavatar    string `json:"originalavatar"`
				BlockTime         string `json:"block_time"`
				Grade             int    `json:"grade"`
				Medal             int    `json:"medal"`
				BlockReason       string `json:"block_reason"`
				Qq                int    `json:"qq"`
				RecommendReason   string `json:"recommendReason"`
				LastUpdateFeedNum int    `json:"lastUpdateFeedNum"`
				Updateinfo        struct {
					Flag int    `json:"flag"`
					Tip  string `json:"tip"`
					Num  int    `json:"num"`
				} `json:"updateinfo"`
				NickUpdatetime     int64  `json:"nick_updatetime"`
				LastDownloadAvatar string `json:"lastDownloadAvatar"`
				RealName           string `json:"realName"`
				PinyinFirst        string `json:"pinyin_first"`
				CertifDesc         string `json:"certif_desc"`
				PrivateInfo        struct {
					PhoneNum string `json:"phone_num"`
					Name     string `json:"name"`
					IDNum    string `json:"id_num"`
				} `json:"privateInfo"`
				ExternInfo struct {
					MpEx struct {
						AuditPriority        string `json:"audit_priority"`
						SubPriority          string `json:"sub_priority"`
						DarenPriority        string `json:"daren_priority"`
						DarenCompany         string `json:"daren_company"`
						LoginResaveOldAvatar string `json:"loginResaveOldAvatar"`
						LoginResaveNewAvatar string `json:"loginResaveNewAvatar"`
						LoginResaveTime      string `json:"loginResaveTime"`
					} `json:"mpEx"`
					BindAcct  []interface{} `json:"bind_acct"`
					BgPicURL  string        `json:"bgPicUrl"`
					LevelInfo struct {
						Level           int `json:"level"`
						Score           int `json:"score"`
						PrevUpgradeTime int `json:"prev_upgrade_time"`
					} `json:"level_info"`
					WeishiID            string `json:"weishiId"`
					WeishiidModifyCount string `json:"weishiid_modify_count"`
					WatermarkType       int    `json:"watermark_type"`
					RealNick            string `json:"real_nick"`
					CmtLevel            struct {
						Level           int `json:"level"`
						Cmtscore        int `json:"cmtscore"`
						Dingscore       int `json:"dingscore"`
						PrevUpgradeTime int `json:"prev_upgrade_time"`
					} `json:"cmt_level"`
					FlexibilityFlag int `json:"flexibility_flag"`
					LiveStatus      int `json:"live_status"`
					NowLiveRoomID   int `json:"now_live_room_id"`
					MedalInfo       struct {
						TotalScore int           `json:"total_score"`
						MedalList  []interface{} `json:"medal_list"`
					} `json:"medal_info"`
					H5HasLogin   int    `json:"h5_has_login"`
					RelationType int    `json:"relation_type"`
					Feedid       string `json:"feedid"`
					CoverType    int    `json:"cover_type"`
					IndustryType int    `json:"industry_type"`
					InditeType   int    `json:"indite_type"`
				} `json:"extern_info"`
				CertifData struct {
					CertifIcon    string `json:"certif_icon"`
					CertifJumpurl string `json:"certif_jumpurl"`
				} `json:"certifData"`
				IsShowPOI    int `json:"isShowPOI"`
				IsShowGender int `json:"isShowGender"`
				FormatAddr   struct {
					Country  string `json:"country"`
					Province string `json:"province"`
					City     string `json:"city"`
				} `json:"formatAddr"`
				AuthorizeTime int `json:"authorize_time"`
				ActivityInfo  struct {
					InvitePersonid string `json:"invitePersonid"`
				} `json:"activityInfo"`
				SpecialIdentity struct {
				} `json:"special_identity"`
				TmpMark      int `json:"tmpMark"`
				PmtMark      int `json:"pmtMark"`
				IndustryInfo struct {
					PrimaryIndustry struct {
						IndustryID   int    `json:"industry_id"`
						IndustryDesc string `json:"industry_desc"`
					} `json:"primary_industry"`
					SecondaryIndustry struct {
						IndustryID   int    `json:"industry_id"`
						IndustryDesc string `json:"industry_desc"`
					} `json:"secondary_industry"`
				} `json:"industry_info"`
				Homeland struct {
					Country  string `json:"country"`
					Province string `json:"province"`
					City     string `json:"city"`
				} `json:"homeland"`
			} `json:"poster"`
			Video struct {
				FileID       string `json:"file_id"`
				FileSize     int    `json:"file_size"`
				Sha1         string `json:"sha1"`
				PlayIndex    int    `json:"play_index"`
				Duration     int    `json:"duration"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				Md5          string `json:"md5"`
				Orientation  int    `json:"orientation"`
				H265Hvc1     int    `json:"h265_hvc1"`
				MaxDb        int    `json:"max_db"`
				VoiceRatio   int    `json:"voice_ratio"`
				Loudnorm     string `json:"loudnorm"`
				MetaLoudnorm struct {
					InputI            string `json:"input_i"`
					InputTp           string `json:"input_tp"`
					InputLra          string `json:"input_lra"`
					InputThresh       string `json:"input_thresh"`
					OutputI           string `json:"output_i"`
					OutputTp          string `json:"output_tp"`
					OutputLra         string `json:"output_lra"`
					OutputThresh      string `json:"output_thresh"`
					NormalizationType string `json:"normalization_type"`
					TargetOffset      string `json:"target_offset"`
					WeishiI           string `json:"weishi_i"`
					WeishiTp          string `json:"weishi_tp"`
					WeishiLra         string `json:"weishi_lra"`
				} `json:"meta_loudnorm"`
			} `json:"video"`
			Images []struct {
				URL          string `json:"url"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				Type         int    `json:"type"`
				SpriteWidth  int    `json:"sprite_width"`
				SpriteHeight int    `json:"sprite_height"`
				SpriteSpan   int    `json:"sprite_span"`
				Priority     int    `json:"priority"`
				PhotoRgb     string `json:"photo_rgb"`
				Format       string `json:"format"`
			} `json:"images"`
			UgcVideoIds      []interface{} `json:"ugc_video_ids"`
			UgcVideos        []interface{} `json:"ugc_videos"`
			Createtime       int           `json:"createtime"`
			Mask             int           `json:"mask"`
			Score            int           `json:"score"`
			DingCount        int           `json:"ding_count"`
			CommentlistID    string        `json:"commentlist_id"`
			TotalCommentNum  int           `json:"total_comment_num"`
			Comments         []interface{} `json:"comments"`
			MaterialID       string        `json:"material_id"`
			MaterialDesc     string        `json:"material_desc"`
			DingHashID       string        `json:"ding_hash_id"`
			IsDing           int           `json:"is_ding"`
			PlayNum          int           `json:"playNum"`
			CharacterID      string        `json:"character_id"`
			FlowerNum        int           `json:"flower_num"`
			SendFlowerNum    int           `json:"send_flower_num"`
			RichFlag         int           `json:"rich_flag"`
			VideoURL         string        `json:"video_url"`
			MaterialThumburl string        `json:"material_thumburl"`
			Platform         int           `json:"platform"`
			Reserve          struct {
				Num2  string `json:"2"`
				Num3  string `json:"3"`
				Num6  string `json:"6"`
				Num36 string `json:"36"`
				Num38 string `json:"38"`
				Num41 string `json:"41"`
				Num47 string `json:"47"`
				Num62 string `json:"62"`
			} `json:"reserve"`
			VideoSpecUrls struct {
				Num0 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						TranscodeSource string `json:"transcode_source"`
						VideoEncoder    string `json:"video_encoder"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"0"`
				Num1 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						TranscodeSource string `json:"transcode_source"`
						VideoEncoder    string `json:"video_encoder"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"1"`
				Num2 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						VideoEncoder    string `json:"video_encoder"`
						TranscodeSource string `json:"transcode_source"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"2"`
				Num5 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						TranscodeSource string `json:"transcode_source"`
						VideoEncoder    string `json:"video_encoder"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"5"`
				Num6 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						VideoEncoder    string `json:"video_encoder"`
						TranscodeSource string `json:"transcode_source"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"6"`
				Num11 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						TranscodeSource string `json:"transcode_source"`
						VideoEncoder    string `json:"video_encoder"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"11"`
				Num12 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						VideoEncoder    string `json:"video_encoder"`
						TranscodeSource string `json:"transcode_source"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"12"`
				Num26 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						TranscodeSource string `json:"transcode_source"`
						VideoEncoder    string `json:"video_encoder"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"26"`
				Num27 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						TranscodeSource string `json:"transcode_source"`
						VideoEncoder    string `json:"video_encoder"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"27"`
				Num999 struct {
					URL           string `json:"url"`
					Size          int    `json:"size"`
					Hardorsoft    int    `json:"hardorsoft"`
					RecommendSpec int    `json:"recommendSpec"`
					HaveWatermark int    `json:"haveWatermark"`
					Width         int    `json:"width"`
					Height        int    `json:"height"`
					VideoCoding   int    `json:"videoCoding"`
					VideoQuality  int    `json:"videoQuality"`
					ExternInfo    struct {
						TranscodeSource string `json:"transcode_source"`
						VideoEncoder    string `json:"video_encoder"`
					} `json:"externInfo"`
					StaticCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"staticCover"`
					Fps int `json:"fps"`
				} `json:"999"`
			} `json:"video_spec_urls"`
			ShareInfo struct {
				JumpURL string `json:"jump_url"`
				BodyMap struct {
					Num0 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"0"`
					Num1 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"1"`
					Num2 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"2"`
					Num3 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"3"`
					Num4 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"4"`
					Num5 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"5"`
				} `json:"body_map"`
				WxMiniProgram struct {
					WebpageURL       string `json:"webpageUrl"`
					UserName         string `json:"userName"`
					Path             string `json:"path"`
					HdImageDataURL   string `json:"hdImageDataURL"`
					WithShareTicket  int    `json:"withShareTicket"`
					MiniProgramType  int    `json:"miniProgramType"`
					Appid            string `json:"appid"`
					VideoUserName    string `json:"videoUserName"`
					VideoSource      string `json:"videoSource"`
					VideoCoverWidth  int    `json:"videoCoverWidth"`
					VideoCoverHeight int    `json:"videoCoverHeight"`
					AppThumbURL      string `json:"appThumbUrl"`
					UniversalLink    string `json:"universalLink"`
					Disableforward   int    `json:"disableforward"`
				} `json:"wx_mini_program"`
				SqArkInfo struct {
					ArkData   string `json:"arkData"`
					ShareBody struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"shareBody"`
					CoverProto string `json:"coverProto"`
				} `json:"sq_ark_info"`
				ShareIconURL   string `json:"share_icon_url"`
				ShareIconTitle string `json:"share_icon_title"`
				BackgroundURL  string `json:"background_url"`
				ActivityType   int    `json:"activity_type"`
				HaibaoJumpURL  string `json:"haibao_jump_url"`
				HaibaoBodyMap  struct {
					Num0 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"0"`
					Num1 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"1"`
					Num2 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"2"`
					Num3 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"3"`
					Num4 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"4"`
					Num5 struct {
						Title    string `json:"title"`
						Desc     string `json:"desc"`
						ImageURL string `json:"image_url"`
						URL      string `json:"url"`
					} `json:"5"`
				} `json:"haibao_body_map"`
				BackgroundTitleColor string `json:"background_title_color"`
				HaibaoDesc           string `json:"haibao_desc"`
				ShareNum             int    `json:"share_num"`
				FeedCoverUpdtime     string `json:"feed_cover_updtime"`
			} `json:"share_info"`
			FeedGift struct {
				FeedGiftList []interface{} `json:"feedGiftList"`
			} `json:"feedGift"`
			GiftRank []interface{} `json:"giftRank"`
			TopicID  string        `json:"topic_id"`
			Topic    struct {
				ID             string        `json:"id"`
				Name           string        `json:"name"`
				ThumbURL1      string        `json:"thumbUrl1"`
				ThumbURL2      string        `json:"thumbUrl2"`
				ThumbURL3      string        `json:"thumbUrl3"`
				Detail         string        `json:"detail"`
				Createtime     int           `json:"createtime"`
				FeedlistTimeID string        `json:"feedlist_time_id"`
				FeedlistHotID  string        `json:"feedlist_hot_id"`
				MaterialIds    []interface{} `json:"material_ids"`
				Mask           int           `json:"mask"`
				Type           int           `json:"type"`
				Reserve        struct {
				} `json:"reserve"`
				ViewNum    int `json:"view_num"`
				StartTime  int `json:"start_time"`
				EndTime    int `json:"end_time"`
				AppVersion int `json:"appVersion"`
				WorkNum    int `json:"workNum"`
				LikeNum    int `json:"likeNum"`
				Person     struct {
					ID                string `json:"id"`
					Type              int    `json:"type"`
					UID               string `json:"uid"`
					Createtime        int    `json:"createtime"`
					Nick              string `json:"nick"`
					Avatar            string `json:"avatar"`
					Sex               int    `json:"sex"`
					FeedlistTimeID    string `json:"feedlist_time_id"`
					FeedlistHotID     string `json:"feedlist_hot_id"`
					RelatedFeedlistID string `json:"related_feedlist_id"`
					FollowerlistID    string `json:"followerlist_id"`
					InteresterlistID  string `json:"interesterlist_id"`
					ChatlistID        string `json:"chatlist_id"`
					RichFlag          int    `json:"rich_flag"`
					Age               int    `json:"age"`
					Address           string `json:"address"`
					Wealth            struct {
						FlowerNum int `json:"flower_num"`
						Score     int `json:"score"`
					} `json:"wealth"`
					Background        string `json:"background"`
					Status            string `json:"status"`
					FollowStatus      int    `json:"followStatus"`
					ChartScore        int    `json:"chartScore"`
					ChartRank         int    `json:"chartRank"`
					FeedGoldNum       int    `json:"feedGoldNum"`
					AvatarUpdatetime  int    `json:"avatar_updatetime"`
					DescFromOperator  string `json:"desc_from_operator"`
					SyncContent       int    `json:"sync_content"`
					FeedlistPraiseID  string `json:"feedlist_praise_id"`
					Settingmask       int    `json:"settingmask"`
					Originalavatar    string `json:"originalavatar"`
					BlockTime         string `json:"block_time"`
					Grade             int    `json:"grade"`
					Medal             int    `json:"medal"`
					BlockReason       string `json:"block_reason"`
					Qq                int    `json:"qq"`
					RecommendReason   string `json:"recommendReason"`
					LastUpdateFeedNum int    `json:"lastUpdateFeedNum"`
					Updateinfo        struct {
						Flag int    `json:"flag"`
						Tip  string `json:"tip"`
						Num  int    `json:"num"`
					} `json:"updateinfo"`
					NickUpdatetime     int    `json:"nick_updatetime"`
					LastDownloadAvatar string `json:"lastDownloadAvatar"`
					RealName           string `json:"realName"`
					PinyinFirst        string `json:"pinyin_first"`
					CertifDesc         string `json:"certif_desc"`
					PrivateInfo        struct {
						PhoneNum string `json:"phone_num"`
						Name     string `json:"name"`
						IDNum    string `json:"id_num"`
					} `json:"privateInfo"`
					ExternInfo struct {
						MpEx struct {
						} `json:"mpEx"`
						BindAcct  []interface{} `json:"bind_acct"`
						BgPicURL  string        `json:"bgPicUrl"`
						LevelInfo struct {
							Level           int `json:"level"`
							Score           int `json:"score"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"level_info"`
						WeishiID            string `json:"weishiId"`
						WeishiidModifyCount string `json:"weishiid_modify_count"`
						WatermarkType       int    `json:"watermark_type"`
						RealNick            string `json:"real_nick"`
						CmtLevel            struct {
							Level           int `json:"level"`
							Cmtscore        int `json:"cmtscore"`
							Dingscore       int `json:"dingscore"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"cmt_level"`
						FlexibilityFlag int `json:"flexibility_flag"`
						LiveStatus      int `json:"live_status"`
						NowLiveRoomID   int `json:"now_live_room_id"`
						MedalInfo       struct {
							TotalScore int           `json:"total_score"`
							MedalList  []interface{} `json:"medal_list"`
						} `json:"medal_info"`
						H5HasLogin   int    `json:"h5_has_login"`
						RelationType int    `json:"relation_type"`
						Feedid       string `json:"feedid"`
						CoverType    int    `json:"cover_type"`
						IndustryType int    `json:"industry_type"`
						InditeType   int    `json:"indite_type"`
					} `json:"extern_info"`
					CertifData struct {
						CertifIcon    string `json:"certif_icon"`
						CertifJumpurl string `json:"certif_jumpurl"`
					} `json:"certifData"`
					IsShowPOI    int `json:"isShowPOI"`
					IsShowGender int `json:"isShowGender"`
					FormatAddr   struct {
						Country  string `json:"country"`
						Province string `json:"province"`
						City     string `json:"city"`
					} `json:"formatAddr"`
					AuthorizeTime int `json:"authorize_time"`
					ActivityInfo  struct {
						InvitePersonid string `json:"invitePersonid"`
					} `json:"activityInfo"`
					SpecialIdentity struct {
					} `json:"special_identity"`
					TmpMark      int `json:"tmpMark"`
					PmtMark      int `json:"pmtMark"`
					IndustryInfo struct {
						PrimaryIndustry struct {
							IndustryID   int    `json:"industry_id"`
							IndustryDesc string `json:"industry_desc"`
						} `json:"primary_industry"`
						SecondaryIndustry struct {
							IndustryID   int    `json:"industry_id"`
							IndustryDesc string `json:"industry_desc"`
						} `json:"secondary_industry"`
					} `json:"industry_info"`
					Homeland struct {
						Country  string `json:"country"`
						Province string `json:"province"`
						City     string `json:"city"`
					} `json:"homeland"`
				} `json:"person"`
				FeedID            string `json:"feed_id"`
				PendantMaterialID string `json:"pendant_material_id"`
				MusicMaterialID   string `json:"music_material_id"`
				MusicInfo         struct {
					ID              string        `json:"id"`
					Name            string        `json:"name"`
					Desc            string        `json:"desc"`
					Type            string        `json:"type"`
					ThumbURL        string        `json:"thumbUrl"`
					Version         int           `json:"version"`
					MiniSptVersion  int           `json:"miniSptVersion"`
					PackageURL      string        `json:"packageUrl"`
					FeedlistTimeID  string        `json:"feedlist_time_id"`
					FeedlistHotID   string        `json:"feedlist_hot_id"`
					TopicIds        []interface{} `json:"topic_ids"`
					Mask            int           `json:"mask"`
					ShortName       string        `json:"shortName"`
					RichFlag        int           `json:"rich_flag"`
					EffectID        string        `json:"effectId"`
					Rgbcolor        string        `json:"rgbcolor"`
					IsCollected     int           `json:"isCollected"`
					BubbleStartTime int           `json:"bubbleStartTime"`
					BubbleEndTime   int           `json:"bubbleEndTime"`
					CollectTime     int           `json:"collectTime"`
					SdkInfo         struct {
						IsSdk            int           `json:"isSdk"`
						SdkMinVersion    int           `json:"sdkMinVersion"`
						SdkMaxVersion    int           `json:"sdkMaxVersion"`
						SdkMinSptVersion int           `json:"sdkMinSptVersion"`
						SdkAndroidGrays  []interface{} `json:"sdkAndroidGrays"`
						SdkMinVersionStr string        `json:"sdkMinVersionStr"`
						SdkMaxVersionStr string        `json:"sdkMaxVersionStr"`
					} `json:"sdkInfo"`
					BigThumbURL string        `json:"bigThumbUrl"`
					Priority    int           `json:"priority"`
					MusicIDs    []interface{} `json:"musicIDs"`
					Platform    string        `json:"platform"`
					Reserve     struct {
					} `json:"reserve"`
					Category                string        `json:"category"`
					ShootingTips            string        `json:"shooting_tips"`
					VecSubcategory          []interface{} `json:"vec_subcategory"`
					DemoVideoList           []interface{} `json:"demo_video_list"`
					RecommendTemplateTags   []interface{} `json:"recommendTemplateTags"`
					InspirationButtonText   string        `json:"inspirationButtonText"`
					InspirationButtonSchema string        `json:"inspirationButtonSchema"`
					RelationMaterialID      string        `json:"relationMaterialId"`
					MoreMaterialLink        string        `json:"moreMaterialLink"`
					StartTime               int           `json:"startTime"`
					EndTime                 int           `json:"endTime"`
					RandomPackageUrls       struct {
					} `json:"randomPackageUrls"`
					HideType   int `json:"hideType"`
					FollowInfo struct {
						IsFollowShotShown         int    `json:"isFollowShotShown"`
						SchemeType                int    `json:"schemeType"`
						Scheme                    string `json:"scheme"`
						IconURL                   string `json:"iconUrl"`
						SharePageMark             string `json:"sharePageMark"`
						ShareButtonMark           string `json:"shareButtonMark"`
						Name                      string `json:"name"`
						FollowNewPagIconURL       string `json:"followNewPagIconUrl"`
						FollowNewDefaultIconURL   string `json:"followNewDefaultIconUrl"`
						FollowNewIconScheme       string `json:"followNewIconScheme"`
						FollowNewNameScheme       string `json:"followNewNameScheme"`
						FollowNewMultiLabelScheme string `json:"followNewMultiLabelScheme"`
						CategoryFollowShotInfo    struct {
							IsFollowShotShown int `json:"isFollowShotShown"`
							Priority          int `json:"priority"`
							FollowShotStyle   int `json:"followShotStyle"`
						} `json:"categoryFollowShotInfo"`
					} `json:"followInfo"`
					State       int    `json:"state"`
					Deleted     int    `json:"deleted"`
					Packages    string `json:"packages"`
					PackageUrls struct {
					} `json:"packageUrls"`
					MaterialPackageUrls struct {
					} `json:"materialPackageUrls"`
					RoundThumbURL    string `json:"roundThumbUrl"`
					BigRoundThumbURL string `json:"bigRoundThumbUrl"`
					Title            string `json:"title"`
					CarID            string `json:"carID"`
					ComposedInfo     struct {
						VideoDuration       int           `json:"videoDuration"`
						IncludeMIDList      []interface{} `json:"includeMIDList"`
						AbilityList         []interface{} `json:"abilityList"`
						SlotInfo            string        `json:"slotInfo"`
						ThumbResolution     string        `json:"thumbResolution"`
						SlotInfoList        []interface{} `json:"slotInfoList"`
						IncludeMaterialInfo struct {
						} `json:"includeMaterialInfo"`
					} `json:"composedInfo"`
					AuthorID             string `json:"authorID"`
					UseCount             int    `json:"useCount"`
					CardID               string `json:"cardID"`
					CompressedPackageURL string `json:"compressedPackageUrl"`
				} `json:"music_info"`
				PendantMaterialIDIos string `json:"pendant_material_id_ios"`
				MediaMaterialURL     string `json:"media_material_url"`
				BubbleStartTime      int    `json:"bubble_start_time"`
				BubbleEndTime        int    `json:"bubble_end_time"`
				BubbleCopywrite      string `json:"bubble_copywrite"`
				Rgbcolor             int    `json:"rgbcolor"`
				Lplaynum             int    `json:"lplaynum"`
				QqMusicInfo          struct {
					AlbumInfo struct {
						UIID    int    `json:"uiId"`
						StrMid  string `json:"strMid"`
						StrName string `json:"strName"`
						StrPic  string `json:"strPic"`
					} `json:"albumInfo"`
					SingerInfo struct {
						UIID         int           `json:"uiId"`
						StrMid       string        `json:"strMid"`
						StrName      string        `json:"strName"`
						StrPic       string        `json:"strPic"`
						StrSchema    string        `json:"strSchema"`
						OtherSingers []interface{} `json:"otherSingers"`
					} `json:"singerInfo"`
					SongInfo struct {
						UIID               int    `json:"uiId"`
						StrMid             string `json:"strMid"`
						StrName            string `json:"strName"`
						StrGenre           string `json:"strGenre"`
						IIsOnly            int    `json:"iIsOnly"`
						StrLanguage        string `json:"strLanguage"`
						IPlayable          int    `json:"iPlayable"`
						ITrySize           int    `json:"iTrySize"`
						ITryBegin          int    `json:"iTryBegin"`
						ITryEnd            int    `json:"iTryEnd"`
						IPlayTime          int    `json:"iPlayTime"`
						StrH5URL           string `json:"strH5Url"`
						StrPlayURL         string `json:"strPlayUrl"`
						StrPlayURLStandard string `json:"strPlayUrlStandard"`
						StrPlayURLHq       string `json:"strPlayUrlHq"`
						StrPlayURLSq       string `json:"strPlayUrlSq"`
						ISize              int    `json:"iSize"`
						ISizeStandard      int    `json:"iSizeStandard"`
						ISizeHq            int    `json:"iSizeHq"`
						ISizeSq            int    `json:"iSizeSq"`
						Copyright          int    `json:"copyright"`
						ISource            int    `json:"iSource"`
						PublicTime         string `json:"publicTime"`
						SongTitle          string `json:"songTitle"`
						SongDescription    string `json:"songDescription"`
						State              int    `json:"state"`
						Deleted            int    `json:"deleted"`
						StartTime          string `json:"startTime"`
						EndTime            string `json:"endTime"`
					} `json:"songInfo"`
					LyricInfo struct {
						UISongID      int    `json:"uiSongId"`
						StrSongMid    string `json:"strSongMid"`
						StrFormat     string `json:"strFormat"`
						StrLyric      string `json:"strLyric"`
						StrMatchLyric string `json:"strMatchLyric"`
					} `json:"lyricInfo"`
					ConfInfo struct {
						IType               int    `json:"iType"`
						IStartPos           int    `json:"iStartPos"`
						StrLabel            string `json:"strLabel"`
						IsCollected         int    `json:"isCollected"`
						CollectTime         int    `json:"collectTime"`
						Exclusive           int    `json:"exclusive"`
						FollowFeed          string `json:"followFeed"`
						UseCount            int    `json:"useCount"`
						TogetherFeed        string `json:"togetherFeed"`
						TogetherType        int    `json:"togetherType"`
						FeedUseType         int    `json:"feedUseType"`
						DefaultFeedPosition int    `json:"defaultFeedPosition"`
						DefaultTogetherFeed int    `json:"defaultTogetherFeed"`
						BubbleStartTime     int    `json:"bubbleStartTime"`
						BubbleEndTime       int    `json:"bubbleEndTime"`
						SongLabels          struct {
						} `json:"songLabels"`
						SongLabelCategory struct {
						} `json:"songLabelCategory"`
						IsStuckPoint      bool   `json:"isStuckPoint"`
						StuckPointJSONURL string `json:"stuckPointJsonUrl"`
						TrackBeatInfo     struct {
							TrackBeatFinished int `json:"trackBeatFinished"`
							Vocal             struct {
								JSONURL      string `json:"jsonURL"`
								AudioFileURL string `json:"audioFileURL"`
							} `json:"vocal"`
							Drums struct {
								JSONURL      string `json:"jsonURL"`
								AudioFileURL string `json:"audioFileURL"`
							} `json:"drums"`
							Accompaniment struct {
								JSONURL      string `json:"jsonURL"`
								AudioFileURL string `json:"audioFileURL"`
							} `json:"accompaniment"`
							Bass struct {
								JSONURL      string `json:"jsonURL"`
								AudioFileURL string `json:"audioFileURL"`
							} `json:"bass"`
						} `json:"trackBeatInfo"`
						ExtraInfo string `json:"extraInfo"`
					} `json:"confInfo"`
					SubtitleInfo struct {
						UISongID      int    `json:"uiSongId"`
						StrSongMid    string `json:"strSongMid"`
						StrFormat     string `json:"strFormat"`
						StrLyric      string `json:"strLyric"`
						StrMatchLyric string `json:"strMatchLyric"`
					} `json:"subtitleInfo"`
					Foreignlyric struct {
						UISongID      int    `json:"uiSongId"`
						StrSongMid    string `json:"strSongMid"`
						StrFormat     string `json:"strFormat"`
						StrLyric      string `json:"strLyric"`
						StrMatchLyric string `json:"strMatchLyric"`
					} `json:"foreignlyric"`
					RecommendInfo struct {
						TraceStr      string `json:"traceStr"`
						AnalyseResult string `json:"analyse_result"`
						RecomReason   string `json:"recom_reason"`
					} `json:"recommendInfo"`
					UnplayableInfo struct {
						UnplayableCode int    `json:"unplayableCode"`
						UnplayableMsg  string `json:"unplayableMsg"`
					} `json:"unplayableInfo"`
					LabelInfo       []interface{} `json:"labelInfo"`
					MusicType       int           `json:"musicType"`
					MusicSrcType    int           `json:"musicSrcType"`
					CacheUpdateTime int           `json:"cacheUpdateTime"`
					State           int           `json:"state"`
				} `json:"qqMusicInfo"`
				TopicType           int    `json:"topicType"`
				TopicSource         int    `json:"topicSource"`
				Creater             string `json:"creater"`
				LastOperator        string `json:"lastOperator"`
				SecurityAuditstate  int    `json:"securityAuditstate"`
				SecurityAuditReason string `json:"securityAuditReason"`
				ManualAuditstate    int    `json:"manualAuditstate"`
				ManualAuditReason   string `json:"manualAuditReason"`
				Status              int    `json:"status"`
				UpdateTime          int    `json:"updateTime"`
				NewAppVersion       string `json:"newAppVersion"`
				TopicMusicName      string `json:"TopicMusicName"`
				PendantMaterialCate string `json:"pendant_material_cate"`
				Schema              string `json:"schema"`
				SchemaType          int    `json:"schemaType"`
				ExternalLink        struct {
					LinkURL  string `json:"link_url"`
					LinkName string `json:"link_name"`
				} `json:"external_link"`
				InteractiveNews struct {
					InteractiveDetails []interface{} `json:"interactive_details"`
				} `json:"interactive_news"`
				ActivityInfo struct {
					Name          string `json:"name"`
					Label         string `json:"label"`
					SubName       string `json:"sub_name"`
					BtnTxt        string `json:"btn_txt"`
					ShowStartTime int    `json:"show_start_time"`
					ShowEndTime   int    `json:"show_end_time"`
					StartTime     int    `json:"start_time"`
					EndTime       int    `json:"end_time"`
					RuleInfo      struct {
						RuleDetails []interface{} `json:"rule_details"`
					} `json:"rule_info"`
					NeedShow          int    `json:"need_show"`
					ResourceBackColor string `json:"resource_back_color"`
					Status            int    `json:"status"`
				} `json:"activity_info"`
				PublishInfo struct {
					BtnStyle          int    `json:"btn_style"`
					BtnText           string `json:"btn_text"`
					BtnPic            string `json:"btn_pic"`
					BlueCollarPublish struct {
						DefaultCallModel    int           `json:"default_call_model"`
						DefaultCamera       int           `json:"default_camera"`
						BlueCollarMaterials []interface{} `json:"blue_collar_materials"`
						PendantID           string        `json:"pendant_id"`
						MusicID             string        `json:"music_id"`
						SongListID          string        `json:"song_list_id"`
						TeleprompterDesc    string        `json:"teleprompter_desc"`
					} `json:"blue_collar_publish"`
				} `json:"publish_info"`
				LatestPublishTime int `json:"latest_publish_time"`
				CollarType        int `json:"collar_type"`
				BlueCollar        struct {
					BlueCollarTags          []interface{} `json:"blue_collar_tags"`
					BlueCollarProfessionTag string        `json:"blue_collar_profession_tag"`
				} `json:"blue_collar"`
				UserGroupID string `json:"user_group_id"`
			} `json:"topic"`
			FlowerNumDb int           `json:"flowerNumDb"`
			FlowerRank  []interface{} `json:"flowerRank"`
			FeedDesc    string        `json:"feed_desc"`
			DescMask    int           `json:"desc_mask"`
			ShieldID    string        `json:"shieldId"`
			VideoCover  struct {
				StaticCover struct {
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
					Priority     int    `json:"priority"`
					PhotoRgb     string `json:"photo_rgb"`
					Format       string `json:"format"`
				} `json:"static_cover"`
				DynamicCover struct {
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
					Priority     int    `json:"priority"`
					PhotoRgb     string `json:"photo_rgb"`
					Format       string `json:"format"`
				} `json:"dynamic_cover"`
				CoverTime     int           `json:"cover_time"`
				VMetaEffect   []interface{} `json:"vMetaEffect"`
				AnimatedCover struct {
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
					Priority     int    `json:"priority"`
					PhotoRgb     string `json:"photo_rgb"`
					Format       string `json:"format"`
				} `json:"animated_cover"`
				SmallAnimatedCover struct {
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
					Priority     int    `json:"priority"`
					PhotoRgb     string `json:"photo_rgb"`
					Format       string `json:"format"`
				} `json:"small_animated_cover"`
				CoverWidth      int `json:"cover_width"`
				CoverHeight     int `json:"cover_height"`
				AnimatedCover5F struct {
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
					Priority     int    `json:"priority"`
					PhotoRgb     string `json:"photo_rgb"`
					Format       string `json:"format"`
				} `json:"animated_cover_5f"`
				SmallAnimatedCover5F struct {
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
					Priority     int    `json:"priority"`
					PhotoRgb     string `json:"photo_rgb"`
					Format       string `json:"format"`
				} `json:"small_animated_cover_5f"`
				SmartCover struct {
					URL          string `json:"url"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Type         int    `json:"type"`
					SpriteWidth  int    `json:"sprite_width"`
					SpriteHeight int    `json:"sprite_height"`
					SpriteSpan   int    `json:"sprite_span"`
					Priority     int    `json:"priority"`
					PhotoRgb     string `json:"photo_rgb"`
					Format       string `json:"format"`
				} `json:"smart_cover"`
			} `json:"video_cover"`
			GeoInfo struct {
				Country   string `json:"country"`
				Province  string `json:"province"`
				City      string `json:"city"`
				Latitude  int    `json:"latitude"`
				Longitude int    `json:"longitude"`
				Altitude  int    `json:"altitude"`
				District  string `json:"district"`
				Name      string `json:"name"`
				Distance  int    `json:"distance"`
				PolyGeoID string `json:"polyGeoID"`
			} `json:"geoInfo"`
			MusicID  string `json:"music_id"`
			VideoBgm struct {
				MusicID    string `json:"music_id"`
				MusicCover string `json:"music_cover"`
				MusicDesc  string `json:"music_desc"`
				Duration   int    `json:"duration"`
				Size       int    `json:"size"`
				FeedID     string `json:"feed_id"`
				MpEx       struct {
					TogetherBgm string `json:"together_bgm"`
					Mp3ID       string `json:"mp3_id"`
				} `json:"mpEx"`
			} `json:"video_bgm"`
			RecgBgm struct {
				MusicID    string `json:"music_id"`
				MusicCover string `json:"music_cover"`
				MusicDesc  string `json:"music_desc"`
				Duration   int    `json:"duration"`
				Size       int    `json:"size"`
				FeedID     string `json:"feed_id"`
				MpEx       struct {
				} `json:"mpEx"`
			} `json:"recg_bgm"`
			EnableRealRcmd      int    `json:"enable_real_rcmd"`
			FeedDescWithat      string `json:"feed_desc_withat"`
			FeedRecommendReason string `json:"feed_recommend_reason"`
			Interaction         struct {
				MpEx struct {
				} `json:"mpEx"`
				Type     int           `json:"type"`
				PersonID string        `json:"person_id"`
				FeedID   string        `json:"feed_id"`
				Score    int           `json:"score"`
				Buttons  []interface{} `json:"buttons"`
			} `json:"interaction"`
			Ornament struct {
				MpEx struct {
				} `json:"mpEx"`
				FilterID    string `json:"filter_id"`
				FilterName  string `json:"filter_name"`
				PendantID   string `json:"pendant_id"`
				PendantCate string `json:"pendant_cate"`
			} `json:"ornament"`
			VideoOrnaments []interface{} `json:"video_ornaments"`
			HaveText       int           `json:"have_text"`
			ExternInfo     struct {
				MpEx struct {
					FeedCover        string `json:"feed_cover"`
					ShowWxShareIcon  string `json:"show_wx_share_icon"`
					ActivityInfo     string `json:"activity_info"`
					ReportJSON       string `json:"report_json"`
					PrepareRecommend string `json:"prepare_recommend"`
					SecurityCheck    string `json:"security_check"`
					FeedSource       string `json:"feed_source"`
				} `json:"mpEx"`
				VisibleType       int `json:"visible_type"`
				ActivityShareInfo struct {
				} `json:"activity_share_info"`
				ActTogetherInfo struct {
					ExtInfo struct {
					} `json:"extInfo"`
					AllowTogether int    `json:"allowTogether"`
					TogetherType  int    `json:"togetherType"`
					PolyID        string `json:"polyId"`
					LastFeedID    string `json:"lastFeedId"`
					SrcFeedID     string `json:"srcFeedId"`
					TogetherCount int    `json:"togetherCount"`
					TogetherSpec  struct {
						Num1 int `json:"1"`
						Num2 int `json:"2"`
						Num3 int `json:"3"`
					} `json:"togetherSpec"`
					TogetherJump string `json:"togetherJump"`
					LastPersonID string `json:"lastPersonId"`
					GhostFeed    int    `json:"ghostFeed"`
					SrcBgmID     string `json:"srcBgmId"`
					FeedPosition struct {
					} `json:"feedPosition"`
					DefaultFeedPosition int `json:"defaultFeedPosition"`
					DefaultTogetherFeed int `json:"defaultTogetherFeed"`
				} `json:"actTogetherInfo"`
				DangerMarker int    `json:"danger_marker"`
				Rowkey       string `json:"rowkey"`
				FeedAdsInfo  struct {
					Icon        string `json:"icon"`
					IconWidth   int    `json:"icon_width"`
					IconHeight  int    `json:"icon_height"`
					Schema      string `json:"schema"`
					CommentType int    `json:"comment_type"`
					AdsGoal     int    `json:"ads_goal"`
					AdsType     int    `json:"ads_type"`
					QbossReport struct {
						Qbossid   int    `json:"qbossid"`
						TaskID    int    `json:"task_id"`
						TraceInfo string `json:"trace_info"`
						Position  int    `json:"position"`
					} `json:"qboss_report"`
					Extra struct {
					} `json:"extra"`
					AdsName string `json:"ads_name"`
				} `json:"feedAdsInfo"`
				ClarifyScore int `json:"clarifyScore"`
				ConcernHint  struct {
					EnableHint  int    `json:"enableHint"`
					BeginSecond int    `json:"beginSecond"`
					EndSecond   int    `json:"endSecond"`
					Thumbnail   string `json:"thumbnail"`
					Hint        string `json:"hint"`
				} `json:"concernHint"`
				RewardNum       int    `json:"reward_num"`
				SubtitleFlag    int    `json:"subtitle_flag"`
				UploadLyricFlag int    `json:"upload_lyric_flag"`
				FriendLikeNum   int    `json:"friend_like_num"`
				SafeVisibleType int    `json:"safe_visible_type"`
				RecommendReason string `json:"recommend_reason"`
				NowLiveRoomID   int    `json:"now_live_room_id"`
				ExtraMask       int    `json:"extra_mask"`
				RecommendMore   int    `json:"recommend_more"`
				QuestionList    struct {
					MaxIndex  int           `json:"max_index"`
					Questions []interface{} `json:"questions"`
				} `json:"question_list"`
				InteractConf struct {
					StickerData struct {
						TimeLines []interface{} `json:"time_lines"`
					} `json:"sticker_data"`
					MagicData struct {
						VideoWidth  int           `json:"video_width"`
						VideoHeight int           `json:"video_height"`
						EventList   []interface{} `json:"event_list"`
						MpEx        struct {
						} `json:"mpEx"`
					} `json:"magic_data"`
					Token           string `json:"token"`
					TemplateTypes   string `json:"template_types"`
					VideoShareCover struct {
						URL          string `json:"url"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						Type         int    `json:"type"`
						SpriteWidth  int    `json:"sprite_width"`
						SpriteHeight int    `json:"sprite_height"`
						SpriteSpan   int    `json:"sprite_span"`
						Priority     int    `json:"priority"`
						PhotoRgb     string `json:"photo_rgb"`
						Format       string `json:"format"`
					} `json:"video_share_cover"`
					QzoneSkin struct {
						SkinID        string `json:"skin_id"`
						SkinType      int    `json:"skin_type"`
						Picurl        string `json:"picurl"`
						Bgcolor       string `json:"bgcolor"`
						GradientBegin string `json:"gradient_begin"`
						GradientEnd   string `json:"gradient_end"`
						PicurlAnd     string `json:"picurl_and"`
					} `json:"qzone_skin"`
					TemplateName      string `json:"template_name"`
					TemplateID        string `json:"template_id"`
					TemplateBusiness  string `json:"template_business"`
					TemplateTitleSkin struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"template_title_skin"`
					Path2Ending struct {
					} `json:"path_2_ending"`
					EndingDetail struct {
					} `json:"ending_detail"`
					Paths []interface{} `json:"paths"`
					MpEx  struct {
					} `json:"mpEx"`
					HbActivityType  int    `json:"hb_activity_type"`
					HbBusinessType  int    `json:"hb_business_type"`
					DatiActivityID  string `json:"datiActivityId"`
					IndexJSONURL    string `json:"index_json_url"`
					WebIndexJSONURL string `json:"web_index_json_url"`
				} `json:"interact_conf"`
				InteractUgcData struct {
					UgcContent       string `json:"ugc_content"`
					HasVote          int    `json:"has_vote"`
					PersonOfficeName string `json:"person_office_name"`
					HasDone          int    `json:"has_done"`
					Score            int    `json:"score"`
					EndingIdx        int    `json:"ending_idx"`
					NeedFakeDesc     bool   `json:"need_fake_desc"`
					UserFlowKey      string `json:"user_flow_key"`
					HasGrab          int    `json:"has_grab"`
				} `json:"interact_ugc_data"`
				SrcFeedID        string        `json:"src_feed_id"`
				VKeyFrame        []interface{} `json:"vKeyFrame"`
				ZanStyleID       int           `json:"zan_style_id"`
				GameBattleReport struct {
					GameType  int    `json:"gameType"`
					VideoType int    `json:"videoType"`
					Vid       string `json:"vid"`
					IconURL   string `json:"iconUrl"`
					Schema    string `json:"schema"`
				} `json:"game_battle_report"`
				BusinessType       int    `json:"business_type"`
				ExternalSdkVersion string `json:"external_sdk_version"`
				BusinessConf       struct {
					Type             string        `json:"type"`
					BusinessInfo     string        `json:"business_info"`
					TemplateBusiness string        `json:"template_business"`
					BusinessVersion  int           `json:"business_version"`
					BusinessIds      []interface{} `json:"business_ids"`
					TaskID           string        `json:"task_id"`
				} `json:"business_conf"`
				Competition struct {
					TrackID       string `json:"track_id"`
					TrackName     string `json:"track_name"`
					TrackDesc     string `json:"track_desc"`
					RemainVote    int    `json:"remain_vote"`
					HostRank      int    `json:"host_rank"`
					RemainWebVote int    `json:"remain_web_vote"`
					TrackState    int    `json:"track_state"`
					Status        int    `json:"status"`
					HostVote      int    `json:"host_vote"`
					VoteBtnText   string `json:"vote_btn_text"`
					VoteBtnSchema string `json:"vote_btn_schema"`
					Ext           struct {
					} `json:"ext"`
					EventRank int `json:"event_rank"`
				} `json:"competition"`
				VideoAdapterInfo struct {
					ShowStatusBar   int    `json:"show_status_bar"`
					StatusBarHeight int    `json:"status_bar_height"`
					EnableCrop      int    `json:"enable_crop"`
					Ratio           string `json:"ratio"`
				} `json:"videoAdapterInfo"`
				GameVideoFeedInfo struct {
					AbInfo []interface{} `json:"abInfo"`
					IsPost bool          `json:"isPost"`
					Schema string        `json:"schema"`
				} `json:"gameVideoFeedInfo"`
				SafeCheckStatus int           `json:"safe_check_status"`
				RelatedVideos   []interface{} `json:"related_videos"`
				OvertComment    struct {
					SimpleComments []interface{} `json:"simpleComments"`
				} `json:"overtComment"`
				Mask                int `json:"mask"`
				HomepageAuditMode   int `json:"homepage_audit_mode"`
				HomepageAuditStatus int `json:"homepage_audit_status"`
				BusinessInfo        struct {
					Category struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"category"`
					SubCategory struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"sub_category"`
				} `json:"business_info"`
				Mark struct {
					MarkType int    `json:"mark_type"`
					Title    string `json:"title"`
				} `json:"mark"`
				IsFavor int `json:"is_favor"`
				FvsInfo struct {
					Type           int    `json:"type"`
					FvsID          string `json:"fvsID"`
					AssociateFVSID string `json:"associateFVSID"`
				} `json:"fvsInfo"`
				MfID      string `json:"mfID"`
				WebSiteID string `json:"webSiteID"`
				OmUTag    string `json:"omUTag"`
				HotRankID string `json:"hotRankID"`
				VideoIDs  struct {
					Vid string `json:"vid"`
					Cid string `json:"cid"`
					Lid string `json:"lid"`
				} `json:"videoIDs"`
				Tab2EventInfo struct {
					EventID   string `json:"event_id"`
					MatchType int    `json:"match_type"`
					IsRec     int    `json:"is_rec"`
				} `json:"tab2_event_info"`
				NewHotEventInfo struct {
					EventID           string `json:"event_id"`
					HotStatus         int    `json:"hot_status"`
					HotSource         int    `json:"hot_source"`
					InterventionLevel int    `json:"intervention_level"`
				} `json:"new_hot_event_info"`
				LegalCheckStatus int `json:"legal_check_status"`
				FavorNum         int `json:"favorNum"`
			} `json:"extern_info"`
			StarRanking struct {
				InRanking   int    `json:"in_ranking"`
				InActivity  int    `json:"in_activity"`
				CallBangImg string `json:"call_bang_img"`
				RankingTips string `json:"ranking_tips"`
			} `json:"starRanking"`
			Tags         []interface{} `json:"tags"`
			CollectionID string        `json:"collectionId"`
			Collection   struct {
				Cid       string `json:"cid"`
				Name      string `json:"name"`
				Cover     string `json:"cover"`
				Desc      string `json:"desc"`
				FeedNum   int    `json:"feedNum"`
				PlayNum   int    `json:"playNum"`
				ShareInfo struct {
					JumpURL string `json:"jump_url"`
					BodyMap struct {
					} `json:"body_map"`
					WxMiniProgram struct {
						WebpageURL       string `json:"webpageUrl"`
						UserName         string `json:"userName"`
						Path             string `json:"path"`
						HdImageDataURL   string `json:"hdImageDataURL"`
						WithShareTicket  int    `json:"withShareTicket"`
						MiniProgramType  int    `json:"miniProgramType"`
						Appid            string `json:"appid"`
						VideoUserName    string `json:"videoUserName"`
						VideoSource      string `json:"videoSource"`
						VideoCoverWidth  int    `json:"videoCoverWidth"`
						VideoCoverHeight int    `json:"videoCoverHeight"`
						AppThumbURL      string `json:"appThumbUrl"`
						UniversalLink    string `json:"universalLink"`
						Disableforward   int    `json:"disableforward"`
					} `json:"wx_mini_program"`
					SqArkInfo struct {
						ArkData   string `json:"arkData"`
						ShareBody struct {
							Title    string `json:"title"`
							Desc     string `json:"desc"`
							ImageURL string `json:"image_url"`
							URL      string `json:"url"`
						} `json:"shareBody"`
						CoverProto string `json:"coverProto"`
					} `json:"sq_ark_info"`
					ShareIconURL   string `json:"share_icon_url"`
					ShareIconTitle string `json:"share_icon_title"`
					BackgroundURL  string `json:"background_url"`
					ActivityType   int    `json:"activity_type"`
					HaibaoJumpURL  string `json:"haibao_jump_url"`
					HaibaoBodyMap  struct {
					} `json:"haibao_body_map"`
					BackgroundTitleColor string `json:"background_title_color"`
					HaibaoDesc           string `json:"haibao_desc"`
					ShareNum             int    `json:"share_num"`
					FeedCoverUpdtime     string `json:"feed_cover_updtime"`
				} `json:"shareInfo"`
				AttachInfo string `json:"attach_info"`
				Poster     struct {
					ID                string `json:"id"`
					Type              int    `json:"type"`
					UID               string `json:"uid"`
					Createtime        int    `json:"createtime"`
					Nick              string `json:"nick"`
					Avatar            string `json:"avatar"`
					Sex               int    `json:"sex"`
					FeedlistTimeID    string `json:"feedlist_time_id"`
					FeedlistHotID     string `json:"feedlist_hot_id"`
					RelatedFeedlistID string `json:"related_feedlist_id"`
					FollowerlistID    string `json:"followerlist_id"`
					InteresterlistID  string `json:"interesterlist_id"`
					ChatlistID        string `json:"chatlist_id"`
					RichFlag          int    `json:"rich_flag"`
					Age               int    `json:"age"`
					Address           string `json:"address"`
					Wealth            struct {
						FlowerNum int `json:"flower_num"`
						Score     int `json:"score"`
					} `json:"wealth"`
					Background        string `json:"background"`
					Status            string `json:"status"`
					FollowStatus      int    `json:"followStatus"`
					ChartScore        int    `json:"chartScore"`
					ChartRank         int    `json:"chartRank"`
					FeedGoldNum       int    `json:"feedGoldNum"`
					AvatarUpdatetime  int    `json:"avatar_updatetime"`
					DescFromOperator  string `json:"desc_from_operator"`
					SyncContent       int    `json:"sync_content"`
					FeedlistPraiseID  string `json:"feedlist_praise_id"`
					Settingmask       int    `json:"settingmask"`
					Originalavatar    string `json:"originalavatar"`
					BlockTime         string `json:"block_time"`
					Grade             int    `json:"grade"`
					Medal             int    `json:"medal"`
					BlockReason       string `json:"block_reason"`
					Qq                int    `json:"qq"`
					RecommendReason   string `json:"recommendReason"`
					LastUpdateFeedNum int    `json:"lastUpdateFeedNum"`
					Updateinfo        struct {
						Flag int    `json:"flag"`
						Tip  string `json:"tip"`
						Num  int    `json:"num"`
					} `json:"updateinfo"`
					NickUpdatetime     int    `json:"nick_updatetime"`
					LastDownloadAvatar string `json:"lastDownloadAvatar"`
					RealName           string `json:"realName"`
					PinyinFirst        string `json:"pinyin_first"`
					CertifDesc         string `json:"certif_desc"`
					PrivateInfo        struct {
						PhoneNum string `json:"phone_num"`
						Name     string `json:"name"`
						IDNum    string `json:"id_num"`
					} `json:"privateInfo"`
					ExternInfo struct {
						MpEx struct {
						} `json:"mpEx"`
						BindAcct  []interface{} `json:"bind_acct"`
						BgPicURL  string        `json:"bgPicUrl"`
						LevelInfo struct {
							Level           int `json:"level"`
							Score           int `json:"score"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"level_info"`
						WeishiID            string `json:"weishiId"`
						WeishiidModifyCount string `json:"weishiid_modify_count"`
						WatermarkType       int    `json:"watermark_type"`
						RealNick            string `json:"real_nick"`
						CmtLevel            struct {
							Level           int `json:"level"`
							Cmtscore        int `json:"cmtscore"`
							Dingscore       int `json:"dingscore"`
							PrevUpgradeTime int `json:"prev_upgrade_time"`
						} `json:"cmt_level"`
						FlexibilityFlag int `json:"flexibility_flag"`
						LiveStatus      int `json:"live_status"`
						NowLiveRoomID   int `json:"now_live_room_id"`
						MedalInfo       struct {
							TotalScore int           `json:"total_score"`
							MedalList  []interface{} `json:"medal_list"`
						} `json:"medal_info"`
						H5HasLogin   int    `json:"h5_has_login"`
						RelationType int    `json:"relation_type"`
						Feedid       string `json:"feedid"`
						CoverType    int    `json:"cover_type"`
						IndustryType int    `json:"industry_type"`
						InditeType   int    `json:"indite_type"`
					} `json:"extern_info"`
					CertifData struct {
						CertifIcon    string `json:"certif_icon"`
						CertifJumpurl string `json:"certif_jumpurl"`
					} `json:"certifData"`
					IsShowPOI    int `json:"isShowPOI"`
					IsShowGender int `json:"isShowGender"`
					FormatAddr   struct {
						Country  string `json:"country"`
						Province string `json:"province"`
						City     string `json:"city"`
					} `json:"formatAddr"`
					AuthorizeTime int `json:"authorize_time"`
					ActivityInfo  struct {
						InvitePersonid string `json:"invitePersonid"`
					} `json:"activityInfo"`
					SpecialIdentity struct {
					} `json:"special_identity"`
					TmpMark      int `json:"tmpMark"`
					PmtMark      int `json:"pmtMark"`
					IndustryInfo struct {
						PrimaryIndustry struct {
							IndustryID   int    `json:"industry_id"`
							IndustryDesc string `json:"industry_desc"`
						} `json:"primary_industry"`
						SecondaryIndustry struct {
							IndustryID   int    `json:"industry_id"`
							IndustryDesc string `json:"industry_desc"`
						} `json:"secondary_industry"`
					} `json:"industry_info"`
					Homeland struct {
						Country  string `json:"country"`
						Province string `json:"province"`
						City     string `json:"city"`
					} `json:"homeland"`
				} `json:"poster"`
				UpdateTime         int    `json:"updateTime"`
				UpdateFeedNum      int    `json:"updateFeedNum"`
				IsFollowed         int    `json:"isFollowed"`
				LikeNum            int    `json:"likeNum"`
				IsHidden           int    `json:"isHidden"`
				CollectionType     int    `json:"collectionType"`
				OrderType          int    `json:"orderType"`
				IsManualCollection int    `json:"isManualCollection"`
				ThemeID            string `json:"themeId"`
				ThemeInfo          struct {
					ThemeID            string `json:"themeId"`
					ThemeName          string `json:"themeName"`
					Category1          string `json:"category1"`
					Category2          string `json:"category2"`
					ThemeType          int    `json:"themeType"`
					ThemeStatus        int    `json:"themeStatus"`
					LatestCollectionID string `json:"latestCollectionId"`
					PosterID           string `json:"posterId"`
					UpdateTime         int    `json:"updateTime"`
					Category1Name      string `json:"category1Name"`
					Category2Name      string `json:"category2Name"`
				} `json:"themeInfo"`
				CollectionSeriesType int    `json:"collectionSeriesType"`
				FeedRelation         int    `json:"feedRelation"`
				SummaryDesc          string `json:"summaryDesc"`
				KeyWord              string `json:"keyWord"`
				Category1            string `json:"category1"`
				Category2            string `json:"category2"`
				ColKeyword           string `json:"colKeyword"`
				IsBusiness           int    `json:"isBusiness"`
				ShowIndexUIType      int    `json:"showIndexUIType"`
				ColAccountType       int    `json:"colAccountType"`
				ColAccountPID        string `json:"colAccountPID"`
			} `json:"collection"`
			MusicBeginTime int `json:"music_begin_time"`
			MusicEndTime   int `json:"music_end_time"`
			MusicInfo      struct {
				AlbumInfo struct {
					UIID    int    `json:"uiId"`
					StrMid  string `json:"strMid"`
					StrName string `json:"strName"`
					StrPic  string `json:"strPic"`
				} `json:"albumInfo"`
				SingerInfo struct {
					UIID         int           `json:"uiId"`
					StrMid       string        `json:"strMid"`
					StrName      string        `json:"strName"`
					StrPic       string        `json:"strPic"`
					StrSchema    string        `json:"strSchema"`
					OtherSingers []interface{} `json:"otherSingers"`
				} `json:"singerInfo"`
				SongInfo struct {
					UIID               int    `json:"uiId"`
					StrMid             string `json:"strMid"`
					StrName            string `json:"strName"`
					StrGenre           string `json:"strGenre"`
					IIsOnly            int    `json:"iIsOnly"`
					StrLanguage        string `json:"strLanguage"`
					IPlayable          int    `json:"iPlayable"`
					ITrySize           int    `json:"iTrySize"`
					ITryBegin          int    `json:"iTryBegin"`
					ITryEnd            int    `json:"iTryEnd"`
					IPlayTime          int    `json:"iPlayTime"`
					StrH5URL           string `json:"strH5Url"`
					StrPlayURL         string `json:"strPlayUrl"`
					StrPlayURLStandard string `json:"strPlayUrlStandard"`
					StrPlayURLHq       string `json:"strPlayUrlHq"`
					StrPlayURLSq       string `json:"strPlayUrlSq"`
					ISize              int    `json:"iSize"`
					ISizeStandard      int    `json:"iSizeStandard"`
					ISizeHq            int    `json:"iSizeHq"`
					ISizeSq            int    `json:"iSizeSq"`
					Copyright          int    `json:"copyright"`
					ISource            int    `json:"iSource"`
					PublicTime         string `json:"publicTime"`
					SongTitle          string `json:"songTitle"`
					SongDescription    string `json:"songDescription"`
					State              int    `json:"state"`
					Deleted            int    `json:"deleted"`
					StartTime          string `json:"startTime"`
					EndTime            string `json:"endTime"`
				} `json:"songInfo"`
				LyricInfo struct {
					UISongID      int    `json:"uiSongId"`
					StrSongMid    string `json:"strSongMid"`
					StrFormat     string `json:"strFormat"`
					StrLyric      string `json:"strLyric"`
					StrMatchLyric string `json:"strMatchLyric"`
				} `json:"lyricInfo"`
				ConfInfo struct {
					IType               int    `json:"iType"`
					IStartPos           int    `json:"iStartPos"`
					StrLabel            string `json:"strLabel"`
					IsCollected         int    `json:"isCollected"`
					CollectTime         int    `json:"collectTime"`
					Exclusive           int    `json:"exclusive"`
					FollowFeed          string `json:"followFeed"`
					UseCount            int    `json:"useCount"`
					TogetherFeed        string `json:"togetherFeed"`
					TogetherType        int    `json:"togetherType"`
					FeedUseType         int    `json:"feedUseType"`
					DefaultFeedPosition int    `json:"defaultFeedPosition"`
					DefaultTogetherFeed int    `json:"defaultTogetherFeed"`
					BubbleStartTime     int    `json:"bubbleStartTime"`
					BubbleEndTime       int    `json:"bubbleEndTime"`
					SongLabels          struct {
					} `json:"songLabels"`
					SongLabelCategory struct {
					} `json:"songLabelCategory"`
					IsStuckPoint      bool   `json:"isStuckPoint"`
					StuckPointJSONURL string `json:"stuckPointJsonUrl"`
					TrackBeatInfo     struct {
						TrackBeatFinished int `json:"trackBeatFinished"`
						Vocal             struct {
							JSONURL      string `json:"jsonURL"`
							AudioFileURL string `json:"audioFileURL"`
						} `json:"vocal"`
						Drums struct {
							JSONURL      string `json:"jsonURL"`
							AudioFileURL string `json:"audioFileURL"`
						} `json:"drums"`
						Accompaniment struct {
							JSONURL      string `json:"jsonURL"`
							AudioFileURL string `json:"audioFileURL"`
						} `json:"accompaniment"`
						Bass struct {
							JSONURL      string `json:"jsonURL"`
							AudioFileURL string `json:"audioFileURL"`
						} `json:"bass"`
					} `json:"trackBeatInfo"`
					ExtraInfo string `json:"extraInfo"`
				} `json:"confInfo"`
				SubtitleInfo struct {
					UISongID      int    `json:"uiSongId"`
					StrSongMid    string `json:"strSongMid"`
					StrFormat     string `json:"strFormat"`
					StrLyric      string `json:"strLyric"`
					StrMatchLyric string `json:"strMatchLyric"`
				} `json:"subtitleInfo"`
				Foreignlyric struct {
					UISongID      int    `json:"uiSongId"`
					StrSongMid    string `json:"strSongMid"`
					StrFormat     string `json:"strFormat"`
					StrLyric      string `json:"strLyric"`
					StrMatchLyric string `json:"strMatchLyric"`
				} `json:"foreignlyric"`
				RecommendInfo struct {
					TraceStr      string `json:"traceStr"`
					AnalyseResult string `json:"analyse_result"`
					RecomReason   string `json:"recom_reason"`
				} `json:"recommendInfo"`
				UnplayableInfo struct {
					UnplayableCode int    `json:"unplayableCode"`
					UnplayableMsg  string `json:"unplayableMsg"`
				} `json:"unplayableInfo"`
				LabelInfo       []interface{} `json:"labelInfo"`
				MusicType       int           `json:"musicType"`
				MusicSrcType    int           `json:"musicSrcType"`
				CacheUpdateTime int           `json:"cacheUpdateTime"`
				State           int           `json:"state"`
			} `json:"music_info"`
			Header struct {
				Active            int    `json:"active"`
				Type              int    `json:"type"`
				Title             string `json:"title"`
				Jumpurl           string `json:"jumpurl"`
				ContentBackground string `json:"content_background"`
				Logo              string `json:"logo"`
				ActivityIcon      string `json:"activity_icon"`
				LeftBackground    string `json:"left_background"`
				Traceid           string `json:"traceid"`
			} `json:"header"`
			RewardInfo struct {
				InRewarding int `json:"in_rewarding"`
			} `json:"rewardInfo"`
			NearbyfeedCoverUrls    []interface{} `json:"nearbyfeed_cover_urls"`
			FingerprintCheckStatus int           `json:"fingerprint_check_status"`
			LiveInfo               struct {
				LiveStatus   int           `json:"live_status"`
				RoomSchema   string        `json:"room_schema"`
				RoomCoverURL string        `json:"room_cover_url"`
				RoomID       int           `json:"room_id"`
				RoomTitle    string        `json:"room_title"`
				AnchorID     int           `json:"anchor_id"`
				AnchorName   string        `json:"anchor_name"`
				AnchorIcon   string        `json:"anchor_icon"`
				StreamInfo   []interface{} `json:"stream_info"`
				ProgramID    string        `json:"program_id"`
				SugFormat    string        `json:"sug_format"`
				FeedStyle    struct {
					Slogan     string `json:"slogan"`
					SloganIcon string `json:"slogan_icon"`
					SloganType int    `json:"slogan_type"`
				} `json:"feed_style"`
				SlideID string `json:"slide_id"`
			} `json:"live_info"`
			FeedTags []interface{} `json:"feed_tags"`
			RichDing struct {
				EmojiListDefault []struct {
					ID      string `json:"id"`
					Count   int    `json:"count"`
					IsLiked bool   `json:"isLiked"`
				} `json:"emojiListDefault"`
				EmojiListEdit []struct {
					ID      string `json:"id"`
					Count   int    `json:"count"`
					IsLiked bool   `json:"isLiked"`
				} `json:"emojiListEdit"`
				ShowThreshold     int `json:"showThreshold"`
				H5RichDingDisplay struct {
					URL   string `json:"url"`
					Text  string `json:"text"`
					Count int    `json:"count"`
				} `json:"h5RichDingDisplay"`
			} `json:"richDing"`
			Category    string `json:"category"`
			SubCategory string `json:"sub_category"`
			ContentTags []struct {
				Name string `json:"name"`
			} `json:"content_tags"`
			LongVideoInfo struct {
				IsAssociateLongVideo bool `json:"is_associate_long_video"`
				Type                 int  `json:"type"`
				VideoIds             struct {
					Vid string `json:"vid"`
					Cid string `json:"cid"`
					Lid string `json:"lid"`
				} `json:"video_ids"`
				ContentID  string `json:"contentID"`
				ManualBind int    `json:"manual_bind"`
				ManualInfo struct {
					Xtype   int           `json:"xtype"`
					ID      string        `json:"id"`
					Cuts    []interface{} `json:"cuts"`
					Similar int           `json:"similar"`
				} `json:"manual_info"`
				AutoBind int `json:"auto_bind"`
				AutoInfo struct {
					Binds []interface{} `json:"binds"`
				} `json:"auto_info"`
			} `json:"long_video_info"`
			HorizontalVideo int `json:"horizontalVideo"`
			CopyrightMask   int `json:"copyrightMask"`
		} `json:"feeds"`
		Isdeleted      int           `json:"isdeleted"`
		Recommendfeeds []interface{} `json:"recommendfeeds"`
		Errmsg         string        `json:"errmsg"`
		CopyRightInfo  struct {
			CopyRights struct {
				SevenWkzwZSDL1On6V1A2 struct {
					Ret      int    `json:"ret"`
					Msg      string `json:"msg"`
					Mask     int    `json:"mask"`
					HasRight bool   `json:"hasRight"`
				} `json:"7wkzwZSDL1On6v1A2"`
			} `json:"copyRights"`
		} `json:"copyRightInfo"`
		FeedResultInfo struct {
			FeedResults struct {
				SevenWkzwZSDL1On6V1A2 struct {
					Ret int    `json:"ret"`
					Msg string `json:"msg"`
				} `json:"7wkzwZSDL1On6v1A2"`
			} `json:"feedResults"`
		} `json:"feedResultInfo"`
		Idc string `json:"_idc"`
	} `json:"data"`
}

func (a WeishiAdapter) GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Sugar.Error(r)
		}
	}()
	re := regexp.MustCompile(`id=(.*)&spid`)
	match := re.FindStringSubmatch(url)
	if len(match) != 2 {
		return nil, errors.New("匹配视频id失败")
	}
	getResp, err := restyHttp.GetMobileHttpRequest().Get("https://h5.weishi.qq.com/webapp/json/weishi/WSH5GetPlayPage?feedid=" + match[1])
	if err != nil {
		return nil, err
	}
	var weishiResponse WeishiAdapterResponse
	err = json.Cjson.Unmarshal(getResp.Body(), &weishiResponse)
	if err != nil {
		return nil, err
	}
	shortVideoInfo := &ShortVideoInfoResponse{
		AuthorName:             weishiResponse.Data.Feeds[0].Poster.Nick,
		AuthorAvatar:           weishiResponse.Data.Feeds[0].Poster.Avatar,
		Title:                  weishiResponse.Data.Feeds[0].FeedDesc,
		Cover:                  weishiResponse.Data.Feeds[0].Images[0].URL,
		NoWatermarkDownloadUrl: weishiResponse.Data.Feeds[0].VideoURL,
	}
	return shortVideoInfo, nil
}
