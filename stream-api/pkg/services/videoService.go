package services

import (
	"io"
	"mime/multipart"
	"stream-api/pkg/daos"
	"stream-api/pkg/models"
)

/* type VideoService struct {
    videoDAO *daos.VideoDAO
} */

func ProcessAndSaveChunks(file multipart.File, fileName string) error {
	buffer := make([]byte, 4*1024*1024) // 4MB chunks
	index := 0

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if bytesRead == 0 {
			break
		}

		chunk := models.VideoChunk{
			Chunk:    buffer[:bytesRead],
			FileName: fileName,
			Index:    index,
		}

		if err := daos.InsertVideo(chunk); err != nil {
			return err
		}
		index++
	}

	return nil
}
