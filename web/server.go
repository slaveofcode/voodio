package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/voodio/web/config"
)

func requestIDGenerator() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// NewServer create new HTTP server
func NewServer(cfg *config.ServerConfig) *http.Server {
	logrusWriter := logrus.New().Writer()
	defer logrusWriter.Close()

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		Handler:      NewRouter(cfg),
		ErrorLog:     log.New(logrusWriter, "", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return server
}
