package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

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
	return
}}

var adminSCAreaAdd = web.Route{"POST", "/admin/area", func(w http.ResponseWriter, r *http.Request) {
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

var adminSCAreaUpdate = web.Route{"POST", "/admin/area/:id", func(w http.ResponseWriter, r *http.Request) {

	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/area"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	area, err := GetOnlyOneSC_AreaById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error finding supply chain area")
		return
	}

	area.Name = r.FormValue("name")
	area.Description = r.FormValue("description")

	if err := UpdateAllSC_AreaById(area.Id, area); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error updating supply chain area")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully updated supply chain area")
	return

}}

var adminSCAreaDelete = web.Route{"POST", "/admin/area/:id/del", func(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/area"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	if err := AreaDeleteHyper(r.FormValue(":id")); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error deleting supply chain area")
		return
	}

	if err := DeleteAllSC_AreaById(r.FormValue(":id")); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error deleting supply chain area")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully deleted Supply chain area")
	return
}}

var adminAreaElement = web.Route{"GET", "/admin/area/:id/element", func(w http.ResponseWriter, r *http.Request) {
	area, err := GetOnlyOneSC_AreaById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area", "Error finding supply chain area")
		return
	}
	elements, err := AreaGetElementsLocatedIn(area.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area", "Error retrieving supply chain elements for supply chain area")
		return
	}

	otherElements, err := AreaGetElementsNotLocatedIn(area.Id)
	if err != nil {
		fmt.Printf("\n>>>>%v\n", err)
		web.SetErrorRedirect(w, r, "/admin/area", "Error getting supply chain elements")
		return
	}

	tmpl.Render(w, r, "admin-area-element.tmpl", web.Model{
		"area":          area,
		"elements":      elements,
		"otherElements": otherElements,
	})
	return
}}

var adminAreaElementAdd = web.Route{"POST", "/admin/area/:id/element/add", func(w http.ResponseWriter, r *http.Request) {
	area, err := GetOnlyOneSC_AreaById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area", "Error finding supply chain area")
		return
	}
	if err := AreaSetLocatedInElement(strings.Split(r.FormValue("elementIds"), ","), area.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/area/"+area.Id+"/element", "Error adding supply chain elements to supply chain area")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/area/"+area.Id+"/element", "Successfully added suply chain elements to supply chain areas")
	return
}}

var adminAreaElementRemove = web.Route{"POST", "/admin/area/:id/element/remove", func(w http.ResponseWriter, r *http.Request) {
	area, err := GetOnlyOneSC_AreaById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area", "Error finding supply chain area")
		return
	}
	if err := AreaRemoveLocatedInElement(r.FormValue("elementId"), area.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/area/"+area.Id+"/element", "Error removing supply chain element from supply chain area")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/area/"+area.Id+"/element", "Successfully removed suply chain element from supply chain area")
	return
}}

