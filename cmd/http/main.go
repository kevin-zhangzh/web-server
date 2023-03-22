package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	story := ""
	hour := t.Hour()
	switch {
	case hour >= 0 && hour < 6:
		story = "夜深了，小动物们都睡了，只有月亮陪伴着你。"
	case hour >= 6 && hour < 12:
		story = "早上好！阳光明媚，鸟儿在欢快地歌唱。"
	case hour >= 12 && hour < 18:
		story = "午后的阳光照在身上，有点热，但是有一阵阵微风，让人感到很舒服。"
	case hour >= 18 && hour < 24:
		story = "晚上好！月亮升起来了，天空变得更加宁静。"
	}
	fmt.Fprintf(w, "你好！%s", story)
}
