package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/knackwurstking/go-tgbwp/pkg/tbot"
	"golang.org/x/exp/slog"
)

var (
	prefix = "test: "
)

type plugin struct {
	*tbot.Bot
}

func (p *plugin) Init(b *tbot.Bot) {
	p.Bot = b
	slog.Debug(prefix + "plugin initialized")
}

func (p *plugin) Register() error {
	p.Bot.RegisterCommand("testpirgb", func(bot *gotgbot.Bot, ctx *ext.Context) error {
		_, err := ctx.EffectiveMessage.Reply(bot, "Use me daddy!", &gotgbot.SendMessageOpts{
			ParseMode: "HTML",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Press me", WebApp: &gotgbot.WebAppInfo{
						Url: "https://rpi-server:50832/sections",
					}},
				}},
			},
		})

		return err
	})

	p.Bot.AddGroupCommand("testpirgb", "open pirgb web app")

	return nil
}

var Plugin plugin
