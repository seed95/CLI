package main

import (
	"log"
	"travel-agency/handler"
	"travel-agency/repo/city"
	"travel-agency/repo/road"
	"travel-agency/service"
)

func main() {

	cityRepo := city.New()
	roadRepo := road.New()

	serviceSetting := service.Setting{
		CityRepo: cityRepo,
		RoadRepo: roadRepo,
	}
	agencyService := service.New(&serviceSetting)

	handlerSetting := handler.Setting{AgencyService: agencyService}
	cli := handler.New(&handlerSetting)

	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}

}
