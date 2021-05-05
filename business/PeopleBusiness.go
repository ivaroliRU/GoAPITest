package business

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"github.com/ivaroliRU/peopleAPI/data"
)

func GetPeople() []models.Person {
	return data.GetData().People
}

func GetPersonByID(id int32) *models.Person {
	people := data.GetData().People

	for _, p := range people {
		if p.ID == id {
			return &p
		}
	}

	return nil
}