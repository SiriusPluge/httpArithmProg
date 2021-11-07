package internal

type Task struct {
	Id            int     `json:"id"`
	IdQueue       int     `json:"id_queue"`       // Number in the queue (integer)
	ExecStatus    string  `json:"exec_status"`    // Status: In progress/Queued/Completed
	N             int     `json:"n"`       // the number of elements (integer)
	D             float64 `json:"d"`              // the delta between the elements of the sequence (real)
	N1            float64 `json:"n_1"`            // the starting value (real)
	I             float64 `json:"i"`              // the interval in seconds between iterations (real)
	TTL           float64 `json:"ttl"`            //  the storage time of the result in seconds (real)
	CurIteration  uint    `json:"cur_iteration"`  // Current iteration
	StatementTime string  `json:"statement_time"` // Task statement time
	StartTime     string  `json:"start_time"`     // Task start time
	EndTime       string  `json:"end_time"`       // The end time of the task (if the task is completed)
}
