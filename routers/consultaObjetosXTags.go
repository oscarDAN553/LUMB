package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/oscarDAN553/LUMB/bd"
)

/*ConsultaObjetosXTag ejecuta la rutina para buscar varios objetos de un tag dado*/
func ConsultaObjetosXTag(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	pag := r.URL.Query().Get("pagina")
	pagina, err := strconv.Atoi(pag)

	if err != nil {
		http.Error(w, "DEBE ENVIAR UN NUMERO DE PAGINA SUPERIOR A 0 ", 400)
		return
	}

	if len(tag) < 0 || tag == "" {
		http.Error(w, "DEBE ENVIAR POR LO MENOS UNA CARACTERISTICA DEL PRODUCTO", 400)
		return
	}

	objetos, err := bd.BuscoXTag(tag, int64(pagina))

	if err != nil {
		http.Error(w, "ERROR EN CONSULTA DE LOS OBJETOS DE ESTE TAG "+err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(objetos)
}
