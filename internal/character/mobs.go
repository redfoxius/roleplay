package character

import (
	"math/rand"
	"strconv"
	"time"
)

// Mob represents a non-player character enemy
type Mob struct {
	ID            string
	Name          string
	Type          string
	Level         int
	Health        int
	MaxHealth     int
	Attributes    Attributes
	SteamPower    int
	MaxSteamPower int
	Experience    int
	Abilities     []Ability
	LootTable     LootTable
	Location      struct {
		X, Y int
	}
}

// MobType represents different categories of mobs
type MobType string

const (
	Mechanical MobType = "Mechanical"
	Biological MobType = "Biological"
	Hybrid     MobType = "Hybrid"
	Elemental  MobType = "Elemental"
	Construct  MobType = "Construct"
)

// NewMob creates a new mob with the given parameters
func NewMob(name string, mobType MobType, level int) *Mob {
	rand.Seed(time.Now().UnixNano())

	baseHealth := 50 + (level * 10)
	baseSteam := 20 + (level * 5)

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
		Location:      struct{ X, Y int }{X: 0, Y: 0},
	}
}

// generateMobID generates a unique ID for a mob
func generateMobID() string {
	return "mob_" + time.Now().Format("20060102150405") + "_" + strconv.Itoa(rand.Intn(1000))
}

// generateMobAttributes generates attributes based on mob type and level
func generateMobAttributes(level int, mobType MobType) Attributes {
	baseValue := 5 + (level / 2)
	attrs := Attributes{
		Strength:            Attribute{Name: "Strength", Value: baseValue},
		Dexterity:           Attribute{Name: "Dexterity", Value: baseValue},
		Constitution:        Attribute{Name: "Constitution", Value: baseValue},
		Intelligence:        Attribute{Name: "Intelligence", Value: baseValue},
		Wisdom:              Attribute{Name: "Wisdom", Value: baseValue},
		Charisma:            Attribute{Name: "Charisma", Value: baseValue},
		TechnicalAptitude:   Attribute{Name: "Technical Aptitude", Value: baseValue},
		SteamPower:          Attribute{Name: "Steam Power", Value: baseValue},
		MechanicalPrecision: Attribute{Name: "Mechanical Precision", Value: baseValue},
		ArcaneKnowledge:     Attribute{Name: "Arcane Knowledge", Value: baseValue},
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

// GenerateMobs generates 50 different mobs
func GenerateMobs() []*Mob {
	mobs := make([]*Mob, 0, 50)

	// Mechanical Mobs
	mobs = append(mobs, []*Mob{
		NewMob("Steam Golem", Mechanical, 1),
		NewMob("Clockwork Spider", Mechanical, 2),
		NewMob("Automaton Guard", Mechanical, 3),
		NewMob("Steam-powered Sentry", Mechanical, 4),
		NewMob("Brass Guardian", Mechanical, 5),
		NewMob("Copper Construct", Mechanical, 6),
		NewMob("Iron Defender", Mechanical, 7),
		NewMob("Steel Automaton", Mechanical, 8),
		NewMob("Bronze Guardian", Mechanical, 9),
		NewMob("Titanium Sentry", Mechanical, 10),
	}...)

	// Biological Mobs
	mobs = append(mobs, []*Mob{
		NewMob("Steam-mutated Rat", Biological, 1),
		NewMob("Chemical Hound", Biological, 2),
		NewMob("Steam-infused Bear", Biological, 3),
		NewMob("Toxic Slime", Biological, 4),
		NewMob("Mutated Wolf", Biological, 5),
		NewMob("Steam-powered Tiger", Biological, 6),
		NewMob("Chemical Lion", Biological, 7),
		NewMob("Toxic Panther", Biological, 8),
		NewMob("Steam-mutated Gorilla", Biological, 9),
		NewMob("Chemical Dragon", Biological, 10),
	}...)

	// Hybrid Mobs
	mobs = append(mobs, []*Mob{
		NewMob("Steam-cyborg", Hybrid, 1),
		NewMob("Mechanical Mutant", Hybrid, 2),
		NewMob("Steam-enhanced Human", Hybrid, 3),
		NewMob("Clockwork Mutant", Hybrid, 4),
		NewMob("Steam-powered Cyborg", Hybrid, 5),
		NewMob("Mechanical Hybrid", Hybrid, 6),
		NewMob("Steam-mutated Cyborg", Hybrid, 7),
		NewMob("Clockwork Human", Hybrid, 8),
		NewMob("Steam-enhanced Mutant", Hybrid, 9),
		NewMob("Mechanical Dragon", Hybrid, 10),
	}...)

	// Elemental Mobs
	mobs = append(mobs, []*Mob{
		NewMob("Steam Elemental", Elemental, 1),
		NewMob("Fire Golem", Elemental, 2),
		NewMob("Steam Spirit", Elemental, 3),
		NewMob("Metal Elemental", Elemental, 4),
		NewMob("Steam Phantom", Elemental, 5),
		NewMob("Iron Elemental", Elemental, 6),
		NewMob("Steam Wraith", Elemental, 7),
		NewMob("Brass Elemental", Elemental, 8),
		NewMob("Steam Specter", Elemental, 9),
		NewMob("Steel Elemental", Elemental, 10),
	}...)

	// Construct Mobs
	mobs = append(mobs, []*Mob{
		NewMob("Steam-powered Tank", Construct, 1),
		NewMob("Clockwork Knight", Construct, 2),
		NewMob("Steam Goliath", Construct, 3),
		NewMob("Mechanical Titan", Construct, 4),
		NewMob("Steam Colossus", Construct, 5),
		NewMob("Clockwork Giant", Construct, 6),
		NewMob("Steam Behemoth", Construct, 7),
		NewMob("Mechanical Leviathan", Construct, 8),
		NewMob("Steam Juggernaut", Construct, 9),
		NewMob("Clockwork Dragon", Construct, 10),
	}...)

	return mobs
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
