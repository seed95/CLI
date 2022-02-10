package handler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"travel-agency/model"
	"travel-agency/service"
)

type (
	bundle struct {
		agencyService service.TravelAgencyService
		scanner       *bufio.Scanner
	}

	Setting struct {
		AgencyService service.TravelAgencyService
	}

	Runner interface {
		Run() error
	}
)

const (
	_ = iota
	HelpAction
	AddAction
	DeleteAction
	PathAction
	ExitAction
)

const (
	_ = iota
	CityModel
	RoadModel
)

func New(s *Setting) Runner {
	return &bundle{
		agencyService: s.AgencyService,
		scanner:       bufio.NewScanner(os.Stdin),
	}
}

func (b *bundle) Run() error {

	printMainMenu()
	var input string
	for b.scanner.Scan() {

		// Get input and check is valid number
		input = b.scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			printInvalidInput()
			continue
		}

		switch num {
		case HelpAction:
			handleHelp()
		case AddAction:
			b.handleAdd()
		case DeleteAction:
			b.handleDelete()
		case PathAction:
		case ExitAction:
			os.Exit(0)
		default:
			printInvalidInput()

		}
		printMainMenu()
	}

	return b.scanner.Err()

}

func printMainMenu() {
	fmt.Println("Main Menu - Select an action:")
	fmt.Println("1. Help")
	fmt.Println("2. Add")
	fmt.Println("3. Delete")
	fmt.Println("4. Path")
	fmt.Println("5. Exit")
}

func printInvalidInput() {
	fmt.Println("Invalid input. Please enter 1 for more info.")
}

func printAddSelectModel() {
	fmt.Println("Select model:")
	fmt.Println("1. City")
	fmt.Println("2. Road")
}

func printAddSuccessful(model string, id int) {
	fmt.Printf("%s with id=%d added!", model, id)
}

func printAddNextAction(model string) {

}

func handleHelp() {
	fmt.Println("Select a number from shown menu and enter. For example 1 is for help.1")
}

func (b *bundle) handleAdd() error {
	printAddSelectModel()
	if b.scanner.Scan() {
		num, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return err
		}

		switch num {
		case CityModel:
			city, err := b.getCityModel()
			if err != nil {
				return err
			}
			b.agencyService.AddCity(city)
			printAddSuccessful("City", city.ID)
		case RoadModel:
			road, err := b.getRoadModel()
			if err != nil {
				return err
			}
			b.agencyService.AddRoad(road)
			printAddSuccessful("Road", road.ID)
		default:
			return errors.New("invalid input")
		}

	}
	return b.scanner.Err()
}

func (b *bundle) getCityModel() (*model.City, error) {
	city := &model.City{}

	fmt.Println("id=?")
	if b.scanner.Scan() {
		id, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return nil, err
		}
		city.ID = id
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("name=?")
	if b.scanner.Scan() {
		city.Name = b.scanner.Text()
	} else {
		return nil, b.scanner.Err()
	}

	return city, nil

}

func (b *bundle) getRoadModel() (*model.Road, error) {
	road := &model.Road{}

	fmt.Println("id=?")
	if b.scanner.Scan() {
		id, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return nil, err
		}
		road.ID = id
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("name=?")
	if b.scanner.Scan() {
		road.Name = b.scanner.Text()
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("from=?")
	if b.scanner.Scan() {
		from, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return nil, err
		}
		road.From = from
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("to=?")
	if b.scanner.Scan() {
		to, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return nil, err
		}
		road.To = to
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("through=?")
	if b.scanner.Scan() {
		splitted := strings.Split(strings.Replace(strings.Replace(b.scanner.Text(), "[", "", -1), "]", "", -1), ",")
		through := make([]int, len(splitted))
		var err error
		for i, v := range splitted {
			through[i], err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		}
		road.Through = through
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("speed_limit=?")
	if b.scanner.Scan() {
		speedLimit, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return nil, err
		}
		road.SpeedLimit = speedLimit
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("length=?")
	if b.scanner.Scan() {
		length, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return nil, err
		}
		road.Length = length
	} else {
		return nil, b.scanner.Err()
	}

	fmt.Println("bi_directional=?")
	if b.scanner.Scan() {
		biDirectional, err := strconv.ParseBool(b.scanner.Text())
		if err != nil {
			return nil, err
		}
		road.BiDirectional = biDirectional
	} else {
		return nil, b.scanner.Err()
	}

	return road, nil

}

func (b *bundle) handleDelete() {

}

func (b *bundle) handlePath() {

}
