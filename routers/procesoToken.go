package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/oscarDAN553/LUMB/bd"
	"github.com/oscarDAN553/LUMB/models"
)

/*Email es el valor de usuario usado en todos los endpoinds*/
var Email string

/*IDUsuario es el valor que se usara en todos los endpoinds*/
var IDUsuario string

/*ProcesoToken extrae los valores del token*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("en nombre de la vida")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
