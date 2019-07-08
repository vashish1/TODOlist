package main

import (
	"TODOlist/task/cmd"
	"TODOlist/task/datab"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbpath := filepath.Join(home, "tasks.db")
	err := datab.Init(dbpath)
	if err != nil {
		panic(err)
	}
	err1 := cmd.RootCmd.Execute()
	if err1 != nil {
		panic(err1)
	}
}
