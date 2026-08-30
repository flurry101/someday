package cli

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"someday/db"
)

func RunAdd(database *sql.DB, rest []string) {
	text:=strings.TrimSpace(strings.Join(rest, " "))
	if text=="" {
		fmt.Fprintln(os.Stderr, `usage: someday add "something you want to do"`)
		os.Exit(1)
	}
	id,err := db.AddItem(database, text)
	if err!=nil {
		Fatal(err)
	}
	fmt.Printf("added #%d: %s\n", id, text)
}

func RunList(database *sql.DB, rest []string) {
	showDone:=len(rest) > 0 && rest[0] == "--all"
	items, err := db.ListItems(database, showDone)
	if err != nil {
		Fatal(err)
	}
	PrintItems(items)
}

func RunRandom(database *sql.DB) {
	it,err := db.RandomItem(database)
	if err!=nil {
		Fatal(err)
	}
	if it==nil {
		fmt.Println("nothing left. add something someday.")
		return
	}
	fmt.Printf("#%d  %s\n", it.ID, it.Text)
}