var adminAreaElementQuestion = web.Route{"GET", "/admin/area/:areaId/element/:elementId/question", func(w http.ResponseWriter, r *http.Request) {
	area, err := GetOnlyOneSC_AreaById(r.FormValue(":areaId"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area", "Error finding supply chain area")
		return
	}
	element, err := GetOnlyOneSC_ElementById(r.FormValue(":elementId"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area/"+area.Id+"/area", "Error finding supply chain element")
		return
	}

	questions, err := ElementGetAsksIn(area.Id, element.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area/"+area.Id+"/element/", "Error getting questions")
		return
	}

	otherQuestions, err := ElementGetNotAsksIn(area.Id, element.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/area/"+area.Id+"/element/", "Error getting questions")
		return
	}

	if len(otherQuestions) == 0 && len(questions) == 0 {
		otherQuestions, err = GetAllQuestion()
		if err != nil {
			web.SetErrorRedirect(w, r, "/admin/area/"+area.Id+"/element/", "Error getting questions")
			return
		}
	}

	tmpl.Render(w, r, "admin-element-area-question.tmpl", web.Model{
		"element":        element,
		"area":           area,
		"questions":      questions,
		"otherQuestions": otherQuestions,
		"fromArea":       true,
	})
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
	return
}}

var adminSCElementAdd = web.Route{"POST", "/admin/element", func(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/element"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}
	element := SC_Element{
		Id:          genId(),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}
	if err := AddSC_Element(element); err != nil {
		log.Printf("\nadminRoutes.go >> adminAddSCElement >> AddSCElement() >> %v\n", err)
		web.SetErrorRedirect(w, r, redirect, "Error adding supply chain element")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfuly added supply chain element")
	return

}}

var adminSCElementUpdate = web.Route{"POST", "/admin/element/:id", func(w http.ResponseWriter, r *http.Request) {

	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/element"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	element, err := GetOnlyOneSC_ElementById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error finding supply chain element")
		return
	}

	element.Name = r.FormValue("name")
	element.Description = r.FormValue("description")

	if err := UpdateAllSC_ElementById(element.Id, element); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error updating supply chain element")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully updated supply chain element")
	return

}}

var adminSCElementDelete = web.Route{"POST", "/admin/element/:id/del", func(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/element"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	if err := ElementDeleteHyper(r.FormValue(":id")); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error deleting supply chain element")
		return
	}

	if err := DeleteAllSC_ElementById(r.FormValue(":id")); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error deleting supply chain element")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully deleted supply chain element")
	return
}}

var adminSCElementArea = web.Route{"GET", "/admin/element/:id/area", func(w http.ResponseWriter, r *http.Request) {
	element, err := GetOnlyOneSC_ElementById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element", "Error finding supply chain element")
		return
	}
	areas, err := ElementGetLocatedInAreas(element.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element", "Error retrieving supply chain areas for supply chain element")
		return
	}

	otherAreas, err := ElementGetNotLocatedInAreas(element.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element", "Error getting supply chain areas")
		return
	}

	tmpl.Render(w, r, "admin-element-area.tmpl", web.Model{
		"element":    element,
		"areas":      areas,
		"otherAreas": otherAreas,
	})
	return
}}

var adminSCElementAreaAdd = web.Route{"POST", "/admin/element/:id/area/add", func(w http.ResponseWriter, r *http.Request) {
	element, err := GetOnlyOneSC_ElementById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element", "Error finding supply chain element")
		return
	}
	if err := ElementSetLocatedIn(strings.Split(r.FormValue("areaIds"), ","), element.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/element/"+element.Id+"/add", "Error adding supply chain element to supply chain areas")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/element/"+element.Id+"/area", "Successfully added suply chain element to supply chain areas")
	return
}}

var adminSCElementAreaRemove = web.Route{"POST", "/admin/element/:id/area/remove", func(w http.ResponseWriter, r *http.Request) {
	element, err := GetOnlyOneSC_ElementById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element", "Error finding supply chain element")
		return
	}
	if err := ElementRemoveLocatedIn(r.FormValue("areaId"), element.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/element/"+element.Id+"/area", "Error removing supply chain element from supply chain area")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/element/"+element.Id+"/area", "Successfully removed suply chain element from supply chain area")
	return
}}

