package models

import "time"

type Timer struct {
	RestTime               time.Time
	WorkTime               time.Time
	IterationDescription   string
	DefaultRestDescription string
	DefaultWorkDescription string
	Foloving	int64
}
