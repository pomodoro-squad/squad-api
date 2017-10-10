package models

import "time"

type Timer struct {
	RestTime               time.Time	`json:"rest_time"`
	WorkTime               time.Time	`json:"work_time"`
	IterationDescription   string	`json:"iteration_description"`
	DefaultRestDescription string	`json:"default_rest_description"`
	DefaultWorkDescription string	`json:"default_work_description"`
}
