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

var NoSC_ElementFound = fmt.Errorf("no sC_Element found")
var MultipleSC_ElementFound = fmt.Errorf("multiple sC_Element found")

func AddSC_Element(sC_Element SC_Element) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (sC_Element:SC_Element { Id:{sC_ElementId}, Name:{sC_ElementName}, Description:{sC_ElementDescription} })", map[string]interface{}{
		"sC_ElementId":          sC_Element.Id,
		"sC_ElementName":        sC_Element.Name,
		"sC_ElementDescription": sC_Element.Description,
	})

	return err
}

func GetAllSC_Element() ([]SC_Element, error) {
	var sC_Elements []SC_Element
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Elements, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element) RETURN sC_Element", nil)
	if err != nil {
		return sC_Elements, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return sC_Elements, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return sC_Elements, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		sC_Element := SC_Element{}
		if Id, ok := node.Properties["Id"].(string); ok {
			sC_Element.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			sC_Element.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			sC_Element.Description = Description
		}

		sC_Elements = append(sC_Elements, sC_Element)
	}

	return sC_Elements, nil
}

func IndexSC_ElementById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :SC_Element(Id)", nil)

	return err
}

func GetSC_ElementById(id string) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Id:{ Id } }) RETURN sC_Element", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetOnlyOneSC_ElementById(id string) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Id:{ Id } }) RETURN sC_Element", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Element, MultipleSC_ElementFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetAllSC_ElementById(id string) ([]SC_Element, error) {
	var sC_Elements []SC_Element
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Elements, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Id:{ Id } }) RETURN sC_Element", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return sC_Elements, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return sC_Elements, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return sC_Elements, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		sC_Element := SC_Element{}
		if Id, ok := node.Properties["Id"].(string); ok {
			sC_Element.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			sC_Element.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			sC_Element.Description = Description
		}

		sC_Elements = append(sC_Elements, sC_Element)
	}

	return sC_Elements, nil
}

func UpdateAllSC_ElementById(id string, sC_Element SC_Element) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Element:SC_Element{ Id:{ Id } }) SET sC_Element += { Id:{sC_ElementId}, Name:{sC_ElementName}, Description:{sC_ElementDescription} }", map[string]interface{}{
		"Id":                    id,
		"sC_ElementId":          sC_Element.Id,
		"sC_ElementName":        sC_Element.Name,
		"sC_ElementDescription": sC_Element.Description,
	})
	return err
}

func DeleteAllSC_ElementById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Element:SC_Element{ Id:{ Id } }) DETACH DELETE sC_Element", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetSC_ElementByName(name string) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Name:{ Name } }) RETURN sC_Element", map[string]interface{}{
		"Name": name,
	})
	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetOnlyOneSC_ElementByName(name string) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Name:{ Name } }) RETURN sC_Element", map[string]interface{}{
		"Name": name,
	})

	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Element, MultipleSC_ElementFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetAllSC_ElementByName(name string) ([]SC_Element, error) {
	var sC_Elements []SC_Element
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Elements, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Name:{ Name } }) RETURN sC_Element", map[string]interface{}{
		"Name": name,
	})

	if err != nil {
		return sC_Elements, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return sC_Elements, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return sC_Elements, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		sC_Element := SC_Element{}
		if Id, ok := node.Properties["Id"].(string); ok {
			sC_Element.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			sC_Element.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			sC_Element.Description = Description
		}

		sC_Elements = append(sC_Elements, sC_Element)
	}

	return sC_Elements, nil
}

func UpdateAllSC_ElementByName(name string, sC_Element SC_Element) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Element:SC_Element{ Name:{ Name } }) SET sC_Element += { Id:{sC_ElementId}, Name:{sC_ElementName}, Description:{sC_ElementDescription} }", map[string]interface{}{
		"Name":                  name,
		"sC_ElementId":          sC_Element.Id,
		"sC_ElementName":        sC_Element.Name,
		"sC_ElementDescription": sC_Element.Description,
	})
	return err
}

func DeleteAllSC_ElementByName(name string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Element:SC_Element{ Name:{ Name } }) DETACH DELETE sC_Element", map[string]interface{}{
		"Name": name,
	})
	return err
}

func GetSC_ElementByDescription(description string) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Description:{ Description } }) RETURN sC_Element", map[string]interface{}{
		"Description": description,
	})
	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetOnlyOneSC_ElementByDescription(description string) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Description:{ Description } }) RETURN sC_Element", map[string]interface{}{
		"Description": description,
	})

	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Element, MultipleSC_ElementFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetAllSC_ElementByDescription(description string) ([]SC_Element, error) {
	var sC_Elements []SC_Element
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Elements, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Element:SC_Element{ Description:{ Description } }) RETURN sC_Element", map[string]interface{}{
		"Description": description,
	})

	if err != nil {
		return sC_Elements, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return sC_Elements, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return sC_Elements, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		sC_Element := SC_Element{}
		if Id, ok := node.Properties["Id"].(string); ok {
			sC_Element.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			sC_Element.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			sC_Element.Description = Description
		}

		sC_Elements = append(sC_Elements, sC_Element)
	}

	return sC_Elements, nil
}

func UpdateAllSC_ElementByDescription(description string, sC_Element SC_Element) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Element:SC_Element{ Description:{ Description } }) SET sC_Element += { Id:{sC_ElementId}, Name:{sC_ElementName}, Description:{sC_ElementDescription} }", map[string]interface{}{
		"Description":           description,
		"sC_ElementId":          sC_Element.Id,
		"sC_ElementName":        sC_Element.Name,
		"sC_ElementDescription": sC_Element.Description,
	})
	return err
}

func DeleteAllSC_ElementByDescription(description string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Element:SC_Element{ Description:{ Description } }) DETACH DELETE sC_Element", map[string]interface{}{
		"Description": description,
	})
	return err
}

func GetSC_ElementByCustom(query map[string]interface{}) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Element:SC_Element{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN sC_Element"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetOnlyOneSC_ElementByCustom(query map[string]interface{}) (SC_Element, error) {
	sC_Element := SC_Element{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Element, err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Element:SC_Element{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN sC_Element"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return sC_Element, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Element, NoSC_ElementFound
	}
	if err != nil {
		return sC_Element, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Element, MultipleSC_ElementFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Element, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Element.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Element.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Element.Description = Description
	}

	return sC_Element, nil
}

func GetAllSC_ElementByCustom(query map[string]interface{}) ([]SC_Element, error) {
	var sC_Elements []SC_Element

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Elements, err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Element:SC_Element{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN sC_Element"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return sC_Elements, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return sC_Elements, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return sC_Elements, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		sC_Element := SC_Element{}
		if Id, ok := node.Properties["Id"].(string); ok {
			sC_Element.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			sC_Element.Name = Name
		}
		if Description, ok := node.Properties["Description"].(string); ok {
			sC_Element.Description = Description
		}

		sC_Elements = append(sC_Elements, sC_Element)
	}

	return sC_Elements, nil
}

func UpdateAllSC_ElementByCustom(params map[string]interface{}, sC_Element SC_Element) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Element:SC_Element{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET sC_Element += { Id:{sC_ElementId}, Name:{sC_ElementName}, Description:{sC_ElementDescription} }"

	params["sC_ElementId"] = sC_Element.Id
	params["sC_ElementName"] = sC_Element.Name
	params["sC_ElementDescription"] = sC_Element.Description

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllSC_ElementByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Element:SC_Element{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE sC_Element"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
