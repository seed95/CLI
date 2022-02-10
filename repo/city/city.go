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

func (r *cityRepo) AddCity(city *model.City) error {
	// Found city with this id
	if _, err := r.GetCityIndexByID(city.ID); err == nil {
		return ErrDuplicate
	}
	r.cities = append(r.cities, *city)
	return nil
}

func (r *cityRepo) UpdateCity(city *model.City) error {
	index, err := r.GetCityIndexByID(city.ID)
	if err != nil {
		return err
	}
	r.cities[index] = *city
	return nil
}

func (r *cityRepo) DeleteCity(id int) error {
	index, err := r.GetCityIndexByID(id)
	if err != nil {
		return err
	}

	r.cities[index] = r.cities[len(r.cities)-1]
	r.cities = r.cities[:len(r.cities)-1]
	return nil
}

func (r *cityRepo) IsExist(id int) bool {
	_, err := r.GetCityIndexByID(id)
	if err != nil {
		return false
	}
	return true
}

func (r *cityRepo) GetCityByID(id int) *model.City {
	for _, c := range r.cities {
		if c.ID == id {
			return &c
		}
	}
	return nil
}

func (r *cityRepo) GetCityIndexByID(id int) (int, error) {
	for index, c := range r.cities {
		if c.ID == id {
			return index, nil
		}
	}
	return -1, ErrNotFound
}
