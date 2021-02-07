package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oscarDAN553/LUMB/middlew"
	"github.com/oscarDAN553/LUMB/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo el puerto,handler y pongo a ecuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/registroproducto", middlew.ChequeoBD(middlew.ValidoJWT(routers.RegistroProducto))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
