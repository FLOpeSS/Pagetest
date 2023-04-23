package Users

import (
	"fmt"
	"unicode"
)

type User struct {
	Name      string
	Password  string
	ValidPass bool
}

func (u *User) ValidatePassword() {
	if len(u.Password) > 8 && ThereIsUpper(u.Password) {
		fmt.Println("Valid password")
		u.ValidPass = true
	} else if len(u.Password) < 8 {
		fmt.Println("Invalid password. Your password must countain at least 8 chars")
		u.ValidPass = false
	} else if ThereIsUpper(u.Password) == false {
		fmt.Println("Invalid password. Your password must countain at least a Uppercase char")
		u.ValidPass = false
	}
}

func ThereIsUpper(s string) bool {
	for _, j := range s {
		if unicode.IsUpper(j) {
			return true
		}
	}
	return false
}
