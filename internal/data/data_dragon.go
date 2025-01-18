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

	championIDSet := make(map[string]struct{})
	for _, id := range championIDs {
		championIDSet[strconv.Itoa(id)] = struct{}{} // Convert integer IDs to strings
	}
		// Create a map of numeric ID to champion name
	numericIDToName := make(map[int]string)
	for _, champion := range allChampions.Data {
		champID, err := strconv.Atoi(champion.ID) // If numeric ID exists in the JSON, use it
		if err != nil {
			continue
		}
		numericIDToName[champID] = champion.ID
	}


	var rotationChampions []ChampionData
	for _, id := range championIDs {
		name, exists := numericIDToName[id]
		if !exists {
			fmt.Printf("No name found for champion ID: %d\n", id)
			continue
		}
		if champion, exists := allChampions.Data[name]; exists {
			rotationChampions = append(rotationChampions, champion)
		}
	}




	return rotationChampions, nil
}
