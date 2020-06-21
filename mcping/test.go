package mcping

import (
	"errors"
	"sort"
	"sync"
)

func Test(url string, ip []string) (string, int64, error) {
	ch := make(chan bool, 10)
	ti := make(chan iptime, 3)
	m := make(map[string]int64)
	ip = append(ip, "")
	t := testip{
		ch:  ch,
		ips: ti,
		url: url,
		m:   m,
	}

	for _, v := range ip {
		t.goget(v)
	}
	go func() {
		t.Wait()
		close(ti)
	}()
	t.tim()
	aip, err := tosort(t.m)
	if err != nil {
		return "", 0, err
	}
	return aip.k, aip.v, nil
}

type testip struct {
	ch  chan bool
	ips chan iptime
	url string
	sync.WaitGroup
	m map[string]int64
}

type iptime struct {
	ip    string
	atime []int64
}

func (ti *testip) goget(ip string) {
	ti.ch <- true
	ti.Add(1)
	go func() {
		tt := make([]int64, 0, 3)
		for i := 0; i < 5; i++ {
			t, err := Ping(ti.url, ip)
			if err != nil {
				t = 99999999
			}
			tt = append(tt, t)
		}
		iptime := iptime{
			ip:    ip,
			atime: tt,
		}
		ti.ips <- iptime
		ti.Done()
		<-ti.ch
	}()
}

func (ti *testip) tim() {
	for {
		iptime, ok := <-ti.ips
		if ok {
			var u int64
			uu := len(iptime.atime)
			uu = uu / 2
			u = iptime.atime[uu]
			ti.m[iptime.ip] = u
		} else {
			break
		}
	}
}

func tosort(m map[string]int64) (tosorts, error) {
	t := make([]tosorts, 0, len(m))
	for k, v := range m {
		t = append(t, tosorts{
			k: k,
			v: v,
		})
	}
	sort.Slice(t, func(i, j int) bool { return t[i].v < t[j].v })
	if t[0].v == 99999999 {
		return t[0], errors.New("can not get ip")
	}
	return t[0], nil
}

type tosorts struct {
	k string
	v int64
}
