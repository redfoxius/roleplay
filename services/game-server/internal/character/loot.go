package character

import (
	"math/rand"
)

// Item represents a game item
type Item struct {
	ID          string
	Name        string
	Type        string // weapon, armor, consumable, material
	Rarity      string // common, uncommon, rare, epic, legendary
	Level       int
	Value       int
	Description string
	Stats       map[string]int
	Stackable   bool
	Quantity    int
}

// LootTable represents a mob's possible drops
type LootTable struct {
	Common     []Item
	Uncommon   []Item
	Rare       []Item
	Epic       []Item
	Legendary  []Item
	SteamCores int
}

// GenerateLootTable creates a loot table for a mob
func GenerateLootTable(mobType MobType, level int) LootTable {
	table := LootTable{
		Common:     generateCommonItems(mobType, level),
		Uncommon:   generateUncommonItems(mobType, level),
		Rare:       generateRareItems(mobType, level),
		Epic:       generateEpicItems(mobType, level),
		Legendary:  generateLegendaryItems(mobType, level),
		SteamCores: calculateSteamCores(mobType, level),
	}
	return table
}

// RollLoot generates loot from a loot table
func (lt *LootTable) RollLoot() []Item {
	var loot []Item

	// Add Steam Cores
	if lt.SteamCores > 0 {
		loot = append(loot, Item{
			ID:          "steam_core",
			Name:        "Steam Core",
			Type:        "currency",
			Rarity:      "common",
			Value:       1,
			Description: "A core of condensed steam power",
			Stackable:   true,
			Quantity:    lt.SteamCores,
		})
	}

	// Roll for items based on rarity
	if rand.Float32() < 0.7 { // 70% chance for common
		loot = append(loot, lt.Common[rand.Intn(len(lt.Common))])
	}
	if rand.Float32() < 0.4 { // 40% chance for uncommon
		loot = append(loot, lt.Uncommon[rand.Intn(len(lt.Uncommon))])
	}
	if rand.Float32() < 0.2 { // 20% chance for rare
		loot = append(loot, lt.Rare[rand.Intn(len(lt.Rare))])
	}
	if rand.Float32() < 0.05 { // 5% chance for epic
		loot = append(loot, lt.Epic[rand.Intn(len(lt.Epic))])
	}
	if rand.Float32() < 0.01 { // 1% chance for legendary
		loot = append(loot, lt.Legendary[rand.Intn(len(lt.Legendary))])
	}

	return loot
}

func calculateSteamCores(mobType MobType, level int) int {
	base := level * 2
	switch mobType {
	case Mechanical, Construct:
		base *= 2 // Mechanical and Construct mobs drop more Steam Cores
	}
	return base
}

func generateCommonItems(mobType MobType, level int) []Item {
	items := []Item{
		{
			ID:          "steam_crystal",
			Name:        "Steam Crystal",
			Type:        "material",
			Rarity:      "common",
			Level:       level,
			Value:       5,
			Description: "A crystal formed from condensed steam",
			Stackable:   true,
			Quantity:    1,
		},
		{
			ID:          "brass_gear",
			Name:        "Brass Gear",
			Type:        "material",
			Rarity:      "common",
			Level:       level,
			Value:       3,
			Description: "A well-crafted brass gear",
			Stackable:   true,
			Quantity:    1,
		},
	}
	return items
}

func generateUncommonItems(mobType MobType, level int) []Item {
	items := []Item{
		{
			ID:          "steam_essence",
			Name:        "Steam Essence",
			Type:        "material",
			Rarity:      "uncommon",
			Level:       level,
			Value:       15,
			Description: "Pure essence of steam power",
			Stackable:   true,
			Quantity:    1,
		},
		{
			ID:          "clockwork_spring",
			Name:        "Clockwork Spring",
			Type:        "material",
			Rarity:      "uncommon",
			Level:       level,
			Value:       12,
			Description: "A precision-crafted spring",
			Stackable:   true,
			Quantity:    1,
		},
	}
	return items
}

func generateRareItems(mobType MobType, level int) []Item {
	items := []Item{
		{
			ID:          "steam_core_fragment",
			Name:        "Steam Core Fragment",
			Type:        "material",
			Rarity:      "rare",
			Level:       level,
			Value:       50,
			Description: "A fragment of a powerful steam core",
			Stackable:   true,
			Quantity:    1,
		},
		{
			ID:          "enchanted_gear",
			Name:        "Enchanted Gear",
			Type:        "material",
			Rarity:      "rare",
			Level:       level,
			Value:       45,
			Description: "A gear infused with arcane energy",
			Stackable:   true,
			Quantity:    1,
		},
	}
	return items
}

func generateEpicItems(mobType MobType, level int) []Item {
	items := []Item{
		{
			ID:          "steam_core_shard",
			Name:        "Steam Core Shard",
			Type:        "material",
			Rarity:      "epic",
			Level:       level,
			Value:       150,
			Description: "A shard of a legendary steam core",
			Stackable:   true,
			Quantity:    1,
		},
		{
			ID:          "ancient_gear",
			Name:        "Ancient Gear",
			Type:        "material",
			Rarity:      "epic",
			Level:       level,
			Value:       120,
			Description: "A gear from an ancient machine",
			Stackable:   true,
			Quantity:    1,
		},
	}
	return items
}

func generateLegendaryItems(mobType MobType, level int) []Item {
	items := []Item{
		{
			ID:          "primal_steam_core",
			Name:        "Primal Steam Core",
			Type:        "material",
			Rarity:      "legendary",
			Level:       level,
			Value:       500,
			Description: "A core of pure, primal steam power",
			Stackable:   true,
			Quantity:    1,
		},
		{
			ID:          "clockwork_heart",
			Name:        "Clockwork Heart",
			Type:        "material",
			Rarity:      "legendary",
			Level:       level,
			Value:       450,
			Description: "The heart of a legendary machine",
			Stackable:   true,
			Quantity:    1,
		},
	}
	return items
}
