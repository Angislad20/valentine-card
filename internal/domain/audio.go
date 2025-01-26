package domain

import (
	"fmt"
	"time"
)

// Audio représente un fichier audio ou musical.
type Audio struct {
	ID         int           // Identifiant unique
	Title      string        // Nom ou titre du fichier
	Format     string        // Format du fichier (mp3, wav, etc.)
	Duration   time.Duration // Durée en nanosecondes
	UploadedAt time.Time     // Date et heure de l'upload
	URL        string        // URL du fichier stocké
	Bitrate    int           // Bitrate en kbps
	SampleRate int           // Fréquence d'échantillonnage (Hz)
	FileSize   int64         // Taille du fichier en octets
	Type       string        // Uploaded ou Recorded
	OwnerID    int           // ID de l'utilisateur ayant uploadé l'audio
}

// NewAudio crée une nouvelle instance d'Audio.
func NewAudio(id int, title, format string, duration time.Duration, url string) *Audio {
	return &Audio{
		ID:         id,
		Title:      title,
		Format:     format,
		Duration:   duration,
		UploadedAt: time.Now(), // Timestamp au moment de la création
		URL:        url,
	}
}

// GetReadableDuration retourne la durée lisible d'un fichier audio.
func (a *Audio) GetReadableDuration() string {
	hours := int(a.Duration.Hours())
	minutes := int(a.Duration.Minutes()) % 60
	seconds := int(a.Duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// IsFileSizeValid vérifie si la taille d'un fichier audio est valide.
func (a *Audio) IsFileSizeValid(maxSize int64) bool {
	return a.FileSize <= maxSize
}

// IsUploaded vérifie si un fichier audio a été uploadé.
func (a *Audio) IsRecorded() bool {
	return a.Type == "Recorded"
}

// IsUploaded vérifie si un fichier audio a été uploadé.
func (a *Audio) IsUploaded() bool {
	return a.Type == "Uploaded"
}

// IsValid vérifie si les attributs d'un fichier audio sont valides.
func (a *Audio) IsValid() bool {
	if a.Title == "" || a.Format == "" || a.Duration <= 0 || a.URL == "" {
		return false
	}
	return true
}
