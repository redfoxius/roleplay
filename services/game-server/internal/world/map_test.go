package world

import (
	"testing"

	"roleplay/internal/common"
)

func TestWorldMapGeneration(t *testing.T) {
	// Create a new world map
	wm := NewWorldMap(100, 100)

	// Test region generation
	if len(wm.Regions) == 0 {
		t.Error("Expected regions to be generated")
	}

	// Test location generation
	if len(wm.Locations) == 0 {
		t.Error("Expected locations to be generated")
	}

	// Test location types
	locationTypes := make(map[LocationType]int)
	for _, location := range wm.Locations {
		locationTypes[location.Type]++
	}

	// Verify that we have a mix of location types
	if locationTypes[LocationTypeNormal] == 0 {
		t.Error("Expected normal locations to be generated")
	}
	if locationTypes[LocationTypeDungeon] == 0 {
		t.Error("Expected dungeon locations to be generated")
	}
	if locationTypes[LocationTypeTown] == 0 {
		t.Error("Expected town locations to be generated")
	}
	if locationTypes[LocationTypeResource] == 0 {
		t.Error("Expected resource locations to be generated")
	}
	if locationTypes[LocationTypeRuin] == 0 {
		t.Error("Expected ruin locations to be generated")
	}
}

func TestLocationProperties(t *testing.T) {
	// Test properties for each location type
	testCases := []struct {
		locationType LocationType
		terrain      TerrainType
		expectedName string
	}{
		{LocationTypeDungeon, Mountain, "Steam Dungeon"},
		{LocationTypeTown, SteamCity, "Steam Town"},
		{LocationTypeResource, Forest, "Resource Node"},
		{LocationTypeRuin, Desert, "Ancient Ruins"},
		{LocationTypeNormal, Plains, "Normal Location"},
	}

	for _, tc := range testCases {
		props := GetLocationProperties(tc.locationType, tc.terrain)
		if props.Name != tc.expectedName {
			t.Errorf("Expected name %s for %s, got %s", tc.expectedName, tc.locationType, props.Name)
		}
		if props.Type != tc.locationType {
			t.Errorf("Expected type %s, got %s", tc.locationType, props.Type)
		}
		if len(props.Resources) == 0 {
			t.Errorf("Expected resources for %s", tc.locationType)
		}
		if len(props.MobTypes) == 0 {
			t.Errorf("Expected mob types for %s", tc.locationType)
		}
	}
}

func TestNearbyLocations(t *testing.T) {
	wm := NewWorldMap(100, 100)
	center := common.Coordinates{X: 50, Y: 50}
	distance := 10

	// Test GetNearbyLocations
	nearby := wm.GetNearbyLocations(center, distance)
	if len(nearby) == 0 {
		t.Error("Expected nearby locations")
	}

	// Test GetNearbyResources
	resources := wm.GetNearbyResources(center, distance)
	if len(resources) == 0 {
		t.Error("Expected nearby resources")
	}

	// Test GetNearbyTowns
	towns := wm.GetNearbyTowns(center, distance)
	if len(towns) == 0 {
		t.Error("Expected nearby towns")
	}

	// Test GetNearbyDungeons
	dungeons := wm.GetNearbyDungeons(center, distance)
	if len(dungeons) == 0 {
		t.Error("Expected nearby dungeons")
	}

	// Test GetNearbyRuins
	ruins := wm.GetNearbyRuins(center, distance)
	if len(ruins) == 0 {
		t.Error("Expected nearby ruins")
	}
}

func TestLocationPropertiesByCoordinates(t *testing.T) {
	wm := NewWorldMap(100, 100)
	coords := common.Coordinates{X: 50, Y: 50}

	// Get location properties
	props := wm.GetLocationProperties(coords)
	if props == nil {
		t.Error("Expected location properties")
	} else {
		// Verify properties
		if props.Name == "" {
			t.Error("Expected non-empty name")
		}
		if props.Description == "" {
			t.Error("Expected non-empty description")
		}
		if len(props.Resources) == 0 {
			t.Error("Expected resources")
		}
		if len(props.MobTypes) == 0 {
			t.Error("Expected mob types")
		}
	}

}

func TestTerrainGeneration(t *testing.T) {
	wm := NewWorldMap(100, 100)

	// Test terrain generation at different coordinates
	testCoords := []common.Coordinates{
		{X: 0, Y: 0},
		{X: 50, Y: 50},
		{X: 99, Y: 99},
	}

	for _, coords := range testCoords {
		location := wm.GetLocationAt(coords)
		if location == nil {
			t.Errorf("Expected location at coordinates %v", coords)
			continue
		}

		// Verify terrain type
		switch location.Terrain {
		case Forest, Mountain, Plains, Desert, Swamp, SteamCity:
			// Valid terrain type
		default:
			t.Errorf("Invalid terrain type %s at coordinates %v", location.Terrain, coords)
		}

		// Verify region
		if location.Region == "" {
			t.Errorf("Expected non-empty region at coordinates %v", coords)
		}
	}
}

func TestWorldMap(t *testing.T) {
	// Initialize a new world map
	worldMap := NewWorldMap(100, 100)

	// Test getting a location
	location := worldMap.GetLocationAt(common.Coordinates{X: 50, Y: 50})
	if location == nil {
		t.Errorf("Expected location at (50, 50), got nil")
	}

	// Test getting nearby locations
	nearbyLocations := worldMap.GetNearbyLocations(common.Coordinates{X: 50, Y: 50}, 10)
	if len(nearbyLocations) == 0 {
		t.Errorf("Expected nearby locations, got none")
	}

	// Test updating the world
	worldMap.UpdateWorld()
	// Add assertions here if needed to verify the update logic
}
