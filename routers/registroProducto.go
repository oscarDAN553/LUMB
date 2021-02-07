package routers

import (
	"encoding/json"
	"net/http"

	"github.com/oscarDAN553/LUMB/bd"
	"github.com/oscarDAN553/LUMB/models"
)

/*RegistroProducto añade el registro de un nuevo producto a la DB*/
func RegistroProducto(w http.ResponseWriter, r *http.Request) {
	var t models.Objeto

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "ERROR AL DECODIFICAR LOS DATOS DEL PRODUCTO"+err.Error(), 400)
		return
	}
	if len(t.Color) == 0 {
		http.Error(w, "SE DEBE ESPECIFICAR EL COLOR DEL PRODUCTO A REGISTRAR ", 400)
		return
	}
	if len(t.CodigoDeBarras) == 0 {
		http.Error(w, "EL CODIGO DE BARRAS DEL PRODUCTO ES REQUERIDO ", 400)
		return
	}
	if len(t.CodigoDeBarras) < 12 {
		http.Error(w, "EL CODIGO DE BARRAS DEBE TENER AL MENOS 12 CARACTERES", 400)
		return
	}

	_, existe, _ := bd.ChequeoYaExisteProducto(t.CodigoDeBarras)

	if existe == true {
		http.Error(w, "ESTE CODIGO DE BARRAS YA FUE REGISTRADO", 400)
	}
	_, status, err := bd.InsertoRegistroProducto(t)

	if err != nil {
		http.Error(w, "ERROR AL REGISTRAR EL PRODUCTO"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "EL PRODUCTO NO SE AGREGO A LA DB", 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
