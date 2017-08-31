package main

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

/*

MATCH
(e:SC_Element{Id:"1504114258945128188"}),
(q:Question{Id:"1504114289214666257"}),
(a:SC_Area{Id:"1504112043762496770"})
MERGE (e)-[:ASKS_IN]->(h:Hyper)-[:ASKS]->(q)
MERGE (h)-[:IN]->(a)

*/

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
