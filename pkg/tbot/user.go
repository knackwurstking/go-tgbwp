package tbot

type User map[string]int64

func (u *User) GetUser(id int64) string {
	for u, nr := range *u {
		if nr == id {
			return u
		}
	}

	return ""
}
