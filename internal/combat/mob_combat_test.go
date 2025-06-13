package combat

import (
	"testing"

	"roleplay/internal/character"
	"roleplay/internal/common"
	"roleplay/internal/mob"
)

func TestNewMobCombat(t *testing.T) {
	char := &character.Character{
		ID:         "test_char",
		Name:       "Test Character",
		Health:     100,
		SteamPower: 50,
		Attributes: common.Attributes{
			Strength: common.Attribute{Name: "Strength", Value: 10},
		},
	}

	m := &mob.Mob{
		ID:         "test_mob",
		Name:       "Test Mob",
		Type:       string(mob.Mechanical),
		Health:     50,
		SteamPower: 30,
		Attributes: common.Attributes{
			Strength: common.Attribute{Name: "Strength", Value: 8},
		},
	}

	combat := NewMobCombat(char, m)

	if combat.Character != char {
		t.Errorf("Expected character %v, got %v", char, combat.Character)
	}
	if combat.Mob != m {
		t.Errorf("Expected mob %v, got %v", m, combat.Mob)
	}
	if combat.Turn != 1 {
		t.Errorf("Expected turn 1, got %d", combat.Turn)
	}
}

func TestCharacterAttack(t *testing.T) {
	char := &character.Character{
		ID:         "test_char",
		Name:       "Test Character",
		Health:     100,
		SteamPower: 50,
		Attributes: common.Attributes{
			Strength: common.Attribute{Name: "Strength", Value: 10},
		},
	}

	m := &mob.Mob{
		ID:         "test_mob",
		Name:       "Test Mob",
		Type:       string(mob.Mechanical),
		Health:     50,
		SteamPower: 30,
		Attributes: common.Attributes{
			Strength: common.Attribute{Name: "Strength", Value: 8},
		},
		Experience: 100,
	}

	combat := NewMobCombat(char, m)

	ability := &common.Ability{
		Name:        "Test Attack",
		Description: "A test attack",
		Type:        "damage",
		Damage:      20,
		SteamCost:   10,
		Cooldown:    0,
	}

	// Test successful attack
	damage, err := combat.CharacterAttack(ability)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if damage != 25 { // 20 base + (10 strength / 2)
		t.Errorf("Expected damage 25, got %d", damage)
	}
	if m.Health != 25 { // 50 - 25
		t.Errorf("Expected mob health 25, got %d", m.Health)
	}
	if char.SteamPower != 40 { // 50 - 10
		t.Errorf("Expected character steam power 40, got %d", char.SteamPower)
	}

	// Test insufficient steam power
	char.SteamPower = 5
	_, err = combat.CharacterAttack(ability)
	if err != ErrInsufficientSteamPower {
		t.Errorf("Expected ErrInsufficientSteamPower, got %v", err)
	}
}

func TestMobAttack(t *testing.T) {
	char := &character.Character{
		ID:         "test_char",
		Name:       "Test Character",
		Health:     100,
		SteamPower: 50,
		Attributes: common.Attributes{
			Strength: common.Attribute{Name: "Strength", Value: 10},
		},
	}

	m := &mob.Mob{
		ID:         "test_mob",
		Name:       "Test Mob",
		Type:       string(mob.Mechanical),
		Health:     50,
		SteamPower: 30,
		Attributes: common.Attributes{
			Strength: common.Attribute{Name: "Strength", Value: 8},
		},
		Abilities: []common.Ability{
			{
				Name:        "Mob Attack",
				Description: "A basic attack",
				Type:        "damage",
				Damage:      15,
				SteamCost:   10,
				Cooldown:    0,
			},
		},
	}

	combat := NewMobCombat(char, m)

	damage, err := combat.MobAttack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if damage != 19 { // 15 base + (8 strength / 2)
		t.Errorf("Expected damage 19, got %d", damage)
	}
	if char.Health != 81 { // 100 - 19
		t.Errorf("Expected character health 81, got %d", char.Health)
	}
	if m.SteamPower != 20 { // 30 - 10
		t.Errorf("Expected mob steam power 20, got %d", m.SteamPower)
	}
}

func TestMobIsCombatOver(t *testing.T) {
	char := &character.Character{
		ID:         "test_char",
		Name:       "Test Character",
		Health:     100,
		SteamPower: 50,
	}

	m := &mob.Mob{
		ID:         "test_mob",
		Name:       "Test Mob",
		Type:       string(mob.Mechanical),
		Health:     50,
		SteamPower: 30,
	}

	combat := NewMobCombat(char, m)

	// Test ongoing combat
	if combat.IsCombatOver() {
		t.Error("Expected combat to be ongoing")
	}

	// Test character defeat
	char.Health = 0
	if !combat.IsCombatOver() {
		t.Error("Expected combat to be over (character defeated)")
	}

	// Test mob defeat
	char.Health = 100
	m.Health = 0
	if !combat.IsCombatOver() {
		t.Error("Expected combat to be over (mob defeated)")
	}
}

func TestGetCombatResult(t *testing.T) {
	char := &character.Character{
		ID:         "test_char",
		Name:       "Test Character",
		Health:     100,
		SteamPower: 50,
	}

	m := &mob.Mob{
		ID:         "test_mob",
		Name:       "Test Mob",
		Type:       string(mob.Mechanical),
		Health:     50,
		SteamPower: 30,
	}

	combat := NewMobCombat(char, m)

	// Test ongoing combat
	if result := combat.GetCombatResult(); result != "ongoing" {
		t.Errorf("Expected result 'ongoing', got '%s'", result)
	}

	// Test character defeat
	char.Health = 0
	if result := combat.GetCombatResult(); result != "defeat" {
		t.Errorf("Expected result 'defeat', got '%s'", result)
	}

	// Test mob defeat
	char.Health = 100
	m.Health = 0
	if result := combat.GetCombatResult(); result != "victory" {
		t.Errorf("Expected result 'victory', got '%s'", result)
	}
}
