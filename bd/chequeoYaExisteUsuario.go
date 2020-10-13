package bd

import (
	"context"
	"time"

	"github.com/FabrizzioRuiz/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya está en la base de datos
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	cond := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx, cond).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
