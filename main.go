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
	"strings"
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
		log.Println("VLINX Java Wrapper 0.1")
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

	// 将setting中的项都赋值出来，不直接在setting中操作
	classpath := setting.Classpath

	// 遍历classpath，将每个相对路径都变为绝对路径
	for index, value := range classpath {

		value = exeDir + string(os.PathSeparator) + value
		classpath[index] = value

	}

	sep := ":"

	if runtime.GOOS == "windows" {
		sep = ";"
	}

	//将classpath转换为字符串
	cp := strings.Join(classpath,sep)

	classpath = append([]string{"-cp"}, cp)

	// java的Arg在前，classpath在后
	args := append(setting.JArgs, classpath...)

	// 将mainClass内容并入参数
	args = append(args, setting.MainClass)

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
