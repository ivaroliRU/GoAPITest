package data

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"encoding/json"
	"io/ioutil"
)

type Data struct {
    People []models.Person
	Jobs []models.Job
}

var data Data = Data{}

func GetData() Data {
	jobsFile, _ := ioutil.ReadFile("data/jobs.json")
	peopleFile, _ := ioutil.ReadFile("data/people.json")

	_ = json.Unmarshal([]byte(peopleFile), &data.People)
	_ = json.Unmarshal([]byte(jobsFile), &data.Jobs)

	return data
}