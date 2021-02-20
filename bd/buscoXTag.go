package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*BuscoXTag busca objetos en la DB por un tag elegido*/
func BuscoXTag(tag string, page int64) ([]*models.Objeto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("LUMB")
	col := db.Collection("objects")

	var objetos []*models.Objeto

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	condicion := bson.M{

		"tags": tag,
	}

	cursor, err := col.Find(ctx, condicion, findOptions)
	if err != nil {
		fmt.Println("NO SE ENCONTRO NINGUN PRODUCTO CON ESE TAG" + err.Error())
		return objetos, err
	}
	for cursor.Next(ctx) {
		var obj models.Objeto

		err := cursor.Decode(&obj)
		if err != nil {
			fmt.Println(err.Error())
			return objetos, err
		}
		objetos = append(objetos, &obj)
	}
	err = cursor.Err()
	if err != nil {
		return objetos, err
	}
	cursor.Close(ctx)
	return objetos, nil
}
