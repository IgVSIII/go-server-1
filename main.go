package main

import (
	"dblib"
	"encoding/json"
	"fmt"
	"loglib"
	"net/http"
	"readconfig"
)

func main() {

	config := readconfig.GetConfig("./config.json")
	http.HandleFunc("/users", accessMiddleware(getUsers, config))
	http.HandleFunc("/adduser", accessMiddleware(setUsers, config))
	fmt.Println("Server is runnig")
	err := http.ListenAndServe(config.Address, nil)
	loglib.CheckFatall(err, "Server error")

}

func accessMiddleware(next http.HandlerFunc, config readconfig.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if key == config.Key {
			fmt.Println("Server access key is correct")
			r.Form.Add("db", config.DBname)
			next(w, r)
		} else {
			fmt.Println("Server access key is not correct")
			w.WriteHeader(http.StatusUnauthorized)
		}

	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := dblib.ReadUsers(r.FormValue("db"))
	fmt.Println("Server response:")
	fmt.Printf("%12s | %12s | %12s\n", "Id", "Name", "Family")
	fmt.Println("------------------------------------------")
	for _, u := range users {
		fmt.Printf("%12d | %12s | %12s\n", u.Id, u.Name, u.Family)
	}
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(b))
	}
}

func setUsers(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	family := r.FormValue("family")

	result := true

	if name != "" || family != "" {
		result = dblib.WriteUser(r.FormValue("db"), name, family)
		fmt.Println("Server get data:")
	} else {
		result = false
	}

	if !result {
		w.WriteHeader(http.StatusNotImplemented)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
