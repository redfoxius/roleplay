package character

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/redfoxius/roleplay/services/game-server/internal/common"
)

// generateID generates a unique ID for characters
func generateID() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Character represents a player character
type Character struct {
	ID                string
	Name              string
	Class             Class
	Level             int
	Experience        int
	Health            int
	MaxHealth         int
	SteamPower        int
	MaxSteamPower     int
	MovementPoints    int
	MaxMovementPoints int
	Position          common.Coordinates
	Inventory         []common.Item
	Equipment         map[string]common.Item
	Abilities         []common.Ability
	Stats             common.Stats
}

// Class represents a character class in the game
type Class string

const (
	Engineer        Class = "Engineer"
	Alchemist       Class = "Alchemist"
	Aeronaut        Class = "Aeronaut"
	ClockworkKnight Class = "ClockworkKnight"
	SteamMage       Class = "SteamMage"
)

// ApplyClassBonuses applies the initial class bonuses to attributes
func (c *Character) ApplyClassBonuses() {
	switch c.Class {
	case Engineer:
		c.Stats.Intelligence += 3
		c.Stats.Dexterity += 3
		c.MaxSteamPower += 10
	case Alchemist:
		c.Stats.Intelligence += 10
	case Aeronaut:
		c.Stats.Dexterity += 5
		c.Stats.SteamPower += 5
	case ClockworkKnight:
		c.Stats.Strength += 5
		c.Stats.Vitality += 5
	case SteamMage:
		c.Stats.Intelligence += 5
		c.Stats.SteamPower += 5
	}
}

// NewCharacter creates a new character
func NewCharacter(name string, class Class) *Character {
	char := &Character{
		ID:                generateID(),
		Name:              name,
		Class:             class,
		Level:             1,
		Experience:        0,
		Health:            100,
		MaxHealth:         100,
		SteamPower:        50,
		MaxSteamPower:     50,
		MovementPoints:    100,
		MaxMovementPoints: 100,
		Position:          common.Coordinates{X: 0, Y: 0},
		Inventory:         make([]common.Item, 0),
		Equipment:         make(map[string]common.Item),
		Abilities:         make([]common.Ability, 0),
		Stats: common.Stats{
			Strength:     10,
			Dexterity:    10,
			Intelligence: 10,
			Vitality:     10,
		},
	}

	// Add starting equipment based on class
	switch class {
	case Engineer:
		char.Equipment["head"] = convertEquipmentToItem(SteamGoggles)
		char.Equipment["body"] = convertEquipmentToItem(SteamVest)
	case Alchemist:
		char.Equipment["head"] = convertEquipmentToItem(SteamGoggles)
		char.Equipment["body"] = convertEquipmentToItem(SteamVest)
	case Aeronaut:
		char.Equipment["head"] = convertEquipmentToItem(SteamGoggles)
		char.Equipment["body"] = convertEquipmentToItem(SteamVest)
	case ClockworkKnight:
		char.Equipment["head"] = convertEquipmentToItem(SteamGoggles)
		char.Equipment["body"] = convertEquipmentToItem(SteamVest)
	case SteamMage:
		char.Equipment["head"] = convertEquipmentToItem(SteamGoggles)
		char.Equipment["body"] = convertEquipmentToItem(SteamVest)
	}

	// Add class-specific abilities
	switch class {
	case Engineer:
		char.Abilities = append(char.Abilities, common.Ability{
			Name:        "Steam Blast",
			Description: "A powerful blast of steam",
			Type:        "mechanical",
			Damage:      20,
			SteamCost:   15,
			Range:       3,
			Cooldown:    2,
		})
		char.Abilities = append(char.Abilities, common.Ability{
			Name:        "Repair",
			Description: "Repair mechanical damage",
			Type:        "healing",
			Healing:     15,
			SteamCost:   10,
			Range:       2,
			Cooldown:    3,
		})
	case Alchemist:
		char.Abilities = append(char.Abilities, common.Ability{
			Name:        "Toxic Cloud",
			Description: "Release a cloud of toxic gas",
			Type:        "chemical",
			Damage:      15,
			SteamCost:   12,
			Range:       2,
			Area:        2,
			Cooldown:    3,
		})
		char.Abilities = append(char.Abilities, common.Ability{
			Name:        "Healing Vapor",
			Description: "Release healing steam",
			Type:        "healing",
			Healing:     20,
			SteamCost:   15,
			Range:       2,
			Cooldown:    4,
		})
	case SteamMage:
		char.Abilities = append(char.Abilities, common.Ability{
			Name:        "Steam Bolt",
			Description: "A bolt of condensed steam",
			Type:        "arcane",
			Damage:      25,
			SteamCost:   20,
			Range:       4,
			Cooldown:    2,
		})
		char.Abilities = append(char.Abilities, common.Ability{
			Name:        "Steam Shield",
			Description: "Create a protective steam barrier",
			Type:        "healing",
			Healing:     10,
			SteamCost:   15,
			Range:       1,
			Cooldown:    3,
		})
	}

	return char
}

// UpdateStats updates character stats based on equipment
func (c *Character) UpdateStats() {
	// Get equipment bonuses
	equipStats := c.Equipment

	// Update attributes
	for _, item := range equipStats {
		c.Stats.Strength += item.Stats.Strength
		c.Stats.Dexterity += item.Stats.Dexterity
		c.Stats.Intelligence += item.Stats.Intelligence
		c.Stats.Vitality += item.Stats.Vitality
	}

	// Update steam power
	c.MaxSteamPower = 50 + (c.Level * 5) + c.Stats.SteamPower
	if c.SteamPower > c.MaxSteamPower {
		c.SteamPower = c.MaxSteamPower
	}
}

