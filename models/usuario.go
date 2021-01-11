package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo para el usuario en la DB de mongo*/
type Usuario struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre            string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos         string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaDeNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email             string             `bson:"email" json:"email"`
	Password          string             `bson:"password" json:"password,omitempty"`
	Avatar            string             `bson:"avatar" json:"avatar,omitempty"`
	Ubicacion         string             `bson:"ubicacion" json:"ubicacion,omitempty"`
}
