package road

import (
	"errors"
	"travel-agency/model"
	"travel-agency/repo"
)

type (
	roadRepo struct {
		roads []model.Road
	}
)

var (
	ErrNotFound  = errors.New("road not found")
	ErrDuplicate = errors.New("road already exist")
)

var _ repo.RoadRepo = (*roadRepo)(nil)

func New() repo.RoadRepo {
	return &roadRepo{}
}

func (r *roadRepo) AddRoad(road *model.Road) error {

	// Found road with this id
	if _, err := r.GetRoadIndexWithID(road.ID); err == nil {
		return ErrDuplicate
	}
	r.roads = append(r.roads, *road)
	return nil
}

func (r *roadRepo) UpdateRoad(road *model.Road) error {
	index, err := r.GetRoadIndexWithID(road.ID)
	if err != nil {
		return err
	}
	r.roads[index] = *road
	return nil
}

func (r *roadRepo) DeleteRoad(id int) error {

	index, err := r.GetRoadIndexWithID(id)
	if err != nil {
		return err
	}

	r.roads[index] = r.roads[len(r.roads)-1]
	r.roads = r.roads[:len(r.roads)-1]
	return nil
}

func (r *roadRepo) GetRoadByID(id int) *model.Road {
	//if road, ok := r.roads[id]; ok {
	//	return road
	//}
	return nil
}

func (r *roadRepo) GetRoadIndexWithID(id int) (int, error) {
	for index, road := range r.roads {
		if road.ID == id {
			return index, nil
		}
	}
	return -1, ErrNotFound
}
