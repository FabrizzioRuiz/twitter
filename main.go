package main

import (
	"log"

	"github.com/FabrizzioRuiz/twitter/bd"
	"github.com/FabrizzioRuiz/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
