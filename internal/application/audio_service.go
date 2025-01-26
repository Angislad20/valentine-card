package application

import (
	"fmt"
	"time"
	"valentine-card/internal/domain"
	"valentine-card/internal/ports"
)

type AudioService struct {
	storage        ports.StoragePort
	audioProcessor ports.AudioProcessor
}

func NewAudioService(storage ports.StoragePort, audioProcessor ports.AudioProcessor) *AudioService {
	return &AudioService{
		storage:        storage,
		audioProcessor: audioProcessor,
	}
}

func (s *AudioService) UploadAudio(audio *domain.Audio, data []byte) error {
	url, err := s.storage.SaveFile(audio.URL, data)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}
	audio.URL = url
	return nil
}

func (s *AudioService) ProcessAudio(inputPath string) (string, error) {
	return s.audioProcessor.ConvertFormat(inputPath, "mp3")
}

func (s *AudioService) TrimAudio(inputPath string, start, duration time.Duration) (string, error) {
	return s.audioProcessor.TrimAudio(inputPath, start, duration)
}
