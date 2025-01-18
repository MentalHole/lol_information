package api

type Region string

const (
	RegionEuropeWest        Region = "euw1"
	RegionRussia			Region = "ru"
	RegionKorea             Region = "kr"
)

type Route string

const (
	RouteAsia     Route = "asia"
	RouteEurope   Route = "europe"
)

var (
	Regions = []Region{
		RegionEuropeWest,
		RegionKorea,
		RegionRussia,
	}

	RegionToRoute = map[Region]Route{
		RegionKorea: RouteAsia,
		RegionEuropeWest: RouteEurope,
		RegionRussia: RouteEurope,
	}
)