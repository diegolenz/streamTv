package api

/* import (
	"encoding/json"
	"net/http"
	"stream-video-api/pkg/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
) */

import (
	"log"
	"net/http"
	"stream-api/pkg/services"
)

/* type ConsumerVideoApi struct {
    videoService *services.VideoService
}

func NewConsumerVideoApi(videoService *services.VideoService) *ConsumerVideoApi {
    return &ConsumerVideoApi{videoService: videoService}
} */

func UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("REQ")
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // Limite de 10 MB
	err := r.ParseMultipartForm(10 << 20)           // Define o limite de 10 MB

	if err != nil {
		http.Error(w, "O arquivo é muito grande", http.StatusRequestEntityTooLarge)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := services.ProcessAndSaveChunks(file, header.Filename); err != nil {
		http.Error(w, "Failed to process video", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Video uploaded successfully"))
}
