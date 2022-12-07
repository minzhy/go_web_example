package template

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func Template() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	// layout.html就是一个模板，对照着模板可以看出，整个数据的结构。
	// 首先需要一个PageTitle作为heading，然后Todos结构体，里面包含Done属性和Title属性
	// 访问http.ListenAndServe(":80", nil)的时候，会调用HandleFunc函数。
	// HandleFunc函数，会调用内层的function，tmpl.Execute(w, data)，这个函数会执行
	// 会根据layout里面的判断，替换相应的部分
	http.ListenAndServe(":80", nil)
}
