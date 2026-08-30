package cli

import (
	"fmt"
	"os"
	"someday/db"
)

func PrintItems(items []db.Item) {
	if len(items)==0 {
		fmt.Println("nothing here yet. someday add \"something\"")
		return
	}
	for _,it := range items {
		mark := " "
		if it.Done {
			mark = "x"
		}
		fmt.Printf("[%s] #%-3d %s\n", mark, it.ID, it.Text)
	}
}

func Fatal(err error) {
	fmt.Fprintln(os.Stderr, "someday:", err)
	os.Exit(1)
}

func PrintHelp() {
	fmt.Print(`someday - a place for things you want to do
usage:
  someday                    open the list
  someday add "text"         add something
  someday list [--all]       print the list (--all includes done)
  someday random             show one random open item
  someday done <id>          mark an item done
  someday rm <id>            remove an item
`)
}
