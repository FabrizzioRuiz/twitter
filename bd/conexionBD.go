package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto creado de conexión a Mongo
var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://frp466:africarubi2020@twitter.2oahp.mongodb.net/<dbname>?retryWrites=true&w=majority")

//ConectarBD Metodo que nos permite conectarnos a la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

//ChequeoConnection nos varifica la conexion
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
