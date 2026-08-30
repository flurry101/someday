package main

import (
	"fmt"
	"os"
	"someday/cli"
	"someday/db"
	"someday/tui"
)

func main() {
	db,err:=db.Open()
	if err!=nil {
		fmt.Fprintln(os.Stderr, "someday:", err)
		os.Exit(1)
	}
	defer db.Close()
	if len(os.Args)>1 {
		cli.Run(db, os.Args[1:])
		return
	}
	if err:=tui.Run(db);err!=nil {
		fmt.Fprintln(os.Stderr,"someday:",err)
		os.Exit(1)
	}
}