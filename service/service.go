package service

import (
	"time"
	"travel-agency/model"
)

type (
	TravelAgencyService interface {
		AddCity(city *model.City) error
		AddRoad(road *model.Road) error
		DeleteCity(id int) error
		DeleteRoad(id int) error
		GetPath(sourceID, destinationID int) ([]Path, error)
	}

	Path struct {
		SourceCityName      string
		DestinationCityName string
		RoadName            string
		TravelTime          time.Duration
	}
)
