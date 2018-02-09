package main

import (
	"encoding/json"
	"github.com/vlinx-io/java-wrapper/settings"
	"github.com/vlinx-io/java-wrapper/utils"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// 这个Wrapper这周先做windows的吧，下周再做linux和mac两个系统的
func main() {

	exeDir := utils.GetExeDir()

	settingsFile := utils.GetBaseName(os.Args[0]) + ".json"
	if !utils.FileExist(settingsFile) {
		settingsFile = exeDir + string(os.PathSeparator) + "wrapper.json"
	}

	data, err := ioutil.ReadFile(settingsFile)
	if err != nil {
		log.Fatalln(err)
	}
	var setting settings.Settings
	err = json.Unmarshal(data, &setting)
	if err != nil {
		log.Fatalln(err)
	}

	if setting.Verbose {
	}

	command := "java"
	if setting.HideConsole && runtime.GOOS == "windows" {
		command = "javaw"
	}

	// 获取可执行的java文件的详细路径
	executable := exeDir + string(os.PathSeparator) + "jre" + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + command

	if setting.Verbose {
		log.Println("Command:", executable)
	}

	args := setting.JArgs

	// 再将用户命令中的参数赋值
	if len(os.Args) > 1 {
		args = append(args, os.Args[1:]...)
	}

	if setting.Verbose {
		log.Println("Args:", args)
	}

	process := exec.Command(executable, args...)

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	err = process.Run()

	if err != nil {
		log.Fatalln(err)
	}

}
