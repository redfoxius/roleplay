package combat

import (
	"math/rand"
	"time"

	"roleplay/internal/character"

	"github.com/google/uuid"
)

// Action represents a combat action
type Action struct {
	Name        string
	Damage      int
	Healing     int
	SteamCost   int
	Description string
}

// CombatState represents the current state of combat
type CombatState struct {
	ID              string
	ActiveCharacter *character.Character
	TurnOrder       []*character.Character
	CurrentTurn     int
	Round           int
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
	}
}

// sortByInitiative sorts characters by their initiative (Dexterity + SteamPower)
func sortByInitiative(participants []*character.Character) {
	for i := 0; i < len(participants)-1; i++ {
		for j := i + 1; j < len(participants); j++ {
			initiativeI := participants[i].Attributes.Dexterity.Value + participants[i].Attributes.SteamPower.Value
			initiativeJ := participants[j].Attributes.Dexterity.Value + participants[j].Attributes.SteamPower.Value
			if initiativeJ > initiativeI {
				participants[i], participants[j] = participants[j], participants[i]
			}
		}
	}
}

// ExecuteAction executes a combat action
func (cs *CombatState) ExecuteAction(action Action, target *character.Character) bool {
	// Check if action is valid
	if !cs.isValidAction(action, target) {
		return false
	}

	// Apply action effects
	if action.Damage > 0 {
		damage := calculateDamage(action.Damage, cs.ActiveCharacter, target)
		target.Health -= damage
	}

	if action.Healing > 0 {
		healing := calculateHealing(action.Healing, cs.ActiveCharacter)
		target.Health = min(target.Health+healing, target.MaxHealth)
	}

	// Move to next turn
	cs.nextTurn()
	return true
}

// isValidAction checks if an action can be performed
func (cs *CombatState) isValidAction(action Action, target *character.Character) bool {
	// Check if target is valid
	if target == nil || target.Health <= 0 {
		return false
	}

	// Check if character has enough steam power
	if action.SteamCost > cs.ActiveCharacter.Attributes.SteamPower.Value {
		return false
	}

	return true
}

// calculateDamage calculates the damage of an action
func calculateDamage(baseDamage int, attacker, defender *character.Character) int {
	// Base damage calculation
	damage := baseDamage

	// Add strength bonus
	damage += attacker.Attributes.Strength.Value / 5

	// Add class-specific bonuses
	switch attacker.Class {
	case character.Engineer:
		damage += attacker.Attributes.TechnicalAptitude.Value / 4
	case character.SteamMage:
		damage += attacker.Attributes.ArcaneKnowledge.Value / 4
	}

	// Apply defense reduction
	defense := defender.Attributes.Constitution.Value / 2
	damage = max(1, damage-defense)

	return damage
}

// calculateHealing calculates the healing amount
func calculateHealing(baseHealing int, healer *character.Character) int {
	healing := baseHealing

	// Add wisdom bonus
	healing += healer.Attributes.Wisdom.Value / 5

	// Add class-specific bonuses
	switch healer.Class {
	case character.Alchemist:
		healing += healer.Attributes.Intelligence.Value / 4
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
