package game

import (
	"fmt"
	"math"
	"sync"

	"roleplay/internal/character"
	"roleplay/internal/combat"
	"roleplay/internal/common"
	"roleplay/internal/database"
	"roleplay/internal/mob"
	"roleplay/internal/world"
)

// GameServer represents the game server
type GameServer struct {
	players     map[string]*character.Character
	activeGames map[string]*combat.CombatState
	mobs        map[string]*mob.Mob
	mutex       sync.RWMutex
	repo        *database.Repository
	worldMap    *world.WorldMap
	spawner     *world.WorldSpawner
}

// NewGameServer creates a new game server
func NewGameServer(repo *database.Repository) *GameServer {
	server := &GameServer{
		players:     make(map[string]*character.Character),
		activeGames: make(map[string]*combat.CombatState),
		mobs:        make(map[string]*mob.Mob),
		repo:        repo,
		worldMap:    world.NewWorldMap(100, 100), // Create a 100x100 world
		spawner:     world.NewWorldSpawner(),
	}

	// Initialize spawn points
	server.spawner.InitializeSpawnPoints(server.worldMap)

	// Load existing data from Redis
	server.loadData()
	return server
}

// GetWorldMap returns the world map
func (gs *GameServer) GetWorldMap() *world.WorldMap {
	return gs.worldMap
}

// GetLocationAt returns the location at the given coordinates
func (gs *GameServer) GetLocationAt(x, y int) *world.Location {
	return gs.worldMap.GetLocationAt(common.Coordinates{X: x, Y: y})
}

// GetNearbyLocations returns locations within a certain distance
func (gs *GameServer) GetNearbyLocations(x, y, distance int) []*world.Location {
	return gs.worldMap.GetNearbyLocations(common.Coordinates{X: x, Y: y}, distance)
}

// UpdateWorld updates the world state
func (gs *GameServer) UpdateWorld() {
	gs.worldMap.UpdateWorld()
	gs.spawner.UpdateSpawner()
}

// GetRegionAt returns the region at the given coordinates
func (gs *GameServer) GetRegionAt(x, y int) string {
	location := gs.GetLocationAt(x, y)
	if location != nil {
		return location.Region
	}
	return ""
}

// GetTerrainProperties returns the properties of the terrain at the given coordinates
func (gs *GameServer) GetTerrainProperties(x, y int) world.TerrainProperties {
	location := gs.GetLocationAt(x, y)
	if location != nil {
		return world.GetTerrainProperties(location.Terrain)
	}
	return world.GetTerrainProperties(world.Plains) // Default to plains if no location found
}

// loadData loads all data from Redis
func (gs *GameServer) loadData() error {
	// Load characters
	characters, err := gs.repo.GetAllCharacters()
	if err != nil {
		return fmt.Errorf("failed to load characters: %v", err)
	}
	for _, char := range characters {
		gs.players[char.ID] = char
	}

	// Load mobs
	mobs, err := gs.repo.GetAllMobs()
	if err != nil {
		return fmt.Errorf("failed to load mobs: %v", err)
	}
	for _, m := range mobs {
		gs.mobs[m.ID] = m
	}

	return nil
}

// CreateCharacter creates a new character
func (gs *GameServer) CreateCharacter(name string, class character.Class) (*character.Character, error) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()

	if _, exists := gs.players[name]; exists {
		return nil, fmt.Errorf("character with name %s already exists", name)
	}

	char := character.NewCharacter(name, class)
	char.ApplyClassBonuses()

	// Save to Redis
	if err := gs.repo.SaveCharacter(char); err != nil {
		return nil, fmt.Errorf("failed to save character: %v", err)
	}

	gs.players[name] = char
	return char, nil
}

// StartCombat initiates a new combat between players
func (gs *GameServer) StartCombat(participantNames []string) (*combat.CombatState, error) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()

	var participants []*character.Character
	for _, name := range participantNames {
		if char, exists := gs.players[name]; exists {
			participants = append(participants, char)
		} else {
			return nil, fmt.Errorf("player %s not found", name)
		}
	}

	if len(participants) < 2 {
		return nil, fmt.Errorf("need at least 2 participants for combat")
	}

	combatState := combat.NewCombatState(participants)

	// Save to Redis
	if err := gs.repo.SaveCombatState(combatState); err != nil {
		return nil, fmt.Errorf("failed to save combat state: %v", err)
	}

	gs.activeGames[combatState.ID] = combatState
	return combatState, nil
}

