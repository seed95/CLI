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
		AddCity(city *model.City) error
		AddRoad(road *model.Road) error
		DeleteCity(id int) error
		DeleteRoad(id int) error
		GetPath(sourceID, destinationID int) error
	}
)

var _ TravelAgencyService = (*agencyService)(nil)

func New(s *Setting) TravelAgencyService {
	return &agencyService{
		cityRepo: s.CityRepo,
		roadRepo: s.RoadRepo,
	}
}

func (a *agencyService) AddCity(city *model.City) error {

	// Update city if already exist
	if a.cityRepo.IsExist(city.ID) {
		return a.cityRepo.UpdateCity(city)
	}

	return a.cityRepo.AddCity(city)
}

func (a *agencyService) AddRoad(road *model.Road) error {

	// Update road if already exist
	if a.cityRepo.IsExist(road.ID) {
		return a.roadRepo.UpdateRoad(road)
	}

	return a.roadRepo.AddRoad(road)
}

func (a *agencyService) DeleteCity(id int) error {
	return a.cityRepo.DeleteCity(id)
}

func (a *agencyService) DeleteRoad(id int) error {
	return a.roadRepo.DeleteRoad(id)
}

func (a *agencyService) GetPath(sourceID, destinationID int) error {
	return nil
}
