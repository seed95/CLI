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
		DeleteRoad(id int)
		DeleteCity(id int)
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
	if mRoad := a.roadRepo.GetRoadByID(road.ID); mRoad != nil {
		a.roadRepo.UpdateRoad(road)
	} else {
		a.roadRepo.AddRoad(road)
	}
}

func (a *agencyService) DeleteRoad(id int) {
	if mRoad := a.roadRepo.GetRoadByID(id); mRoad != nil {
		a.roadRepo.DeleteRoad(id)
	}
}

func (a *agencyService) AddCity(city *model.City) {
	if mCity := a.cityRepo.GetCityByID(city.ID); mCity != nil {
		a.cityRepo.UpdateCity(city)
	} else {
		a.cityRepo.AddCity(city)
	}
}

func (a *agencyService) DeleteCity(id int) {
	if mCity := a.cityRepo.GetCityByID(id); mCity != nil {
		a.cityRepo.DeleteCity(id)
	}
}

func (a *agencyService) GetPath() {

}
