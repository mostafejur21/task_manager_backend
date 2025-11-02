package domain

type Task struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Status      *string `json:"status" db:"status"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}
