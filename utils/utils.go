package utils

import (
	"os"
	"path/filepath"

	log "github.com/ngaut/logging"
)

func GetExecPath() string {
	fileDir := filepath.Dir(os.Args[0])
	execPath, err := filepath.Abs(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	return execPath
}
