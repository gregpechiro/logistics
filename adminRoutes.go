package main

import (
	"log"
	"net/http"

	"github.com/cagnosolutions/web"
)

var adminHome = web.Route{"GET", "/admin", func(w http.ResponseWriter, r *http.Request) {
	/*id := web.GetId(r)
	var user User
	if !db.Get("user", id, &user) {
		web.Logout(w)
		web.SetErrorRedirect(w, r, "/login", "Error retrieving user")
		return
	}
	tmpl.Render(w, r, "admin.tmpl", web.Model{
		"user": user,
	})
	*/
	tmpl.Render(w, r, "admin.tmpl", nil)
}}

var adminSCArea = web.Route{"GET", "/admin/area", func(w http.ResponseWriter, r *http.Request) {
	areas, err := GetAllSC_Area()
	if err != nil {
		log.Printf("\nadminRoutes.go >> adminSCArea >> GetAllSC_Area() >> %v\n", err)
	}

	tmpl.Render(w, r, "admin-area.tmpl", web.Model{
		"areas": areas,
	})
}}

var adminAddSCArea = web.Route{"POST", "/admin/area", func(w http.ResponseWriter, r *http.Request) {
	area := SC_Area{
		Id:          genId(),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}
	if err := AddSC_Area(area); err != nil {
		log.Printf("\nadminRoutes.go >> adminAddSCArea >> AddSCArea() >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin/area", "Error adding supply chain area")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/area", "Successfuly added supply chain area")
	return

}}

var adminSCElement = web.Route{"GET", "/admin/element", func(w http.ResponseWriter, r *http.Request) {
	elements, err := GetAllSC_Element()
	if err != nil {
		log.Printf("\nadminRoutes.go >> adminSCElememt >> GetAllSC_Element() >> %v\n", err)
	}

	tmpl.Render(w, r, "admin-element.tmpl", web.Model{
		"elements": elements,
	})
}}

var adminAddSCElement = web.Route{"POST", "/admin/element", func(w http.ResponseWriter, r *http.Request) {
	element := SC_Element{
		Id:          genId(),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}
	if err := AddSC_Element(element); err != nil {
		log.Printf("\nadminRoutes.go >> adminAddSCElement >> AddSCElement() >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin/element", "Error adding supply chain element")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/element", "Successfuly added supply chain element")
	return

}}

var adminResponse = web.Route{"GET", "/admin/response", func(w http.ResponseWriter, r *http.Request) {
	responses, err := GetAllResponse()
	if err != nil {
		log.Printf("\nadminRoutes.go >> adminResponses >> GetAllResponse() >> %v\n", err)
	}

	tmpl.Render(w, r, "admin-response.tmpl", web.Model{
		"responses": responses,
	})
}}

var adminAddResponse = web.Route{"POST", "/admin/response", func(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Id: genId(),
		R:  r.FormValue("r"),
	}
	if err := AddResponse(response); err != nil {
		log.Printf("\nadminRoutes.go >> adminAddResponse >> AddResponse() >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin/response", "Error adding response")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/response", "Successfuly added response")
	return

}}

var adminQuestion = web.Route{"GET", "/admin/question", func(w http.ResponseWriter, r *http.Request) {
	questions, err := GetAllQuestion()
	if err != nil {
		log.Printf("\nadminRoutes.go >> adminQuestion >> GetAllQuestion() >> %v\n", err)
	}

	tmpl.Render(w, r, "admin-question.tmpl", web.Model{
		"questions": questions,
	})
}}

var adminAddQuestion = web.Route{"POST", "/admin/question", func(w http.ResponseWriter, r *http.Request) {
	question := Question{
		Id: genId(),
		Q:  r.FormValue("q"),
	}
	if err := AddQuestion(question); err != nil {
		log.Printf("\nadminRoutes.go >> adminAddQuestion >> AddQuestion() >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin/question", "Error adding question")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/question", "Successfuly added question")
	return

}}
