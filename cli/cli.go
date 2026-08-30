package cli

import (
	"database/sql"
	"fmt"
	"os"
)

func Run(database *sql.DB, args []string) {
	cmd:=args[0]
	rest:=args[1:]
	switch cmd {
	case "add":
		RunAdd(database,rest)
	case "list","ls":
		RunList(database,rest)
	case "random":
		RunRandom(database)
	case "done":
		RunDone(database,rest)
	case "rm","remove":
		RunRemove(database,rest)
	case "help","-h","--help":
		PrintHelp()
	default:
		fmt.Fprintf(os.Stderr, "someday: unknown command %q\n\n", cmd)
		PrintHelp()
		os.Exit(1)
	}
}
