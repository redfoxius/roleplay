package character

import (
	"fmt"
	"math/rand"
	"time"
)

// SpawnPoint represents a location where mobs can spawn
type SpawnPoint struct {
	ID          string
	Location    string
	MobTypes    []MobType
	MinLevel    int
	MaxLevel    int
	SpawnRadius int
	RespawnTime int // in seconds
	LastSpawn   time.Time
	ActiveMobs  []*Mob
}

// SpawnManager manages mob spawning
type SpawnManager struct {
	spawnPoints map[string]*SpawnPoint
	mobs        map[string]*Mob
}

// NewSpawnManager creates a new spawn manager
func NewSpawnManager() *SpawnManager {
	return &SpawnManager{
		spawnPoints: make(map[string]*SpawnPoint),
		mobs:        make(map[string]*Mob),
	}
}

// AddSpawnPoint adds a new spawn point
func (sm *SpawnManager) AddSpawnPoint(sp *SpawnPoint) {
	sm.spawnPoints[sp.ID] = sp
}

// SpawnMob spawns a mob at the given spawn point
func (sm *SpawnManager) SpawnMob(spawnPointID string) (*Mob, error) {
	sp, exists := sm.spawnPoints[spawnPointID]
	if !exists {
		return nil, fmt.Errorf("spawn point %s not found", spawnPointID)
	}

	// Check if we can spawn (based on respawn time)
	if time.Since(sp.LastSpawn).Seconds() < float64(sp.RespawnTime) {
		return nil, fmt.Errorf("spawn point %s is on cooldown", spawnPointID)
	}

	// Select random mob type
	mobType := sp.MobTypes[rand.Intn(len(sp.MobTypes))]

	// Generate level within range
	level := sp.MinLevel
	if sp.MaxLevel > sp.MinLevel {
		level += rand.Intn(sp.MaxLevel - sp.MinLevel + 1)
	}

	// Create new mob
	mob := NewMob(generateMobName(mobType), mobType, level)
	mob.Abilities = GetMobAbilities(mobType, level)
	mob.LootTable = GenerateLootTable(mobType, level)

	// Update spawn point
	sp.LastSpawn = time.Now()
	sp.ActiveMobs = append(sp.ActiveMobs, mob)
	sm.mobs[mob.ID] = mob

	return mob, nil
}

// RemoveMob removes a mob from the spawn point
func (sm *SpawnManager) RemoveMob(mobID string) {
	_, exists := sm.mobs[mobID]
	if !exists {
		return
	}

	// Find and remove from spawn point
	for _, sp := range sm.spawnPoints {
		for i, m := range sp.ActiveMobs {
			if m.ID == mobID {
				sp.ActiveMobs = append(sp.ActiveMobs[:i], sp.ActiveMobs[i+1:]...)
				break
			}
		}
	}

	delete(sm.mobs, mobID)
}

// generateMobName generates a name for a mob based on its type
func generateMobName(mobType MobType) string {
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

// GetActiveMobs returns all currently active mobs
func (sm *SpawnManager) GetActiveMobs() []*Mob {
	mobs := make([]*Mob, 0, len(sm.mobs))
	for _, mob := range sm.mobs {
		mobs = append(mobs, mob)
	}
	return mobs
}

// GetSpawnPoint returns a spawn point by ID
func (sm *SpawnManager) GetSpawnPoint(id string) (*SpawnPoint, error) {
	sp, exists := sm.spawnPoints[id]
	if !exists {
		return nil, fmt.Errorf("spawn point %s not found", id)
	}
	return sp, nil
}
