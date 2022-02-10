package main

import (
	"log"
	"travel-agency/handler"
	"travel-agency/model"
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

	var tehran = new(model.City)
	tehran.ID = 1
	tehran.Name = "Tehran"

	var qom = new(model.City)
	qom.ID = 2
	qom.Name = "Qom"

	var arak = new(model.City)
	arak.ID = 3
	arak.Name = "Arak"

	var borujerd = new(model.City)
	borujerd.ID = 4
	borujerd.Name = "Borujerd"

	var khorramabad = new(model.City)
	khorramabad.ID = 5
	khorramabad.Name = "KhorramAbad"

	var kashan = new(model.City)
	kashan.ID = 6
	kashan.Name = "Kashan"

	var khorramabadArak = new(model.Road)
	khorramabadArak.ID = 1
	khorramabadArak.Name = "Khorramabad-Arak"
	khorramabadArak.From = 5
	khorramabadArak.To = 3
	khorramabadArak.Through = []int{4}
	khorramabadArak.SpeedLimit = 120
	khorramabadArak.Length = 100
	khorramabadArak.BiDirectional = 1

	var arakTehran = new(model.Road)
	arakTehran.ID = 2
	arakTehran.Name = "Arak-Tehran"
	arakTehran.From = 3
	arakTehran.To = 1
	arakTehran.Through = []int{2}
	arakTehran.SpeedLimit = 110
	arakTehran.Length = 240
	arakTehran.BiDirectional = 1

	var qomTehran = new(model.Road)
	qomTehran.ID = 3
	qomTehran.Name = "Qom-Tehran"
	qomTehran.From = 1
	qomTehran.To = 2
	qomTehran.Through = []int{}
	qomTehran.SpeedLimit = 120
	qomTehran.Length = 140
	qomTehran.BiDirectional = 1

	var qomKashan = new(model.Road)
	qomKashan.ID = 4
	qomKashan.Name = "Qom-Kashan"
	qomKashan.From = 2
	qomKashan.To = 6
	qomKashan.Through = []int{}
	qomKashan.SpeedLimit = 100
	qomKashan.Length = 90
	qomKashan.BiDirectional = 1

	agencyService.AddCity(tehran)
	agencyService.AddCity(qom)
	agencyService.AddCity(arak)
	agencyService.AddCity(borujerd)
	agencyService.AddCity(khorramabad)
	agencyService.AddCity(kashan)

	agencyService.AddRoad(khorramabadArak)
	agencyService.AddRoad(arakTehran)
	agencyService.AddRoad(qomTehran)
	agencyService.AddRoad(qomKashan)

	agencyService.GetPath(1, 2)
	agencyService.GetPath(1, 3)
	agencyService.GetPath(2, 3)
	agencyService.GetPath(4, 5)
	agencyService.GetPath(2, 5)
	agencyService.GetPath(4, 6)
	agencyService.GetPath(2, 6)

	handlerSetting := handler.Setting{AgencyService: agencyService}
	cli := handler.New(&handlerSetting)

	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}

}