// AddToInventory adds an item to the character's inventory
func (c *Character) AddToInventory(item common.Item) {
	c.Inventory = append(c.Inventory, item)
}

// RemoveFromInventory removes an item from the character's inventory
func (c *Character) RemoveFromInventory(itemID string) bool {
	for i, item := range c.Inventory {
		if item.ID == itemID {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

// ExperienceToNextLevel returns the experience needed for the next level
func (c *Character) ExperienceToNextLevel() int {
	return c.Level * 1000
}

// AddExperience adds experience to the character and handles leveling
func (c *Character) AddExperience(exp int) {
	c.Experience += exp
	for c.Experience >= c.ExperienceToNextLevel() {
		c.LevelUp()
	}
}

// LevelUp handles character leveling
func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.MaxSteamPower += 5
	c.SteamPower = c.MaxSteamPower
	c.MaxMovementPoints += 5
	c.MovementPoints = c.MaxMovementPoints

	// Increase attributes based on class
	switch c.Class {
	case Engineer:
		c.Stats.Strength += 2
		c.Stats.Intelligence += 1
	case Alchemist:
		c.Stats.Intelligence += 2
		c.Stats.SteamPower += 1
	case Aeronaut:
		c.Stats.Dexterity += 2
		c.Stats.SteamPower += 1
	case ClockworkKnight:
		c.Stats.Strength += 2
		c.Stats.Vitality += 1
	case SteamMage:
		c.Stats.Intelligence += 2
		c.Stats.SteamPower += 1
	}
}

// MoveTo moves the character to a new location
func (c *Character) MoveTo(x, y int) {
	c.Position = common.Coordinates{X: x, Y: y}
}

// GetLocation returns the character's current location
func (c *Character) GetLocation() common.Coordinates {
	return c.Position
}

// AddMoney adds money to the character's purse
func (c *Character) AddMoney(money common.Currency) {
	// Implementation needed
}

// RemoveMoney removes money from the character's purse
func (c *Character) RemoveMoney(money common.Currency) bool {
	// Implementation needed
	return false
}

// GetMoney returns the character's current money
func (c *Character) GetMoney() common.Currency {
	// Implementation needed
	return common.Currency{}
}

// GetAbilities returns the character's abilities
func (c *Character) GetAbilities() []common.Ability {
	return c.Abilities
}

// UseAbility uses an ability and returns the damage dealt
func (c *Character) UseAbility(abilityName string) (int, error) {
	for _, ability := range c.Abilities {
		if ability.Name == abilityName {
			if c.SteamPower < ability.SteamCost {
				return 0, fmt.Errorf("insufficient steam power")
			}
			c.SteamPower -= ability.SteamCost
			return ability.Damage, nil
		}
	}
	return 0, fmt.Errorf("ability not found")
}

// AddItem adds an item to the character's inventory
func (c *Character) AddItem(item common.Item) {
	c.Inventory = append(c.Inventory, item)
}

// RemoveItem removes an item from the character's inventory
func (c *Character) RemoveItem(itemID string) {
	for i, item := range c.Inventory {
		if item.ID == itemID {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			break
		}
	}
}

// EquipItem equips an item from the inventory
func (c *Character) EquipItem(itemID string) error {
	var item common.Item
	var itemIndex int
	found := false

	// Find the item in inventory
	for i, invItem := range c.Inventory {
		if invItem.ID == itemID {
			item = invItem
			itemIndex = i
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("item not found in inventory")
	}

	// Remove from inventory
	c.Inventory = append(c.Inventory[:itemIndex], c.Inventory[itemIndex+1:]...)

	// If there's already an item equipped in that slot, unequip it first
	if existingItem, exists := c.Equipment[item.Slot]; exists {
		c.Inventory = append(c.Inventory, existingItem)
	}

	// Equip the new item
	c.Equipment[item.Slot] = item

	// Apply item stats
	c.Stats.Strength += item.Stats.Strength
	c.Stats.Dexterity += item.Stats.Dexterity
	c.Stats.Intelligence += item.Stats.Intelligence
	c.Stats.Vitality += item.Stats.Vitality

	return nil
}

// UnequipItem unequips an item and puts it back in inventory
func (c *Character) UnequipItem(slot string) error {
	item, exists := c.Equipment[slot]
	if !exists {
		return fmt.Errorf("no item equipped in that slot")
	}

	// Remove item stats
	c.Stats.Strength -= item.Stats.Strength
	c.Stats.Dexterity -= item.Stats.Dexterity
	c.Stats.Intelligence -= item.Stats.Intelligence
	c.Stats.Vitality -= item.Stats.Vitality

	// Add to inventory
	c.Inventory = append(c.Inventory, item)

	// Remove from equipment
	delete(c.Equipment, slot)

	return nil
}

// convertEquipmentToItem converts Equipment to common.Item
func convertEquipmentToItem(eq *Equipment) common.Item {
	return common.Item{
		ID:          eq.ID,
		Name:        eq.Name,
		Description: eq.Description,
		Type:        "equipment",
		Slot:        string(eq.Slot),
		Stats: common.Stats{
			Strength:     eq.Stats["Strength"],
			Dexterity:    eq.Stats["Dexterity"],
			Intelligence: eq.Stats["Intelligence"],
			Vitality:     eq.Stats["Constitution"],
			SteamPower:   eq.SteamPower,
		},
		Value: eq.Level * 10,
	}
}
