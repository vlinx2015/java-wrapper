package main

import (
	"os"
	"github.com/vlinx-io/go-logging"
	"os/exec"
	"io/ioutil"
	"github.com/vlinx-io/wrapper/settings"
	"encoding/json"
	"github.com/vlinx-io/wrapper/utils"
	"fmt"
)

var logger = logging.New("info.log","error.log")

func main() {

	exeDir := utils.GetExeDir()

	settingsFile := exeDir + string(os.PathSeparator) + "settings.json"

	data, err := ioutil.ReadFile(settingsFile)

	if err!=nil {
		logger.Error(err)
		os.Exit(1)
	}

	var setting settings.Settings

	err = json.Unmarshal(data,&setting)

	if err!=nil {
		logger.Error(err)
		os.Exit(1)
	}

	executable := exeDir + string(os.PathSeparator) + setting.Command
	log := fmt.Sprint("Executable",executable)
	logger.Info(log)

	var args []string

	if len(os.Args) > 1 {
		args = append(setting.Args,os.Args[1:]...)
	}else{
		args = setting.Args
	}

	logger.Info("Args",args)

	process := exec.Command(executable,args...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	err = process.Run()

	var arr = []interface{}{"hello","world"}

	str := fmt.Sprint(arr...)

	fmt.Println(str)

	if err!=nil {
		logger.Error(err)
	}

}