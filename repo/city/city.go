package city

import (
	"travel-agency/model"
	"travel-agency/repo"
)

type (
	cityRepo struct {
		cities map[int]*model.City
	}
)

var _ repo.CityRepo = (*cityRepo)(nil)

func New() repo.CityRepo {
	return &cityRepo{}
}

func (c *cityRepo) AddCity(city *model.City) {
	c.cities = make(map[int]*model.City)
	c.cities[city.ID] = city
}

func (c *cityRepo) UpdateCity(city *model.City) {
	c.cities[city.ID] = city
}

func (c *cityRepo) DeleteCity(id int) {
	delete(c.cities, id)
}

func (c *cityRepo) GetCityByID(id int) *model.City {
	if city, ok := c.cities[id]; ok {
		return city
	}
	return nil
}
