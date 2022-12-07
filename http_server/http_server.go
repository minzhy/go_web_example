package http_server

import (
	"fmt"
	"net/http"
)

// Process dynamic requests: Process incoming requests from users who browse the website, log into their accounts or post images.
// Serve static assets: Serve JavaScript, CSS and images to browsers to create a dynamic experience for the user.
// Accept connections: The HTTP Server must listen on a specific port to be able to accept connections from the internet.
// 这个例子是一个请求静态资源的

func Http_server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// 所有的请求的url的信息，其实都记录在http类内
	// http.StripPrefix函数的作用之一，就是在将请求定向到你通过参数指定的请求处理处之前，将特定的prefix从URL中过滤出去。下面是一个浏览器或HTTP客户端请求资源的例子：
	// 访问：……/static/example.txt，这时候会调用http.Handle，这个函数和HandleFunc其实是一样的，只是传递的参数不同。（因为这个更具体，所以会优先这个
	// StripPrefix 函数将会过滤掉/static/，并将修改过的请求定向到http.FileServer所返回的Handler中去，因此请求的资源将会是：
	// /example.txt
	// http.FileServer 返回的Handler将会进行查找，并将与文件夹或文件系统有关的内容以参数的形式返回给你（在这里你将"static"作为静态文件的根目录）。
	// 因为你的"example.txt"文件在静态目录中，你必须定义一个相对路径去获得正确的文件路径。
	// 因为这个函数是在main.go里面调用的，所以要放在main的目录下，而不是http_server的目录下

	http.ListenAndServe(":80", nil)
}
