package db

import (
	"database/sql"
	"time"
)

type Item struct {
	ID      int
	Text    string
	Done    bool
	AddedAt time.Time
	DoneAt  sql.NullTime
}
