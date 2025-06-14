package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"roleplay/internal/character"
)

func main() {
	// Create a new character
	createCharReq := struct {
		Name  string          `json:"name"`
		Class character.Class `json:"class"`
	}{
		Name:  "SteamKnight",
		Class: character.ClockworkKnight,
	}

	reqBody, err := json.Marshal(createCharReq)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/api/character/create", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var createCharResp struct {
		Character *character.Character `json:"character"`
		Error     string               `json:"error,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&createCharResp); err != nil {
		log.Fatal(err)
	}

	if createCharResp.Error != "" {
		log.Fatal("Error creating character:", createCharResp.Error)
	}

	fmt.Printf("Created character: %+v\n", createCharResp.Character)

	// Create another character for combat
	createCharReq.Name = "SteamMage"
	createCharReq.Class = character.SteamMage

	reqBody, err = json.Marshal(createCharReq)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Post("http://localhost:8080/api/character/create", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&createCharResp); err != nil {
		log.Fatal(err)
	}

	if createCharResp.Error != "" {
		log.Fatal("Error creating character:", createCharResp.Error)
	}

	fmt.Printf("Created character: %+v\n", createCharResp.Character)

	// Start combat between the two characters
	startCombatReq := struct {
		Participants []string `json:"participants"`
	}{
		Participants: []string{"SteamKnight", "SteamMage"},
	}

	reqBody, err = json.Marshal(startCombatReq)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = http.Post("http://localhost:8080/api/combat/start", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var startCombatResp struct {
		GameID string `json:"game_id"`
		State  struct {
			ActiveCharacter *character.Character `json:"active_character"`
			Round           int                  `json:"round"`
		} `json:"state"`
		Error string `json:"error,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&startCombatResp); err != nil {
		log.Fatal(err)
	}

	if startCombatResp.Error != "" {
		log.Fatal("Error starting combat:", startCombatResp.Error)
	}

	log.Printf("Started combat: %v", startCombatResp)

	// fmt.Printf("Started combat: GameID=%s, ActiveCharacter=%s, Round=%d\n",
	// 	startCombatResp.GameID,
	// 	startCombatResp.State.ActiveCharacter.Name,
	// 	startCombatResp.State.Round)
}
