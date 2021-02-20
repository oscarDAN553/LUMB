package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/oscarDAN553/LUMB/bd"
)

/*ConsultaObjetosXCaracteristicas realiza la busqueda de objetos con caracteristicas en comun*/
func ConsultaObjetosXCaracteristicas(w http.ResponseWriter, r *http.Request) {
	caract := r.URL.Query().Get("caracteristica")
	pag := r.URL.Query().Get("pagina")

	pagina, err := strconv.Atoi(pag)

	if err != nil {
		http.Error(w, "DEBE ENVIAR UN NUMERO DE PAGINA SUPERIOR A 0 ", 400)
		return
	}
	if caract == "" {
		http.Error(w, "NO HAS ENVIADO UNA CARACTERISTICA DEL PRODUCTO", 400)
		return
	}
	if len(caract) <= 2 {
		http.Error(w, "DESCRIPCION MUY CORTA!!, DANOS MAS DETALLES", 400)
		return
	}

	objetos, err := bd.BuscoXCaracteristicas(caract, int64(pagina))

	if err != nil {
		http.Error(w, " SIN COINCIDENCIAS "+err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(objetos)
}
