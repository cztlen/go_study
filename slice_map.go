package main

import "fmt"

func main() {
	data := map[string]interface{}{}
	info := []map[string]string{}
	info = append(info,
		map[string]string{"name": "dd", "color": "red"},
		map[string]string{"name": "ff", "color": "back"})
	info = append(info,
		map[string]string{"name": "dd1", "color": "red1"},
		map[string]string{"name": "ff1", "color": "back1"})
	data["info"] = info

	for _, v := range data["info"].([]map[string]string) {
		// fmt.Println(k)
		fmt.Println(v)
	}
}
