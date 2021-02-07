package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Objeto es el modelo de un producto en la DB*/
type Objeto struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
	//Nombre string  `bson:"nombre" json:"nombre,omitempty"`
	Marca           string    `bson:"marca" json:"marca,omitempty"`
	CodigoDeBarras  string    `bson:"codigoDeBarras" json:"codigoDeBarras"`
	Color           string    `bson:"color" json:"color"`
	Contenido       string    `bson:"contenido" json:"contenido,omitempty"`
	Empaque         string    `bson:"empaque" json:"empaque"`
	Capacidad       string    `bson:"capacidad" json:"capacidad,omitempty"`
	Material        string    `bson:"material" json:"material,omitempty"`
	FechaDeRegistro time.Time `bson:"fechaDeRegistro" json:"fechaDeRegistro,omitempty"`
	Imagen          string    `bson:"imagen" json:"imagen"`
}
