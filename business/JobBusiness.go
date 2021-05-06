package business

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"github.com/ivaroliRU/peopleAPI/data"
	"fmt"
)

type JobsBusiness struct {
	Data *data.Data
}

func (j *JobsBusiness) GetJobs() []models.Job {
	return j.Data.GetData().Jobs
}

func (j *JobsBusiness) GetJobsByID(id int32) *models.Job {
	jobs := j.Data.GetData().Jobs

	fmt.Println(j.Data.GetData())

	for _, j := range jobs {
		if j.JobID == id {
			return &j
		}
	}

	return nil
}

func (j *JobsBusiness) GetPeopleWithJob(id int32) []models.Person {
	people := j.Data.GetData().People
	returnPeople := []models.Person{}

	for _, p := range people {
		if p.JobID == id {
			returnPeople = append(returnPeople, p)
		}
	}

	return returnPeople
}