package cmd

import (
	"fmt"
	"os"

	"github.com/kirkbyers/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong getting tasks", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vactaion? 🏖")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, val := range tasks {
			fmt.Printf("%d. %s\n", i+1, val.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
