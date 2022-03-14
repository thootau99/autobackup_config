package main


import (
	"fmt"
	"syscall"
	"strings"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"bytes"
	"time"
)

type Env struct {
	name string
	path string
}

type Exec struct {
	app  string
	args []string
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
			if (envName[len(envName) - 1] == "PWD") {
				continue
			} 
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
	fmt.Println("update finish")
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

func PushToGit() {
	var out bytes.Buffer
	addExe := exec.Command("git", "add", ".")
    addExe.Stdout = &out

    addErr := addExe.Run()
    if nil != addErr {
    	fmt.Println("ERROR!")
    	return
    }

    nowTime := time.Now()
    commitMessage := nowTime.Format("2006-01-02 03:04:05") + " auto backup finished"
	commitExe := exec.Command("git", "commit", "--allow-empty", "-m", commitMessage)
	commitExe.Stdout = &out
    commitErr := commitExe.Run()
    if nil != commitErr {
    	fmt.Println("ERROR!", commitErr)
    	return
    }

	pushExe := exec.Command("git", "push")
	pushExe.Stdout = &out

    pushErr := pushExe.Run()
    if nil != pushErr {
    	fmt.Println("ERROR!")
    	return
    }
    fmt.Println(out.String())
}