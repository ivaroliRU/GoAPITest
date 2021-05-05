package models

type Job struct {
	JobID        int64     `json:"JobID"`
	JobTitle     string    `json:"JobTitle"`
	Salary       float32   `json:"Salary"`
}