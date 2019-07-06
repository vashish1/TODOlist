package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd is used to call execute
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a cli task manager",
}
