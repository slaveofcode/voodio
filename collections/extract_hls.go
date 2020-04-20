package collections

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/slaveofcode/pms/repository/models"
)

const (
	// ExtractionDirName store directory name of extracted HLS files
	ExtractionDirName = "generated_hls"
)

func cmdHLS360p(movieFilePath, destDir string) []string {
	// fixed width not divisible by 2
	// see https://superuser.com/questions/624563/how-to-resize-a-video-to-make-it-smaller-with-ffmpeg
	// One downside of scale when using libx264 is that this encoder requires even values
	// and scale may automatically choose an odd value resulting in an error: width or height not divisible by 2.
	// You can tell scale to choose an even value for a given height
	return []string{
		"-hide_banner",
		"-y",
		"-i", movieFilePath,
		"-vf", "scale=trunc(oh*a/2)*2:360",
		"-c:a", "aac",
		"-ar", "48000",
		"-c:v", "h264", // codec:video output format
		"-profile:v", "main",
		"-crf", "20", // video quality 1 the best, 51 worst
		"-sc_threshold", "0",
		"-g", "48",
		"-keyint_min", "48",
		"-hls_time", "10",
		"-hls_playlist_type", "vod",
		"-b:v", "800k",
		"-maxrate", "856k",
		"-bufsize", "1200k",
		"-b:a", "96k",
		"-hls_segment_filename", filepath.Join(destDir, "360p_%03d.ts"),
		filepath.Join(destDir, "360p.m3u8"),
	}
}

func cmdHLS480p(movieFilePath, destDir string) []string {
	return []string{
		"-hide_banner",
		"-y",
		"-i", movieFilePath,
		"-vf", "scale=trunc(oh*a/2)*2:480",
		"-c:a", "aac",
		"-ar", "48000",
		"-c:v", "h264",
		"-profile:v", "main",
		"-crf", "20",
		"-sc_threshold", "0",
		"-g", "48",
		"-keyint_min", "48",
		"-hls_time", "10",
		"-hls_playlist_type", "vod",
		"-b:v", "1400k",
		"-maxrate", "1498k",
		"-bufsize", "2100k",
		"-b:a", "128k",
		"-preset", "ultrafast",
		"-hls_segment_filename", filepath.Join(destDir, "480p_%03d.ts"),
		filepath.Join(destDir, "480p.m3u8"),
	}
}

func cmdHLS720p(movieFilePath, destDir string) []string {
	return []string{
		"-hide_banner",
		"-y",
		"-i", movieFilePath,
		"-vf", "scale=trunc(oh*a/2)*2:720",
		"-c:a", "aac",
		"-ar", "48000",
		"-c:v", "h264",
		"-profile:v", "main",
		"-crf", "20",
		"-sc_threshold", "0",
		"-g", "48",
		"-keyint_min", "48",
		"-hls_time", "10",
		"-hls_playlist_type", "vod",
		"-b:v", "2800k",
		"-maxrate", "2996k",
		"-bufsize", "4200k",
		"-b:a", "128k",
		"-preset", "ultrafast",
		"-hls_segment_filename", filepath.Join(destDir, "720p_%03d.ts"),
		filepath.Join(destDir, "720p.m3u8"),
	}
}

func cmdHLS1080p(movieFilePath, destDir string) []string {
	return []string{
		"-hide_banner",
		"-y",
		"-i", movieFilePath,
		"-vf", "scale=trunc(oh*a/2)*2:1080",
		"-c:a", "aac",
		"-ar", "48000",
		"-c:v", "h264",
		"-profile:v", "main",
		"-crf", "20",
		"-sc_threshold", "0",
		"-g", "48",
		"-keyint_min", "48",
		"-hls_time", "10",
		"-hls_playlist_type", "vod",
		"-b:v", "5000k",
		"-maxrate", "5350k",
		"-bufsize", "7500k",
		"-b:a", "192k",
		"-preset", "ultrafast",
		"-hls_segment_filename", filepath.Join(destDir, "1080p_%03d.ts"),
		filepath.Join(destDir, "1080p.m3u8"),
	}
}

// ExtractMovHLS will generate HLS files
func ExtractMovHLS(movieFilePath, destDir string) error {
	resolutions := map[string]func(string, string) []string{
		"360p":  cmdHLS360p,
		"480p":  cmdHLS480p,
		"720p":  cmdHLS720p,
		"1080p": cmdHLS1080p,
	}

	output := make(chan error, len(resolutions))
	for _, cmdStrings := range resolutions {
		go func(out chan<- error, commandProducer func(string, string) []string) {
			cmd := exec.Command("ffmpeg", commandProducer(movieFilePath, destDir)...)
			stdout, err := cmd.CombinedOutput()

			log.Println("Args:", strings.Join(cmd.Args, " "))
			log.Println("output:", string(stdout))
			if err != nil {
				out <- err
				return
			}

			out <- nil
		}(output, cmdStrings)
	}

	var errors []error
	for out := range output {
		errors = append(errors, out)
	}

	log.Println(errors)

	close(output)

	return nil
}

// DoExtraction will do extract HLS files
func DoExtraction(movie *models.Movie, appDir string) error {
	extractionDirName := filepath.Join(appDir, ExtractionDirName, movie.DirName)

	if err := os.MkdirAll(extractionDirName, 0777); err != nil {
		return err
	}

	return ExtractMovHLS(filepath.Join(movie.DirPath, movie.BaseName), extractionDirName)
}
