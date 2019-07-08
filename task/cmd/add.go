package cmd

import (
	"TODOlist/task/datab"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the do command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "to add a task in your task list",

	Run: func(cmd *cobra.Command, args []string) {
		str := strings.Join(args, " ")
		_, err := datab.CreateTask(str)
		if err != nil {
			fmt.Println("something went wrong:", err)
			return
		}
		fmt.Printf("task \"%s\" added in the list.\n ", str)

	},
}

func init() {
	RootCmd.AddCommand(addCmd)

}
