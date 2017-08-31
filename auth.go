package main

import "github.com/cagnosolutions/web"

var ADMIN = web.Auth{
	Roles:    []string{"ADMIN"},
	Redirect: "/login",
	Msg:      "You must be an admin to view this page",
}

var USER = web.Auth{
	Roles:    []string{"ADMIN", "USER"},
	Redirect: "/login",
	Msg:      "Please login",
}
