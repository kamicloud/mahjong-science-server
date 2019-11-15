package models

import (
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
}
