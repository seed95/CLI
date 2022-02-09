package road

import (
	"travel-agency/model"
	"travel-agency/repo"
)

type (
	roadRepo struct {
	}
)

var _ repo.RoadRepo = (*roadRepo)(nil)

func New() repo.RoadRepo {
	return &roadRepo{}
}

func (c *roadRepo) AddRoad(road *model.Road) {

}

func (c *roadRepo) DeleteRoad() {

}

func (c *roadRepo) UpdateRoad() {

}
