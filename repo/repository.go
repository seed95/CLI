package repo

import "travel-agency/model"

type RoadRepo interface {
	AddRoad(road *model.Road)
	DeleteRoad()
	UpdateRoad()
}

type CityRepo interface {
	AddCity(city *model.City)
	DeleteCity()
	UpdateCity()
}
