package model

import (
	"fmt"
)

type User struct {
	name               string
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.name)
}