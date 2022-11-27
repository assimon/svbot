package videos

import (
	"github.com/assimon/svbot/videos/adapter"
	"strings"
)

func GetShortVideoAdapter(url string) adapter.IVideosInterface {
	if strings.Contains(url, "pipix") {
		return adapter.PipixiaAdapter{}
	}
	if strings.Contains(url, "douyin") {
		return adapter.DouyinAdapter{}
	}
	if strings.Contains(url, "huoshan") {
		return adapter.HuoshanAdapter{}
	}
	if strings.Contains(url, "h5.weishi") || strings.Contains(url, "isee.weishi") {
		return adapter.WeishiAdapter{}
	}
	if strings.Contains(url, "weibo.com") {
		return adapter.WeiboAdapter{}
	}
	if strings.Contains(url, "oasis.weibo") {
		return adapter.LvZhouAdapter{}
	}
	if strings.Contains(url, "zuiyou") {
		return adapter.ZuiyouAdapter{}
	}
	if strings.Contains(url, "kuaishou") {
		return adapter.KuaiShouAdapter{}
	}
	if strings.Contains(url, "ixigua.com") {
		return adapter.XiGuaAdapter{}
	}
	return nil
}
