package models

type Category struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
}