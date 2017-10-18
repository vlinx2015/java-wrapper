package utils

import (

	"os"
	"github.com/kardianos/osext"
	"github.com/vlinx-io/go-logging"
)

var logger = logging.New("info.log","error.log")

func GetExeDir() string {

	path, err := osext.ExecutableFolder()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	return path
}

