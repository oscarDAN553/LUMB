package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Hortaliza es el modelo de registro de las recompensas*/
type Hortaliza struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
	Nombre              string             `bson:"nombre" json:"nombre,omitempty"`
	Tipo                string             `bson:"tipo" json:"tipo,omitempty"`
	TiempoDeCrecimiento int                `bson:"tiempoDeCrecimiento" json:"tiempoDeCrecimiento,omitempty"`
	Descripcion         string             `bson:"descripcion" json:"descripcion,omitempty"`
	PuntosRequeridos    int                `bson:"puntosRequeridos" json:"puntosRequeridos,omitempty"`
	Imagen              string             `bson:"imagen" json:"imagen,omitempty"`
	TimeLapse           string             `bson:"timeLapse" json:"timeLapse,omitempty"`
}
