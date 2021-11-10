package foundation

import (
	"github.com/withmandala/go-log"
	"os"
)

func App() {
	logger := log.New(os.Stderr)
	logger.Info("Hi, this is your logger")
}
