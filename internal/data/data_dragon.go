package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const dataDragonBaseURL = "https://ddragon.leagueoflegends.com/cdn"

func GetChampionsInRotation(version string, championIDs []int) ([]ChampionData, error) {
	url := fmt.Sprintf("%s/%s/data/en_US/champion.json", dataDragonBaseURL, version)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var allChampions ChampionListResponse
	if err := json.Unmarshal(body, &allChampions); err != nil {
		return nil, err
	}

	// Map champion keys (ID as string) to champion data
	championsMap := make(map[int]ChampionData)
	for _, champ := range allChampions.Data {
		id, err := strconv.Atoi(champ.ID)
		if err == nil {
			championsMap[id] = champ
		}
	}

	// Filter champions in rotation
	var rotationChampions []ChampionData
	for _, id := range championIDs {
		if champ, exists := championsMap[id]; exists {
			rotationChampions = append(rotationChampions, champ)
		}
	}

	return rotationChampions, nil
}
