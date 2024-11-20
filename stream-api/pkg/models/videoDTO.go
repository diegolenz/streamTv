package models

type Video struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VideoChunk struct {
	ID       string `bson:"_id,omitempty"`
	Chunk    []byte `bson:"chunk"`
	FileName string `bson:"fileName"`
	Index    int    `bson:"index"`
}
