package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mcping/mcping"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	read := bufio.NewScanner(os.Stdin)
	b := getjson()
	fmt.Println("此程序会尝试测试 mojang 正版验证相关的网站的 ip 的速度，然后尝试修改 hosts 中的内容来解决无法正版登录之类的问题。")
	fmt.Println("如果需要使用自动设置 hosts 的功能，请以右键以管理员身份运行")
	fmt.Print("按回车键继续")
	read.Scan()
	fmt.Println("正在测试，耐心等待")
	u := mcping.JSON(b)
	m := make(map[string]string)
	for k, v := range u {
		ip, atime, err := mcping.Test(k, v)
		if err != nil {
			fmt.Print(k, ":")
			fmt.Println("所有 ip 均不可用")
		} else {
			if ip == "" {
				fmt.Println(k, "无需更改")
			} else {
				fmt.Println(k, ": 测试所有 ip 中延迟最低的为", ip, "延迟为", atime)
			}
			m[k] = ip
		}
	}
	fmt.Println("测试完毕，按下回车键将尝试更改 hosts 。会尝试将已有的 hosts 备份，可能导致文件损坏的后果请自行承担。")
	read.Scan()
	w := bytes.NewBuffer(nil)
	w.WriteString("\n")
	for k, v := range m {
		if v != "" {
			s := strings.Split(k, "/")
			w.WriteString(v)
			w.WriteString(" ")
			w.WriteString(s[2])
			w.WriteString("\n")
		}
	}
	err := write(w.Bytes())
	if err != nil {
		fmt.Println(err)
		fmt.Println("设置失败，请尝试右键以管理员身份运行")
		fmt.Println("文件保存在此程序同一目录下，可自行查阅有关资料自行设置")
		fff, err := os.Create(`hosts`)
		defer fff.Close()
		_, err = fff.Write(w.Bytes())
		if err != nil {
			fmt.Println(err)
			fmt.Println("依然保存失败，你的电脑有问题。")
			fmt.Println(w.String())
		}
		read.Scan()
	} else {
		fmt.Println("设置成功")
		cmd := exec.Command("ipconfig", "/flushdns")
		cmd.Run()
		read.Scan()
	}
}

func getjson() []byte {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	rep, err := http.NewRequest("GET", "https://ping.xmdhs.top/ip.json", nil)
	reps, err := c.Do(rep)
	var b []byte
	if err != nil {
		fmt.Println(err)
		fmt.Println("大概是网络问题，使用内置 ip 列表")
		return nil
	}
	b, err = ioutil.ReadAll(reps.Body)
	defer reps.Body.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Println("大概是网络问题，使用内置 ip 列表")
		return nil
	}
	return b
}

func write(b []byte) error {
	f, err := os.Open(`C:\Windows\System32\drivers\etc\hosts`)
	defer f.Close()
	if err != nil {
		return err
	}
	ff, err := os.Create(`C:\Windows\System32\drivers\etc\hosts.mcping.bak`)
	defer ff.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(ff, f)
	if err != nil {
		return err
	}
	fff, err := os.OpenFile(`C:\Windows\System32\drivers\etc\hosts`, os.O_APPEND|os.O_WRONLY, 777)
	defer fff.Close()
	if err != nil {
		return err
	}
	_, err = fff.Write(b)
	if err != nil {
		return err
	}
	return nil
}
