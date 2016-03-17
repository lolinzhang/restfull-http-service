package model

import (
    // "fmt"
)

type User struct {
	Id int
	Name string
	Type string
}

type Relationship struct {
	Id int
	UserId int
	OtherId int
	State string
	Type string
}