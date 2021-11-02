package tasks

import "time"

type TasQ struct {
	ID        string    `json:"task_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	QueueName           string  `json:"queue_name"`
	JobName             string  `json:"job_name"`
	ArgsMap             ArgsMap `json:"args_map"`
	CurrentHandlerIndex int     `json:"current_handler_index"`
	OriginalArgsMap     ArgsMap `json:"original_args_map"`
	ResultLog           string  `json:"result_log"`
}
