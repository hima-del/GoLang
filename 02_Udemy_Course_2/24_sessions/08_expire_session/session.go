package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		SID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: SID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	//if the user exists already, get the user

	var u user
	s, ok := dbSessions[c.Value]
	if ok == true {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println(err)
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	//refreash session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func cleanSessions() {
	fmt.Println("before clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("after clean")
	showSessions()
}

func showSessions() {
	fmt.Println("******")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
