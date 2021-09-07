package main

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"strings"
	"text/template"

	app "github.com/ozonva/ova-person-api/internal/app"
	repo "github.com/ozonva/ova-person-api/internal/repo"
	utils "github.com/ozonva/ova-person-api/internal/utils"
	desc "github.com/ozonva/ova-person-api/pkg"
)

func main() {
	configPath := utils.GetConfigPath()
	fmt.Println(configPath)

	set, _ := utils.GetConfig(configPath)
	fmt.Println("Server is starting...")

	grpcPort := format(`{{index . "host"}}:{{index . "port"}}`, set)
	url := format(`{{index . "url"}}`, set)

	if err := run(grpcPort, url); err != nil {
		log.Err(err)
		log.Fatal()
	}
}

func createRepository(url string) repo.PersonRepo {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Print(err)
		log.Fatal()
	}
	log.Printf("Connection with db %v is successful", url)
	return repo.NewContext(*db)
}

func run(grpcPort string, url string) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		log.Fatal()
	}

	server := grpc.NewServer()
	repository := createRepository(url)
	desc.RegisterPersonApiServiceServer(server, app.NewPersonApi(repository))

	if err := server.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
		log.Fatal()
	}

	return nil
}

func format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	err := template.Must(t.Parse(s)).Execute(b, v)
	if err != nil {
		log.Err(err)
		return ""
	}
	return b.String()
}
