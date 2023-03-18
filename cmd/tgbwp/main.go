package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/exp/slog"

	"github.com/knackwurstking/go-tgbwp/pkg/plugins"
	"github.com/knackwurstking/go-tgbwp/pkg/tbot"
)

var (
	ApplicationName = "tgbwp"
	PluginsPath     = getPluginsPath()
)

type Config struct {
	ID tbot.ID `json:"id"`
}

func init() {
	flag.StringVar(&PluginsPath, "p", PluginsPath,
		"change the default plugins path")

	// Custom usage string (info about all the env vars)
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`usage: %s [OPTIONS]

Environment:
  export TGBWP_TOKEN=<api-token>
  export TGBWP_CHAT_ID=<telegram-chat-id1,...> (multiple possible)
  export TGBWP_USER_ID=<telegram-user-id1,...> (multiple possible)

Options:
`,
			ApplicationName,
		)

		flag.PrintDefaults()
	}

	flag.Parse()

	// Setup logger
	o := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if os.Getenv("DEBUG") == "true" {
		o.Level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(o.NewTextHandler(os.Stderr)))
}

func main() {
	var (
		bot *tbot.Bot
		err error
	)

	for {
		// Get a new bot
		bot, err = tbot.NewBot(
			os.Getenv("TGBWP_TOKEN"),
			tbot.DefaultBotOpts,
			tbot.DefaultUpdaterOpts,
		)
		if err != nil {
			// Try again on a temp. failure error (sleep 10s before try again)
			if strings.Contains(err.Error(), "Temporary failure in name resolution") {
				time.Sleep(time.Second * 10)
				continue
			}

			// Error and exit
			slog.Error("Creating a bot failed!",
				"error", err)
			os.Exit(1)
		}

		break
	}

	// load config from "~/.config/tgbwp/config.json" (of "config.json")
	config, err := readConfigJSON()
	if err != nil {
		slog.Warn("Failed to read user config!",
			"error", err)
	}
	bot.ID = config.ID

	// plugin manager - load *.so files (plugins) get plugin struct (type Plugins)...
	pm := plugins.NewManager(PluginsPath)

	slog.Info("Plugins path: \"%s\"", pm.Path)

	for _, p := range pm.List() {
		slog.Info(fmt.Sprintf("Load plugin: \"%s\"", p))

		if err = pm.Register(p, bot); err != nil {
			slog.Warn(fmt.Sprintf("Load plugin \"%s\" failed!", p),
				"error", err)
		}
	}

	if err = bot.Start(); err != nil {
		slog.Error("Starting bot failed!",
			"error", err)
		os.Exit(1)
	}

	setMyCommands(bot)

	bot.Wait()
}

func getPluginsPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("Failed to get user home dir!",
			"error", err)
		os.Exit(1)
	}

	return path.Join(home, ".local", "share", ApplicationName, "plugins")
}

func readConfigJSON() (Config, error) {
	var c Config

	// Get reader (os just read the file), parse json and return.
	home, err := os.UserHomeDir()
	if err != nil {
		return c, fmt.Errorf("Failed to get user home dir: %s", err.Error())
	}

	f, err := os.Open("config.json")
	if err != nil {
		f, err = os.Open(filepath.Join(home, ".config", ApplicationName, "config.json"))
	}
	err = json.NewDecoder(f).Decode(&c)

	return c, err
}

func setMyCommands(b *tbot.Bot) {
	// Set global scoped bot commands
	ok, err := b.SetCommands()
	if err != nil {
		slog.Error("Set global scoped commands failed!", "error", err)
	} else if !ok {
		slog.Warn("Set global scoped commands failed!")
	}

	// Set all group scoped bot commands
	ok, err = b.SetGroupCommands()
	if err != nil {
		slog.Error("Set global group scoped commands failed!", "error", err)
	} else if !ok {
		slog.Warn("Set global group scoped commands failed!")
	}

	// Set chat scoped bot commands
	err = b.SetUserCommands()
	if err != nil {
		slog.Warn(err.Error())
	}

	// Set user scoped bot commands
	err = b.SetChatCommands()
	if err != nil {
		slog.Warn(err.Error())
	}
}
