package cmd

import (
	"fmt"
	"strconv"

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
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
