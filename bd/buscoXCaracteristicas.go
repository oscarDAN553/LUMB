package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*BuscoXCaracteristicas realiza una busqueda en la DB de varios objetos por palabras clave*/
func BuscoXCaracteristicas(caract string, page int64) ([]*models.Objeto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("LUMB")
	col := db.Collection("objects")

	var objetos []*models.Objeto

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	//condicion := bson.M {or: [{ "marca": caract}, {"material": caract}]}
	condicion := bson.M{"$or": []bson.M{
		{"marca": bson.M{"$regex": `(?i)` + caract}},
		{"material": bson.M{"$regex": `(?i)` + caract}},
		{"color": bson.M{"$regex": `(?i)` + caract}},
		{"contenido": bson.M{"$regex": `(?i)` + caract}},
		{"capacidad": bson.M{"$regex": `(?i)` + caract}},
		{"codigoDeBarras": bson.M{"$regex": `(?i)` + caract}},
		{"empaque": bson.M{"$regex": `(?i)` + caract}},
		{"tags": bson.M{"$regex": `(?i)` + caract}},
	}}
	//"marca": bson.M{"$regex": `(?i)` + caract},
	//"color": bson.M{"$regex": `(?i)` + caract},

	// "contenido":      caract,
	// "empaque":        caract,
	// "capacidad":      caract,
	// "codigoDeBarras": caract,
	// "tags":           caract,
	//}

	cursor, err := col.Find(ctx, condicion, findOptions)

	if err != nil {
		fmt.Println("NO TENEMOS REGISTRO DE PRODUCTOS CON ESA DESCRIPCION")
		return objetos, err
	}
	for cursor.Next(ctx) {
		var obj models.Objeto
		errDecode := cursor.Decode(&obj)

		if errDecode != nil {
			fmt.Println("NO SE DECODIFICARON LOS PRODUCTOS ENCONTRADOS " + err.Error())
			return objetos, errDecode
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
