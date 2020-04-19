package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/slaveofcode/pms/logger"
	"github.com/slaveofcode/pms/repository"
)

const (
	appDirName = "pmsapp"
	dbFileName = "pms.db"
)

var cacheDir, _ = os.UserCacheDir()

func getAppDir() string {
	return filepath.FromSlash(cacheDir + "/" + appDirName)
}

func getDBPath() string {
	// safe path for cross OS
	return filepath.FromSlash(getAppDir() + "/" + dbFileName)
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
}
