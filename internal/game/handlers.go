package game

import (
	"encoding/json"
	"net/http"

	"roleplay/internal/character"
	"roleplay/internal/combat"
	"roleplay/internal/common"
)

// Handler handles HTTP requests for the game server
type Handler struct {
	server *GameServer
}

// NewHandler creates a new game handler
func NewHandler(server *GameServer) *Handler {
	return &Handler{server: server}
}

// CreateCharacterRequest represents a request to create a new character
type CreateCharacterRequest struct {
	Name  string          `json:"name"`
	Class character.Class `json:"class"`
}

// CreateCharacterResponse represents the response for character creation
type CreateCharacterResponse struct {
	Character *character.Character `json:"character"`
	Error     string               `json:"error,omitempty"`
}

// StartCombatRequest represents a request to start combat
type StartCombatRequest struct {
	Participants []string `json:"participants"`
}

// StartCombatResponse represents the response for starting combat
type StartCombatResponse struct {
	GameID string              `json:"game_id"`
	State  *combat.CombatState `json:"state"`
	Error  string              `json:"error,omitempty"`
}

// StartMobCombatRequest represents a request to start combat with a mob
type StartMobCombatRequest struct {
	CharacterID string `json:"character_id"`
	MobID       string `json:"mob_id"`
}

// StartMobCombatResponse represents the response for starting mob combat
type StartMobCombatResponse struct {
	Combat *combat.MobCombat `json:"combat"`
	Error  string            `json:"error,omitempty"`
}

// MobCombatActionRequest represents a request to perform a combat action
type MobCombatActionRequest struct {
	CharacterID string          `json:"character_id"`
	MobID       string          `json:"mob_id"`
	Ability     *common.Ability `json:"ability"`
}

// MobCombatActionResponse represents the response for a combat action
type MobCombatActionResponse struct {
	CharacterDamage int    `json:"character_damage"`
	MobDamage       int    `json:"mob_damage"`
	Result          string `json:"result"`
	Error           string `json:"error,omitempty"`
}

// RegisterRoutes registers all HTTP routes
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/character/create", h.handleCreateCharacter)
	mux.HandleFunc("/api/combat/start", h.handleStartCombat)
	mux.HandleFunc("/api/mob-combat/start", h.handleStartMobCombat)
	mux.HandleFunc("/api/mob-combat/action", h.handleMobCombatAction)
}

func (h *Handler) handleCreateCharacter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	char, err := h.server.CreateCharacter(req.Name, req.Class)
	response := CreateCharacterResponse{
		Character: char,
	}
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) handleStartCombat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StartCombatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	state, err := h.server.StartCombat(req.Participants)
	response := StartCombatResponse{
		State: state,
	}
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) handleStartMobCombat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StartMobCombatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	combat, err := h.server.StartMobCombat(req.CharacterID, req.MobID)
	response := StartMobCombatResponse{
		Combat: combat,
	}
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) handleMobCombatAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req MobCombatActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	combat, err := h.server.StartMobCombat(req.CharacterID, req.MobID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	charDamage, mobDamage, err := h.server.HandleMobCombatAction(combat, req.Ability)
	response := MobCombatActionResponse{
		CharacterDamage: charDamage,
		MobDamage:       mobDamage,
		Result:          combat.GetCombatResult(),
	}
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
