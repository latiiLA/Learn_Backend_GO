package models

import "time"

// Task represents a task with its properties
type Task struct{
	ID 	string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
	Status string `json:"status"`
}