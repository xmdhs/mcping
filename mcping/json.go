package mcping

import "encoding/json"

const j = `{"urlip":{"https://minecraft.net":["99.84.227.53","54.192.172.51","99.84.227.49","13.227.231.59","13.225.123.59","99.84.227.63","13.224.155.57","13.227.48.59","13.249.163.64","13.249.163.58","13.227.231.41","99.86.219.58","13.225.171.39","13.225.171.63","13.225.92.35","99.86.203.39","13.227.231.37","99.86.203.65","13.227.48.35","13.249.163.49","99.86.219.36","13.227.48.49","13.227.48.44","99.84.227.44","52.85.191.254","13.227.231.55","13.249.163.48","13.225.171.40","13.224.155.51","99.86.219.38","13.225.92.48","13.224.155.60","13.33.200.48","54.230.150.51","13.227.55.60","13.225.187.36","13.225.92.58","13.33.200.56","13.33.200.55","13.224.155.36","13.227.55.46","13.35.91.57","99.86.171.51","13.225.187.55","13.225.92.51"],"https://textures.minecraft.net":["143.204.131.92","143.204.131.76","143.204.131.87","13.33.244.75","143.204.131.70","13.33.244.78","13.227.21.20","13.226.217.110","52.85.242.104","13.249.11.68","13.226.217.117","54.192.73.86","13.35.242.12","13.225.146.2","13.225.146.71","13.226.217.54","143.204.222.48","13.225.253.20","13.225.146.68","13.224.13.35","13.224.95.120","13.225.198.113","13.227.21.105","13.249.109.106","13.33.235.22","52.84.140.65","13.224.13.30"],"https://authserver.mojang.com":["99.84.227.68","13.225.107.68","52.85.117.68","13.225.155.69","13.33.200.67","13.35.41.67","99.86.219.68","13.225.92.68","13.33.8.69","52.85.191.41","54.192.172.69","13.224.155.68","99.86.203.68","13.225.187.68","54.230.84.68","54.230.150.67","13.35.91.67","13.227.64.68","13.224.251.68","13.35.25.67"],"https://api.mojang.com":["13.35.40.174","13.35.166.182","143.204.132.179","13.225.125.168","52.85.229.95","13.224.146.169","99.86.207.169","143.204.83.175","13.224.153.169","99.84.138.181","13.33.181.97","52.85.192.79","13.225.98.168","13.225.199.168","99.84.198.181","13.227.74.103","54.192.20.105","13.227.251.106","13.249.111.180","13.35.31.176","54.192.148.176"],"https://sessionserver.mojang.com":["99.84.233.186","13.249.156.175","13.227.59.108","13.225.125.168","13.35.166.182","13.224.153.169","13.249.174.171","13.225.98.168","52.85.192.79","99.84.138.181","52.85.229.95","13.225.199.168","13.35.31.176","99.86.207.169","54.192.148.176","54.192.20.105","13.227.74.103","13.227.251.106","99.84.198.181"]}}`

type urlip struct {
	IP map[string][]string `json:"urlip"`
}

func JSON(b []byte) urlip {
	u := urlip{}
	if b == nil {
		b = []byte(j)
		json.Unmarshal(b, &u)
	} else {
		err := json.Unmarshal(b, &u)
		if err != nil {
			b = []byte(j)
			json.Unmarshal(b, &u)
		}
	}
	return u
}