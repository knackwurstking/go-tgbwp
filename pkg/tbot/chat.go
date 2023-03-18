package tbot

type Chat map[string]int64

func (u *Chat) GetName(id int64) string {
	for u, n := range *u {
		if n == id {
			return u
		}
	}

	return ""
}
