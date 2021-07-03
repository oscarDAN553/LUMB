package bd

import (
	"context"
	"errors"
	"fmt"

	//"strings"
	"time"

	"github.com/oscarDAN553/LUMB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*BuscoXCaracteristicas realiza una busqueda en la DB de varios objetos por palabras clave*/
func BuscoXCaracteristicas(caract string, page int64) ([]*models.Objeto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("LUMB")
	col := db.Collection("objectss")
	//s := col.CreateIndex( bson.M{"$marca" : "text"})
	//s := options.Create
	//model := mongo.IndexModel()
	//op := options.CreateIndexesOptions(bson.M{"$marca" : "text"},)
	// f := mongo.NewIndexOptionsBuilder().TextVersion(3)
	// col.Indexes().CreateMany()
	var objetos []*models.Objeto
	//var caracteristicas []string
	//var condicionx []bson.M

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	// res:= col.Indexes()
	// ok, err := res.CreateOne(ctx,{"marca": 1,findOptions})

	//mod := mongo.IndexModel{Keys: bson.M{"marca": "text"}, Options: nil}

	// opt := options.Index()
	// opt.SetUnique(false)
	//opt2 := options.DefaultIndex()
	opt3 := options.Index()
	opt3.SetDefaultLanguage("es")
	//opt3.SetSparse(true)

	index := mongo.IndexModel{Keys: bson.M{
		//wildcard ayuda a realizar un seach text en todos los field que contengan strings de la DB
		"color":     "text",
		"contenido": "text",
		"material":  "text",
	}, Options: opt3}

	_, err := col.Indexes().CreateOne(ctx, index)

	// splitCaract := strings.Split(caract, " ")

	// for i := 0; i < len(splitCaract); i++ {
	// 	//var caracteri string
	// 	trimCaract := strings.TrimSpace(splitCaract[i])

	// 	caracteristicas = append(caracteristicas, trimCaract)
	// }

	/*condiciony funciona pero se comporta como AND/ la opcion x quita espacios y la opcion i permite mayusculas*/
	//condiciony := bson.M{"marca": bson.M{"$regex": caract, "$options": "xi"}}

	/*condicionx funciona, permite obtener docs con varias caracteristicas pero no permite regex para mayusculas y palabras incompletas*/
	//condicionx := bson.M{"marca": bson.M{"$in": caracteristicas}}

	condicionw := bson.M{"$text": bson.M{"$search": caract, "$language": "es"}}

	/*condicionz no funciona se busco fucionar el comportamiento OR y opcion de mayusculas y palablas incompletas*/
	//condicionz := bson.M{"marca": bson.M{"$regex": bson.M{"$in": caracteristicas}, "$options": "i"}}

	//condicionx = append(condicionx, condicion)

	//condicion := bson.M {or: [{ "marca": caract}, {"material": caract}]}
	/*intento de escanedo de slide de conciciones no funciono*/
	// for i := 0; i < len(caracteristicas); i++ {

	// 	condicion := bson.M{"$or": []bson.M{
	// 		//{"marca": bson.M{"$or": []bson.M{{caracteristicas[0]},{caracteristicas[1]},}}},
	// 		{"$or": []bson.M{{"marca": bson.M{"$regex": `(?i)` + caracteristicas[i]}}}},
	// 		// {"material": bson.M{"$regex": `(?i)` + caracteristicas[0]}},
	// 		// {"color": bson.M{"$regex": `(?i)` + caract}},
	// 		// {"contenido": bson.M{"$regex": `(?i)` + caract}},
	// 		// {"capacidad": bson.M{"$regex": `(?i)` + caract}},
	// 		// {"codigoDeBarras": bson.M{"$regex": `(?i)` + caract}},
	// 		// {"empaque": bson.M{"$regex": `(?i)` + caract}},
	// 		//{"tags": bson.M{"$regex": `(?i)` + caracteristicas[c]}},
	// 	}}
	// 	condicionx = append(condicionx, condicion)
	// }

	// // for x := 0; x < len(condicionx); x++ {

	// // }
	// cond := bson.M{"$or": condicionx}
	// //"marca": bson.M{"$regex": `(?i)` + caract},
	// //"color": bson.M{"$regex": `(?i)` + caract},

	// // "contenido":      caract,
	// // "empaque":        caract,
	// // "capacidad":      caract,
	// // "codigoDeBarras": caract,
	// // "tags":           caract,
	// //}
	//index := col.Indexes("$marca": "text")
	cursor, err := col.Find(ctx, condicionw, findOptions)

	if err != nil {
		//fmt.Println("NO TENEMOS REGISTRO DE PRODUCTOS CON ESA DESCRIPCION")
		return objetos, errors.New("NO TENEMOS REGISTRO DE PRODUCTOS CON ESA DESCRIPCION")
	}
	if cursor.RemainingBatchLength() == 0 {
		return objetos, errors.New("NO TENEMOS REGISTRO DE PRODUCTOS CON ESA DESCRIPCION")
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
		return objetos, errors.New("NO TENEMOS REGISTRO DE PRODUCTOS CON ESA DESCRIPCION")
	}
	cursor.Close(ctx)
	return objetos, nil

}
