package main

import (
	"log"

	"github.com/oscarDAN553/LUMB/bd"
	"github.com/oscarDAN553/LUMB/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}
