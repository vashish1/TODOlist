package main

import (
	"TODOlist/task/cmd"
	"TODOlist/task/datab"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbpath := filepath.Join(home, "tasks.db")
	must(datab.Init(dbpath))
	must(cmd.RootCmd.Execute())

}
func must(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
