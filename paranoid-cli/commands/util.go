package commands

import (
	"github.com/cpssd/paranoid/logger"
	"os"
)

var Log *logger.ParanoidLogger

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
