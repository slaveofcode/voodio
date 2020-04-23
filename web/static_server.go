package web

import (
	"net/http"

	"github.com/rakyll/statik/fs"
	// import generated statik file
	_ "github.com/slaveofcode/voodio/statik"
)

// NewStaticServer will spawn new static server
func NewStaticServer(port string) {
	statikFS, err := fs.New()
	if err != nil {
		panic("Unable to spawn Web UI" + err.Error())
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":"+port, nil)
}
