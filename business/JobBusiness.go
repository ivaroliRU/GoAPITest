package business

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"github.com/ivaroliRU/peopleAPI/data"
	"fmt"
)

type JobsBusiness struct {
	Data *data.Data
}

// get all jobs
func (j *JobsBusiness) GetJobs() []models.Job {
	return j.Data.GetData().Jobs
}

// get a specific job with an id
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

// get people with a specific job
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

// add a job to our datasource
func (j *JobsBusiness) AddJob(job models.Job) {
	j.Data.Jobs = append(j.Data.Jobs, job)
}

// update a job in our datasource