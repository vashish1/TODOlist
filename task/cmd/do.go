package cmd

import (
	"TODOlist/task/datab"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, i := range args {
			id, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("could not parse the argument")
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := datab.Alltasks()
		if err != nil {
			fmt.Println("something went wrong:", err)
			return
		}
		for _, i := range ids {
			if i <= 0 || i > len(tasks) {
				fmt.Println("invalid task number!!")
				continue
			}
			task := tasks[i-1]
			err := datab.Deletetask(task.Key)
			if err != nil {
				fmt.Println("failed to mark as completed:\n", err)

			} else {
				fmt.Printf("marked %s as completed", task.Value)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