// StartMobCombat starts combat between a character and a mob
func (s *GameServer) StartMobCombat(characterID, mobID string) (*combat.MobCombat, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	char, exists := s.players[characterID]
	if !exists {
		return nil, fmt.Errorf("character not found")
	}

	m, exists := s.mobs[mobID]
	if !exists {
		return nil, fmt.Errorf("mob not found")
	}

	combat := combat.NewMobCombat(char, m)
	return combat, nil
}

// HandleMobCombatAction handles a combat action between a character and mob
func (s *GameServer) HandleMobCombatAction(combat *combat.MobCombat, ability *common.Ability) (int, int, error) {
	// Character attacks
	charDamage, err := combat.CharacterAttack(ability)
	if err != nil {
		return 0, 0, err
	}

	// Save character state
	if err := s.repo.SaveCharacter(combat.Character); err != nil {
		return 0, 0, fmt.Errorf("failed to save character state: %v", err)
	}

	// Check if combat is over after character attack
	if combat.IsCombatOver() {
		return charDamage, 0, nil
	}

	// Mob attacks
	mobDamage, err := combat.MobAttack()
	if err != nil {
		return charDamage, 0, err
	}

	// Save character state again after mob attack
	if err := s.repo.SaveCharacter(combat.Character); err != nil {
		return 0, 0, fmt.Errorf("failed to save character state: %v", err)
	}

	return charDamage, mobDamage, nil
}

// SaveGameState saves the current game state to Redis
func (gs *GameServer) SaveGameState() error {
	gs.mutex.RLock()
	defer gs.mutex.RUnlock()

	// Save all characters
	for _, char := range gs.players {
		if err := gs.repo.SaveCharacter(char); err != nil {
			return fmt.Errorf("failed to save character %s: %v", char.ID, err)
		}
	}

	// Save all mobs
	for _, mob := range gs.mobs {
		if err := gs.repo.SaveMob(mob); err != nil {
			return fmt.Errorf("failed to save mob %s: %v", mob.ID, err)
		}
	}

	// Save all active combat states
	for _, state := range gs.activeGames {
		if err := gs.repo.SaveCombatState(state); err != nil {
			return fmt.Errorf("failed to save combat state %s: %v", state.ID, err)
		}
	}

	return nil
}

// MoveCharacter moves a character to a new location
func (gs *GameServer) MoveCharacter(characterID string, x, y int) error {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()

	character, exists := gs.players[characterID]
	if !exists {
		return fmt.Errorf("character not found")
	}

	// Check if the target location is valid
	targetLocation := gs.GetLocationAt(x, y)
	if targetLocation == nil {
		return fmt.Errorf("invalid location")
	}

	// Get terrain properties
	terrainProps := gs.GetTerrainProperties(x, y)

	// Check if the character has enough steam power to move
	if character.SteamPower < terrainProps.MovementCost {
		return fmt.Errorf("insufficient steam power")
	}

	// Move the character
	character.MoveTo(x, y)
	character.SteamPower -= terrainProps.MovementCost

	// Apply terrain effects
	character.SteamPower += terrainProps.SteamPowerBonus
	if character.SteamPower > character.MaxSteamPower {
		character.SteamPower = character.MaxSteamPower
	}

	return nil
}

// GetCharacterLocation returns a character's current location
func (gs *GameServer) GetCharacterLocation(characterID string) (common.Coordinates, error) {
	gs.mutex.RLock()
	defer gs.mutex.RUnlock()

	character, exists := gs.players[characterID]
	if !exists {
		return common.Coordinates{}, fmt.Errorf("character not found")
	}

	return character.GetLocation(), nil
}

// GetNearbyCharacters returns all characters within a certain distance
func (gs *GameServer) GetNearbyCharacters(x, y, distance int) []*character.Character {
	gs.mutex.RLock()
	defer gs.mutex.RUnlock()

	var nearby []*character.Character
	for _, char := range gs.players {
		charLoc := char.GetLocation()
		if distanceBetween(common.Coordinates{X: x, Y: y}, charLoc) <= distance {
			nearby = append(nearby, char)
		}
	}
	return nearby
}

// distanceBetween calculates the distance between two coordinates
func distanceBetween(a, b common.Coordinates) int {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	return int(math.Sqrt(dx*dx + dy*dy))
}

// GetNearbyMobs returns all mobs within a certain distance
func (gs *GameServer) GetNearbyMobs(x, y, distance int) []*mob.Mob {
	return gs.spawner.GetMobsInArea(common.Coordinates{X: x, Y: y}, distance)
}

// GetAllMobs returns all currently active mobs
func (gs *GameServer) GetAllMobs() []*mob.Mob {
	return gs.spawner.GetAllMobs()
}

// RemoveMob removes a mob from the world
func (gs *GameServer) RemoveMob(mobID string) {
	gs.spawner.RemoveMob(mobID)
}
