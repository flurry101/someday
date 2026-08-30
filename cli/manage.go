package cli

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"someday/db"
)

func RunDone(database *sql.DB, rest []string) {
	if len(rest)==0 {
		fmt.Fprintln(os.Stderr, "usage: someday done <id>")
		os.Exit(1)
	}
	id, err := strconv.Atoi(rest[0])
	if err != nil {
		Fatal(fmt.Errorf("'%s' isn't a valid id", rest[0]))
	}
	if err := db.MarkDone(database, id); err != nil {
		Fatal(err)
	}
	fmt.Printf("done: #%d\n", id)
}

func RunRemove(database *sql.DB, rest []string) {
	if len(rest)==0 {
		fmt.Fprintln(os.Stderr, "usage: someday rm <id>")
		os.Exit(1)
	}
	id, err := strconv.Atoi(rest[0])
	if err != nil {
		Fatal(fmt.Errorf("'%s' isn't a valid id", rest[0]))
	}
	if err := db.RemoveItem(database, id); err != nil {
		Fatal(err)
	}
	fmt.Printf("removed #%d\n", id)
}
