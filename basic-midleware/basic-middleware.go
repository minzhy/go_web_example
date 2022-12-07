// basic-middleware.go
package basicmidleware

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		fmt.Println(r.URL.Path)
		// 如果访问：http://localhost:8080/bar
		// r.URL.Path返回的是：/bar
		// log.Println 会在/bar前面加上今天的日期和时间
		f(w, r)
		// 这个函数相当于只执行了：fmt.Fprintln(w, "foo")
		// 会往客户端写一个foo，而上面的fmt.Println是在服务端写
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func BasicMidleware() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	// logging其实就是自己定义的一个函数，与hello-world.go的程序是没有区别的

	http.ListenAndServe(":8080", nil)
}
