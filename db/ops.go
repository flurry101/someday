package db

import (
	"database/sql"
	"fmt"
	"time"
)

func AddItem(database *sql.DB, text string) (int64, error) {
	res, err := database.Exec(`INSERT INTO items (text, done, added_at) VALUES (?, 0, ?)`,
		text, time.Now())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func MarkDone(database *sql.DB, id int) error {
	res, err := database.Exec(`UPDATE items SET done = 1, done_at = ? WHERE id = ? AND done = 0`,
		time.Now(), id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("no open item with id %d", id)
	}
	return nil
}

func RemoveItem(database *sql.DB, id int) error {
	res, err := database.Exec(`DELETE FROM items WHERE id = ?`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("no item with id %d", id)
	}
	return nil
}
