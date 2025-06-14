package combat

import (
	"math"
	"math/rand"
	"time"

	"github.com/redfoxius/roleplay/services/game-server/internal/character"
	"github.com/redfoxius/roleplay/services/game-server/internal/common"

	"github.com/google/uuid"
)

// Action represents a combat action
type Action struct {
	Name        string
	Type        string // Can be: "mechanical", "chemical", "arcane", "ranged", "melee"
	Damage      int
	Healing     int
	SteamCost   int
	Description string
	Range       int // Range of the action in tiles
	Area        int // Area of effect in tiles (0 for single target)
	Cooldown    int // Cooldown in turns
	LastUsed    int // Last round this action was used
}

// CombatState represents the current state of combat
type CombatState struct {
	ID              string
	ActiveCharacter *character.Character
	TurnOrder       []*character.Character
	CurrentTurn     int
	Round           int
	CombatLog       []CombatLogEntry
	StatusEffects   map[string][]StatusEffect
	Terrain         string
	Weather         string
}

// CombatLogEntry represents a single entry in the combat log
type CombatLogEntry struct {
	Round     int
	Turn      int
	Character string
	Action    string
	Target    string
	Damage    int
	Healing   int
	Effects   []string
}

// StatusEffect represents a temporary effect on a character
type StatusEffect struct {
	Name        string
	Duration    int
	Value       int
	Description string
}

// NewCombatState creates a new combat state
func NewCombatState(participants []*character.Character) *CombatState {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Sort participants by initiative (Dexterity + SteamPower)
	sortByInitiative(participants)

	return &CombatState{
		ID:              uuid.New().String(),
		ActiveCharacter: participants[0],
		TurnOrder:       participants,
		CurrentTurn:     0,
		Round:           1,
		CombatLog:       make([]CombatLogEntry, 0),
		StatusEffects:   make(map[string][]StatusEffect),
		Terrain:         "neutral", // Can be: neutral, steam-rich, mechanical, toxic, etc.
		Weather:         "clear",   // Can be: clear, steam-fog, acid-rain, etc.
	}
}

// sortByInitiative sorts characters by their initiative (Dexterity + SteamPower)
func sortByInitiative(participants []*character.Character) {
	for i := 0; i < len(participants)-1; i++ {
		for j := i + 1; j < len(participants); j++ {
			initiativeI := participants[i].Stats.Dexterity + participants[i].Stats.SteamPower
			initiativeJ := participants[j].Stats.Dexterity + participants[j].Stats.SteamPower
			if initiativeJ > initiativeI {
				participants[i], participants[j] = participants[j], participants[i]
			}
		}
	}
}

// ExecuteAction executes a combat action with enhanced effects
func (cs *CombatState) ExecuteAction(action Action, target *character.Character) bool {
	// Check if action is valid
	if !cs.isValidAction(action, target) {
		return false
	}

	// Calculate terrain and weather effects
	terrainBonus := cs.calculateTerrainBonus(action)
	weatherPenalty := cs.calculateWeatherPenalty(action)

	// Apply action effects
	var damage, healing int
	if action.Damage > 0 {
		damage = calculateDamage(action.Damage, cs.ActiveCharacter, target)
		damage = int(float64(damage) * terrainBonus * weatherPenalty)
		target.Health -= damage
	}

	if action.Healing > 0 {
		healing = calculateHealing(action.Healing, cs.ActiveCharacter)
		healing = int(float64(healing) * terrainBonus * weatherPenalty)
		target.Health = min(target.Health+healing, target.MaxHealth)
	}

	// Apply status effects
	cs.applyStatusEffects(cs.ActiveCharacter, target, action)

	// Log the action
	cs.logAction(action, target, damage, healing)

	// Move to next turn
	cs.nextTurn()
	return true
}

// calculateTerrainBonus calculates bonus damage/healing based on terrain
func (cs *CombatState) calculateTerrainBonus(action Action) float64 {
	bonus := 1.0
	switch cs.Terrain {
	case "steam-rich":
		if action.SteamCost > 0 {
			bonus = 1.2
		}
	case "mechanical":
		if cs.ActiveCharacter.Class == character.Engineer {
			bonus = 1.15
		}
	case "toxic":
		if cs.ActiveCharacter.Class == character.Alchemist {
			bonus = 1.15
		}
	}
	return bonus
}

// calculateWeatherPenalty calculates penalties based on weather
func (cs *CombatState) calculateWeatherPenalty(action Action) float64 {
	penalty := 1.0
	switch cs.Weather {
	case "steam-fog":
		if action.Type == "ranged" {
			penalty = 0.8
		}
	case "acid-rain":
		penalty = 0.9
	}
	return penalty
}

