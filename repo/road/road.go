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
	ErrNoRoad    = errors.New("no road between source & dest")
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

func (r *roadRepo) IsExist(id int) bool {
	_, err := r.GetRoadIndexWithID(id)
	if err != nil {
		return false
	}
	return true
}

func (r *roadRepo) GetRoadByID(id int) *model.Road {
	for _, r := range r.roads {
		if r.ID == id {
			return &r
		}
	}
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

func (r *roadRepo) GetRoadByCities(sourceID, destID int) ([]model.Road, error) {
	var availableRoads []model.Road
	for _, road := range r.roads {
		through := road.Through
		if !contains(through, road.From) {
			through = append(through, road.From)
		}
		if !contains(through, road.To) {
			through = append(through, road.To)
		}
		if contains(through, sourceID) && contains(through, destID) {
			if road.BiDirectional == 1 {
				availableRoads = append(availableRoads, road)
			} else {
				if index(through, destID) > index(through, sourceID) {
					availableRoads = append(availableRoads, road)
				}
			}
		}
	}
	if len(availableRoads) > 0 {
		return availableRoads, nil
	}
	return nil, ErrNoRoad
}

func contains(a []int, b int) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func index(a []int, b int) int {
	for i, v := range a {
		if v == b {
			return i
		}
	}
	return -1
}
