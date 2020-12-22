package collections

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/voodio/repository/models"
	"github.com/slaveofcode/voodio/web/config"
)

const (
	// ExtractionDirName store directory name of extracted HLS files
	ExtractionDirName = "generated_hls"
)

// TranscodingError will hold the error information after transcoding process finished
type TranscodingError struct {
	Resolution string
	Error      error
}

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

func createm3u8Playlist(path string, res []string) {
	f, _ := os.Create(filepath.Join(path, "playlist.m3u8"))
	defer f.Close()

	c := `#EXTM3U
#EXT-X-VERSION:3
`

	for _, r := range res {
		if r == "360p" {
			c = c + `#EXT-X-STREAM-INF:BANDWIDTH=800000,RESOLUTION=640x360
360p.m3u8
`
		}

		if r == "480p" {
			c = c + `#EXT-X-STREAM-INF:BANDWIDTH=1400000,RESOLUTION=842x480
480p.m3u8
`
		}

		if r == "720p" {
			c = c + `#EXT-X-STREAM-INF:BANDWIDTH=2800000,RESOLUTION=1280x720
720p.m3u8
`
		}

		if r == "1080p" {
			c = c + `#EXT-X-STREAM-INF:BANDWIDTH=5000000,RESOLUTION=1920x1080
1080p.m3u8
`
		}
	}

	logrus.Debugln(c)

	f.Write([]byte(c))
}

// ExtractMovHLS will generate HLS files
func ExtractMovHLS(movieFilePath, destDir string, reso []string) (bool, []TranscodingError) {
	availableResolutions := map[string]func(string, string) []string{
		"360p":  cmdHLS360p,
		"480p":  cmdHLS480p,
		"720p":  cmdHLS720p,
		"1080p": cmdHLS1080p,
	}

	logrus.Debugln("reso", reso)

	resolutions := make(map[string]func(string, string) []string)
	for _, r := range reso {
		if availableResolutions[r] != nil {
			resolutions[r] = availableResolutions[r]
		}
	}

	logrus.Debugln("generated reso", resolutions)

	createm3u8Playlist(destDir, reso)

	output := make(chan TranscodingError, len(resolutions))
	for reso, cmdStrings := range resolutions {
		go func(out chan<- TranscodingError, commandProducer func(string, string) []string, resolution string) {
			cmd := exec.Command("ffmpeg", commandProducer(movieFilePath, destDir)...)
			cmd.Stdout = logrus.New().Out
			cmd.Stderr = logrus.New().Out

			logrus.Infoln("Exec:", strings.Join(cmd.Args, " "))

			err := cmd.Start()

			if err != nil {
				logrus.Errorln("Exec error:", err)
				out <- TranscodingError{
					Resolution: resolution,
					Error:      err,
				}
				return
			}

			out <- TranscodingError{
				Resolution: resolution,
				Error:      nil,
			}
		}(output, cmdStrings, reso)
	}

	var hasError bool
	errors := make([]TranscodingError, len(resolutions))

	iter := 0
	for out := range output {
		if out.Error != nil && !hasError {
			hasError = true
		}
		errors = append(errors, out)
		iter++

		if iter == len(resolutions) {
			close(output)
		}
	}

	return hasError, errors
}

func getExtractionMovieDir(appDir string, movieID uint) string {
	return filepath.Join(appDir, ExtractionDirName, strconv.Itoa(int(movieID)))
}

func createWriteableDir(path string) error {
	return os.MkdirAll(path, 0777)
}

// DoExtraction will do extract HLS files
func DoExtraction(movie *models.Movie, cfg *config.ServerConfig) (bool, []TranscodingError) {
	extractionDirName := getExtractionMovieDir(cfg.AppDir, movie.ID)

	if err := createWriteableDir(extractionDirName); err != nil {
		return true, nil
	}

	return ExtractMovHLS(
		filepath.Join(movie.DirPath, movie.BaseName),
		extractionDirName,
		cfg.ScreenResolutions,
	)
}
