package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*BuscoXTag busca objetos en la DB por un tag elegido*/
func BuscoXTag(tag string) ([]models.Objeto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("LUMB")
	col := db.Collection("objects")

	var objetos []models.Objeto

	condicion := bson.M{
		"tag": tag,
	}

	cursor, err := col.Find(ctx, condicion)
	if err != nil {
		fmt.Println("NO SE ENCONTRO NINGUN PRODUCTO CON ESE TAG")
		return objetos, err
	}
	errDecode := cursor.Decode(&objetos)
	if errDecode != nil {
		fmt.Println("NO SE DECODIFICARON LOS OBJETOS ENCONTADOS POR ESE TAG")
		return objetos, errDecode
	}
	return objetos, nil

}
