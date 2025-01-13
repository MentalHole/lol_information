package main

import (
	"fmt"
	"log"
	"lol_legends_information/internal/api"
	"lol_legends_information/internal/config"
	"lol_legends_information/internal/data"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Fetch free champion IDs
	freeChampionIDs, err := api.GetFreeChampionIds(cfg.APIKey)
	if err != nil {
		log.Fatalf("Failed to fetch free champion IDs: %v", err)
	}

	fmt.Println("Free rotation: ", freeChampionIDs)

	// Fetch champion details for rotation
	rotationChampions, err := data.GetChampionsInRotation(cfg.Version, freeChampionIDs)
	if err != nil {
		log.Fatalf("Failed to fetch champions in rotation: %v", err)
	}

	// Print champions in rotation
	fmt.Println("Champions in Current Rotation:")
	for _, champ := range rotationChampions {	
		fmt.Println(champ)
	}

}
