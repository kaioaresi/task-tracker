package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreateAt    time.Time    `json:"createAt"`
	DeleteAt    time.Time `json:"deleteAt"`
	UpdateAt    time.Time `json:"updateAt"`
}