var adminSCElementAreaQuestion = web.Route{"GET", "/admin/element/:elementId/area/:areaId/question", func(w http.ResponseWriter, r *http.Request) {
	element, err := GetOnlyOneSC_ElementById(r.FormValue(":elementId"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element", "Error finding supply chain element")
		return
	}
	area, err := GetOnlyOneSC_AreaById(r.FormValue(":areaId"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element/"+element.Id+"/area", "Error finding supply chain area")
		return
	}

	questions, err := ElementGetAsksIn(area.Id, element.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element/"+element.Id+"/area/", "Error getting questions")
		return
	}

	otherQuestions, err := ElementGetNotAsksIn(area.Id, element.Id)
	if err != nil {
		fmt.Printf(">>>>\n%v\n", err)
		web.SetErrorRedirect(w, r, "/admin/element/"+element.Id+"/area/", "Error getting questions")
		return
	}

	if len(otherQuestions) == 0 && len(questions) == 0 {
		otherQuestions, err = GetAllQuestion()
		if err != nil {
			fmt.Printf(">>>>\n%v\n", err)
			web.SetErrorRedirect(w, r, "/admin/element/"+element.Id+"/area/", "Error getting questions")
			return
		}
	}

	tmpl.Render(w, r, "admin-element-area-question.tmpl", web.Model{
		"element":        element,
		"area":           area,
		"questions":      questions,
		"otherQuestions": otherQuestions,
	})
	return

}}

var adminSCElementAreaQuestionAdd = web.Route{"POST", "/admin/element/:elementId/area/:areaId/question/add", func(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/element"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}
	element, err := GetOnlyOneSC_ElementById(r.FormValue(":elementId"))
	if err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error finding supply chain element")
		return
	}
	area, err := GetOnlyOneSC_AreaById(r.FormValue(":areaId"))
	if err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error finding supply chain area")
		return
	}
	if err := ElementSetAsksIn(area.Id, element.Id, strings.Split(r.FormValue("questionIds"), ",")); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error adding question to supply chain element")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully added question to supply chain element")
	return
}}

var adminSCElementAreaQuestionRemove = web.Route{"POST", "/admin/element/:elementId/area/:areaId/question/remove", func(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/element"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}
	element, err := GetOnlyOneSC_ElementById(r.FormValue(":elementId"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/element", "Error finding supply chain element")
		return
	}
	area, err := GetOnlyOneSC_AreaById(r.FormValue(":areaId"))
	if err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error finding supply chain area")
		return
	}
	if err := ElementRemoveAsksIn(area.Id, r.FormValue("questionId"), element.Id); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error removing question from supply chain element")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully removed question from supply chain element")
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
	return
}}

var adminQuestionAdd = web.Route{"POST", "/admin/question", func(w http.ResponseWriter, r *http.Request) {

	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/question"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	question := Question{
		Id: genId(),
		Q:  r.FormValue("q"),
	}
	if err := AddQuestion(question); err != nil {
		log.Printf("\nadminRoutes.go >> adminAddQuestion >> AddQuestion() >> %v\n", err)
		web.SetErrorRedirect(w, r, redirect, "Error adding question")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfuly added question")
	return

}}

var adminQuestionUpdate = web.Route{"POST", "/admin/question/:id", func(w http.ResponseWriter, r *http.Request) {

	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/question"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	question, err := GetOnlyOneQuestionById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error finding question")
		return
	}

	question.Q = r.FormValue("q")

	if err := UpdateAllQuestionById(question.Id, question); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error updating question")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully updated question")
	return

}}

var adminQuestionDelete = web.Route{"POST", "/admin/question/:id/del", func(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/question"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	if err := DeleteAllQuestionById(r.FormValue(":id")); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error deleting question")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully deleted question")
	return
}}

var adminQuestionResponse = web.Route{"GET", "/admin/question/:id/response", func(w http.ResponseWriter, r *http.Request) {
	question, err := GetOnlyOneQuestionById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/question", "Error finding question")
		return
	}
	responses, err := QuestionGetReceivedResponses(question.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/question", "Error getting responses")
		return
	}
	otherResponses, err := QuestionGetNotReceivedResponses(question.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/question", "Error getting responses")
	}
	tmpl.Render(w, r, "admin-question-response.tmpl", web.Model{
		"question":       question,
		"responses":      responses,
		"otherResponses": otherResponses,
	})

	return
}}

