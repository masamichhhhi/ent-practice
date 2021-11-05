package main

import (
	"testing"

	"github.com/masamichhhhi/ent-grpc-example/ent/proto/entpb"
)

func TestUserProto(t *testing.T) {
	user := entpb.User{
		Name:         "masamichi",
		EmailAddress: "masamichi@example.com",
	}
	if user.GetName() != "masamichi" {
		t.Fatal("expected user name to be masamichi")
	}
	if user.GetEmailAddress() != "masamichi@example.com" {
		t.Fatal("expected email address to be masamichi@example.com")
	}
}
