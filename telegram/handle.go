package telegram

import (
	"github.com/assimon/svbot/videos"
	tb "gopkg.in/telebot.v3"
	"regexp"
	"sync"
)

var (
	CurrentLimitingLockMap sync.Map
)

func OnTextHandle(c tb.Context) error {
	_, ok := CurrentLimitingLockMap.Load(c.Sender().ID)
	if ok {
		return c.Send("请等待上次解析任务完成⏰")
	}
	text := c.Text()
	if text == "/start" {
		return nil
	}
	re := regexp.MustCompile("http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\\(\\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+")
	urls := re.FindStringSubmatch(text)
	if len(urls) <= 0 {
		return c.Send("❌未检测到短视频地址，请检查")
	}
	shortVideoUri := urls[0]
	adapter := videos.GetShortVideoAdapter(shortVideoUri)
	if adapter == nil {
		return c.Send("❌未匹配到该视频链接解析器，无法解析~")
	}
	CurrentLimitingLockMap.Store(c.Sender().ID, 1)
	defer CurrentLimitingLockMap.Delete(c.Sender().ID)
	c.Send("🎬视频解析中，请等待.....")
	videoInfo, err := adapter.GetShortVideoInfo(shortVideoUri)
	if err != nil {
		return c.Send("😭解析异常，服务器内部错误~ 请稍后重试")
	}
	if videoInfo.NoWatermarkDownloadUrl == "" {
		return c.Send("😭未能成功解析视频地址，可能是引擎失效或服务器异常~")
	}
	return c.Send("🎉解析成功，直链地址：" + videoInfo.NoWatermarkDownloadUrl)
}
