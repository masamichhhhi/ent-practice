package main

import (
	"context"
	"log"
	"net"

	"github.com/go-sql-driver/mysql"
	"github.com/masamichhhhi/ent-grpc-example/ent"
	"github.com/masamichhhhi/ent-grpc-example/ent/proto/entpb"
	"google.golang.org/grpc"
)

func main() {
	entOption := []ent.Option{}
	entOption = append(entOption, ent.Debug())

	mc := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + "33306",
		DBName:               "test",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOption...)
	if err != nil {
		log.Fatalf("Error open mysql ent client: %v\n", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	svc := entpb.NewUserService(client)

	server := grpc.NewServer()

	entpb.RegisterUserServiceServer(server, svc)

	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
