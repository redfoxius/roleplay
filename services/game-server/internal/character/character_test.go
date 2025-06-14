package character

import (
	"testing"

	"github.com/redfoxius/roleplay/services/game-server/internal/common"
)

func TestNewCharacter(t *testing.T) {
	char := NewCharacter("Test Character", Engineer)

	// Test character creation
	if char.Name != "Test Character" {
		t.Errorf("Expected name 'Test Character', got '%s'", char.Name)
	}
	if char.Class != Engineer {
		t.Errorf("Expected class 'Engineer', got '%s'", char.Class)
	}
	if char.Level != 1 {
		t.Errorf("Expected level 1, got %d", char.Level)
	}
	if char.Health != 100 {
		t.Errorf("Expected health 100, got %d", char.Health)
	}
	if char.MaxHealth != 100 {
		t.Errorf("Expected max health 100, got %d", char.MaxHealth)
	}
	if char.SteamPower != 50 {
		t.Errorf("Expected steam power 50, got %d", char.SteamPower)
	}
	if char.MaxSteamPower != 50 {
		t.Errorf("Expected max steam power 50, got %d", char.MaxSteamPower)
	}
}

func TestApplyClassBonuses(t *testing.T) {
	char := NewCharacter("Test Character", Engineer)
	char.ApplyClassBonuses()

	// Test attribute bonuses
	if char.Stats.Intelligence != 15 {
		t.Errorf("Expected TechnicalAptitude 15, got %d", char.Stats.Intelligence)
	}
	if char.Stats.SteamPower != 15 {
		t.Errorf("Expected MechanicalPrecision 15, got %d", char.Stats.SteamPower)
	}
}

func TestAddToInventory(t *testing.T) {
	char := NewCharacter("Test Character", Engineer)
	item := common.Item{ID: "item1", Name: "Test Item"}

	// Test adding an item to inventory
	char.AddToInventory(item)
	if len(char.Inventory) != 1 {
		t.Errorf("Expected inventory length 1, got %d", len(char.Inventory))
	}
	if char.Inventory[0].ID != "item1" {
		t.Errorf("Expected item ID 'item1', got '%s'", char.Inventory[0].ID)
	}
}

func TestRemoveFromInventory(t *testing.T) {
	char := NewCharacter("Test Character", Engineer)
	item := common.Item{ID: "item1", Name: "Test Item"}
	char.AddToInventory(item)

	// Test removing an item from inventory
	removed := char.RemoveFromInventory("item1")
	if !removed {
		t.Errorf("Expected item to be removed")
	}
	if len(char.Inventory) != 0 {
		t.Errorf("Expected inventory length 0, got %d", len(char.Inventory))
	}
}

func TestAddExperience(t *testing.T) {
	char := NewCharacter("Test Character", Engineer)

	// Test adding experience
	char.AddExperience(500)
	if char.Experience != 500 {
		t.Errorf("Expected experience 500, got %d", char.Experience)
	}

	// Test leveling up
	char.AddExperience(1000)
	if char.Level != 2 {
		t.Errorf("Expected level 2, got %d", char.Level)
	}
	if char.Experience != 1500 {
		t.Errorf("Expected experience 1500, got %d", char.Experience)
	}
}

func TestInventoryManagement(t *testing.T) {
	char := NewCharacter("Test Char", Engineer)

	// Test adding items
	item1 := common.Item{
		ID:          "item1",
		Name:        "Test Item 1",
		Type:        "weapon",
		Description: "A test weapon",
	}

	item2 := common.Item{
		ID:          "item2",
		Name:        "Test Item 2",
		Type:        "consumable",
		Description: "A test consumable",
	}

	char.AddToInventory(item1)
	char.AddToInventory(item2)

	if len(char.Inventory) != 2 {
		t.Errorf("Expected 2 items in inventory, got %d", len(char.Inventory))
	}

	// Test removing items
	if !char.RemoveFromInventory("item1") {
		t.Error("Expected successful removal of item1")
	}
	if len(char.Inventory) != 1 {
		t.Errorf("Expected 1 item in inventory after removal, got %d", len(char.Inventory))
	}

	// Test removing non-existent item
	if char.RemoveFromInventory("nonexistent") {
		t.Error("Expected failed removal of nonexistent item")
	}
}

func TestEquipmentManagement(t *testing.T) {
	char := NewCharacter("Test Char", Engineer)

	// Test equipping items
	err := char.EquipItem(SteamGoggles.ID)
	if err != nil {
		t.Error("Expected successful equipment of SteamGoggles")
	}

	// Test equipping to same slot
	err = char.EquipItem(SteamGoggles.ID)
	if err != nil {
		t.Error("Expected successful equipment of duplicate item")
	}

	// Test unequipping items
	err = char.UnequipItem(string(Head))
	if err != nil {
		t.Error("Expected successful unequip of SteamGoggles")
	}

	// Test unequipping empty slot
	err = char.UnequipItem(string(Head))
	if err == nil {
		t.Error("Expected error when unequipping empty slot")
	}
}

func TestUpdateStats(t *testing.T) {
	char := NewCharacter("Test Char", Engineer)

	// Equip some items
	err := char.EquipItem(SteamGoggles.ID)
	if err != nil {
		t.Error("Expected successful equipment of SteamGoggles")
	}
	err = char.EquipItem(SteamVest.ID)
	if err != nil {
		t.Error("Expected successful equipment of SteamVest")
	}

	// Update stats
	char.UpdateStats()

	// Test steam power update
	expectedSteamPower := 50 + (char.Level * 5) + 10 // Base + Level bonus + Equipment bonus
	if char.MaxSteamPower != expectedSteamPower {
		t.Errorf("Expected max steam power %d, got %d", expectedSteamPower, char.MaxSteamPower)
	}

	// Test attribute updates from equipment
	equipStats := make(map[string]int)
	for _, item := range char.Equipment {
		equipStats["Intelligence"] += item.Stats.Intelligence
		equipStats["SteamPower"] += item.Stats.SteamPower
	}
	for stat, value := range equipStats {
		switch stat {
		case "Intelligence":
			if char.Stats.Intelligence != 13+value { // Base + Equipment
				t.Errorf("Expected Intelligence %d, got %d", 13+value, char.Stats.Intelligence)
			}
		case "SteamPower":
			if char.Stats.SteamPower != 10+value { // Base + Equipment
				t.Errorf("Expected SteamPower %d, got %d", 10+value, char.Stats.SteamPower)
			}
		}
	}
}
