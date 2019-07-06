package main

import "TODOlist/task/cmd"

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
