package models

// Example model
type Example struct {
	ID     string `db:"id"`
	Item   string `db:"item"`
	Qty    int    `db:"qty"`
	Tags   string `db:"tags"`
	Status string `db:"status"`
	Size   string `db:"size"`
}
