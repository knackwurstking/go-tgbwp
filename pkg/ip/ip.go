package ip

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Ipify struct {
	Response *http.Response `json:"-"`
	IP       string         `json:"ip"`
}

func (i *Ipify) GetIP() error {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		i.Response = resp
		i.IP = strings.Trim(string(data), "\n ")
		return nil
	}

	return fmt.Errorf("status: %v\n", resp.Status)
}
