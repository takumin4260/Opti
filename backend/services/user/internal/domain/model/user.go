package model

import "github.com/kinoshitatakumi/opti/pkg/domain/value"

type User struct {
	ID    string
	Email value.Email
	Name  string
}
