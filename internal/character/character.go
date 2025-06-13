package character

import (
	"math/rand"
	"time"

	"roleplay/internal/common"
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
	ID            string
	Name          string
	Class         Class
	Level         int
	Experience    int
	Health        int
	MaxHealth     int
	SteamPower    int
	MaxSteamPower int
	Attributes    common.Attributes
	Inventory     []Item
	Equipment     *EquipmentSet
	CombatStats   *CombatStats
	Location      common.Coordinates
	Money         common.Currency
}

// NewCharacter creates a new character
func NewCharacter(name string, class Class) *Character {
	baseAttrs := getBaseAttributes(class)
	char := &Character{
		ID:            generateID(),
		Name:          name,
		Class:         class,
		Level:         1,
		Experience:    0,
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
		Attributes:    baseAttrs,
		Inventory:     make([]Item, 0),
		Equipment:     NewEquipmentSet(),
		CombatStats:   NewCombatStats(),
		Location:      common.Coordinates{X: 0, Y: 0},   // Start at origin
		Money:         common.NewCurrency(100, 0, 0, 0), // Start with 100 copper
	}

	// Add starting equipment based on class
	switch class {
	case Engineer:
		char.Equipment.EquipItem(SteamGoggles)
		char.Equipment.EquipItem(SteamVest)
	case Alchemist:
		char.Equipment.EquipItem(SteamGoggles)
		char.Equipment.EquipItem(SteamVest)
	case Aeronaut:
		char.Equipment.EquipItem(SteamGoggles)
		char.Equipment.EquipItem(SteamVest)
	case ClockworkKnight:
		char.Equipment.EquipItem(SteamGoggles)
		char.Equipment.EquipItem(SteamVest)
	case SteamMage:
		char.Equipment.EquipItem(SteamGoggles)
		char.Equipment.EquipItem(SteamVest)
	}

	return char
}

// UpdateStats updates character stats based on equipment
func (c *Character) UpdateStats() {
	// Get equipment bonuses
	equipStats := c.Equipment.GetTotalStats()
	steamBonus := c.Equipment.GetTotalSteamPower()

	// Update attributes
	for stat, value := range equipStats {
		switch stat {
		case "Strength":
			c.Attributes.Strength.Value += value
		case "Dexterity":
			c.Attributes.Dexterity.Value += value
		case "Constitution":
			c.Attributes.Constitution.Value += value
		case "Intelligence":
			c.Attributes.Intelligence.Value += value
		case "Wisdom":
			c.Attributes.Wisdom.Value += value
		case "Charisma":
			c.Attributes.Charisma.Value += value
		case "TechnicalAptitude":
			c.Attributes.TechnicalAptitude.Value += value
		case "SteamPower":
			c.Attributes.SteamPower.Value += value
		case "MechanicalPrecision":
			c.Attributes.MechanicalPrecision.Value += value
		case "ArcaneKnowledge":
			c.Attributes.ArcaneKnowledge.Value += value
		}
	}

	// Update steam power
	c.MaxSteamPower = 50 + (c.Level * 5) + steamBonus
	if c.SteamPower > c.MaxSteamPower {
		c.SteamPower = c.MaxSteamPower
	}
}

// AddToInventory adds an item to the character's inventory
func (c *Character) AddToInventory(item Item) {
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

	// Increase attributes based on class
	switch c.Class {
	case Engineer:
		c.Attributes.TechnicalAptitude.Value += 2
		c.Attributes.MechanicalPrecision.Value += 1
	case Alchemist:
		c.Attributes.Intelligence.Value += 2
		c.Attributes.ArcaneKnowledge.Value += 1
	case Aeronaut:
		c.Attributes.Dexterity.Value += 2
		c.Attributes.SteamPower.Value += 1
	case ClockworkKnight:
		c.Attributes.Strength.Value += 2
		c.Attributes.Constitution.Value += 1
	case SteamMage:
		c.Attributes.ArcaneKnowledge.Value += 2
		c.Attributes.SteamPower.Value += 1
	}
}

// MoveTo moves the character to a new location
func (c *Character) MoveTo(x, y int) {
	c.Location = common.Coordinates{X: x, Y: y}
}

// GetLocation returns the character's current location
func (c *Character) GetLocation() common.Coordinates {
	return c.Location
}

// AddMoney adds money to the character's purse
func (c *Character) AddMoney(money common.Currency) {
	c.Money.Add(money)
}

// RemoveMoney removes money from the character's purse
func (c *Character) RemoveMoney(money common.Currency) bool {
	return c.Money.Subtract(money)
}

// GetMoney returns the character's current money
func (c *Character) GetMoney() common.Currency {
	return c.Money
}
