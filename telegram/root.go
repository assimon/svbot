package telegram

import (
	"github.com/assimon/svbot/internal/config"
	"github.com/assimon/svbot/internal/log"
	tb "gopkg.in/telebot.v3"
	"time"
)

var (
	TeleBot *tb.Bot
)

func Start() {
	var err error
	setting := tb.Settings{
		Token:   config.GetTelegramInstance().ApiToken,
		Updates: 500,
		Poller:  &tb.LongPoller{Timeout: 60 * time.Second},
		OnError: func(err error, context tb.Context) {
			log.Sugar.Error(err)
		},
	}
	TeleBot, err = tb.NewBot(setting)
	if err != nil {
		panic(err)
	}
	RegisterHandle()
	TeleBot.Start()
}

func RegisterHandle() {
	global := TeleBot.Group()
	global.Use(func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			if !c.Message().Private() {
				return nil
			}
			return next(c)
		}
	})
	global.Handle(tb.OnText, OnTextHandle)
	global.Handle("/start", func(c tb.Context) error {
		return c.Send("👋🏻Hi，我目前支持：|抖音/火山/快手/绿洲/皮皮虾/微博/微视/西瓜/最右|的短视频无水印解析，请发送短视频链接给我吧")
	})
}
