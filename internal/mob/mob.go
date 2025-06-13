package mob

import (
	"math/rand"
	"strconv"
	"time"

	"roleplay/internal/common"
)

// Mob represents a non-player character enemy
type Mob struct {
	ID            string
	Name          string
	Type          string
	Level         int
	Health        int
	MaxHealth     int
	Attributes    common.Attributes
	SteamPower    int
	MaxSteamPower int
	Experience    int
	Abilities     []common.Ability
	LootTable     common.LootTable
	Location      common.Coordinates
	MoneyDrop     common.Currency
}

// MobType represents different categories of mobs
type MobType string

const (
	Mechanical MobType = "mechanical"
	Biological MobType = "biological"
	Hybrid     MobType = "hybrid"
	Elemental  MobType = "elemental"
	Construct  MobType = "construct"
)

// NewMob creates a new mob with the given parameters
func NewMob(name string, mobType MobType, level int) *Mob {
	rand.Seed(time.Now().UnixNano())

	baseHealth := 50 + (level * 10)
	baseSteam := 20 + (level * 5)

	// Generate money drop based on level and type
	moneyDrop := generateMoneyDrop(level, mobType)

	return &Mob{
		ID:            generateMobID(),
		Name:          name,
		Type:          string(mobType),
		Level:         level,
		Health:        baseHealth,
		MaxHealth:     baseHealth,
		SteamPower:    baseSteam,
		MaxSteamPower: baseSteam,
		Experience:    level * 100,
		Attributes:    generateMobAttributes(level, mobType),
		Location:      common.Coordinates{X: 0, Y: 0},
		MoneyDrop:     moneyDrop,
	}
}

// generateMobID generates a unique ID for a mob
func generateMobID() string {
	return "mob_" + time.Now().Format("20060102150405") + "_" + strconv.Itoa(rand.Intn(1000))
}

// generateMobAttributes generates attributes based on mob type and level
func generateMobAttributes(level int, mobType MobType) common.Attributes {
	baseValue := 5 + (level / 2)
	attrs := common.Attributes{
		Strength:            common.Attribute{Name: "Strength", Value: baseValue},
		Dexterity:           common.Attribute{Name: "Dexterity", Value: baseValue},
		Constitution:        common.Attribute{Name: "Constitution", Value: baseValue},
		Intelligence:        common.Attribute{Name: "Intelligence", Value: baseValue},
		Wisdom:              common.Attribute{Name: "Wisdom", Value: baseValue},
		Charisma:            common.Attribute{Name: "Charisma", Value: baseValue},
		TechnicalAptitude:   common.Attribute{Name: "Technical Aptitude", Value: baseValue},
		SteamPower:          common.Attribute{Name: "Steam Power", Value: baseValue},
		MechanicalPrecision: common.Attribute{Name: "Mechanical Precision", Value: baseValue},
		ArcaneKnowledge:     common.Attribute{Name: "Arcane Knowledge", Value: baseValue},
	}

	// Apply type-specific bonuses
	switch mobType {
	case Mechanical:
		attrs.TechnicalAptitude.Value += level
		attrs.MechanicalPrecision.Value += level
	case Biological:
		attrs.Strength.Value += level
		attrs.Constitution.Value += level
	case Hybrid:
		attrs.SteamPower.Value += level
		attrs.ArcaneKnowledge.Value += level
	case Elemental:
		attrs.ArcaneKnowledge.Value += level * 2
	case Construct:
		attrs.MechanicalPrecision.Value += level
		attrs.Constitution.Value += level
	}

	return attrs
}

// generateMoneyDrop generates a money drop based on mob level and type
func generateMoneyDrop(level int, mobType MobType) common.Currency {
	// Base copper amount based on level
	baseCopper := level * 10

	// Type-specific multipliers
	multiplier := 1.0
	switch mobType {
	case Mechanical:
		multiplier = 1.2 // Mechanical mobs drop more money
	case Biological:
		multiplier = 0.8 // Biological mobs drop less money
	case Hybrid:
		multiplier = 1.0
	case Elemental:
		multiplier = 1.5 // Elemental mobs drop the most money
	case Construct:
		multiplier = 1.3 // Constructs drop more money
	}

	// Add some randomness
	randomFactor := 0.8 + (rand.Float64() * 0.4) // 0.8 to 1.2
	totalCopper := int(float64(baseCopper) * multiplier * randomFactor)

	// Convert to appropriate denominations
	return common.NewCurrency(totalCopper, 0, 0, 0)
}

// GenerateMobName generates a name for a mob based on its type
func GenerateMobName(mobType MobType) string {
	prefixes := map[MobType][]string{
		Mechanical: {"Steam", "Clockwork", "Brass", "Copper", "Iron"},
		Biological: {"Mutated", "Toxic", "Chemical", "Steam-infused", "Enhanced"},
		Hybrid:     {"Steam-cyborg", "Mechanical", "Enhanced", "Steam-powered", "Clockwork"},
		Elemental:  {"Steam", "Fire", "Metal", "Iron", "Brass"},
		Construct:  {"Steam-powered", "Clockwork", "Mechanical", "Steam", "Brass"},
	}

	suffixes := map[MobType][]string{
		Mechanical: {"Golem", "Spider", "Guard", "Sentry", "Guardian"},
		Biological: {"Rat", "Hound", "Bear", "Slime", "Wolf"},
		Hybrid:     {"Cyborg", "Mutant", "Human", "Hybrid", "Dragon"},
		Elemental:  {"Elemental", "Golem", "Spirit", "Phantom", "Wraith"},
		Construct:  {"Tank", "Knight", "Goliath", "Titan", "Colossus"},
	}

	prefix := prefixes[mobType][rand.Intn(len(prefixes[mobType]))]
	suffix := suffixes[mobType][rand.Intn(len(suffixes[mobType]))]

	return prefix + " " + suffix
}

// GetMobAbilities returns the abilities for a mob type and level
func GetMobAbilities(mobType MobType, level int) []common.Ability {
	// TODO: Implement mob abilities
	return []common.Ability{}
}

// GenerateLootTable generates a loot table for a mob type and level
func GenerateLootTable(mobType MobType, level int) common.LootTable {
	// TODO: Implement loot table generation
	return common.LootTable{}
}
