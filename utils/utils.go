package utils

import (
	"os"
	"github.com/kardianos/osext"
	"github.com/vlinx-io/go-logging"
	"path/filepath"
)

var logger = logging.New("info.log","error.log")

func FileExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func GetExeDir() string {

	path, err := osext.ExecutableFolder()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	return path
}

func GetBaseName(filename string) string {
	extension := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(extension)]
	return name
}


