package bd

import (
	"context"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistroProducto añade el producto a la DB*/
func InsertoRegistroProducto(t models.Objeto) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("LUMB")
	col := bd.Collection("products")

	result, err := col.InsertOne(ctx, t)

	if err != nil {
		return "", false, err
	}
	objID := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
