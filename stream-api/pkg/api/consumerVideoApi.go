package api

/* import (
	"encoding/json"
	"net/http"
	"stream-video-api/pkg/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
) */

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"stream-api/pkg/daos"

	"github.com/gorilla/mux"
)

/* type ConsumerVideoApi struct {
    videoService *services.VideoService
}

func NewConsumerVideoApi(videoService *services.VideoService) *ConsumerVideoApi {
    return &ConsumerVideoApi{videoService: videoService}
} */

func GetVideoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("REQ video")

	vars := mux.Vars(r)
	index, err := strconv.Atoi(vars["index"])

	if err != nil {
		http.Error(w, "Falha ao converter parametro", http.StatusInternalServerError)
		log.Println("Erro ao buscar chunk:", err)
		return
	}

	chunk, err := daos.FindChunkByIndex(index) //fileName, index)
	if err != nil {
		http.Error(w, "Erro ao buscar o chunk de vídeo", http.StatusInternalServerError)
		log.Println("Erro ao buscar chunk:", err)
		return
	}

	if chunk == nil {
		http.Error(w, "Chunk não encontrado", http.StatusNotFound)
		return
	}

	// Configura o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Permite todas as origens
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	// Serializa o objeto chunk em JSON e envia como resposta
	if err := json.NewEncoder(w).Encode(chunk); err != nil {
		http.Error(w, "Erro ao serializar a resposta JSON", http.StatusInternalServerError)
		log.Println("Erro ao serializar JSON:", err)
	}
}
