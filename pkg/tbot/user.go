package tbot

import "golang.org/x/exp/slog"

type User map[string]int64

func (u User) GetUser(id int64) string {
	slog.Debug("Get (user) name", "id", id, "u", u)
	for user, nr := range u {
		if nr == id {
			return user
		}
	}

	return ""
}
