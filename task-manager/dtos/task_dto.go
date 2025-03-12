package dtos

type TaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"` // Formato: YYYY-MM-DD
}