var adminQuestionResponseAdd = web.Route{"POST", "/admin/question/:id/response/add", func(w http.ResponseWriter, r *http.Request) {
	question, err := GetOnlyOneQuestionById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/question", "Error finding question")
		return
	}
	if err := QuestionSetReceived(strings.Split(r.FormValue("responseIds"), ","), question.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/question/"+question.Id+"/response", "Error adding response to question")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/question/"+question.Id+"/response", "Successfully added response to question")

	return
}}

var adminQuestionResponseRemove = web.Route{"POST", "/admin/question/:id/response/remove", func(w http.ResponseWriter, r *http.Request) {
	question, err := GetOnlyOneQuestionById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/question", "Error finding question")
		return
	}
	if err := QuestionRemoveReceived(r.FormValue("responseId"), question.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/question/"+question.Id+"/response", "Error removing response from question")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/question/"+question.Id+"/response", "Successfully removing response from question")

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

var adminResponseAdd = web.Route{"POST", "/admin/response", func(w http.ResponseWriter, r *http.Request) {

	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/response"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	response := Response{
		Id: genId(),
		R:  r.FormValue("r"),
	}
	if err := AddResponse(response); err != nil {
		log.Printf("\nadminRoutes.go >> adminAddResponse >> AddResponse() >> %v\n", err)
		web.SetErrorRedirect(w, r, redirect, "Error adding response")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfuly added response")
	return

}}

var adminResponseUpdate = web.Route{"POST", "/admin/response/:id", func(w http.ResponseWriter, r *http.Request) {

	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/response"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	response, err := GetOnlyOneResponseById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error finding response")
		return
	}

	response.R = r.FormValue("r")

	if err := UpdateAllResponseById(response.Id, response); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error updating response")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully updated response")
	return

}}

var adminResponseDelete = web.Route{"POST", "/admin/response/:id/del", func(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect")
	if redirect == "" {
		redirect = "/admin/response"
		rUrl, err := url.Parse(r.Referer())
		if err == nil {
			redirect = rUrl.Path
		}
	}

	if err := DeleteAllResponseById(r.FormValue(":id")); err != nil {
		web.SetErrorRedirect(w, r, redirect, "Error deleting response")
		return
	}

	web.SetSuccessRedirect(w, r, redirect, "Successfully deleted response")
	return
}}

var adminResponseQuestion = web.Route{"GET", "/admin/response/:id/question", func(w http.ResponseWriter, r *http.Request) {
	response, err := GetOnlyOneResponseById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/response", "Error finding response")
		return
	}
	questions, err := ResponseGetFollowUpQuestions(response.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/response", "Error getting questions")
		return
	}
	otherquestions, err := ResponseGetNotFollowUpQuestions(response.Id)
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/response", "Error getting questions")
	}
	tmpl.Render(w, r, "admin-response-question.tmpl", web.Model{
		"response":       response,
		"questions":      questions,
		"otherQuestions": otherquestions,
	})

	return
}}

var adminResponseQuestionAdd = web.Route{"POST", "/admin/response/:id/question/add", func(w http.ResponseWriter, r *http.Request) {
	response, err := GetOnlyOneResponseById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/response", "Error finding response")
		return
	}
	if err := ResponseSetFollowUp(strings.Split(r.FormValue("questionIds"), ","), response.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/response/"+response.Id+"/question", "Error adding follow up question to response")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/response/"+response.Id+"/question", "Successfully added follow up question to response")

	return
}}

var adminResponseQuestionRemove = web.Route{"POST", "/admin/response/:id/question/remove", func(w http.ResponseWriter, r *http.Request) {
	response, err := GetOnlyOneResponseById(r.FormValue(":id"))
	if err != nil {
		web.SetErrorRedirect(w, r, "/admin/response", "Error finding response")
		return
	}
	if err := ResponseRemoveFollowUp(r.FormValue("questionId"), response.Id); err != nil {
		web.SetErrorRedirect(w, r, "/admin/response/"+response.Id+"/question", "Error removing follow up question from response")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin/response/"+response.Id+"/question", "Successfully removed follow up question from response")

	return
}}
