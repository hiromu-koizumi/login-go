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
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/sample_docker_compose")
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
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ID, user.Name)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from Docker container!")
}
