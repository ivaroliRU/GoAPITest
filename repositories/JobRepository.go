package repositories

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"github.com/ivaroliRU/peopleAPI/repositories/data"
)

func GetJobs() []models.Job {
	return data.GetData().Jobs
}