package collections

import (
	"path/filepath"

	"strings"

	"github.com/asticode/go-astisub"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/voodio/repository/models"
)

// ProcessSrt will detect .srt files on DB and convert it into .vtt to extracted movie dir
func ProcessSrt(db *gorm.DB, movie *models.Movie, appDir string) error {
	destPath := filepath.Join(getExtractionMovieDir(appDir, movie.ID), "subs")

	if err := createWriteableDir(destPath); err != nil {
		return err
	}

	// get subs based on path movie
	var subs []models.Subtitle
	db.Where(&models.Subtitle{
		DirPath: movie.DirPath,
	}).Find(&subs)

	for _, sub := range subs {
		// convert & save to app dir
		s, err := astisub.OpenFile(filepath.Join(sub.DirPath, sub.BaseName))
		if err != nil {
			logrus.Errorln("Couldn't read the SRT file", sub.BaseName, err)
		}

		s.Write(filepath.Join(destPath, GetVTTFileName(sub.BaseName)))
	}

	return nil
}

// GetVTTFileName will return .vtt file based on given file name with ext
func GetVTTFileName(srtFileName string) string {
	return strings.Split(srtFileName, ".")[0] + ".vtt"
}
