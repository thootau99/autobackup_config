package main

import (
	"os"
)


func init_folder(envs []Env) {
	for _, env := range envs {
		if _, err := os.Stat("./" + env.name ); os.IsNotExist(err) {
			os.Mkdir("./" + env.name, 0755)
		}
	}
}