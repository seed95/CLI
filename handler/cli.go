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

var (
	ErrInvalidInput = errors.New("invalid input")
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
			err = b.handleAdd()
			if err != nil {
				return err
			}
		case DeleteAction:
			err = b.handleDelete()
			if err != nil {
				return err
			}
		case PathAction:
			err = b.handlePath()
			if err != nil {
				return err
			}
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

func printSelectModel() {
	fmt.Println("Select model:")
	fmt.Println("1. City")
	fmt.Println("2. Road")
}

// handleSelectModel print menu and return selected model
func (b *bundle) handleSelectModel() (int, error) {

	printSelectModel()

	if b.scanner.Scan() {
		num, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return 0, err
		}

		if num == CityModel || num == RoadModel {
			return num, nil
		} else {
			return 0, ErrInvalidInput
		}

	}
	return 0, b.scanner.Err()

}

func (b *bundle) handleAddSuccessful(model string, id int) (int, error) {
	fmt.Printf("%s with id=%d added!\n", model, id)
	fmt.Println("1. Add another", model)
	fmt.Println("2. Main Menu")

	if b.scanner.Scan() {
		num, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return 0, err
		}

		if num == 1 || num == 2 {
			return num, nil
		} else {
			return 0, ErrInvalidInput
		}

	}
	return 0, b.scanner.Err()
}

func handleHelp() {
	fmt.Println("Select a number from shown menu and enter. For example 1 is for help.1")
}

func (b *bundle) handleAdd() error {

	selectedModel, err := b.handleSelectModel()
	if err != nil {
		return err
	}

	switch selectedModel {
	case CityModel:
		return b.handleAddCity()
	case RoadModel:
		return b.handleAddRoad()
	}
	return errors.New("invalid input")
}

func (b *bundle) handleAddCity() error {
	city, err := b.getCityModel()
	if err != nil {
		return err
	}
	err = b.agencyService.AddCity(city)
	if err != nil {
		return err
	}
	nextStep, err := b.handleAddSuccessful("City", city.ID)
	if err != nil {
		return err
	}
	if nextStep == 1 { // Add another City selected
		return b.handleAddCity()
	} else { // Main menu selected
		return nil
	}
}

func (b *bundle) getCityModel() (city *model.City, err error) {
	city = &model.City{}

	fmt.Println("id=?")
	city.ID, err = b.readIntInput()
	if err != nil {
		return nil, err
	}

	fmt.Println("name=?")
	city.Name, err = b.readStringInput()
	if err != nil {
		return nil, err
	}

	return city, nil

}

func (b *bundle) handleAddRoad() error {
	road, err := b.getRoadModel()
	if err != nil {
		return err
	}
	err = b.agencyService.AddRoad(road)
	if err != nil {
		return err
	}
	nextStep, err := b.handleAddSuccessful("Road", road.ID)
	if err != nil {
		return err
	}
	if nextStep == 1 { // Add another Road selected
		return b.handleAddRoad()
	} else { // Main menu selected
		return nil
	}
}

func (b *bundle) getRoadModel() (road *model.Road, err error) {
	road = &model.Road{}

	fmt.Println("id=?")
	road.ID, err = b.readIntInput()
	if err != nil {
		return nil, err
	}

	fmt.Println("name=?")
	road.Name, err = b.readStringInput()
	if err != nil {
		return nil, err
	}

	fmt.Println("from=?")
	road.From, err = b.readIntInput()
	if err != nil {
		return nil, err
	}

	fmt.Println("to=?")
	road.To, err = b.readIntInput()
	if err != nil {
		return nil, err
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
	road.SpeedLimit, err = b.readIntInput()
	if err != nil {
		return nil, err
	}

	fmt.Println("length=?")
	road.Length, err = b.readIntInput()
	if err != nil {
		return nil, err
	}

	fmt.Println("bi_directional=?")
	road.BiDirectional, err = b.readIntInput()
	if err != nil {
		return nil, err
	}
	if road.BiDirectional != 0 && road.BiDirectional != 1 {
		return nil, ErrInvalidInput
	}

	return road, nil

}

func (b *bundle) readStringInput() (string, error) {
	if b.scanner.Scan() {
		return b.scanner.Text(), nil
	}
	return "", b.scanner.Err()
}

func (b *bundle) readIntInput() (int, error) {

	if b.scanner.Scan() {
		num, err := strconv.Atoi(b.scanner.Text())
		if err != nil {
			return 0, err
		}
		return num, nil
	}
	return 0, b.scanner.Err()
}

func (b *bundle) handleDelete() error {

	selectedModel, err := b.handleSelectModel()
	if err != nil {
		return err
	}

	id, err := b.readIntInput()
	if err != nil {
		return err
	}

	switch selectedModel {
	case CityModel:
		err = b.agencyService.DeleteCity(id)
		if err != nil {
			fmt.Printf("City with id %d not found!\n", id)
		} else {
			fmt.Printf("City:%d deleted!\n", id)
		}
	case RoadModel:
		err = b.agencyService.DeleteRoad(id)
		if err != nil {
			fmt.Printf("Road with id %d not found!\n", id)
		} else {
			fmt.Printf("Road:%d deleted!\n", id)
		}
	default:
		return errors.New("invalid input")
	}

	return nil

}

func (b *bundle) handlePath() error {
	input, err := b.readStringInput()
	if err != nil {
		return err
	}

	splitInput := strings.Split(input, ":")
	if len(splitInput) != 2 {
		return ErrInvalidInput
	}
	sourceID, err := strconv.Atoi(splitInput[0])
	if err != nil {
		return err
	}
	destinationID, err := strconv.Atoi(splitInput[1])
	if err != nil {
		return err
	}
	fmt.Println("source id:", sourceID, ", destination id:", destinationID)
	return b.agencyService.GetPath(sourceID, destinationID)
}
