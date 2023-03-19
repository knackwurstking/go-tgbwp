package tbot

import (
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"golang.org/x/exp/slog"
)

var (
	DefaultBotOpts = gotgbot.BotOpts{
		Client: http.Client{},
	}

	DefaultUpdaterOpts = ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				slog.Error("Bot Dispatcher", "error", err)
				return ext.DispatcherActionNoop
			},
			ErrorLog:    nil,
			MaxRoutines: ext.DefaultMaxRoutines,
		},
	}
)

type Bot struct {
	*gotgbot.Bot
	Updater ext.Updater

	ID ID // ID contains user and chat ids ("~/.config/tgbwp/config.json")

	Commands      []gotgbot.BotCommand
	GroupCommands []gotgbot.BotCommand
}

func NewBot(token string, botOpts gotgbot.BotOpts, updaterOpts ext.UpdaterOpts) (*Bot, error) {
	bot, err := gotgbot.NewBot(token, &botOpts)
	if err != nil {
		return nil, err
	}

	updater := ext.NewUpdater(&updaterOpts)

	return &Bot{
		Bot:     bot,
		Updater: updater,
	}, err
}

func (bot *Bot) Dispatcher() *ext.Dispatcher {
	return bot.Updater.Dispatcher
}

func (bot *Bot) RegisterCommand(command string, handler func(bot *gotgbot.Bot, ctx *ext.Context) error) {
	dispatcher := bot.Dispatcher()
	dispatcher.AddHandler(handlers.NewCommand(command, handler))
}

func (bot *Bot) RegisterCallback(filter filters.CallbackQuery, handler func(bot *gotgbot.Bot, ctx *ext.Context) error) {
	dispatcher := bot.Dispatcher()
	dispatcher.AddHandler(handlers.NewCallback(filter, handler))
}

func (bot *Bot) AddCommand(command, desc string) {
	bot.Commands = append(bot.Commands, gotgbot.BotCommand{
		Command:     command,
		Description: desc,
	})
}

func (b *Bot) SetCommands() (ok bool, err error) {
	return b.SetMyCommands(b.Commands, &gotgbot.SetMyCommandsOpts{
		Scope: gotgbot.BotCommandScopeDefault{},
	})
}

func (bot *Bot) AddGroupCommand(command, desc string) {
	bot.GroupCommands = append(bot.GroupCommands, gotgbot.BotCommand{
		Command:     command,
		Description: desc,
	})
}

func (b *Bot) SetGroupCommands() (ok bool, err error) {
	return b.SetMyCommands(b.GroupCommands, &gotgbot.SetMyCommandsOpts{
		Scope: gotgbot.BotCommandScopeAllGroupChats{},
	})
}

func (bot *Bot) Start() error {
	return bot.Updater.StartPolling(bot.Bot, &ext.PollingOpts{
		DropPendingUpdates: true,
	})
}

func (bot *Bot) Wait() {
	bot.Updater.Idle()
}
