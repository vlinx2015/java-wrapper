package main

import (
	"os"
	"github.com/vlinx-io/go-logging"
	"os/exec"
	"io/ioutil"
	"github.com/vlinx-io/java-wrapper/settings"
	"encoding/json"
	"github.com/vlinx-io/java-wrapper/utils"
	"fmt"
	"runtime"
)

var logger = logging.New("info.log","error.log")

// 这个Wrapper这周先做windows的吧，下周再做linux和mac两个系统的
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

	command := "java"

	if setting.HideConsole && runtime.GOOS == "windows" {
		command = "javaw"
	}

	// 获取可执行的java文件的详细路径
	executable := exeDir + string(os.PathSeparator) + "jre" + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + command
	log := fmt.Sprint("Executable",executable)
	logger.Info(log)

	// 将setting中的项都赋值出来，不直接在setting中操作
	classpath := setting.Classpath

	// 遍历classpath，将每个相对路径都变为绝对路径
	for index,value := range classpath {

		value = exeDir + string(os.PathSeparator) + value
		classpath[index] = value

	}

	classpath = append([]string{"-cp"},classpath...)

	// java的Arg在前，classpath在后
	args := append(setting.JArgs,classpath...)

	// 将mainClass内容并入参数
	args = append(args,setting.MainClass)
	
	// 再将用户命令中的参数赋值
	if len(os.Args) > 1 {
		args = append(args,os.Args[1:]...)
	}

	logger.Info("Args",args)

	process := exec.Command(executable,args...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	err = process.Run()

	if err!=nil {
		logger.Error(err)
	}

}