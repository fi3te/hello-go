package main

type User struct {
	Id       int64  `json:"id"`
	Forename string `json:"forename"`
	Surname  string `json:"surename"`
}
