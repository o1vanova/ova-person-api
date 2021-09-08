package main

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/ozonva/ova-person-api/internal/app"
	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-person-api/internal/utils"
)

func main() {
	configPath := utils.GetConfigPath()
	fmt.Println(configPath)

	set, _ := utils.GetConfig(configPath)
	fmt.Println("Server is starting...")

	grpcPort := utils.Format(`{{index . "host"}}:{{index . "port"}}`, set)
	dsn := utils.Format(`{{index . "dsn"}}`, set)

	if err := app.RunPersonServer(grpcPort, dsn); err != nil {
		log.Err(err)
		log.Fatal()
	}
}
