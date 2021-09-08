package app

import (
	desc "github.com/ozonva/ova-person-api/pkg"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-person-api/internal/repo"
)

func RunPersonServer(grpcPort string, dsn string) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		log.Fatal()
	}

	server := grpc.NewServer()
	repository := createRepository(dsn)
	desc.RegisterPersonApiServiceServer(server, NewPersonApi(repository))

	log.Printf("Server is started!!! Port: %v", grpcPort)

	if err := server.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
		log.Fatal()
	}

	return nil
}

func createRepository(dsn string) repo.PersonRepo {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Print(err)
		log.Fatal()
	}
	log.Printf("Connection with db %v is successful", dsn)
	return repo.NewContext(*db)
}
