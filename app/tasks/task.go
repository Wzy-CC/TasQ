package tasks

import "time"

type TasQ struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Status     int       `json:"status"`
	Result     string    `json:"result"`
	OwnerQueue string    `json:"owner_queue"`
	ArgsMap    ArgsMap   `json:"args_map"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type ArgsMap map[string]interface{}
