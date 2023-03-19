package tbot

import "golang.org/x/exp/slog"

type Chat map[string]int64

func (c Chat) GetName(id int64) string {
	slog.Debug("Get (chat) name", "id", id, "c", c)
	for u, nr := range c {
		if nr == id {
			return u
		}
	}

	return ""
}
