package mcping

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	ips := []string{
		"122.246.10.93",
		"119.167.216.140",
		"121.12.123.201",
	}
	m, err := test("https://www.mcbbs.net/", ips)
	if err != nil {
		t.Fatal(err)
	}
	if m == "" {
		fmt.Println("无需更改")
	}
	fmt.Println(m)
}
