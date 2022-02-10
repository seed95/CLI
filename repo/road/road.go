package road

import (
	"travel-agency/model"
	"travel-agency/repo"
)

type (
	roadRepo struct {
		roads map[int]*model.Road
	}
)

var _ repo.RoadRepo = (*roadRepo)(nil)

func New() repo.RoadRepo {
	return &roadRepo{}
}

func (r *roadRepo) AddRoad(road *model.Road) {
	r.roads = make(map[int]*model.Road)
	r.roads[road.ID] = road
}

func (r *roadRepo) UpdateRoad(road *model.Road) {
	r.roads[road.ID] = road
}

func (r *roadRepo) DeleteRoad(id int) {
	delete(r.roads, id)
}

func (r *roadRepo) GetRoadByID(id int) *model.Road {
	if road, ok := r.roads[id]; ok {
		return road
	}
	return nil
}
