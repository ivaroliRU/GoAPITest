package data

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Data struct {
    People []models.Person
	Jobs []models.Job
}

func (d *Data) GetData() *Data {
	jobsFile, _ := ioutil.ReadFile("data/jobs.json")
	peopleFile, _ := ioutil.ReadFile("data/people.json")

	_ = json.Unmarshal([]byte(peopleFile), &d.People)
	_ = json.Unmarshal([]byte(jobsFile), &d.Jobs)

	fmt.Println(d)
	
	return d
}