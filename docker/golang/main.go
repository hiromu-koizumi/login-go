package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Mail string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}



func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(mysql)/sample_docker_compose")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Mail)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ID, user.Name, user.Mail)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	stmtInsert, err := db.Prepare("INSERT INTO users(name,mail) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	//Exec()に入る言葉を保存する
	result, err := stmtInsert.Exec("進次郎","kkk")
	if err != nil {
		panic(err.Error())
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastInsertID)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, from Docker container!")
// }