// applyStatusEffects applies status effects to characters
func (cs *CombatState) applyStatusEffects(attacker, target *character.Character, action Action) {
	// Apply effects based on character class and action type
	switch attacker.Class {
	case character.Engineer:
		if action.Type == "mechanical" {
			cs.addStatusEffect(target.ID, StatusEffect{
				Name:        "Steam Burn",
				Duration:    2,
				Value:       5,
				Description: "Takes 5 damage per turn",
			})
		}
	case character.Alchemist:
		if action.Type == "chemical" {
			cs.addStatusEffect(target.ID, StatusEffect{
				Name:        "Poison",
				Duration:    3,
				Value:       3,
				Description: "Takes 3 damage per turn",
			})
		}
	case character.SteamMage:
		if action.Type == "arcane" {
			cs.addStatusEffect(target.ID, StatusEffect{
				Name:        "Steam Weakness",
				Duration:    2,
				Value:       -10,
				Description: "Steam power reduced by 10",
			})
		}
	}
}

// addStatusEffect adds a status effect to a character
func (cs *CombatState) addStatusEffect(characterID string, effect StatusEffect) {
	cs.StatusEffects[characterID] = append(cs.StatusEffects[characterID], effect)
}

// logAction adds an entry to the combat log
func (cs *CombatState) logAction(action Action, target *character.Character, damage, healing int) {
	entry := CombatLogEntry{
		Round:     cs.Round,
		Turn:      cs.CurrentTurn,
		Character: cs.ActiveCharacter.Name,
		Action:    action.Name,
		Target:    target.Name,
		Damage:    damage,
		Healing:   healing,
		Effects:   make([]string, 0),
	}

	// Add status effects to log
	for _, effect := range cs.StatusEffects[target.ID] {
		entry.Effects = append(entry.Effects, effect.Name)
	}

	cs.CombatLog = append(cs.CombatLog, entry)
}

// GetCombatLog returns the combat log
func (cs *CombatState) GetCombatLog() []CombatLogEntry {
	return cs.CombatLog
}

// GetStatusEffects returns all status effects for a character
func (cs *CombatState) GetStatusEffects(characterID string) []StatusEffect {
	return cs.StatusEffects[characterID]
}

// UpdateStatusEffects updates all status effects
func (cs *CombatState) UpdateStatusEffects() {
	for charID, effects := range cs.StatusEffects {
		var remainingEffects []StatusEffect
		for _, effect := range effects {
			effect.Duration--
			if effect.Duration > 0 {
				remainingEffects = append(remainingEffects, effect)
			}
		}
		cs.StatusEffects[charID] = remainingEffects
	}
}

// isValidAction checks if an action can be performed
func (cs *CombatState) isValidAction(action Action, target *character.Character) bool {
	// Check if target is valid
	if target == nil || target.Health <= 0 {
		return false
	}

	// Check if character has enough steam power
	if action.SteamCost > cs.ActiveCharacter.Stats.SteamPower {
		return false
	}

	// Check cooldown
	if action.Cooldown > 0 && cs.Round-action.LastUsed < action.Cooldown {
		return false
	}

	// Check range
	if action.Range > 0 {
		attackerLoc := cs.ActiveCharacter.GetLocation()
		targetLoc := target.GetLocation()
		distance := distanceBetween(attackerLoc, targetLoc)
		if distance > action.Range {
			return false
		}
	}

	return true
}

// calculateDamage calculates the damage of an action
func calculateDamage(baseDamage int, attacker, defender *character.Character) int {
	// Base damage calculation
	damage := baseDamage

	// Add strength bonus
	damage += attacker.Stats.Strength / 5

	// Add class-specific bonuses
	switch attacker.Class {
	case character.Engineer:
		damage += attacker.Stats.Intelligence / 4
	case character.SteamMage:
		damage += attacker.Stats.SteamPower / 4
	}

	// Apply defense reduction
	defense := defender.Stats.Vitality / 2
	damage = max(1, damage-defense)

	return damage
}

// calculateHealing calculates the healing amount
func calculateHealing(baseHealing int, healer *character.Character) int {
	healing := baseHealing

	// Add intelligence bonus
	healing += healer.Stats.Intelligence / 5

	// Add class-specific bonuses
	switch healer.Class {
	case character.Alchemist:
		healing += healer.Stats.Intelligence / 4
	}

	return healing
}

// nextTurn advances to the next turn
func (cs *CombatState) nextTurn() {
	cs.CurrentTurn++
	if cs.CurrentTurn >= len(cs.TurnOrder) {
		cs.CurrentTurn = 0
		cs.Round++
	}
	cs.ActiveCharacter = cs.TurnOrder[cs.CurrentTurn]
}

// IsCombatOver checks if the combat is finished
func (cs *CombatState) IsCombatOver() bool {
	aliveCount := 0
	for _, char := range cs.TurnOrder {
		if char.Health > 0 {
			aliveCount++
		}
	}
	return aliveCount <= 1
}

// GetWinner returns the winning character if combat is over
func (cs *CombatState) GetWinner() *character.Character {
	if !cs.IsCombatOver() {
		return nil
	}
	for _, char := range cs.TurnOrder {
		if char.Health > 0 {
			return char
		}
	}
	return nil
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// distanceBetween calculates the distance between two coordinates
func distanceBetween(a, b common.Coordinates) int {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	return int(math.Sqrt(dx*dx + dy*dy))
}
