package service

import (
	"fmt"
	"log"
)

type User struct {
	id       int
	name     string
	email    string
	userName string
	password string
}

func (u *User) Get(id int) (User, Error) {

}

func (u *User) Remove(id int) Error {

}

func (u *User) List() ([]User, Error) {

}
