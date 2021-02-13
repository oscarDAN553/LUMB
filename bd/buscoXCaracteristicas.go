package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*BuscoXCaracteristicas realiza una busqueda en la DB de varios objetos por palabras clave*/
func BuscoXCaracteristicas(caract string) ([]models.Objeto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("LUMB")
	col := db.Collection("objects")

	var objetos []models.Objeto

	condicion := bson.M{

		"marca":          caract,
		"color":          caract,
		"contenido":      caract,
		"empaque":        caract,
		"capacidad":      caract,
		"material":       caract,
		"codigoDeBarras": caract,
		"tags":           caract,
	}

	cursor, err := col.Find(ctx, condicion)
	//err := col.Find(ctx, condicion).Decode(&objeto)
	if err != nil {
		fmt.Println("NO SE ENCONTRO NINGUN PRODUCTO CON ESA DESCRIPCION")
		return objetos, err
	}
	errDecode := cursor.Decode(&objetos)

	if errDecode != nil {
		fmt.Println("NO SE DECODIFICARON LOS PRODUCTOS ENCONTRADOS ")
		return objetos, errDecode
	}
	return objetos, nil

}
