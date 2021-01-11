package routers

import (
	"encoding/json"
	"net/http"

	"github.com/oscarDAN553/LUMB/bd"
	"github.com/oscarDAN553/LUMB/models"
)

/*Registro es la funcion para crear el registro del usuario en la DB*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "ERROR EN LOS DATOS RECIBIDOS "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "EL EMAIL DE USUARIO ES REQUERIDO ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "DEBE USAR UN PASSWOR DE AL MENOS 6 CARACTERES", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "ESTE EMAIL YA FUE REGISTRADO", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "OCURRIO UN ERROR AL INTENTAR REALIZAR EL REGISTRO DEL USUARIO"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "NO SE LOGRO INSERTAR EL REGISTRO DE USUARIO", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
