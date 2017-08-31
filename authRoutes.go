package main

import (
	"net/http"
	"time"

	"github.com/cagnosolutions/adb"
	"github.com/cagnosolutions/web"
)

var register = web.Route{"POST", "/register", func(w http.ResponseWriter, r *http.Request) {
	var user User
	r.ParseForm()

	// check form for errors
	if errs, ok := web.FormToStruct(&user, r.Form, "register"); !ok {
		web.SetFormErrors(w, errs)
		web.SetErrorRedirect(w, r, "/login", "Error registering")
		return
	}

	// check for uniqueness
	var users []User
	db.TestQuery("user", &user, adb.Eq("email", user.Email), adb.Ne("id", `"`+user.Id+`"`))
	if len(users) > 0 {
		web.SetErrorRedirect(w, r, "/login", "Error registering. Email is already in use.")
		return
	}

	// create user
	user.Active = true
	user.Id = genId()
	user.Role = "USER"
	user.Created = time.Now().Unix()
	user.LastSeen = user.Created

	// add to database and check for success
	if !db.Add("user", user.Id, user) {
		web.SetErrorRedirect(w, r, "/", "Error registering. Please try again")
		return
	}

	// login user
	sess := web.Login(w, r, user.Role)
	sess.PutId(w, user.Id)
	sess["EMAIL"] = user.Email
	web.PutMultiSess(w, r, sess)

	// redirect with message
	web.SetSuccessRedirect(w, r, "/account", "Welcome "+user.FirstName)
	return

}}

var login = web.Route{"GET", "/login", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "login.tmpl", nil)
	return
}}

var loginPost = web.Route{"POST", "/login", func(w http.ResponseWriter, r *http.Request) {
	var user User
	r.ParseForm()

	// check for form errors
	if errs, ok := web.FormToStruct(&user, r.Form, "login"); !ok {
		web.SetFormErrors(w, errs)
		web.SetErrorRedirect(w, r, "/", "Error logging in")
		return
	}

	// look for user in database
	if !db.Auth("user", user.Email, user.Password, &user) {
		web.SetErrorRedirect(w, r, "/login", "Incorrect email or password")
		return
	}

	// login user
	sess := web.Login(w, r, user.Role)
	sess.PutId(w, user.Id)
	sess["EMAIL"] = user.Email
	web.PutMultiSess(w, r, sess)

	// redirect with message
	web.SetSuccessRedirect(w, r, "/account", "Welcome "+user.FirstName)
	return

}}

var logout = web.Route{"GET", "/logout", func(w http.ResponseWriter, r *http.Request) {
	web.Logout(w)
	web.SetSuccessRedirect(w, r, "/login", "See you next time")
	return
}}
