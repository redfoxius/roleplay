package common

import (
	"fmt"
	"strings"
)

// Coordinates represents a position in the world
type Coordinates struct {
	X, Y int
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

// Ability represents a character or mob ability
type Ability struct {
	Name        string
	Description string
	Type        string // "damage", "healing", "mechanical", "chemical", "arcane", "ranged", "melee"
	Damage      int
	Healing     int
	SteamCost   int
	Range       int // Range of the ability in tiles
	Area        int // Area of effect in tiles (0 for single target)
	Cooldown    int // Cooldown in turns
	LastUsed    int // Last round this ability was used
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

// Currency represents the game's currency
type Currency struct {
	Copper   int
	Silver   int
	Gold     int
	Platinum int
}

// NewCurrency creates a new currency with the given amounts
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
	c.Normalize()
}

// Subtract subtracts another currency from this one
func (c *Currency) Subtract(other Currency) bool {
	// Convert everything to copper for comparison
	totalCopper := c.ToCopper()
	otherCopper := other.ToCopper()

	if totalCopper < otherCopper {
		return false
	}

	// Convert back to original currency
	c.Copper = totalCopper - otherCopper
	c.Normalize()
	return true
}

// ToCopper converts all currency to copper
func (c Currency) ToCopper() int {
	return c.Copper +
		(c.Silver * 100) +
		(c.Gold * 10000) +
		(c.Platinum * 1000000)
}

// Normalize converts excess coins to higher denominations
func (c *Currency) Normalize() {
	// Convert copper to silver
	c.Silver += c.Copper / 100
	c.Copper = c.Copper % 100

	// Convert silver to gold
	c.Gold += c.Silver / 100
	c.Silver = c.Silver % 100

	// Convert gold to platinum
	c.Platinum += c.Gold / 100
	c.Gold = c.Gold % 100
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
