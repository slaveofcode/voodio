package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/slaveofcode/pms/collections"
	"github.com/slaveofcode/pms/logger"
	"github.com/slaveofcode/pms/repository"
	"github.com/slaveofcode/pms/repository/models"
)

const (
	appDirName = "pmsapp"
	dbFileName = "pms.db"
)

var cacheDir, _ = os.UserCacheDir()

func getAppDir() string {
	return filepath.Join(cacheDir, appDirName)
}

func getDBPath() string {
	return filepath.Join(getAppDir(), dbFileName)
}

func init() {
	// First of all, logger...
	logger.Setup()

	// create working app dir
	appDirPath := getAppDir()
	if _, err := os.Stat(appDirPath); os.IsNotExist(err) {
		err = os.MkdirAll(appDirPath, 0777)
		if err != nil {
			log.Errorln("Unable to create App Dir on ", appDirPath)
			os.Exit(1)
		}
		log.Infoln("Created App dir at", appDirPath)
	}

	// remove old database if exist
	dbPath := getDBPath()
	_, err := os.Stat(dbPath)
	if !os.IsNotExist(err) {
		log.Infoln("Obsolete DB detected, removing...")
		if err = os.Remove(dbPath); err != nil {
			log.Errorln("Unable removing obsolete DB")
			os.Exit(1)
		}
	}

	_, err = os.Create(dbPath)
	if err != nil {
		log.Errorln("Unable to init db file", err)
		os.Exit(1)
	}

	log.Infoln("DB initialized at", dbPath)
}

func cleanup() {
	log.Infoln("Cleaning up artifacts")
	os.RemoveAll(getAppDir())
}

func main() {
	parentMoviePath := flag.String("path", "", "Path string of parent movie directory")
	flag.Parse()

	if len(*parentMoviePath) == 0 {
		log.Errorln("No movie path directory provided, exited")
		cleanup()
		os.Exit(1)
	}

	fmt.Printf("Searching around %s path", *parentMoviePath)

	dbConn, err := repository.OpenDB(getDBPath())
	if err != nil {
		log.Errorln("Unable to create DB connection")
		os.Exit(1)
	}

	defer dbConn.Close()

	log.Infoln("Preparing database...")
	repository.Migrate(dbConn)
	log.Infoln("Database prepared")

	// Scan movies inside given path
	movies, err := collections.ScanDir(*parentMoviePath)
	if err != nil {
		log.Errorln(err)
	}

	// inserts all detected movies
	for _, movie := range movies {
		dbConn.Create(&models.Movie{
			DirPath:    movie.Dir,
			DirName:    filepath.Base(movie.Dir),
			FileSize:   movie.MovieSize,
			BaseName:   movie.MovieFile,
			MimeType:   movie.MimeType,
			IsGroupDir: false,
			IsPrepared: false,
		})
	}

	// Find duplicate directories, means it's kinda serial movie
	var movieGroups []models.Movie
	dbConn.Table("movies").Select("dir_name, dir_path, COUNT(*) count").Group("dir_name, dir_path").Having("count > ?", 1).Find(&movieGroups)

	for _, mg := range movieGroups {
		// find related movie with same dir_name & dir_path
		var movieList []models.Movie
		dbConn.Where(&models.Movie{
			DirName: mg.DirName,
			DirPath: mg.DirPath,
		}).Find(&movieList)

		for _, m := range movieList {
			dbConn.Model(&m).Update(&models.Movie{
				IsGroupDir: true,
			})
		}
	}

}
