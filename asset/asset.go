// static-files.go
package asset

import "net/http"

func Asset() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// 请求中间件相关，其实语法根http server一样

	http.ListenAndServe(":8080", nil)
}
