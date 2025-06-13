package world

import (
	"math"
	"math/rand"
	"time"

	"roleplay/internal/common"
)

// Location represents a point of interest in the world
type Location struct {
	Name        string
	Description string
	Terrain     TerrainType
	Region      string
	Type        LocationType // New field for special locations
}

// LocationType represents different types of special locations
type LocationType string

const (
	LocationTypeNormal   LocationType = "normal"
	LocationTypeDungeon  LocationType = "dungeon"
	LocationTypeTown     LocationType = "town"
	LocationTypeResource LocationType = "resource"
	LocationTypeRuin     LocationType = "ruin"
)

// Region represents a larger area of the world
type Region struct {
	Name        string
	Description string
	Terrain     TerrainType
	Locations   []Location
	Level       int // Difficulty level of the region
}

// WorldMap represents the game world
type WorldMap struct {
	Width     int
	Height    int
	Regions   map[string]*Region
	Locations map[common.Coordinates]*Location
	Seed      int64 // Seed for consistent generation
}

// NewWorldMap creates a new world map
func NewWorldMap(width, height int) *WorldMap {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	wm := &WorldMap{
		Width:     width,
		Height:    height,
		Regions:   make(map[string]*Region),
		Locations: make(map[common.Coordinates]*Location),
		Seed:      seed,
	}
	wm.generateWorld()
	return wm
}

// generateWorld generates the world map
func (wm *WorldMap) generateWorld() {
	wm.generateRegions()
	wm.generateLocations()
}

// generateRegions generates the regions of the world
func (wm *WorldMap) generateRegions() {
	// Create predefined regions with difficulty levels
	regions := []*Region{
		{
			Name:        "Steam City",
			Description: "A bustling metropolis powered by steam technology",
			Terrain:     SteamCity,
			Level:       1,
		},
		{
			Name:        "Mountain Range",
			Description: "A treacherous mountain range with valuable resources",
			Terrain:     Mountain,
			Level:       3,
		},
		{
			Name:        "Ancient Forest",
			Description: "A mysterious forest with ancient technology",
			Terrain:     Forest,
			Level:       2,
		},
		{
			Name:        "Steam Desert",
			Description: "A scorching desert with steam-powered sandstorms",
			Terrain:     Desert,
			Level:       4,
		},
		{
			Name:        "Steam Swamp",
			Description: "A misty swamp with steam-powered research facilities",
			Terrain:     Swamp,
			Level:       3,
		},
		{
			Name:        "Steam Plains",
			Description: "Vast plains dotted with steam-powered windmills",
			Terrain:     Plains,
			Level:       1,
		},
	}

	// Add regions to the world map
	for _, region := range regions {
		wm.Regions[region.Name] = region
	}
}

// generateLocations generates locations within regions
func (wm *WorldMap) generateLocations() {
	// Generate random locations
	for i := 0; i < 100; i++ {
		x := rand.Intn(wm.Width)
		y := rand.Intn(wm.Height)

		// Determine terrain type using Perlin noise
		terrain := wm.getTerrainAt(x, y)

		// Determine location type
		locationType := wm.determineLocationType(terrain)

		// Get location properties
		props := GetLocationProperties(locationType, terrain)

		// Create location
		location := &Location{
			Name:        props.Name,
			Description: props.Description,
			Terrain:     terrain,
			Region:      wm.getRegionAt(x, y),
			Type:        locationType,
		}

		// Add location to the world map
		wm.Locations[common.Coordinates{X: x, Y: y}] = location
	}
}

// getTerrainAt returns the terrain type at the given coordinates using Perlin noise
func (wm *WorldMap) getTerrainAt(x, y int) TerrainType {
	// Use Perlin noise for more natural terrain generation
	noise := perlinNoise(float64(x)/50.0, float64(y)/50.0, wm.Seed)

	switch {
	case noise < 0.2:
		return Forest
	case noise < 0.4:
		return Mountain
	case noise < 0.6:
		return Plains
	case noise < 0.8:
		return Desert
	default:
		return Swamp
	}
}

// determineLocationType determines the type of location based on terrain and random chance
func (wm *WorldMap) determineLocationType(terrain TerrainType) LocationType {
	chance := rand.Float64()

	switch terrain {
	case SteamCity:
		if chance < 0.3 {
			return LocationTypeTown
		}
	case Mountain:
		if chance < 0.4 {
			return LocationTypeDungeon
		} else if chance < 0.6 {
			return LocationTypeResource
		}
	case Forest:
		if chance < 0.3 {
			return LocationTypeRuin
		} else if chance < 0.5 {
			return LocationTypeResource
		}
	case Desert:
		if chance < 0.4 {
			return LocationTypeRuin
		} else if chance < 0.6 {
			return LocationTypeResource
		}
	case Swamp:
		if chance < 0.3 {
			return LocationTypeDungeon
		} else if chance < 0.5 {
			return LocationTypeResource
		}
	}

	return LocationTypeNormal
}

// perlinNoise generates Perlin noise for terrain generation
func perlinNoise(x, y float64, seed int64) float64 {
	// Simple implementation of Perlin noise
	// In a real implementation, you would use a proper Perlin noise library
	rand.Seed(seed + int64(x*1000) + int64(y*1000))
	return rand.Float64()
}

// getRegionAt returns the region at the given coordinates
func (wm *WorldMap) getRegionAt(x, y int) string {
	// Simple region assignment based on coordinates
	if x < wm.Width/3 {
		return "Steam City"
	} else if x < 2*wm.Width/3 {
		return "Mountain Range"
	} else {
		return "Ancient Forest"
	}
}

