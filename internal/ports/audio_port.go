package ports

import "time"

type AudioProcessor interface {
	ConvertFormat(inputPath string, outputFormat string) (string, error)
	TrimAudio(inputPath string, startTime, duration time.Duration) (string, error)
}
