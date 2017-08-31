package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cagnosolutions/adb"
	"github.com/cagnosolutions/web"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var tmpl *web.TmplCache
var mux *web.Mux
var db *adb.DB = adb.NewDB()
var driver bolt.DriverPool

func init() {
	db.AddStore("user")

	web.SESSDUR = 15 * time.Minute
	mux = web.NewMux()

	// unsecure routes
	mux.AddRoutes(home, register, login, logout, loginPost)

	// user routes
	mux.AddSecureRoutes(USER, account)

	// admin routes
	// mux.AddSecureRoutes(ADMIN, adminHome)

	mux.AddRoutes(adminHome, adminSCArea, adminSCElement, adminQuestion, adminResponse)
	mux.AddRoutes(adminAddSCArea, adminAddSCElement, adminAddQuestion, adminAddResponse)

	tmpl = web.NewTmplCache()

	defaultUsers()

	var err error
	driver, err = bolt.NewDriverPool("bolt://neo4j:admin@localhost:7687", 10)
	if err != nil {
		panic(err)
	}

	/*if err := IndexQuestionById(); err != nil {
		panic(err)
	}
	if err := IndexSC_AreaById(); err != nil {
		panic(err)
	}
	if err := IndexSC_ElementById(); err != nil {
		panic(err)
	}
	if err := IndexResponseById(); err != nil {
		panic(err)
	}*/
}

func main() {
	fmt.Println(">>> DID YOU REGISTER ANY NEW ROUTES <<<")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

var home = web.Route{"GET", "/", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "home.tmpl", nil)
	return
}}