package cmd

import (
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
	  for _,i:= range args{
		  id,err:=strconv.Atoi(i)
		  if err!=nil{
			  fmt.Println("could not parse the argument")
            }else{
				ids=append(ids,id)
			}
	  }
	  fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
