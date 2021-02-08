package bd

import (
	"context"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteProducto verifica en la DB si el producto ya fue registrado*/
func ChequeoYaExisteProducto(barras string) (models.Objeto, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("LUMB")
	col := bd.Collection("objets")

	condicion := bson.M{"codigoDeBarras": barras}

	var resultado models.Objeto

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	productID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, productID
	}
	return resultado, true, productID
}
