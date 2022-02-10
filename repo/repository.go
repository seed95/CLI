package repo

import "travel-agency/model"

type RoadRepo interface {
	AddRoad(road *model.Road)
	UpdateRoad(road *model.Road)
	DeleteRoad(id int)
	GetRoadByID(id int) *model.Road
}

type CityRepo interface {
	AddCity(city *model.City)
	UpdateCity(city *model.City)
	DeleteCity(id int)
	GetCityByID(id int) *model.City
}
