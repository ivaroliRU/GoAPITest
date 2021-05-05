package models

type Person struct {
	ID            int32     `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Gender        string    `json:"gender"`
	JobID         int32     `json:"JobID"`
}