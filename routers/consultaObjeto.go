package routers

import (
	"encoding/json"
	"net/http"

	"github.com/oscarDAN553/LUMB/bd"
)

/*ConsultaObjeto retorna el json de un objeto */
func ConsultaObjeto(w http.ResponseWriter, r *http.Request) {
	cBarras := r.URL.Query().Get("codigoDeBarras")

	if len(cBarras) < 12 {
		http.Error(w, "EL CODIGO DE BARRAS DEBE CONTENER ALMENOS 12 CARACTERES", 400)
		return
	}
	objeto, err := bd.BuscoObjeto(cBarras)

	if err != nil {
		http.Error(w, "ERROR EN LA CONSULTA DEL OBJETO"+err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(objeto)
}
