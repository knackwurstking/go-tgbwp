package main

import (
	"github.com/knackwurstking/go-tgbwp/pkg/tbot"
	"golang.org/x/exp/slog"
)

type plugin struct {
	*tbot.Bot
}

func (p *plugin) Init(b *tbot.Bot) {
	p.Bot = b
	slog.Debug("test-menu: plugin initialized")
}

func (p *plugin) Register() error {

	return nil
}

var Plugin plugin
