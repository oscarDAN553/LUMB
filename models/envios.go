package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Envios es el modelo de envios del usuario registrados en la DB*/
type Envios struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"Id"`
	CanastaID      string             `bson:"canastaId" json:"canastaId"`
	FechaSolicitud time.Time          `bson:"fechaSolicitud" json:"fechaSolicitud,omitempty"`
	FechaRecepcion time.Time          `bson:"fechaRecepcion" json:"fechaRecepcion,omitempty"`
	UsuarioID      string             `bson:"usuarioId" json:"usuarioId"`
	Recibido       bool               `bson:"recibido" json:"recibido,omitempty"`
	Direccion      string             `bson:"direccion" json:"direccion"`
	RecolectorID   string             `bson:"recolectorId" json:"recolectorId"`
	Puntos         int                `bson:"puntos" json:"puntos,omitempty"`
}
