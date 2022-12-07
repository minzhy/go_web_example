// advanced-middleware.go
package advancedmiddleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// 这个函数的执行比较复杂
func AdvancedMiddleware() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	// http.HandleFunc相当于一个注册，注册一个http.HandlerFunc的函数。
	// 首先会先执行最里面的函数，先会执行Method("GET")，返回一个func(f http.HandlerFunc) http.HandlerFunc函数，用于后续嵌套
	// 然后执行Logging()，同上
	// 然后执行Chain函数，第一个m是Method("GET")返回的函数，然后把Hello函数放入，返回一个http.HandlerFunc函数
	// tips：func(w http.ResponseWriter, r *http.Request)其实就是http.HandlerFunc的定义
	// 然后range到第二个m是Logging()，但这时候，f已经是包装了Hello的Method("GET")返回的函数了，再进行包装，放入Logging中，这时候，相当于三层嵌套了，然后返回三层嵌套的f
	// 访问的时候，就会一层层的执行
	http.ListenAndServe(":8080", nil)
}

// 创建中间件的模板
func createNewMiddleware() Middleware {

	// Create a new Middleware
	middleware := func(next http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc which is called by the server eventually
		handler := func(w http.ResponseWriter, r *http.Request) {

			// ... do middleware things

			// Call the next middleware/handler in chain
			next(w, r)
		}

		// Return newly created handler
		return handler
	}

	// Return newly created middleware
	return middleware
}
