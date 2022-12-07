// sessions.go
package session

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
	// 首先调用NewCookieStore初始化一个store，同时传入一个secret key用来对session进行认证。初始化session
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	// 检查这个key是否存在，同时是否是bool性，是否为true
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// 在Handler中，调用store.Get()获取一个已经存在的cookie-name或（如果不存在）创建一个新的。
	// 这里应该是不存在的，因为这是第一次登陆，所以是由服务器创建一个cookie值给到客户端

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true // 在session的机制里，设置一个map，key为authenticated，值为true
	session.Save(r, w)
	// 调用session.Save()将session保存到响应中，返回给客户端（这个session里面包括了传过来的cookie值
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func Session() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}

// session存在服务端，cookie是客户端，cookie是session的一种机制

// 例子：
// $ curl -s http://localhost:8080/secret
// Forbidden

// $ curl -s -I http://localhost:8080/login
// Set-Cookie: cookie-name=MTQ4NzE5Mz...
// -I 参数是为了返回Set-Cookie值！否则不会在客户端返回Cookie的值

// $ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
// The cake is a lie!
