package fsclient

import (
	"errors"
	"sync"
	"os"
)

var ErrClosedClient = errors.New("use of closed fs client")

type Client struct {
	sync.Mutex

	RootDir string
	DataDir  string
	TempDir  string
	LockFile string

	lockfd *os.File
	closed bool
}
