package helper

import (
	"regexp"
)

func HelloWorld(name string) string {
	return "Hello World " + name
}

func HelloWorldSub(name string) string {
	return "Hello World Sub " + name
}

func ValidatePassword(password string) bool {
	match, _ := regexp.MatchString(`[^a-zA-Z0-9]`, password)

	if len(password) < 8 || password == "" {
		return false
	} else if !match {
		return false
	} else {
		return true
	}
}
