package hello_world

import (
	"fmt"
	"net/http"
)

// http://localhost/fds/fds 返回 Hello, you've requested: /fds/fds
func Hello_world() {
	// "/"应该是默认的http请求位置。
	// 如果不用"/"，换成"/haha"，那么只有请求http://localhost/haha的时候能够响应，其他的地址都是404not page
	// 如果是"/"的话，后面接什么都默认返回"/"的接收
	// 但如果有具体的"/"后面接什么的话，就返回具体的！
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/haha", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested haha: %s\n", r.URL.Path)
	})

	// 监听80端口，http的一般端口号
	http.ListenAndServe(":80", nil)
}
