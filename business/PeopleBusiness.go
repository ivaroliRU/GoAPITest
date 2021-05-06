package business

import (
	"github.com/ivaroliRU/peopleAPI/models"
	"github.com/ivaroliRU/peopleAPI/data"
)

type PeopleBusiness struct {
	Data *data.Data
}

func (p *PeopleBusiness) GetPeople() []models.Person {
	return p.Data.GetData().People
}

func (p *PeopleBusiness) GetPersonByID(id int32) *models.Person {
	people := p.Data.GetData().People

	for _, p := range people {
		if p.ID == id {
			return &p
		}
	}

	return nil
}