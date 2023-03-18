package plugins

import (
	"fmt"
	"os"
	"path"
	"plugin"
	"regexp"

	"github.com/knackwurstking/go-tgbwp/pkg/tbot"
)

type Plugins interface {
	Init(bot *tbot.Bot)
	Register() error
}

// PluginManager handles plugins
type Manager struct {
	Path string // path to plugins to manage
}

func NewManager(pluginsPath string) *Manager {
	return &Manager{
		Path: pluginsPath,
	}
}

func (pm *Manager) List() (plugins []string) {
	dir, err := os.ReadDir(pm.Path)
	if err != nil {
		return
	}

	re := regexp.MustCompile(`.*\.so`)
	for _, entry := range dir {
		if !entry.IsDir() {
			// check file type (*.so)
			if re.MatchString(entry.Name()) {
				plugins = append(plugins, entry.Name())
			}
		}
	}

	return
}

func (pm *Manager) Register(pluginPath string, bot *tbot.Bot) error {
	p, err := plugin.Open(path.Join(pm.Path, pluginPath))
	if err != nil {
		return err
	}

	v, err := p.Lookup("Plugin")
	if err != nil {
		return err
	}

	plug, ok := v.(Plugins)
	if !ok {
		return fmt.Errorf("%s: Plugin is not from type Plugins (interface)", pluginPath)
	}

	plug.Init(bot)
	return plug.Register()
}
