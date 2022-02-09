package city

import (
	"travel-agency/model"
	"travel-agency/repo"
)

type (
	cityRepo struct {
	}
)

var _ repo.CityRepo = (*cityRepo)(nil)

func New() repo.CityRepo {
	return &cityRepo{}
}

func (c *cityRepo) AddCity(city *model.City) {

}

func (c *cityRepo) UpdateCity() {

}

func (c *cityRepo) DeleteCity() {

}
