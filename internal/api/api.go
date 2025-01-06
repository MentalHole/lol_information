package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const riotBaseURL = "https://euw1.api.riotgames.com"

// GetFreeChampionIds fetches the free champion IDs from the Riot API
func GetFreeChampionIds(apiKey string) ([]int, error) {
	url := fmt.Sprintf("%s/lol/platform/v3/champion-rotations", riotBaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Riot-Token", apiKey)

	res, err := http.DefaultClient.Do(req)
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

	var response FreeChampionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.FreeChampionIds, nil
}
