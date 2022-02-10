package service

import (
	"fmt"
	"time"
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
	if a.roadRepo.IsExist(road.ID) {
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

func (a *agencyService) GetPath(sourceID, destinationID int) ([]Path, error) {
	roads, err := a.roadRepo.GetRoadByCities(sourceID, destinationID)
	if err != nil {
		return nil, err
	}
	var availableRoads []Path
	for _, road := range roads {
		travelTime, _ := time.ParseDuration(fmt.Sprintf("%dm", int64(float32(road.Length)/float32(road.SpeedLimit)*60)))
		path := Path{
			SourceCityName:      a.cityRepo.GetCityByID(sourceID).Name,
			DestinationCityName: a.cityRepo.GetCityByID(destinationID).Name,
			RoadName:            road.Name,
			TravelTime:          travelTime,
		}
		availableRoads = append(availableRoads, path)
	}
	return availableRoads, nil
}