// generateLocationName generates a name for a location based on its terrain
func generateLocationName(terrain TerrainType) string {
	prefixes := map[TerrainType][]string{
		Forest:    {"Ancient", "Mysterious", "Enchanted", "Hidden", "Sacred"},
		Mountain:  {"Steep", "Towering", "Majestic", "Forbidden", "Lost"},
		Plains:    {"Vast", "Endless", "Golden", "Peaceful", "Serene"},
		Desert:    {"Burning", "Barren", "Forgotten", "Ancient", "Mysterious"},
		Swamp:     {"Misty", "Dark", "Haunted", "Forgotten", "Ancient"},
		SteamCity: {"Steam", "Clockwork", "Brass", "Copper", "Iron"},
	}

	suffixes := map[TerrainType][]string{
		Forest:    {"Grove", "Woods", "Forest", "Garden", "Sanctuary"},
		Mountain:  {"Peak", "Range", "Summit", "Cliff", "Ridge"},
		Plains:    {"Plains", "Meadow", "Field", "Valley", "Plateau"},
		Desert:    {"Desert", "Wastes", "Dunes", "Oasis", "Ruins"},
		Swamp:     {"Swamp", "Marsh", "Bog", "Fen", "Mire"},
		SteamCity: {"City", "District", "Quarter", "Square", "Hub"},
	}

	prefix := prefixes[terrain][rand.Intn(len(prefixes[terrain]))]
	suffix := suffixes[terrain][rand.Intn(len(suffixes[terrain]))]

	return prefix + " " + suffix
}

// generateLocationDescription generates a description for a location based on its terrain
func generateLocationDescription(terrain TerrainType) string {
	descriptions := map[TerrainType][]string{
		Forest: {
			"An ancient forest with towering trees and mysterious ruins.",
			"A dense forest filled with strange mechanical creatures.",
			"A sacred grove where nature and technology coexist.",
		},
		Mountain: {
			"A treacherous mountain range with valuable steam-powered mines.",
			"Majestic peaks hiding ancient steam-powered temples.",
			"A forbidden mountain range with dangerous mechanical guardians.",
		},
		Plains: {
			"Vast plains dotted with steam-powered windmills.",
			"Golden fields where steam-powered harvesters work tirelessly.",
			"Peaceful meadows with mechanical wildlife.",
		},
		Desert: {
			"A burning desert with ancient steam-powered ruins.",
			"Barren wasteland hiding valuable steam technology.",
			"Forgotten desert with mysterious mechanical artifacts.",
		},
		Swamp: {
			"A misty swamp with steam-powered research facilities.",
			"Dark marshland with dangerous mechanical experiments.",
			"Haunted bog with strange steam-powered creatures.",
		},
		SteamCity: {
			"A bustling metropolis powered by steam technology.",
			"A city of brass and copper, where steam powers everything.",
			"A technological marvel where steam and magic coexist.",
		},
	}

	return descriptions[terrain][rand.Intn(len(descriptions[terrain]))]
}

// GetLocationAt returns the location at the given coordinates
func (wm *WorldMap) GetLocationAt(coords common.Coordinates) *Location {
	return wm.Locations[coords]
}

// GetNearbyLocations returns all locations within a certain distance
func (wm *WorldMap) GetNearbyLocations(coords common.Coordinates, distance int) []*Location {
	var nearby []*Location
	for locCoords, location := range wm.Locations {
		if distanceBetween(coords, locCoords) <= distance {
			nearby = append(nearby, location)
		}
	}
	return nearby
}

// GetNearbyLocationsByType returns all locations of a specific type within a certain distance
func (wm *WorldMap) GetNearbyLocationsByType(coords common.Coordinates, distance int, locationType LocationType) []*Location {
	var nearby []*Location
	for locCoords, location := range wm.Locations {
		if distanceBetween(coords, locCoords) <= distance && location.Type == locationType {
			nearby = append(nearby, location)
		}
	}
	return nearby
}

// GetNearbyResources returns all resource locations within a certain distance
func (wm *WorldMap) GetNearbyResources(coords common.Coordinates, distance int) []*Location {
	return wm.GetNearbyLocationsByType(coords, distance, LocationTypeResource)
}

// GetNearbyTowns returns all town locations within a certain distance
func (wm *WorldMap) GetNearbyTowns(coords common.Coordinates, distance int) []*Location {
	return wm.GetNearbyLocationsByType(coords, distance, LocationTypeTown)
}

// GetNearbyDungeons returns all dungeon locations within a certain distance
func (wm *WorldMap) GetNearbyDungeons(coords common.Coordinates, distance int) []*Location {
	return wm.GetNearbyLocationsByType(coords, distance, LocationTypeDungeon)
}

// GetNearbyRuins returns all ruin locations within a certain distance
func (wm *WorldMap) GetNearbyRuins(coords common.Coordinates, distance int) []*Location {
	return wm.GetNearbyLocationsByType(coords, distance, LocationTypeRuin)
}

// GetLocationProperties returns the properties of a location
func (wm *WorldMap) GetLocationProperties(coords common.Coordinates) *LocationProperties {
	location := wm.GetLocationAt(coords)
	if location == nil {
		return nil
	}

	props := GetLocationProperties(location.Type, location.Terrain)
	return &props
}

// UpdateWorld updates the world state
func (wm *WorldMap) UpdateWorld() {
	// Update terrain effects
	for _, location := range wm.Locations {
		switch location.Terrain {
		case Forest:
			// Forest slowly regenerates steam power
			// TODO: Implement steam power regeneration
		case Mountain:
			// Mountains have strong winds that reduce steam power
			// TODO: Implement steam power reduction
		}
	}
}

// distanceBetween calculates the distance between two coordinates
func distanceBetween(a, b common.Coordinates) int {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	return int(math.Sqrt(dx*dx + dy*dy))
}
