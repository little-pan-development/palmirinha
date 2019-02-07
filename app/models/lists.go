package models

import (
	"fmt"
	"strconv"
)

// List of items
type List struct {
	ID          uint
	Description string
	Item        []Item
}

// OpenList ...
func OpenList(description, channelID string) string {
	query := `INSERT cart SET description = ?, status = ?, channel_id = ?`
	stmt, err := app.Connection.Prepare(query)
	if err != nil {
		fmt.Println("Model OpenList [prepare]: ", err)
	}

	res, err := stmt.Exec(description, 1, channelID)
	if err != nil {
		fmt.Println("Model OpenList [exec]: ", err)
	}

	id, _ := res.LastInsertId()
	idToString := strconv.FormatInt(int64(id), 10)

	return idToString
}

// CloseList ...
func CloseList(channelID string) bool {

	query := `UPDATE cart SET status = ? WHERE status = ? AND channel_id = ?`
	stmt, err := app.Connection.Prepare(query)
	if err != nil {
		fmt.Println("Model CloseList [prepare]: ", err)
		return false
	}

	_, err = stmt.Exec(0, 1, channelID)
	if err != nil {
		fmt.Println("Model CloseList [exec]: ", err)
		return false
	}

	return true
}

// CountOpenList ...
func CountOpenList(channelID string) uint {

	var count uint
	query := `SELECT COUNT(*) FROM cart WHERE status = 1 and channel_id = ?`
	rows, err := app.Connection.Query(query, app.Message.ChannelID)

	if err != nil {
		fmt.Println(err)
	}

	if rows.Next() {
		rows.Scan(&count)
	}

	return count
}
