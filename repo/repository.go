package repo

import "travel-agency/model"

type RoadRepo interface {
	AddRoad(road *model.Road) error
	UpdateRoad(road *model.Road) error
	DeleteRoad(id int) error
	GetRoadByID(id int) *model.Road
}

type CityRepo interface {
	AddCity(city *model.City) error
	UpdateCity(city *model.City) error
	DeleteCity(id int) error
	GetCityByID(id int) *model.City
}
