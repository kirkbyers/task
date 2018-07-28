package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kirkbyers/task/cmd"
	"github.com/kirkbyers/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	must(err)
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
