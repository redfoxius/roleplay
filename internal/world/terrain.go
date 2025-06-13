package world

// TerrainType represents different types of terrain
type TerrainType string

const (
	Forest    TerrainType = "forest"
	Mountain  TerrainType = "mountain"
	Water     TerrainType = "water"
	Plains    TerrainType = "plains"
	Desert    TerrainType = "desert"
	Swamp     TerrainType = "swamp"
	SteamCity TerrainType = "steam_city"
)

// TerrainProperties defines the characteristics of each terrain type
type TerrainProperties struct {
	Type            TerrainType
	Name            string
	Description     string
	MovementCost    int
	SteamPowerBonus int
	Visibility      int
	Resources       []string
	MobTypes        []string
}

// GetTerrainProperties returns the properties for a given terrain type
func GetTerrainProperties(terrainType TerrainType) TerrainProperties {
	switch terrainType {
	case Forest:
		return TerrainProperties{
			Type:            Forest,
			Name:            "Steam Forest",
			Description:     "A dense forest with steam-powered trees and mechanical wildlife",
			MovementCost:    2,
			SteamPowerBonus: 5,
			Visibility:      3,
			Resources:       []string{"wood", "steam_crystal", "mechanical_parts"},
			MobTypes:        []string{"Mechanical", "Biological", "Hybrid"},
		}
	case Mountain:
		return TerrainProperties{
			Type:            Mountain,
			Name:            "Steam Mountains",
			Description:     "Towering mountains with steam vents and mechanical caves",
			MovementCost:    3,
			SteamPowerBonus: 10,
			Visibility:      5,
			Resources:       []string{"ore", "steam_crystal", "mechanical_parts"},
			MobTypes:        []string{"Mechanical", "Construct", "Elemental"},
		}
	case Water:
		return TerrainProperties{
			Type:            Water,
			Name:            "Steam Lakes",
			Description:     "Boiling lakes with steam-powered currents and mechanical fish",
			MovementCost:    4,
			SteamPowerBonus: 15,
			Visibility:      2,
			Resources:       []string{"steam_crystal", "mechanical_parts", "steam_essence"},
			MobTypes:        []string{"Mechanical", "Elemental", "Hybrid"},
		}
	case Plains:
		return TerrainProperties{
			Type:            Plains,
			Name:            "Steam Plains",
			Description:     "Vast plains with steam-powered windmills and mechanical herds",
			MovementCost:    1,
			SteamPowerBonus: 0,
			Visibility:      8,
			Resources:       []string{"mechanical_parts", "steam_essence"},
			MobTypes:        []string{"Mechanical", "Biological", "Construct"},
		}
	case Desert:
		return TerrainProperties{
			Type:            Desert,
			Name:            "Steam Desert",
			Description:     "A scorching desert with steam-powered sandstorms and mechanical scorpions",
			MovementCost:    2,
			SteamPowerBonus: -5,
			Visibility:      6,
			Resources:       []string{"steam_crystal", "mechanical_parts"},
			MobTypes:        []string{"Mechanical", "Elemental", "Construct"},
		}
	case Swamp:
		return TerrainProperties{
			Type:            Swamp,
			Name:            "Steam Swamp",
			Description:     "A misty swamp with steam-powered plants and mechanical reptiles",
			MovementCost:    3,
			SteamPowerBonus: 8,
			Visibility:      2,
			Resources:       []string{"steam_essence", "mechanical_parts", "steam_crystal"},
			MobTypes:        []string{"Biological", "Hybrid", "Elemental"},
		}
	case SteamCity:
		return TerrainProperties{
			Type:            SteamCity,
			Name:            "Steam City",
			Description:     "A bustling city powered by steam technology and mechanical wonders",
			MovementCost:    1,
			SteamPowerBonus: 20,
			Visibility:      4,
			Resources:       []string{"mechanical_parts", "steam_crystal", "steam_essence"},
			MobTypes:        []string{"Mechanical", "Construct", "Hybrid"},
		}
	default:
		return TerrainProperties{
			Type:            Plains,
			Name:            "Unknown Terrain",
			Description:     "An unexplored area",
			MovementCost:    1,
			SteamPowerBonus: 0,
			Visibility:      5,
			Resources:       []string{},
			MobTypes:        []string{},
		}
	}
}
