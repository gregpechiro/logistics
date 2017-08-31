package main

import (
	"net/http"

	"github.com/cagnosolutions/web"
)

var account = web.Route{"GET", "/account", func(w http.ResponseWriter, r *http.Request) {
	id := web.GetId(r)
	var user User
	if !db.Get("user", id, &user) {
		web.Logout(w)
		web.SetErrorRedirect(w, r, "/login", "Error retrieving user")
	}
	tmpl.Render(w, r, "account.tmpl", web.Model{
		"user": user,
	})
	return
}}
