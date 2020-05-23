package dblib

import (
	"database/sql"
	"fmt"
	"loglib"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id     int    `json: "id"`
	Name   string `json: "name"`
	Family string `json: "family"`
}

func ReadUsers(dbName string) []User {

	dbConnect, err := sql.Open("sqlite3", dbName)
	loglib.CheckFatall(err, "Db connect error")
	defer dbConnect.Close()

	rows, err := dbConnect.Query("select * from user")
	loglib.CheckFatall(err, "Db query error")

	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Family)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users

}

func WriteUser(dbName string, name string, family string) bool {

	dbConnect, err := sql.Open("sqlite3", dbName)
	loglib.CheckFatall(err, "Db connect error")
	defer dbConnect.Close()
	_, errq := dbConnect.Exec("insert into user (name, family) values ($1, $2)", name, family)

	if errq != nil {
		fmt.Println(errq)
		return false
	}

	return true

}
