package foundation

import (
	"github.com/withmandala/go-log"
	"os"
)

func App() {
	logger := log.New(os.Stderr)
	logger.Info("New Info Log from App")
}
