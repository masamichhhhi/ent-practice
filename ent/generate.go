package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("email_adress").Unique(),
	}
}
