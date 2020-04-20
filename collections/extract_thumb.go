package collections

import (
	"os/exec"
	"strconv"
)

// ExtractThumb will extract thumbnail image from video file
func ExtractThumb(moviePath, thumbPath string, seconds int) error {
	cmd := exec.Command(
		"ffmpeg",
		"-hide_banner",
		"-y",
		"-i", moviePath,
		"-vframes", "1",
		"-an", // disable audio
		// "-s", "400x200",
		"-ss", strconv.Itoa(seconds),
		thumbPath)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	cmd.Run()

	return nil
}
