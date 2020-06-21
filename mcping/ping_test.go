package mcping

import (
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	list := map[string]string{
		"www.mcbbs.net/":    "110.53.246.70",
		"www.minecraft.net": "99.84.227.53",
		"xmdhs.top/":        "110.53.246.70", //Fail
	}
	for k, v := range list {
		i, err := Ping(k, v)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(i)
	}

}
