package combat

import (
	"testing"

	"roleplay/internal/character"
	"roleplay/internal/common"
)

func TestNewCombatState(t *testing.T) {
	// Create test characters with different initiative values
	char1 := &character.Character{
		ID:         "char1",
		Name:       "Character 1",
		Health:     100,
		SteamPower: 50,
		Attributes: common.Attributes{
			Dexterity:  common.Attribute{Name: "Dexterity", Value: 15},
			SteamPower: common.Attribute{Name: "Steam Power", Value: 10},
		},
	}

	char2 := &character.Character{
		ID:         "char2",
		Name:       "Character 2",
		Health:     100,
		SteamPower: 50,
		Attributes: common.Attributes{
			Dexterity:  common.Attribute{Name: "Dexterity", Value: 10},
			SteamPower: common.Attribute{Name: "Steam Power", Value: 5},
		},
	}

	participants := []*character.Character{char2, char1} // char2 first to test sorting
	combat := NewCombatState(participants)

	// Test turn order (should be sorted by initiative)
	if combat.TurnOrder[0] != char1 {
		t.Error("Expected char1 to be first in turn order (higher initiative)")
	}
	if combat.TurnOrder[1] != char2 {
		t.Error("Expected char2 to be second in turn order (lower initiative)")
	}

	// Test initial state
	if combat.ActiveCharacter != char1 {
		t.Error("Expected char1 to be active character")
	}
	if combat.CurrentTurn != 0 {
		t.Errorf("Expected current turn 0, got %d", combat.CurrentTurn)
	}
	if combat.Round != 1 {
		t.Errorf("Expected round 1, got %d", combat.Round)
	}
}

func TestExecuteAction(t *testing.T) {
	char1 := &character.Character{
		ID:         "char1",
		Name:       "Character 1",
		Health:     100,
		SteamPower: 50,
		Attributes: common.Attributes{
			Strength: common.Attribute{Name: "Strength", Value: 10},
		},
	}

	char2 := &character.Character{
		ID:         "char2",
		Name:       "Character 2",
		Health:     100,
		SteamPower: 50,
		Attributes: common.Attributes{
			Constitution: common.Attribute{Name: "Constitution", Value: 8},
		},
	}

	combat := NewCombatState([]*character.Character{char1, char2})

	// Test damage action
	damageAction := Action{
		Name:        "Test Attack",
		Damage:      20,
		SteamCost:   10,
		Description: "Test damage action",
	}

	success := combat.ExecuteAction(damageAction, char2)
	if !success {
		t.Error("Expected action to succeed")
	}
	if char2.Health != 86 { // 100 - (20 + 10/5 - 8/2)
		t.Errorf("Expected char2 health 86, got %d", char2.Health)
	}

	// Test healing action
	healingAction := Action{
		Name:        "Test Heal",
		Healing:     30,
		SteamCost:   15,
		Description: "Test healing action",
	}

	success = combat.ExecuteAction(healingAction, char2)
	if !success {
		t.Error("Expected action to succeed")
	}
	if char2.Health != 100 { // 86 + 30 (capped at max health)
		t.Errorf("Expected char2 health 100, got %d", char2.Health)
	}

	// Test insufficient steam power
	char1.SteamPower = 5
	success = combat.ExecuteAction(damageAction, char2)
	if success {
		t.Error("Expected action to fail due to insufficient steam power")
	}
}

func TestIsCombatOver(t *testing.T) {
	char1 := &character.Character{
		ID:         "char1",
		Name:       "Character 1",
		Health:     100,
		SteamPower: 50,
	}

	char2 := &character.Character{
		ID:         "char2",
		Name:       "Character 2",
		Health:     100,
		SteamPower: 50,
	}

	combat := NewCombatState([]*character.Character{char1, char2})

	// Test ongoing combat
	if combat.IsCombatOver() {
		t.Error("Expected combat to be ongoing")
	}

	// Test one character defeated
	char1.Health = 0
	if !combat.IsCombatOver() {
		t.Error("Expected combat to be over (one character defeated)")
	}

	// Test both characters defeated
	char2.Health = 0
	if !combat.IsCombatOver() {
		t.Error("Expected combat to be over (both characters defeated)")
	}
}

func TestGetWinner(t *testing.T) {
	char1 := &character.Character{
		ID:         "char1",
		Name:       "Character 1",
		Health:     100,
		SteamPower: 50,
	}

	char2 := &character.Character{
		ID:         "char2",
		Name:       "Character 2",
		Health:     100,
		SteamPower: 50,
	}

	combat := NewCombatState([]*character.Character{char1, char2})

	// Test ongoing combat
	if winner := combat.GetWinner(); winner != nil {
		t.Error("Expected no winner for ongoing combat")
	}

	// Test char1 wins
	char2.Health = 0
	if winner := combat.GetWinner(); winner != char1 {
		t.Error("Expected char1 to be winner")
	}

	// Test no winner (both defeated)
	char1.Health = 0
	if winner := combat.GetWinner(); winner != nil {
		t.Error("Expected no winner when both characters are defeated")
	}
}
