package main


import (
	"fmt"
)

func main() {
	env := ReadEnv();
	fmt.Println(env)
	init_folder(env);
	ReadConfig(env);
	PushToGit()
}