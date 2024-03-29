package mcping

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func ping(url, ip string) (int64, error) {
	var c http.Client
	transport := http.DefaultTransport.(*http.Transport).Clone()
	if ip != "" {
		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			_, port, err := net.SplitHostPort(addr)
			if err != nil {
				panic(err)
			}
			return dialer.DialContext(ctx, network, net.JoinHostPort(ip, port))
		}
	}
	c = http.Client{
		Timeout:   5 * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	t := time.Now().UnixNano()
	h, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	h.Header.Set("Accept", "*/*")
	h.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	h.Close = true
	rep, err := c.Do(h)
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return 0, fmt.Errorf("ping: %w", err)
	}
	_, err = io.Copy(ioutil.Discard, rep.Body)
	if err != nil {
		return 0, fmt.Errorf("ping: %w", err)
	}
	return time.Now().UnixNano() - t, nil
}
