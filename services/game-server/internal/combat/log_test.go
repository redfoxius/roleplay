package combat

import (
	"testing"
)

func TestCombatLog(t *testing.T) {
	log := NewCombatLog()

	// Test adding an attack entry
	log.LogAttack("Player1", "Mob1", 10)
	entries := log.GetRecentEntries(1)
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(entries))
	}
	if entries[0].Type != "attack" {
		t.Errorf("Expected type 'attack', got '%s'", entries[0].Type)
	}
	if entries[0].Message != "Player1 attacks Mob1 for 10" {
		t.Errorf("Expected message 'Player1 attacks Mob1 for 10', got '%s'", entries[0].Message)
	}

	// Test adding an ability entry
	log.LogAbility("Player1", "Mob1", "Steam Blast", 15)
	entries = log.GetRecentEntries(1)
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(entries))
	}
	if entries[0].Type != "ability" {
		t.Errorf("Expected type 'ability', got '%s'", entries[0].Type)
	}
	if entries[0].Message != "Player1 uses Steam Blast Mob1 for 15" {
		t.Errorf("Expected message 'Player1 uses Steam Blast Mob1 for 15', got '%s'", entries[0].Message)
	}

	// Test adding an effect entry
	log.LogEffect("Player1", "Mob1", "Poison", 5)
	entries = log.GetRecentEntries(1)
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(entries))
	}
	if entries[0].Type != "effect" {
		t.Errorf("Expected type 'effect', got '%s'", entries[0].Type)
	}
	if entries[0].Message != "Player1 is affected by Poison Mob1 for 5" {
		t.Errorf("Expected message 'Player1 is affected by Poison Mob1 for 5', got '%s'", entries[0].Message)
	}

	// Test adding a death entry
	log.LogDeath("Mob1")
	entries = log.GetRecentEntries(1)
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(entries))
	}
	if entries[0].Type != "death" {
		t.Errorf("Expected type 'death', got '%s'", entries[0].Type)
	}
	if entries[0].Message != "Mob1 has been defeated" {
		t.Errorf("Expected message 'Mob1 has been defeated', got '%s'", entries[0].Message)
	}

	// Test adding a loot entry
	log.LogLoot("Player1", "Gold", 100)
	entries = log.GetRecentEntries(1)
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(entries))
	}
	if entries[0].Type != "loot" {
		t.Errorf("Expected type 'loot', got '%s'", entries[0].Type)
	}
	if entries[0].Message != "Player1 receives 100 Gold" {
		t.Errorf("Expected message 'Player1 receives 100 Gold', got '%s'", entries[0].Message)
	}
}
