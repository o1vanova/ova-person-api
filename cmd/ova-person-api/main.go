package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"strings"
	"text/template"

	app "github.com/ozonva/ova-person-api/internal/app"
	utils "github.com/ozonva/ova-person-api/internal/utils"
	desc "github.com/ozonva/ova-person-api/pkg"
)

func main() {
	configPath := utils.GetConfigPath()
	fmt.Println(configPath)

	set, _ := utils.GetConfig(configPath)
	fmt.Println("Server is starting...")

	if err := run(format(`{{index . "host"}}:{{index . "port"}}`, set)); err != nil {
		log.Print(err)
		log.Fatal()
	}
}

func run(grpcPort string) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		log.Fatal()
	}

	server := grpc.NewServer()
	desc.RegisterPersonApiServiceServer(server, app.NewPersonApi())

	if err := server.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
		log.Fatal()
	}

	return nil
}

func format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(b, v)
	return b.String()
}
