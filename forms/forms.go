// forms.go
package forms

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func Forms() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	// 创建模板

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Print("haha\n")
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			// 一开始访问这个网址的时候，先掉进这里，因为一开始不是post方法，应该是get
			// 这里是nil，因为用的是http，而不是mux
			return
		}
		// 点击submit之后，再次进来，执行下面的操作。

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		// 从r里面去取，submit的值

		// do something with details
		_ = details
		// 这里其实是没写的。你可以实现一个逻辑，用得到的值实现一个逻辑

		tmpl.Execute(w, struct{ Success bool }{true})
		// 在template中，将Success置为true；
	})
	// 上面这个函数是当网页被访问一次，则会执行一次，所以如果点击submit按钮
	// 则会在终端多输出一次haha，同时，如果刷新页面也会在终端多输出一次haha

	http.ListenAndServe(":8080", nil)
}
