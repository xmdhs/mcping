package mcping

import "encoding/json"

func JSON(b []byte) map[string][]string {
	u := make(map[string][]string)
	err := json.Unmarshal(b, &u)
	if err != nil {
		panic(err)
	}
	return u
}
