package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

func getUser() (*User, error) {
	return nil, errors.New("User == nil")
}

func main() {
	if u, err := getUser(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.Name)
	}
}
