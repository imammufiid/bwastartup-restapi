package helper

import "golang.org/x/text/message"

// 1. create object response 
type Response struct {
	Meta Meta
	Data interface{} // why interface{}? bcoz value of the data can change
}

type Meta struct {
	Message string
	Code    int
	Status  string
}

