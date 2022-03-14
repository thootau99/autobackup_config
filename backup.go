package main


import (
	"fmt"
)

func main() {
	env := ReadEnv();
	fmt.Println(env, "123")
	init_folder(env);
	ReadConfig(env);
	PushToGit()
}