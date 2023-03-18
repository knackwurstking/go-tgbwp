package main

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"golang.org/x/exp/slog"

	"github.com/knackwurstking/go-tgbwp/pkg/ip"
	"github.com/knackwurstking/go-tgbwp/pkg/tbot"
)

// Plugin struct struct will be load from the tgbwp plugin manager
type plugin struct {
	*tbot.Bot

	ipify ip.Ipify
}

func (p *plugin) Init(b *tbot.Bot) {
	p.Bot = b // NOTE: Always need to do this for all plugins
	slog.Debug("ip: plugin initialized")
}

// Load the plugin - register commands ...
func (p *plugin) Register() error {
	// register ip command handler...
	p.Bot.RegisterCommand("ip", func(bot *gotgbot.Bot, ctx *ext.Context) error {
		slog.Info("Register plugin: /ip " + strings.Join(ctx.Args()[1:], " "))

		// first check if id is in chat ids
		if p.ID.Chat.GetName(ctx.EffectiveChat.Id) != "" {
			return fmt.Errorf("disallowed, invalid chat id! (%d) for user %d",
				ctx.EffectiveChat.Id, ctx.EffectiveUser.Id)
		}

		// get the ip (return error if not nil)
		if err := p.ipify.GetIP(); err != nil {
			return err
		}

		// send message back
		_, err := ctx.EffectiveMessage.Reply(
			bot, fmt.Sprintf("`%s`", p.ipify.IP),
			&gotgbot.SendMessageOpts{
				ParseMode: "MarkdownV2",
			},
		)
		if err != nil {
			return fmt.Errorf("reply: %s", err.Error())
		}

		return nil
	})

	p.Bot.AddChatCommand("ip", "get (rpi-)server ip")

	return nil
}

var Plugin plugin
