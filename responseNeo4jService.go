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

var NoResponseFound = fmt.Errorf("no response found")
var MultipleResponseFound = fmt.Errorf("multiple response found")

func AddResponse(response Response) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (response:Response { Id:{responseId}, R:{responseR} })", map[string]interface{}{
		"responseId": response.Id,
		"responseR":  response.R,
	})

	return err
}

func GetAllResponse() ([]Response, error) {
	var responses []Response
	conn, err := driver.OpenPool()
	if err != nil {
		return responses, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (response:Response) RETURN response", nil)
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
		if Id, ok := node.Properties["Id"].(string); ok {
			response.Id = Id
		}
		if R, ok := node.Properties["R"].(string); ok {
			response.R = R
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func IndexResponseById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :Response(Id)", nil)

	return err
}

func GetResponseById(id string) (Response, error) {
	response := Response{}

	conn, err := driver.OpenPool()
	if err != nil {
		return response, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (response:Response{ Id:{ Id } }) RETURN response", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return response, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return response, NoResponseFound
	}
	if err != nil {
		return response, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return response, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		response.Id = Id
	}
	if R, ok := node.Properties["R"].(string); ok {
		response.R = R
	}

	return response, nil
}

func GetOnlyOneResponseById(id string) (Response, error) {
	response := Response{}

	conn, err := driver.OpenPool()
	if err != nil {
		return response, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (response:Response{ Id:{ Id } }) RETURN response", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return response, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return response, NoResponseFound
	}
	if err != nil {
		return response, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return response, MultipleResponseFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return response, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		response.Id = Id
	}
	if R, ok := node.Properties["R"].(string); ok {
		response.R = R
	}

	return response, nil
}

func GetAllResponseById(id string) ([]Response, error) {
	var responses []Response
	conn, err := driver.OpenPool()
	if err != nil {
		return responses, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (response:Response{ Id:{ Id } }) RETURN response", map[string]interface{}{
		"Id": id,
	})

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
		if Id, ok := node.Properties["Id"].(string); ok {
			response.Id = Id
		}
		if R, ok := node.Properties["R"].(string); ok {
			response.R = R
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func UpdateAllResponseById(id string, response Response) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (response:Response{ Id:{ Id } }) SET response += { Id:{responseId}, R:{responseR} }", map[string]interface{}{
		"Id":         id,
		"responseId": response.Id,
		"responseR":  response.R,
	})
	return err
}

func DeleteAllResponseById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (response:Response{ Id:{ Id }) DETACH DELETE response", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetResponseByR(r string) (Response, error) {
	response := Response{}

	conn, err := driver.OpenPool()
	if err != nil {
		return response, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (response:Response{ R:{ R } }) RETURN response", map[string]interface{}{
		"R": r,
	})
	if err != nil {
		return response, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return response, NoResponseFound
	}
	if err != nil {
		return response, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return response, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		response.Id = Id
	}
	if R, ok := node.Properties["R"].(string); ok {
		response.R = R
	}

	return response, nil
}

func GetOnlyOneResponseByR(r string) (Response, error) {
	response := Response{}

	conn, err := driver.OpenPool()
	if err != nil {
		return response, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (response:Response{ R:{ R } }) RETURN response", map[string]interface{}{
		"R": r,
	})

	if err != nil {
		return response, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return response, NoResponseFound
	}
	if err != nil {
		return response, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return response, MultipleResponseFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return response, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		response.Id = Id
	}
	if R, ok := node.Properties["R"].(string); ok {
		response.R = R
	}

	return response, nil
}

func GetAllResponseByR(r string) ([]Response, error) {
	var responses []Response
	conn, err := driver.OpenPool()
	if err != nil {
		return responses, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (response:Response{ R:{ R } }) RETURN response", map[string]interface{}{
		"R": r,
	})

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
		if Id, ok := node.Properties["Id"].(string); ok {
			response.Id = Id
		}
		if R, ok := node.Properties["R"].(string); ok {
			response.R = R
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func UpdateAllResponseByR(r string, response Response) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (response:Response{ R:{ R } }) SET response += { Id:{responseId}, R:{responseR} }", map[string]interface{}{
		"R":          r,
		"responseId": response.Id,
		"responseR":  response.R,
	})
	return err
}

func DeleteAllResponseByR(r string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (response:Response{ R:{ R }) DETACH DELETE response", map[string]interface{}{
		"R": r,
	})
	return err
}

func GetResponseByCustom(query map[string]interface{}) (Response, error) {
	response := Response{}

	conn, err := driver.OpenPool()
	if err != nil {
		return response, err
	}
	defer conn.Close()

	queryStr := "MATCH (response:Response{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN response"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return response, NoResponseFound
	}
	if err != nil {
		return response, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return response, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		response.Id = Id
	}
	if R, ok := node.Properties["R"].(string); ok {
		response.R = R
	}

	return response, nil
}

func GetOnlyOneResponseByCustom(query map[string]interface{}) (Response, error) {
	response := Response{}

	conn, err := driver.OpenPool()
	if err != nil {
		return response, err
	}
	defer conn.Close()

	queryStr := "MATCH (response:Response{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN response"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return response, NoResponseFound
	}
	if err != nil {
		return response, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return response, MultipleResponseFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return response, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		response.Id = Id
	}
	if R, ok := node.Properties["R"].(string); ok {
		response.R = R
	}

	return response, nil
}

func GetAllResponseByCustom(query map[string]interface{}) ([]Response, error) {
	var responses []Response

	conn, err := driver.OpenPool()
	if err != nil {
		return responses, err
	}
	defer conn.Close()

	queryStr := "MATCH (response:Response{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN response"

	rows, err := conn.QueryNeo(queryStr, query)
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
		if Id, ok := node.Properties["Id"].(string); ok {
			response.Id = Id
		}
		if R, ok := node.Properties["R"].(string); ok {
			response.R = R
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func UpdateAllResponseByCustom(params map[string]interface{}, response Response) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (response:Response{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET response += { Id:{responseId}, R:{responseR} }"

	params["responseId"] = response.Id
	params["responseR"] = response.R

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllResponseByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (response:Response{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE response"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
