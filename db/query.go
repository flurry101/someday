package db

import "database/sql"

func ListItems(database *sql.DB, includeDone bool) ([]Item, error) {
	q := `SELECT id, text, done, added_at, done_at FROM items`
	if !includeDone {
		q += ` WHERE done = 0`
	}
	q += ` ORDER BY added_at ASC`
	rows, err := database.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var it Item
		var doneInt int
		if err := rows.Scan(&it.ID, &it.Text, &doneInt, &it.AddedAt, &it.DoneAt); err != nil {
			return nil, err
		}
		it.Done = doneInt != 0
		items = append(items, it)
	}
	return items, rows.Err()
}

func RandomItem(database *sql.DB) (*Item, error) {
	row := database.QueryRow(`SELECT id, text, done, added_at, done_at FROM items WHERE done = 0 ORDER BY RANDOM() LIMIT 1`)
	var it Item
	var doneInt int
	if err := row.Scan(&it.ID, &it.Text, &doneInt, &it.AddedAt, &it.DoneAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	it.Done = doneInt != 0
	return &it, nil
}
