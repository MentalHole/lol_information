package data


// ChampionData represents the structure of a single champion's information from Data Dragon.
type ChampionData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ChampionListResponse represents the response from Data Dragon for all champions.
type ChampionListResponse struct {
	Data map[string]ChampionData `json:"data"`
}