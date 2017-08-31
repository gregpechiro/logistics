/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/neo4jGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

var NoQuestionFound = fmt.Errorf("no question found")
var MultipleQuestionFound = fmt.Errorf("multiple question found")

func AddQuestion(question Question) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (question:Question { Id:{questionId}, Q:{questionQ} })", map[string]interface{}{
		"questionId": question.Id,
		"questionQ":  question.Q,
	})

	return err
}

func GetAllQuestion() ([]Question, error) {
	var questions []Question
	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (question:Question) RETURN question", nil)
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
		if Q, ok := node.Properties["Q"].(string); ok {
			question.Q = Q
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func IndexQuestionById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :Question(Id)", nil)

	return err
}

func GetQuestionById(id string) (Question, error) {
	question := Question{}

	conn, err := driver.OpenPool()
	if err != nil {
		return question, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (question:Question{ Id:{ Id } }) RETURN question", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return question, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return question, NoQuestionFound
	}
	if err != nil {
		return question, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return question, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		question.Id = Id
	}
	if Q, ok := node.Properties["Q"].(string); ok {
		question.Q = Q
	}

	return question, nil
}

func GetOnlyOneQuestionById(id string) (Question, error) {
	question := Question{}

	conn, err := driver.OpenPool()
	if err != nil {
		return question, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (question:Question{ Id:{ Id } }) RETURN question", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return question, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return question, NoQuestionFound
	}
	if err != nil {
		return question, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return question, MultipleQuestionFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return question, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		question.Id = Id
	}
	if Q, ok := node.Properties["Q"].(string); ok {
		question.Q = Q
	}

	return question, nil
}

func GetAllQuestionById(id string) ([]Question, error) {
	var questions []Question
	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (question:Question{ Id:{ Id } }) RETURN question", map[string]interface{}{
		"Id": id,
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
		if Q, ok := node.Properties["Q"].(string); ok {
			question.Q = Q
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func UpdateAllQuestionById(id string, question Question) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (question:Question{ Id:{ Id } }) SET question += { Id:{questionId}, Q:{questionQ} }", map[string]interface{}{
		"Id":         id,
		"questionId": question.Id,
		"questionQ":  question.Q,
	})
	return err
}

func DeleteAllQuestionById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (question:Question{ Id:{ Id }) DETACH DELETE question", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetQuestionByQ(q string) (Question, error) {
	question := Question{}

	conn, err := driver.OpenPool()
	if err != nil {
		return question, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (question:Question{ Q:{ Q } }) RETURN question", map[string]interface{}{
		"Q": q,
	})
	if err != nil {
		return question, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return question, NoQuestionFound
	}
	if err != nil {
		return question, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return question, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		question.Id = Id
	}
	if Q, ok := node.Properties["Q"].(string); ok {
		question.Q = Q
	}

	return question, nil
}

func GetOnlyOneQuestionByQ(q string) (Question, error) {
	question := Question{}

	conn, err := driver.OpenPool()
	if err != nil {
		return question, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (question:Question{ Q:{ Q } }) RETURN question", map[string]interface{}{
		"Q": q,
	})

	if err != nil {
		return question, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return question, NoQuestionFound
	}
	if err != nil {
		return question, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return question, MultipleQuestionFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return question, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		question.Id = Id
	}
	if Q, ok := node.Properties["Q"].(string); ok {
		question.Q = Q
	}

	return question, nil
}

func GetAllQuestionByQ(q string) ([]Question, error) {
	var questions []Question
	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (question:Question{ Q:{ Q } }) RETURN question", map[string]interface{}{
		"Q": q,
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
		if Q, ok := node.Properties["Q"].(string); ok {
			question.Q = Q
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func UpdateAllQuestionByQ(q string, question Question) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (question:Question{ Q:{ Q } }) SET question += { Id:{questionId}, Q:{questionQ} }", map[string]interface{}{
		"Q":          q,
		"questionId": question.Id,
		"questionQ":  question.Q,
	})
	return err
}

func DeleteAllQuestionByQ(q string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (question:Question{ Q:{ Q }) DETACH DELETE question", map[string]interface{}{
		"Q": q,
	})
	return err
}

func GetQuestionByCustom(query map[string]interface{}) (Question, error) {
	question := Question{}

	conn, err := driver.OpenPool()
	if err != nil {
		return question, err
	}
	defer conn.Close()

	queryStr := "MATCH (question:Question{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN question"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return question, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return question, NoQuestionFound
	}
	if err != nil {
		return question, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return question, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		question.Id = Id
	}
	if Q, ok := node.Properties["Q"].(string); ok {
		question.Q = Q
	}

	return question, nil
}

func GetOnlyOneQuestionByCustom(query map[string]interface{}) (Question, error) {
	question := Question{}

	conn, err := driver.OpenPool()
	if err != nil {
		return question, err
	}
	defer conn.Close()

	queryStr := "MATCH (question:Question{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN question"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return question, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return question, NoQuestionFound
	}
	if err != nil {
		return question, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return question, MultipleQuestionFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return question, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		question.Id = Id
	}
	if Q, ok := node.Properties["Q"].(string); ok {
		question.Q = Q
	}

	return question, nil
}

func GetAllQuestionByCustom(query map[string]interface{}) ([]Question, error) {
	var questions []Question

	conn, err := driver.OpenPool()
	if err != nil {
		return questions, err
	}
	defer conn.Close()

	queryStr := "MATCH (question:Question{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN question"

	rows, err := conn.QueryNeo(queryStr, query)
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
		if Q, ok := node.Properties["Q"].(string); ok {
			question.Q = Q
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func UpdateAllQuestionByCustom(params map[string]interface{}, question Question) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (question:Question{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET question += { Id:{questionId}, Q:{questionQ} }"

	params["questionId"] = question.Id
	params["questionQ"] = question.Q

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllQuestionByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (question:Question{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE question"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
