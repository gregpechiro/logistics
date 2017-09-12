package main

import (
	"fmt"
	"io"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

type Params = map[string]interface{}

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email,omitempty" auth:"username" required:"register, login"`
	Password  string `json:"password,omitempty" auth:"password" required:"register, login"`
	Active    bool   `json:"active" auth:"active"`
	Role      string `json:"role,omitempty"`
	FirstName string `json:"firstName,omitempty" required:"register"`
	LastName  string `json:"lastName,omitempty" required:"register"`
	Created   int64  `json:"created,omitempty"`
	LastSeen  int64  `json:"lastSeen,omitempty"`
}

//go:generate neo4jGenerator Question
type Question struct {
	Id string `neo4j:"index"`
	Q  string
}

//go:generate neo4jGenerator SC_Area
type SC_Area struct {
	Id          string `neo4j:"index"`
	Name        string
	Description string
}

//go:generate neo4jGenerator SC_Element
type SC_Element struct {
	Id          string `neo4j:"index"`
	Name        string
	Description string
}

//go:generate neo4jGenerator Response
type Response struct {
	Id string `neo4j:"index"`
	R  string
}

func AreaGetElementsLocatedIn(areaId string) ([]SC_Element, error) {
	var elements []SC_Element
	conn, err := driver.OpenPool()
	if err != nil {
		return elements, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (e:SC_Element)-[:LOCATED_IN]->(a:SC_Area{ Id:{ areaId } }) RETURN e", Params{"areaId": areaId})
	if err != nil {
		return elements, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return elements, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return elements, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		element := SC_Element{}
		if Id, ok := node.Properties["Id"].(string); ok {
			element.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			element.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			element.Description = Description
		}

		elements = append(elements, element)
	}

	return elements, nil
}

func AreaGetElementsNotLocatedIn(areaId string) ([]SC_Element, error) {
	var elements []SC_Element
	conn, err := driver.OpenPool()

	if err != nil {
		return elements, err

	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (e:SC_Element) WHERE NOT exists( (e)-[:LOCATED_IN]->(:SC_Area{ Id:{ areaId } }) ) RETURN e", Params{"areaId": areaId})
	if err != nil {
		return elements, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return elements, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return elements, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		element := SC_Element{}
		if Id, ok := node.Properties["Id"].(string); ok {
			element.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			element.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			element.Description = Description
		}

		elements = append(elements, element)
	}

	return elements, nil
}

func ElementGetLocatedInAreas(elementId string) ([]SC_Area, error) {
	var sC_Areas []SC_Area
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Areas, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (e:SC_Element{ Id:{ elementId } })-[:LOCATED_IN]->(a:SC_Area) RETURN a", Params{"elementId": elementId})
	if err != nil {
		return sC_Areas, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return sC_Areas, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return sC_Areas, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		sC_Area := SC_Area{}
		if Id, ok := node.Properties["Id"].(string); ok {
			sC_Area.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			sC_Area.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			sC_Area.Description = Description
		}

		sC_Areas = append(sC_Areas, sC_Area)
	}

	return sC_Areas, nil
}

func ElementGetNotLocatedInAreas(elementId string) ([]SC_Area, error) {
	var sC_Areas []SC_Area
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Areas, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (a:SC_Area) WHERE NOT exists( (:SC_Element{ Id:{ elementId } })-[:LOCATED_IN]->(a) ) RETURN a", Params{"elementId": elementId})
	if err != nil {
		return sC_Areas, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return sC_Areas, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return sC_Areas, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		sC_Area := SC_Area{}
		if Id, ok := node.Properties["Id"].(string); ok {
			sC_Area.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			sC_Area.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			sC_Area.Description = Description
		}

		sC_Areas = append(sC_Areas, sC_Area)
	}

	return sC_Areas, nil
}

func ElementSetLocatedIn(areaIds []string, elementId string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	var queries []string
	var params []Params
	for _, areaId := range areaIds {
		queries = append(queries, "MATCH (a:SC_Area{Id:{ areaId }}) MATCH (e:SC_Element{Id:{ elementId }}) MERGE (e)-[:LOCATED_IN]->(a)")
		params = append(params, Params{
			"areaId":    areaId,
			"elementId": elementId,
		})
	}

	if _, err := conn.ExecPipeline(queries, params...); err != nil {
		return err
	}

	return nil
}

func ElementRemoveLocatedIn(areaId string, elementId string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecNeo("MATCH (:SC_Element{Id:{ elementId }})-[r:LOCATED_IN]->(:SC_Area{Id:{ areaId }}) DELETE r", Params{
		"areaId":    areaId,
		"elementId": elementId,
	}); err != nil {
		fmt.Printf(">>>>>\n%v\n", err)
		return err
	}

	return nil
}

func ElementSetAsksIn(areaId, elementId string, questionIds []string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	var queries []string
	var params []Params
	for _, questionId := range questionIds {
		queries = append(queries, "MATCH (a:SC_Area{Id:{ areaId }}) MATCH (e:SC_Element{Id:{ elementId }}) MATCH (q:Question{Id:{ questionId }}) MERGE (e)-[:ASKS_IN]->(h:Hyper)-[:IN]->(a) MERGE (h)-[:ASKS]->(q)")
		params = append(params, Params{
			"areaId":     areaId,
			"elementId":  elementId,
			"questionId": questionId,
		})
	}

	if _, err := conn.ExecPipeline(queries, params...); err != nil {
		return err
	}

	return nil
}

func ElementRemoveAsksIn(areaId, questionId, elementId string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecNeo("MATCH (:SC_Element{Id:{ elementId }})-[:ASKS_IN]->(h:Hyper)-[:IN]->(:SC_Area{Id:{ areaId }}), (h)-[r:ASKS]->(:Question{Id:{ questionId }}) DELETE r", Params{
		"areaId":     areaId,
		"elementId":  elementId,
		"questionId": questionId,
	}); err != nil {
		return err
	}

	return nil
}

func ElementGetAsksIn(areaId, elementId string) ([]Question, error) {
	var questions []Question
	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (:SC_Element{Id:{ elementId }})-[:ASKS_IN]->(h:Hyper)-[:IN]->(:SC_Area{Id:{ areaId }}), (h)-[:ASKS]->(q:Question) RETURN q", Params{
		"areaId":    areaId,
		"elementId": elementId,
	})
	if err != nil {
		return questions, err
	}

	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return questions, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return questions, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		question := Question{}
		if Id, ok := node.Properties["Id"].(string); ok {
			question.Id = Id
		}
		if q, ok := node.Properties["Q"].(string); ok {
			question.Q = q
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func ElementGetNotAsksIn(areaId, elementId string) ([]Question, error) {
	var questions []Question
	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (:SC_Element{Id:{ elementId }})-[:ASKS_IN]->(h:Hyper)-[:IN]->(:SC_Area{Id: { areaId }}) MATCH (q:Question) WHERE NOT exists((h)-[:ASKS]->(q)) RETURN q", Params{
		"areaId":    areaId,
		"elementId": elementId,
	})
	if err != nil {
		return questions, err
	}

	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return questions, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return questions, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		question := Question{}
		if Id, ok := node.Properties["Id"].(string); ok {
			question.Id = Id
		}
		if q, ok := node.Properties["Q"].(string); ok {
			question.Q = q
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func QuestionSetReceived(responseIds []string, questionId string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	var queries []string
	var params []Params
	for _, responseId := range responseIds {
		queries = append(queries, "MATCH (r:Response{Id:{ responseId }}) MATCH (q:Question{Id:{ questionId }}) MERGE (q)-[:RECEIVED]->(r)")
		params = append(params, Params{
			"responseId": responseId,
			"questionId": questionId,
		})
	}

	if _, err := conn.ExecPipeline(queries, params...); err != nil {
		return err
	}

	return nil
}

func QuestionRemoveReceived(responseId string, questionId string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecNeo("MATCH (:Question{Id:{ questionId }})-[r:RECEIVED]->(:Response{Id:{ responseId }}) DELETE r", Params{
		"responseId": responseId,
		"questionId": questionId,
	}); err != nil {
		return err
	}

	return nil
}

func QuestionGetReceivedResponses(questionId string) ([]Response, error) {
	var responses []Response
	conn, err := driver.OpenPool()
	if err != nil {
		return responses, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (q:Question{ Id:{ questionId } })-[:RECEIVED]->(r:Response) RETURN r", Params{"questionId": questionId})
	if err != nil {
		return responses, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return responses, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return responses, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		response := Response{}
		if id, ok := node.Properties["Id"].(string); ok {
			response.Id = id
		}
		if r, ok := node.Properties["R"].(string); ok {
			response.R = r
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func QuestionGetNotReceivedResponses(questionId string) ([]Response, error) {
	var responses []Response
	conn, err := driver.OpenPool()
	if err != nil {
		return responses, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (r:Response) WHERE NOT exists( (:Question{ Id:{ questionId } })-[:RECEIVED]->(r) ) RETURN r", Params{"questionId": questionId})
	if err != nil {
		return responses, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return responses, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return responses, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		response := Response{}
		if id, ok := node.Properties["Id"].(string); ok {
			response.Id = id
		}
		if r, ok := node.Properties["R"].(string); ok {
			response.R = r
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func ResponseSetFollowUp(questionIds []string, responseId string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	var queries []string
	var params []Params
	for _, questionId := range questionIds {
		queries = append(queries, "MATCH (r:Response{Id:{ responseId }}) MATCH (q:Question{Id:{ questionId }}) MERGE (r)-[:FOLLOW_UP]->(q)")
		params = append(params, Params{
			"responseId": responseId,
			"questionId": questionId,
		})
	}

	if _, err := conn.ExecPipeline(queries, params...); err != nil {
		return err
	}

	return nil
}

func ResponseRemoveFollowUp(responseId string, questionId string) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecNeo("MATCH (:Response{Id:{ responseId }})-[r:FOLLOW_UP]->(:Question{Id:{ questionId }}) DELETE r", Params{
		"responseId": responseId,
		"questionId": questionId,
	}); err != nil {
		return err
	}

	return nil
}

func ResponseGetFollowUpQuestions(responseId string) ([]Question, error) {
	var questions []Question
	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (:Response{ Id:{ responseId } })-[:FOLLOW_UP]->(q:Question) RETURN q", Params{"responseId": responseId})
	if err != nil {
		return questions, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return questions, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return questions, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		question := Question{}
		if id, ok := node.Properties["Id"].(string); ok {
			question.Id = id
		}
		if q, ok := node.Properties["Q"].(string); ok {
			question.Q = q
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func ResponseGetNotFollowUpQuestions(responseId string) ([]Question, error) {
	var questions []Question
	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (q:Question) WHERE NOT exists( (:Response{ Id:{ responseId } })-[:FOLLOW_UP]->(q) ) RETURN q", Params{"responseId": responseId})
	if err != nil {
		return questions, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return questions, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return questions, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		question := Question{}
		if id, ok := node.Properties["Id"].(string); ok {
			question.Id = id
		}
		if r, ok := node.Properties["Q"].(string); ok {
			question.Q = r
		}

		questions = append(questions, question)
	}

	return questions, nil
}
