package service

import (
	"travel-agency/model"
	"travel-agency/repo"
)

type (
	agencyService struct {
		cityRepo repo.CityRepo
		roadRepo repo.RoadRepo
	}

	Setting struct {
		CityRepo repo.CityRepo
		RoadRepo repo.RoadRepo
	}

	TravelAgencyService interface {
		AddRoad(road *model.Road)
		AddCity(city *model.City)
		DeleteRoad()
		DeleteCity()
		GetPath()
	}
)

var _ TravelAgencyService = (*agencyService)(nil)

func New(s *Setting) TravelAgencyService {
	return &agencyService{
		cityRepo: s.CityRepo,
		roadRepo: s.RoadRepo,
	}
}

func (a *agencyService) AddRoad(road *model.Road) {
}

func (a *agencyService) DeleteRoad() {

}

func (a *agencyService) AddCity(city *model.City) {

}

func (a *agencyService) DeleteCity() {

}

func (a *agencyService) GetPath() {

}
