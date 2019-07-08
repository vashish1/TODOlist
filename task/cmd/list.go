package cmd

import (
	"TODOlist/task/datab"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "to list the tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := datab.Alltasks()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)

		}
		if len(tasks) == 0 {
			fmt.Println("you have no tasks to complete! why don't you plan a vacation")
			os.Exit(1)
		}
		for i, t := range tasks {
			fmt.Printf("%d. %s \n", i+1, t.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
