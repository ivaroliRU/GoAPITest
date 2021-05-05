package repositories

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"github.com/ivaroliRU/peopleAPI/repositories/data"
)

func GetPeople() []models.Person {
	return data.GetData().People
}