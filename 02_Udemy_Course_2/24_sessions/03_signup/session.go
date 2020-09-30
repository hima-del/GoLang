package main

import (
	"fmt"
	"net/http"
)

func getUser(req *http.Request) user {
	var u user
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println(err)
	}
	un, ok := dbsessions[c.Value]
	if ok == true {
		u = dbusers[un]
	}
	return u
}

func alreadyloggedin(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println(err)
	}
	un := dbsessions[c.Value]
	_, ok := dbusers[un]
	return ok
}
