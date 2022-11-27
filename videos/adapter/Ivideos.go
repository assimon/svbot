package adapter

type IVideosInterface interface {
	GetShortVideoInfo(url string) (*ShortVideoInfoResponse, error)
}

type ShortVideoInfoResponse struct {
	AuthorName             string `json:"author_name"`
	AuthorAvatar           string `json:"author_avatar"`
	Title                  string `json:"title"`
	Cover                  string `json:"cover"`
	CreatedAt              string `json:"created_at"`
	NoWatermarkDownloadUrl string `json:"no_watermark_download_url"`
}
