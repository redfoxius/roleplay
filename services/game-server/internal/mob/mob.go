package mob

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/redfoxius/roleplay/services/game-server/internal/common"
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
	Position      common.Coordinates
	Stats         common.Stats
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

	mob := &Mob{
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
		Position:      common.Coordinates{X: 0, Y: 0},
		Stats: common.Stats{
			Strength:     10 * level,
			Dexterity:    10 * level,
			Intelligence: 10 * level,
			Vitality:     10 * level,
		},
		Abilities: make([]common.Ability, 0),
		LootTable: common.LootTable{},
	}

	// Add type-specific abilities
	switch mobType {
	case Mechanical:
		mob.Abilities = append(mob.Abilities, common.Ability{
			Name:        "Steam Punch",
			Damage:      20 * level,
			SteamCost:   10 * level,
			Description: "A powerful steam-powered punch",
		})
	case Biological:
		mob.Abilities = append(mob.Abilities, common.Ability{
			Name:        "Web Shot",
			Damage:      15 * level,
			SteamCost:   8 * level,
			Description: "Shoots a web to slow down enemies",
		})
	case Hybrid:
		mob.Abilities = append(mob.Abilities, common.Ability{
			Name:        "Steam Burst",
			Damage:      25 * level,
			SteamCost:   15 * level,
			Description: "Releases a burst of scalding steam",
		})
	case Elemental:
		// No additional abilities for elemental mobs
	case Construct:
		// No additional abilities for construct mobs
	}

	return mob
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

// MoveTo moves the mob to a new location
func (m *Mob) MoveTo(x, y int) {
	m.Position = common.Coordinates{X: x, Y: y}
}

// GetLocation returns the mob's current location
func (m *Mob) GetLocation() common.Coordinates {
	return m.Position
}

// UseAbility uses an ability and returns the damage dealt
func (m *Mob) UseAbility(abilityName string) (int, error) {
	for _, ability := range m.Abilities {
		if ability.Name == abilityName {
			if m.SteamPower < ability.SteamCost {
				return 0, fmt.Errorf("insufficient steam power")
			}
			m.SteamPower -= ability.SteamCost
			return ability.Damage, nil
		}
	}
	return 0, fmt.Errorf("ability not found")
}

// AddToLootTable adds an item to the mob's loot table
func (m *Mob) AddToLootTable(item common.Item) {
	lootItem := common.LootItem{
		ID:       item.ID,
		Name:     item.Name,
		Quantity: 1,
		Chance:   1.0,
	}
	m.LootTable.Items = append(m.LootTable.Items, lootItem)
}

// GetLoot returns a random item from the loot table
func (m *Mob) GetLoot() *common.Item {
	if len(m.LootTable.Items) == 0 {
		return nil
	}
	// TODO: Implement proper loot table with drop rates
	return &common.Item{
		ID:   m.LootTable.Items[0].ID,
		Name: m.LootTable.Items[0].Name,
	}
}

// TakeDamage reduces the mob's health
func (m *Mob) TakeDamage(damage int) {
	m.Health -= damage
	if m.Health < 0 {
		m.Health = 0
	}
}

// IsDead returns true if the mob's health is 0
func (m *Mob) IsDead() bool {
	return m.Health <= 0
}

// Heal restores health to the mob
func (m *Mob) Heal(amount int) {
	m.Health += amount
	if m.Health > m.MaxHealth {
		m.Health = m.MaxHealth
	}
}

// RestoreSteamPower restores steam power to the mob
func (m *Mob) RestoreSteamPower(amount int) {
	m.SteamPower += amount
	if m.SteamPower > m.MaxSteamPower {
		m.SteamPower = m.MaxSteamPower
	}
}
