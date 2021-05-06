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

var JOBS_DATASOURCE string = "data/jobs.json"
var PEOPLE_DATASOURCE string = "data/jobs.json"

func (d *Data) GetData() *Data {
	jobsFile, _ := ioutil.ReadFile(JOBS_DATASOURCE)
	peopleFile, _ := ioutil.ReadFile(PEOPLE_DATASOURCE)

	_ = json.Unmarshal([]byte(peopleFile), &d.People)
	_ = json.Unmarshal([]byte(jobsFile), &d.Jobs)

	fmt.Println(d)
	
	return d
}