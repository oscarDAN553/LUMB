package middlew

import (
	"net/http"

	"github.com/oscarDAN553/LUMB/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la DB*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la DB", 500)
			return
		}
		next.ServeHTTP(w, r)

	}

}
