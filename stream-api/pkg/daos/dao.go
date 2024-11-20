package daos

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB struct para a conexão
type MongoDbType struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var DbConection *MongoDbType

// NewMongoDB cria uma nova instância de conexão MongoDB
func Connect() (*MongoDbType, error) {

	//uri := "mongodb://diegomatheuslenzz:viPbgSi9LyNNaZ19@177.136.220.63/?retryWrites=true&w=majority"
	uri := "mongodb+srv://diegomatheuslenzz:viPbgSi9LyNNaZ19@cluster0.souyc.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("falha ao verificar conexão com MongoDB: %v", err)
	}

	log.Println("Conexão com MongoDB Atlas estabelecida com sucesso.")

	db := client.Database("stream")

	DbConection = &MongoDbType{
		Client:   client,
		Database: db,
	}

	return DbConection, nil
}

// Fecha a conexão
func (m *MongoDbType) Disconnect() {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		log.Printf("Erro ao desconectar do MongoDB: %v", err)
	} else {
		log.Println("Desconexão do MongoDB Atlas concluída.")
	}
}
