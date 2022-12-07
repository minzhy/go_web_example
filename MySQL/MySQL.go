package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//import 下划线（如：import _ github/demo）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，
//然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包。
// 说白了，其实没有倒入

func MySQL() {
	db, err := sql.Open("mysql", "root:123321@tcp(127.0.0.1:3306)/test?parseTime=true")
	// 运行之前，需要打开mysql server，同时查看mysql的端口：进入数据库，执行show global variables like 'port';
	// 同时，需要先在mysql中创建数据库，test。 create database test;
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Create a new table
		query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Insert a new user
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}
		// 查询某些行，where id = ？，？用1来替代 db.QueryRow(query, 1)，同时，查询的结果写入&id, &username, &password, &createdAt里

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		//查询所有用户
		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print("[Done]\n")
}
