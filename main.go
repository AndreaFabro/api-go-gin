package main

import (
	"github.com/AndreaFabro/api-go-gin/routes"
	"github.com/gAndreaFabro/api-go-gin/database"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
