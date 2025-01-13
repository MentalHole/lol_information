package data


type ChampionData struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
}

// ChampionListResponse represents the response from Data Dragon for all champions.
type ChampionListResponse struct {
	Data map[string]ChampionData `json:"data"`
}