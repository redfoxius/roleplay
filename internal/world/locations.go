package world

import (
	"math/rand"
)

// LocationProperties defines the characteristics of each location type
type LocationProperties struct {
	Type            LocationType
	Name            string
	Description     string
	Difficulty      int
	SteamPowerBonus int
	Resources       []string
	MobTypes        []string
	SpecialFeatures []string
}

// GetLocationProperties returns the properties for a given location type
func GetLocationProperties(locationType LocationType, terrain TerrainType) LocationProperties {
	baseProps := getBaseProperties(locationType)
	terrainProps := GetTerrainProperties(terrain)

	// Combine base properties with terrain properties
	props := LocationProperties{
		Type:            locationType,
		Name:            baseProps.Name,
		Description:     baseProps.Description,
		Difficulty:      baseProps.Difficulty + terrainProps.MovementCost,
		SteamPowerBonus: baseProps.SteamPowerBonus + terrainProps.SteamPowerBonus,
		Resources:       append(baseProps.Resources, terrainProps.Resources...),
		MobTypes:        append(baseProps.MobTypes, terrainProps.MobTypes...),
		SpecialFeatures: baseProps.SpecialFeatures,
	}

	return props
}

// getBaseProperties returns the base properties for a location type
func getBaseProperties(locationType LocationType) LocationProperties {
	switch locationType {
	case LocationTypeDungeon:
		return LocationProperties{
			Type:            LocationTypeDungeon,
			Name:            "Steam Dungeon",
			Description:     "An ancient dungeon filled with steam-powered traps and guardians",
			Difficulty:      5,
			SteamPowerBonus: 15,
			Resources:       []string{"steam_crystal", "mechanical_parts", "ancient_artifacts"},
			MobTypes:        []string{"Mechanical", "Construct", "Elemental"},
			SpecialFeatures: []string{"traps", "treasure_rooms", "boss_room"},
		}
	case LocationTypeTown:
		return LocationProperties{
			Type:            LocationTypeTown,
			Name:            "Steam Town",
			Description:     "A bustling town powered by steam technology",
			Difficulty:      1,
			SteamPowerBonus: 20,
			Resources:       []string{"mechanical_parts", "steam_essence", "trade_goods"},
			MobTypes:        []string{"Mechanical", "Biological"},
			SpecialFeatures: []string{"shops", "inn", "workshop"},
		}
	case LocationTypeResource:
		return LocationProperties{
			Type:            LocationTypeResource,
			Name:            "Resource Node",
			Description:     "A rich source of valuable resources",
			Difficulty:      2,
			SteamPowerBonus: 10,
			Resources:       []string{"steam_crystal", "mechanical_parts", "raw_materials"},
			MobTypes:        []string{"Mechanical", "Biological"},
			SpecialFeatures: []string{"gathering_point", "defense_turrets"},
		}
	case LocationTypeRuin:
		return LocationProperties{
			Type:            LocationTypeRuin,
			Name:            "Ancient Ruins",
			Description:     "The remains of an ancient steam-powered civilization",
			Difficulty:      3,
			SteamPowerBonus: 5,
			Resources:       []string{"ancient_artifacts", "steam_crystal", "mechanical_parts"},
			MobTypes:        []string{"Mechanical", "Construct", "Elemental"},
			SpecialFeatures: []string{"ancient_machinery", "hidden_chambers"},
		}
	default:
		return LocationProperties{
			Type:            LocationTypeNormal,
			Name:            "Normal Location",
			Description:     "A typical location in the world",
			Difficulty:      1,
			SteamPowerBonus: 0,
			Resources:       []string{},
			MobTypes:        []string{},
			SpecialFeatures: []string{},
		}
	}
}

// GenerateLocationName generates a name for a special location
func GenerateLocationName(locationType LocationType, terrain TerrainType) string {
	prefixes := map[LocationType][]string{
		LocationTypeDungeon:  {"Ancient", "Forgotten", "Cursed", "Mysterious", "Hidden"},
		LocationTypeTown:     {"Steam", "Clockwork", "Brass", "Copper", "Iron"},
		LocationTypeResource: {"Rich", "Valuable", "Abundant", "Precious", "Rare"},
		LocationTypeRuin:     {"Ancient", "Forgotten", "Lost", "Mysterious", "Sacred"},
	}

	suffixes := map[LocationType][]string{
		LocationTypeDungeon:  {"Dungeon", "Lair", "Vault", "Tomb", "Crypt"},
		LocationTypeTown:     {"Town", "Village", "Settlement", "Outpost", "Hub"},
		LocationTypeResource: {"Mine", "Quarry", "Deposit", "Vein", "Source"},
		LocationTypeRuin:     {"Ruins", "Remains", "Relics", "Vestiges", "Remnants"},
	}

	prefix := prefixes[locationType][rand.Intn(len(prefixes[locationType]))]
	suffix := suffixes[locationType][rand.Intn(len(suffixes[locationType]))]

	return prefix + " " + suffix
}

// GenerateLocationDescription generates a description for a special location
func GenerateLocationDescription(locationType LocationType, terrain TerrainType) string {
	descriptions := map[LocationType][]string{
		LocationTypeDungeon: {
			"An ancient dungeon filled with steam-powered traps and guardians.",
			"A mysterious underground complex with dangerous mechanical defenses.",
			"A forgotten vault containing powerful steam technology.",
		},
		LocationTypeTown: {
			"A bustling town powered by steam technology.",
			"A peaceful settlement where steam powers everyday life.",
			"A trading hub where steam technology is bought and sold.",
		},
		LocationTypeResource: {
			"A rich source of valuable steam-powered resources.",
			"An abundant deposit of mechanical parts and steam crystals.",
			"A precious vein of steam-powered materials.",
		},
		LocationTypeRuin: {
			"The remains of an ancient steam-powered civilization.",
			"Forgotten ruins containing mysterious steam technology.",
			"Ancient remnants of a once-great steam-powered society.",
		},
	}

	return descriptions[locationType][rand.Intn(len(descriptions[locationType]))]
}
