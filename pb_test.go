package main

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/masamichhhhi/ent-grpc-example/ent/category"
	"github.com/masamichhhhi/ent-grpc-example/ent/enttest"
	"github.com/masamichhhhi/ent-grpc-example/ent/proto/entpb"
	"github.com/masamichhhhi/ent-grpc-example/ent/user"
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

func TestServiceWithEdges(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	svc := entpb.NewUserService(client)

	cat := client.Category.Create().SetName("cat_1").SaveX(ctx)

	create, err := svc.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Name:         "user",
			EmailAddress: "user@service.code",
			Administered: []*entpb.Category{
				{Id: int32(cat.ID)},
			},
		},
	})

	if err != nil {
		t.Fatal("failed creating user using UserService", err)
	}

	count, err := client.Category.
		Query().
		Where(
			category.HasAdminWith(
				user.ID(int(create.Id)),
			),
		).Count(ctx)

	if err != nil {
		t.Fatal("failed counting categories admin by created user", err)
	}

	if count != 1 {
		t.Fatal("expected exactly one group to managed by the created user")
	}
}
