package main

import (
	"valentine-card/internal/adapters/audio"
	"valentine-card/internal/adapters/storage"
	"valentine-card/internal/application"
	"valentine-card/internal/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Vraies implémentations pour StoragePort et AudioProcessor
	storageImpl := storage.NewLocalStorage()         // Exemple : stockage local
	audioProcessorImpl := audio.NewFfmpegProcessor() // Exemple : utilisation de FFmpeg

	audioService := application.NewAudioService(storageImpl, audioProcessorImpl)

	// Endpoint pour uploader un fichier audio
	r.POST("/upload", func(c *gin.Context) {
		var a domain.Audio
		if err := c.ShouldBindJSON(&a); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Simuler des données de fichier (à remplacer par une lecture réelle)
		fileData := []byte("file data")
		err := audioService.UploadAudio(&a, fileData)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"url": a.URL})
	})

	// Endpoint pour traiter un fichier audio
	r.POST("/process", func(c *gin.Context) {
		var input struct {
			InputPath string `json:"input_path"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		outputPath, err := audioService.ProcessAudio(input.InputPath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"output_path": outputPath})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	} // Lancer le serveur sur le port 8080
}
