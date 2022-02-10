package city

import (
	"errors"
	"travel-agency/model"
	"travel-agency/repo"
)

type (
	cityRepo struct {
		cities []model.City
	}
)

var (
	ErrNotFound  = errors.New("city not found")
	ErrDuplicate = errors.New("city already exist")
)

var _ repo.CityRepo = (*cityRepo)(nil)

func New() repo.CityRepo {
	return &cityRepo{}
}

func (c *cityRepo) AddCity(city *model.City) error {
	// Found city with this id
	if _, err := c.GetCityIndexByID(city.ID); err == nil {
		return ErrDuplicate
	}
	c.cities = append(c.cities, *city)
	return nil
}

func (c *cityRepo) UpdateCity(city *model.City) error {
	index, err := c.GetCityIndexByID(city.ID)
	if err != nil {
		return err
	}
	c.cities[index] = *city
	return nil
}

func (c *cityRepo) DeleteCity(id int) error {
	index, err := c.GetCityIndexByID(id)
	if err != nil {
		return err
	}

	c.cities[index] = c.cities[len(c.cities)-1]
	c.cities = c.cities[:len(c.cities)-1]
	return nil
}

func (c *cityRepo) GetCityByID(id int) *model.City {
	for _, c := range c.cities {
		if c.ID == id {
			return &c
		}
	}
	return nil
}

func (c *cityRepo) GetCityIndexByID(id int) (int, error) {
	for index, c := range c.cities {
		if c.ID == id {
			return index, nil
		}
	}
	return -1, ErrNotFound
}
