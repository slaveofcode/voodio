package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	parsetorrentname "github.com/middelink/go-parse-torrent-name"
	log "github.com/sirupsen/logrus"
	"github.com/slaveofcode/voodio/collections"
	"github.com/slaveofcode/voodio/logger"
	"github.com/slaveofcode/voodio/repository"
	"github.com/slaveofcode/voodio/repository/models"
	"github.com/slaveofcode/voodio/web"
	"github.com/slaveofcode/voodio/web/config"
)

const (
	appDirName = "voodioapp"
	dbFileName = "voodio.db"
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
			panic("Unable to create App Dir on " + appDirPath)
		}
		log.Infoln("Created App dir at", appDirPath)
	}

	// remove old database if exist
	dbPath := getDBPath()
	_, err := os.Stat(dbPath)
	if !os.IsNotExist(err) {
		log.Infoln("Obsolete DB detected, removing...")
		if err = os.Remove(dbPath); err != nil {
			panic("Unable removing obsolete DB")
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
	serverPort := flag.Int("port", 1818, "Server port number")
	flag.Parse()

	if len(*parentMoviePath) == 0 {
		cleanup()
		panic("No movie path directory provided, exited")
	}

	dbConn, err := repository.OpenDB(getDBPath())
	if err != nil {
		cleanup()
		panic("Unable to create DB connection")
	}

	defer dbConn.Close()

	log.Infoln("Preparing database...")
	repository.Migrate(dbConn)
	log.Infoln("Database prepared")

	// Scan movies inside given path
	movies, err := collections.ScanDir(*parentMoviePath)
	if err != nil {
		cleanup()
		panic("Error while scanning path " + err.Error())
	}

	// inserts all detected movies
	for _, movie := range movies {
		dirName := filepath.Base(movie.Dir)
		dirNameParsedInfo, err := parsetorrentname.Parse(filepath.Base(movie.Dir))
		cleanDirName := ""
		if err == nil {
			cleanDirName = dirNameParsedInfo.Title
		}

		baseNameParsedInfo, _ := parsetorrentname.Parse(movie.MovieFile)
		cleanBaseName := ""
		if err == nil {
			cleanBaseName = baseNameParsedInfo.Title
		}

		dbConn.Create(&models.Movie{
			DirPath:       movie.Dir,
			DirName:       dirName,
			CleanDirName:  cleanDirName,
			FileSize:      movie.MovieSize,
			BaseName:      movie.MovieFile,
			CleanBaseName: cleanBaseName,
			MimeType:      movie.MimeType,
			IsGroupDir:    false,
			IsPrepared:    false,
		})
	}

	// Find duplicate directory names, kinda serial movie
	var movieGroups []models.Movie
	dbConn.Table("movies").
		Select("dir_name, dir_path, COUNT(*) count").
		Group("dir_name, dir_path").
		Having("count > ?", 1).
		Find(&movieGroups)

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

	// create simple webserver
	webServer := web.NewServer(&config.ServerConfig{
		DB:     dbConn,
		Port:   *serverPort,
		AppDir: getAppDir(),
	})

	closeSignal := make(chan os.Signal, 1)
	signal.Notify(closeSignal, os.Interrupt)

	serverDone := make(chan bool)

	go func() {
		<-closeSignal
		log.Infoln("got close signal")

		// Waiting for current process server to finish with 30 secs timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		webServer.SetKeepAlivesEnabled(false)
		if err := webServer.Shutdown(ctx); err != nil {
			log.Errorln("Couldn't gracefully shutdown")
		}

		serverDone <- true
	}()

	log.Infoln("Activate Web UI Server")
	go web.NewStaticServer("8080")

	log.Infoln("Activate API Server")
	log.Infoln("Server is alive on port", *serverPort)
	if err = webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Errorln("Unable to start server on port", *serverPort)
	}

	<-serverDone

	cleanup()

	log.Infoln("Server closed")
}
