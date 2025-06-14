package common

import (
	"fmt"
	"strings"
)

// Coordinates represents a position in the game world
type Coordinates struct {
	X int
	Y int
}

// Attributes represents character attributes
type Attributes struct {
	Strength            Attribute
	Dexterity           Attribute
	Constitution        Attribute
	Intelligence        Attribute
	Wisdom              Attribute
	Charisma            Attribute
	TechnicalAptitude   Attribute
	SteamPower          Attribute
	MechanicalPrecision Attribute
	ArcaneKnowledge     Attribute
}

// Attribute represents a single attribute
type Attribute struct {
	Name  string
	Value int
}

// Stats represents character/mob statistics
type Stats struct {
	Strength     int
	Dexterity    int
	Intelligence int
	Vitality     int
	SteamPower   int
}

// Item represents an item in the game
type Item struct {
	ID          string
	Name        string
	Description string
	Type        string
	Slot        string
	Stats       Stats
	Value       int
}

// Ability represents a character/mob ability
type Ability struct {
	Name        string
	Description string
	Type        string
	Damage      int
	Healing     int
	SteamCost   int
	Range       int
	Area        int
	Cooldown    int
}

// LootTable represents a table of possible loot drops
type LootTable struct {
	Items []LootItem
}

// LootItem represents an item that can be dropped
type LootItem struct {
	ID       string
	Name     string
	Quantity int
	Chance   float64
}

// Currency represents money in the game
type Currency struct {
	Copper   int
	Silver   int
	Gold     int
	Platinum int
}

// NewCurrency creates a new currency with the specified amounts
func NewCurrency(copper, silver, gold, platinum int) Currency {
	return Currency{
		Copper:   copper,
		Silver:   silver,
		Gold:     gold,
		Platinum: platinum,
	}
}

// Add adds another currency to this one
func (c *Currency) Add(other Currency) {
	c.Copper += other.Copper
	c.Silver += other.Silver
	c.Gold += other.Gold
	c.Platinum += other.Platinum
}

// Subtract subtracts another currency from this one
func (c *Currency) Subtract(other Currency) bool {
	// Convert everything to copper for comparison
	totalCopper := c.Copper + (c.Silver * 100) + (c.Gold * 10000) + (c.Platinum * 1000000)
	otherCopper := other.Copper + (other.Silver * 100) + (other.Gold * 10000) + (other.Platinum * 1000000)

	if totalCopper < otherCopper {
		return false
	}

	// Convert back to original denominations
	remaining := totalCopper - otherCopper
	c.Platinum = remaining / 1000000
	remaining %= 1000000
	c.Gold = remaining / 10000
	remaining %= 10000
	c.Silver = remaining / 100
	remaining %= 100
	c.Copper = remaining

	return true
}

// GetTotalValue returns the total value in copper pieces
func (c *Currency) GetTotalValue() int {
	return c.Copper + (c.Silver * 100) + (c.Gold * 10000) + (c.Platinum * 1000000)
}

// String returns a string representation of the currency
func (c Currency) String() string {
	var result string
	if c.Platinum > 0 {
		result += fmt.Sprintf("%dp ", c.Platinum)
	}
	if c.Gold > 0 {
		result += fmt.Sprintf("%dg ", c.Gold)
	}
	if c.Silver > 0 {
		result += fmt.Sprintf("%ds ", c.Silver)
	}
	if c.Copper > 0 {
		result += fmt.Sprintf("%dc", c.Copper)
	}
	return strings.TrimSpace(result)
}
