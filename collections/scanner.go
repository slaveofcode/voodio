package collections

import (
	"io/ioutil"
	"math"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

// MovieDirInfo is a holder to keep movie directory information
type MovieDirInfo struct {
	Dir       string
	MovieFile string
	MovieSize float64
	Info      os.FileInfo
	MimeType  string
}

// SubDirInfo is a holder to subtitle information
type SubDirInfo struct {
	Dir     string
	SubFile string
	Info    os.FileInfo
}

// ScanDir will return flat list of Movie directory information
func ScanDir(path string) ([]MovieDirInfo, []SubDirInfo, error) {
	var listMovie []MovieDirInfo
	var listSub []SubDirInfo

	listItems, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range listItems {
		detectedMovies, detectedSubs, err := Identify(path, item)
		if err != nil {
			return nil, nil, err
		}

		listMovie = append(listMovie, detectedMovies...)
		listSub = append(listSub, detectedSubs...)
	}

	return listMovie, listSub, nil
}

// Identify single directory of movie
func Identify(basePath string, file os.FileInfo) ([]MovieDirInfo, []SubDirInfo, error) {
	path := filepath.Join(basePath, file.Name())

	var listMovie []MovieDirInfo
	var listSub []SubDirInfo

	if file.IsDir() {
		// recursive call, create other segment to continue the scan operation
		listMovie, listSub, err := ScanDir(path)
		return listMovie, listSub, err
	}

	if mime, validVideo := isVideo(path); validVideo {
		listMovie = append(listMovie, MovieDirInfo{
			Dir:       basePath,
			MovieFile: file.Name(),
			MovieSize: getFileSizeInMB(file),
			Info:      file,
			MimeType:  mime,
		})
	}

	if validSub := IsTextSrt(path); validSub {
		listSub = append(listSub, SubDirInfo{
			Dir:     basePath,
			SubFile: file.Name(),
			Info:    file,
		})
	}

	return listMovie, listSub, nil
}

func getFileSizeInMB(file os.FileInfo) float64 {
	sizeInBytes := file.Size()
	sizeInMB := float64(sizeInBytes) / 1000000 // 1mio bytes = 1 MB
	sizeInMBRounded := math.Round(sizeInMB*100) / 100
	return sizeInMBRounded
}

func isVideo(path string) (string, bool) {
	validMimes := []string{
		"video/webm",       //.webm
		"video/3gpp",       // .3gp
		"video/3gpp2",      // .3g2
		"video/ogg",        // .ogv
		"video/mpeg",       // .mpeg or .mp4
		"video/x-msvideo",  // .avi
		"video/x-ms-wmv",   // .wmv
		"video/x-flv",      // .flv
		"video/mp4",        // .mp4
		"video/quicktime",  // .mov
		"video/x-matroska", // .mkv
		"video/x-ms-asf",   // .asf
		"video/x-m4v",      // .m4v
	}

	mime, err := mimetype.DetectFile(path)
	if err != nil {
		return "", false
	}

	listChecker := make(map[string]bool)
	for _, validMime := range validMimes {
		listChecker[validMime] = true
	}

	return mime.String(), listChecker[mime.String()]
}

// IsTextSrt will return valid status detection of .srt file
func IsTextSrt(path string) bool {
	mime, err := mimetype.DetectFile(path)
	if err != nil {
		return false
	}

	validMime := mime.String() == "text/plain; charset=utf-8"
	validExt := filepath.Ext(path) == ".srt" || filepath.Ext(path) == ".SRT"

	return validMime && validExt
}
