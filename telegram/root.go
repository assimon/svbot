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
	defer func() {
		if err := recover(); err != nil {
			log.Sugar.Error(err)
		}
	}()
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
		return c.Send("ðð»Hiï¼æç®åæ¯æï¼|æé³/ç«å±±/å¿«æ/ç»¿æ´²/ç®ç®è¾/å¾®å/å¾®è§/è¥¿ç/æå³|çç­è§é¢æ æ°´å°è§£æï¼è¯·åéç­è§é¢é¾æ¥ç»æå§")
	})
}
