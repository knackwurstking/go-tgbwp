package tbot

type User map[string]int64

func (u *User) GetUser(id int64) string {
	for u, n := range *u {
		if n == id {
			return u
		}
	}

	return ""
}
