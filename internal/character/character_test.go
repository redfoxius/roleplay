package character

import (
	"testing"
)

func TestNewCharacter(t *testing.T) {
	// Test Engineer creation
	engineer := NewCharacter("Test Engineer", Engineer)
	if engineer.Name != "Test Engineer" {
		t.Errorf("Expected name 'Test Engineer', got '%s'", engineer.Name)
	}
	if engineer.Class != Engineer {
		t.Errorf("Expected class Engineer, got %s", engineer.Class)
	}
	if engineer.Level != 1 {
		t.Errorf("Expected level 1, got %d", engineer.Level)
	}
	if engineer.Health != 100 {
		t.Errorf("Expected health 100, got %d", engineer.Health)
	}
	if engineer.SteamPower != 50 {
		t.Errorf("Expected steam power 50, got %d", engineer.SteamPower)
	}

	// Test class-specific attribute bonuses
	if engineer.Attributes.TechnicalAptitude.Value != 13 { // 10 base + 3 bonus
		t.Errorf("Expected TechnicalAptitude 13, got %d", engineer.Attributes.TechnicalAptitude.Value)
	}
	if engineer.Attributes.MechanicalPrecision.Value != 12 { // 10 base + 2 bonus
		t.Errorf("Expected MechanicalPrecision 12, got %d", engineer.Attributes.MechanicalPrecision.Value)
	}
	if engineer.Attributes.Intelligence.Value != 11 { // 10 base + 1 bonus
		t.Errorf("Expected Intelligence 11, got %d", engineer.Attributes.Intelligence.Value)
	}
}

func TestAddExperience(t *testing.T) {
	char := NewCharacter("Test Char", Engineer)

	// Test level 1 to 2
	char.AddExperience(1000) // Should level up
	if char.Level != 2 {
		t.Errorf("Expected level 2, got %d", char.Level)
	}
	if char.MaxHealth != 110 { // 100 + 10
		t.Errorf("Expected max health 110, got %d", char.MaxHealth)
	}
	if char.MaxSteamPower != 55 { // 50 + 5
		t.Errorf("Expected max steam power 55, got %d", char.MaxSteamPower)
	}

	// Test multiple level ups
	char.AddExperience(3000) // Should level up to 3
	if char.Level != 3 {
		t.Errorf("Expected level 3, got %d", char.Level)
	}
	if char.MaxHealth != 120 { // 110 + 10
		t.Errorf("Expected max health 120, got %d", char.MaxHealth)
	}
	if char.MaxSteamPower != 60 { // 55 + 5
		t.Errorf("Expected max steam power 60, got %d", char.MaxSteamPower)
	}
}

func TestInventoryManagement(t *testing.T) {
	char := NewCharacter("Test Char", Engineer)

	// Test adding items
	item1 := Item{
		ID:          "item1",
		Name:        "Test Item 1",
		Type:        "weapon",
		Description: "A test weapon",
		Stackable:   false,
		Quantity:    1,
	}

	item2 := Item{
		ID:          "item2",
		Name:        "Test Item 2",
		Type:        "consumable",
		Description: "A test consumable",
		Stackable:   true,
		Quantity:    5,
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
	oldItem, err := char.Equipment.EquipItem(SteamGoggles)
	if err != nil {
		t.Error("Expected successful equipment of SteamGoggles")
	}
	if oldItem != nil {
		t.Error("Expected no previous item in slot")
	}

	// Test equipping to same slot
	oldItem, err = char.Equipment.EquipItem(SteamGoggles)
	if err != nil {
		t.Error("Expected successful equipment of duplicate item")
	}
	if oldItem != SteamGoggles {
		t.Error("Expected previous item to be returned")
	}

	// Test unequipping items
	item, err := char.Equipment.UnequipItem(Head)
	if err != nil {
		t.Error("Expected successful unequip of SteamGoggles")
	}
	if item != SteamGoggles {
		t.Error("Expected SteamGoggles to be returned")
	}

	// Test unequipping empty slot
	item, err = char.Equipment.UnequipItem(Head)
	if err == nil {
		t.Error("Expected error when unequipping empty slot")
	}
	if item != nil {
		t.Error("Expected nil item when unequipping empty slot")
	}
}

func TestUpdateStats(t *testing.T) {
	char := NewCharacter("Test Char", Engineer)

	// Equip some items
	_, err := char.Equipment.EquipItem(SteamGoggles)
	if err != nil {
		t.Error("Expected successful equipment of SteamGoggles")
	}
	_, err = char.Equipment.EquipItem(SteamVest)
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
	equipStats := char.Equipment.GetTotalStats()
	for stat, value := range equipStats {
		switch stat {
		case "TechnicalAptitude":
			if char.Attributes.TechnicalAptitude.Value != 13+value { // Base + Equipment
				t.Errorf("Expected TechnicalAptitude %d, got %d", 13+value, char.Attributes.TechnicalAptitude.Value)
			}
		case "SteamPower":
			if char.Attributes.SteamPower.Value != 10+value { // Base + Equipment
				t.Errorf("Expected SteamPower %d, got %d", 10+value, char.Attributes.SteamPower.Value)
			}
		}
	}
}
