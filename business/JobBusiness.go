package business

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"github.com/ivaroliRU/peopleAPI/data"
)

func GetJobs() []models.Job {
	return data.GetData().Jobs
}

func GetJobsByID(id int32) *models.Job {
	jobs := data.GetData().Jobs

	for _, j := range jobs {
		if j.JobID == id {
			return &j
		}
	}

	return nil
}

func GetPeopleWithJob(id int32) []models.Person {
	people := data.GetData().People
	returnPeople := []models.Person{}

	for _, p := range people {
		if p.JobID == id {
			returnPeople = append(returnPeople, p)
		}
	}

	return returnPeople
}