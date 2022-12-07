package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Routing() {
	// go get -u github.com/gorilla/mux -u意思是update的意思，最新的
	// 路由就是增强版的http库。http.HandleFunc(...)这些方法都是一样的。
	// 例子：/books/go-programming-blueprint/page/10
	// 这个url有两个动态部分！用router取提取（splite

	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
	// 这里不同于前，是r，如果是nil的话，就是默认用自带的http库，但这里用了routing
}

// router的一些特性：限制访问
// 限制方法
// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

// 限制域名：
// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

// 限制协议
// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

// 将请求处理程序限制为特定的路径前缀。
// bookrouter := r.PathPrefix("/books").Subrouter()
// bookrouter.HandleFunc("/", AllBooks)
// bookrouter.HandleFunc("/{title}", GetBook)
