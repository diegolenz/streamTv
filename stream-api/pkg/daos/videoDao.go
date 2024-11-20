package daos

import (
	"context"
	"errors"
	"log"
	"stream-api/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Title  string
	Author string
}

/* type VideoDAO struct {
    collection *mongo.Collection
} */

/*
var collectionVideo *mongo.Collection

	func NewVideoDAO() {
	    collectionVideo := DbConection.Database.Collection("videos")
	}
*/
func InsertVideo(chunk models.VideoChunk) error {

	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	coll := DbConection.Database.Collection("stream_movies")
	//doc := Book{Title: "Atonement", Author: "Ian McEwan"}
	_, err := coll.InsertOne(context.TODO(), chunk)

	//_, err := DbConection.Database.Collection("stream_movies").InsertOne(ctx, chunk)
	if err != nil {
		log.Println("Failed to insert chunk:", err)
	}
	return err
}

func FindChunkByIndex(index int) (*models.VideoChunk, error) {

	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	coll := DbConection.Database.Collection("stream_movies")
	//doc := Book{Title: "Atonement", Author: "Ian McEwan"}

	filter := bson.M{"index": index}

	var videoChunk models.VideoChunk
	err := coll.FindOne(context.TODO(), filter).Decode(&videoChunk)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Nenhum documento encontrado
		}
		log.Println("Erro ao buscar chunk:", err)
		return nil, err
	}

	return &videoChunk, nil
}
