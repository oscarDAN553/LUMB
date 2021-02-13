package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*BuscoObjeto busca en la DB un objeto con el parametro codigo de barras*/
func BuscoObjeto(cBarras string) (models.Objeto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("LUMB")
	col := db.Collection("objects")

	var objeto models.Objeto

	condicion := bson.M{
		"codigoDeBarras": cBarras,
	}

	err := col.FindOne(ctx, condicion).Decode(&objeto)
	if err != nil {
		fmt.Println("PRODUCTO NO ENCONTRADO")
		return objeto, err
	}
	return objeto, nil
}
