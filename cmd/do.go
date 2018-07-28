package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kirkbyers/task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, val := range args {
			id, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Failed to parse the argument:", val)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong getting all task:", err)
			os.Exit(1)
		}
		for _, val := range ids {
			if val <= 0 || val > len(tasks) {
				fmt.Printf("Task id %d is out of range to \"do\"\n", val)
				continue
			}
			err := db.DeleteTask(tasks[val-1].Key)
			if err != nil {
				fmt.Printf("Something went wrong deleting task \"%d. %s\".\n %+v", val, tasks[val-1].Value, err)
				os.Exit(1)
			}
			fmt.Printf("Completed task \"%d. %s\".\n", val, tasks[val-1].Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
