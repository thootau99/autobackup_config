package main


import (
	"fmt"
	"syscall"
	"strings"
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	name string
	path string
}

func ReadEnv() []Env {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR IN LOAD ENV.");
	}
	var envPath []Env;
	envs := syscall.Environ();
	for _, element := range envs {
		if strings.Contains(element, "backup_config") {
			envArray := strings.Split(element, "=");
			envName := strings.Split(envArray[0], "backup_config_")
			env := Env{
				name: envName[len(envName) - 1],
				path: envArray[len(envArray) - 1],
			}
			envPath = append(envPath, env)
		}
	}
	return envPath
}

func ReadConfig(envs []Env) {
	for _, env := range envs {
		data, err := os.ReadFile(env.path)
		if nil != err {
			fmt.Println("error!", err)
			return
		}
		fileName := GetFileName(env.path)
		filePath := "./" + env.name + "/" + fileName
		WriteData(string(data), filePath)
	}
}

func WriteData(data string, path string) {
	file, err := os.Create(path)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, _err := file.WriteString(data)
	if nil != _err {
		fmt.Println(_err)
		return
	}
}

func GetFileName(path string) string {
	var splitPath []string = strings.Split(path, "/")
	return splitPath[len(splitPath) - 1]
}