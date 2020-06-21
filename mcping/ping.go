package mcping

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func Ping(url, ip string) (int64, error) {
	var c http.Client
	if ip != "" {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			s := strings.Split(addr, ":")
			addr = strings.ReplaceAll(addr, s[0], ip)
			return dialer.DialContext(ctx, network, addr)
		}
		c = http.Client{
			Timeout:   5 * time.Second,
			Transport: transport,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	} else {
		c = http.Client{
			Timeout: 5 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}
	t := time.Now().UnixNano()
	h, err := http.NewRequest("GET", "https://"+url, nil)
	rep, err := c.Do(h)
	if err != nil {
		return 0, err
	}
	defer rep.Body.Close()
	_, err = ioutil.ReadAll(rep.Body)
	if err != nil {
		return 0, err
	}
	return time.Now().UnixNano() - t, nil
}
