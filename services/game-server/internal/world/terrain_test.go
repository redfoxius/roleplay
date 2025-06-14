package world

import (
	"testing"
)

func TestGetTerrainProperties(t *testing.T) {
	// Test terrain properties for different terrain types
	terrainTypes := []TerrainType{Plains, Forest, Swamp, Desert}

	for _, terrainType := range terrainTypes {
		properties := GetTerrainProperties(terrainType)

		// Check that properties are valid
		if properties.MovementCost <= 0 {
			t.Errorf("Expected positive movement cost for terrain type '%s', got %d", terrainType, properties.MovementCost)
		}

		// Check that steam power bonus is non-negative
		if properties.SteamPowerBonus < 0 {
			t.Errorf("Expected non-negative steam power bonus for terrain type '%s', got %d", terrainType, properties.SteamPowerBonus)
		}
	}
}
