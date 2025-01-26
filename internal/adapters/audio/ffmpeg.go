package audio

import (
	"fmt"
	"os/exec"
	"time"
)

type FfmpegProcessor struct{}

func NewFfmpegProcessor() *FfmpegProcessor {
	return &FfmpegProcessor{}
}

func (p *FfmpegProcessor) ConvertFormat(inputPath, outputFormat string) (string, error) {
	outputPath := fmt.Sprintf("%s_converted.%s", inputPath, outputFormat)
	cmd := exec.Command("ffmpeg", "-i", inputPath, outputPath)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return outputPath, nil
}

func (p *FfmpegProcessor) TrimAudio(inputPath string, start, duration time.Duration) (string, error) {
	outputPath := fmt.Sprintf("%s_trimmed.mp3", inputPath)
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ss", fmt.Sprintf("%.0f", start.Seconds()), "-t", fmt.Sprintf("%.0f", duration.Seconds()), outputPath)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return outputPath, nil
}
