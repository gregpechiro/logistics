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

var NoSC_AreaFound = fmt.Errorf("no sC_Area found")
var MultipleSC_AreaFound = fmt.Errorf("multiple sC_Area found")

func AddSC_Area(sC_Area SC_Area) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (sC_Area:SC_Area { Id:{sC_AreaId}, Name:{sC_AreaName}, Description:{sC_AreaDescription} })", map[string]interface{}{
		"sC_AreaId":          sC_Area.Id,
		"sC_AreaName":        sC_Area.Name,
		"sC_AreaDescription": sC_Area.Description,
	})

	return err
}

func GetAllSC_Area() ([]SC_Area, error) {
	var sC_Areas []SC_Area
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Areas, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area) RETURN sC_Area", nil)
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

func IndexSC_AreaById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :SC_Area(Id)", nil)

	return err
}

func GetSC_AreaById(id string) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Id:{ Id } }) RETURN sC_Area", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetOnlyOneSC_AreaById(id string) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Id:{ Id } }) RETURN sC_Area", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Area, MultipleSC_AreaFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetAllSC_AreaById(id string) ([]SC_Area, error) {
	var sC_Areas []SC_Area
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Areas, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Id:{ Id } }) RETURN sC_Area", map[string]interface{}{
		"Id": id,
	})

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

func UpdateAllSC_AreaById(id string, sC_Area SC_Area) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Area:SC_Area{ Id:{ Id } }) SET sC_Area += { Id:{sC_AreaId}, Name:{sC_AreaName}, Description:{sC_AreaDescription} }", map[string]interface{}{
		"Id":                 id,
		"sC_AreaId":          sC_Area.Id,
		"sC_AreaName":        sC_Area.Name,
		"sC_AreaDescription": sC_Area.Description,
	})
	return err
}

func DeleteAllSC_AreaById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Area:SC_Area{ Id:{ Id }) DETACH DELETE sC_Area", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetSC_AreaByName(name string) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Name:{ Name } }) RETURN sC_Area", map[string]interface{}{
		"Name": name,
	})
	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetOnlyOneSC_AreaByName(name string) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Name:{ Name } }) RETURN sC_Area", map[string]interface{}{
		"Name": name,
	})

	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Area, MultipleSC_AreaFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetAllSC_AreaByName(name string) ([]SC_Area, error) {
	var sC_Areas []SC_Area
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Areas, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Name:{ Name } }) RETURN sC_Area", map[string]interface{}{
		"Name": name,
	})

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

func UpdateAllSC_AreaByName(name string, sC_Area SC_Area) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Area:SC_Area{ Name:{ Name } }) SET sC_Area += { Id:{sC_AreaId}, Name:{sC_AreaName}, Description:{sC_AreaDescription} }", map[string]interface{}{
		"Name":               name,
		"sC_AreaId":          sC_Area.Id,
		"sC_AreaName":        sC_Area.Name,
		"sC_AreaDescription": sC_Area.Description,
	})
	return err
}

func DeleteAllSC_AreaByName(name string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Area:SC_Area{ Name:{ Name }) DETACH DELETE sC_Area", map[string]interface{}{
		"Name": name,
	})
	return err
}

func GetSC_AreaByDescription(description string) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Description:{ Description } }) RETURN sC_Area", map[string]interface{}{
		"Description": description,
	})
	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetOnlyOneSC_AreaByDescription(description string) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Description:{ Description } }) RETURN sC_Area", map[string]interface{}{
		"Description": description,
	})

	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Area, MultipleSC_AreaFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetAllSC_AreaByDescription(description string) ([]SC_Area, error) {
	var sC_Areas []SC_Area
	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Areas, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (sC_Area:SC_Area{ Description:{ Description } }) RETURN sC_Area", map[string]interface{}{
		"Description": description,
	})

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

func UpdateAllSC_AreaByDescription(description string, sC_Area SC_Area) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Area:SC_Area{ Description:{ Description } }) SET sC_Area += { Id:{sC_AreaId}, Name:{sC_AreaName}, Description:{sC_AreaDescription} }", map[string]interface{}{
		"Description":        description,
		"sC_AreaId":          sC_Area.Id,
		"sC_AreaName":        sC_Area.Name,
		"sC_AreaDescription": sC_Area.Description,
	})
	return err
}

func DeleteAllSC_AreaByDescription(description string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (sC_Area:SC_Area{ Description:{ Description }) DETACH DELETE sC_Area", map[string]interface{}{
		"Description": description,
	})
	return err
}

func GetSC_AreaByCustom(query map[string]interface{}) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Area:SC_Area{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN sC_Area"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetOnlyOneSC_AreaByCustom(query map[string]interface{}) (SC_Area, error) {
	sC_Area := SC_Area{}

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Area, err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Area:SC_Area{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN sC_Area"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return sC_Area, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return sC_Area, NoSC_AreaFound
	}
	if err != nil {
		return sC_Area, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return sC_Area, MultipleSC_AreaFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return sC_Area, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		sC_Area.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		sC_Area.Name = Name
	}
	if Description, ok := node.Properties["Description"].(string); ok {
		sC_Area.Description = Description
	}

	return sC_Area, nil
}

func GetAllSC_AreaByCustom(query map[string]interface{}) ([]SC_Area, error) {
	var sC_Areas []SC_Area

	conn, err := driver.OpenPool()
	if err != nil {
		return sC_Areas, err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Area:SC_Area{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN sC_Area"

	rows, err := conn.QueryNeo(queryStr, query)
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

func UpdateAllSC_AreaByCustom(params map[string]interface{}, sC_Area SC_Area) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Area:SC_Area{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET sC_Area += { Id:{sC_AreaId}, Name:{sC_AreaName}, Description:{sC_AreaDescription} }"

	params["sC_AreaId"] = sC_Area.Id
	params["sC_AreaName"] = sC_Area.Name
	params["sC_AreaDescription"] = sC_Area.Description

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllSC_AreaByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (sC_Area:SC_Area{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE sC_Area"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
