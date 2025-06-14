package world

import (
	"math"
	"math/rand"
	"time"

	"roleplay/internal/common"
	"roleplay/internal/mob"
)

// WorldSpawner manages mob spawning in the world
type WorldSpawner struct {
	spawnPoints map[common.Coordinates]*SpawnPoint
	mobs        map[string]*mob.Mob
}

// SpawnPoint represents a location where mobs can spawn
type SpawnPoint struct {
	Location    common.Coordinates
	Terrain     TerrainType
	MobTypes    []mob.MobType
	MinLevel    int
	MaxLevel    int
	SpawnRadius int
	RespawnTime int // in seconds
	LastSpawn   time.Time
	ActiveMobs  []*mob.Mob
}

// NewWorldSpawner creates a new world spawner
func NewWorldSpawner() *WorldSpawner {
	rand.Seed(time.Now().UnixNano())
	return &WorldSpawner{
		spawnPoints: make(map[common.Coordinates]*SpawnPoint),
		mobs:        make(map[string]*mob.Mob),
	}
}

// InitializeSpawnPoints initializes spawn points based on terrain
func (ws *WorldSpawner) InitializeSpawnPoints(wm *WorldMap) {
	for coords, location := range wm.Locations {
		// Get terrain properties
		props := GetTerrainProperties(location.Terrain)

		// Create spawn point if terrain has mob types
		if len(props.MobTypes) > 0 {
			sp := &SpawnPoint{
				Location:    coords,
				Terrain:     location.Terrain,
				MobTypes:    make([]mob.MobType, 0),
				MinLevel:    1,
				MaxLevel:    10,
				SpawnRadius: 5,
				RespawnTime: 300, // 5 minutes
				LastSpawn:   time.Now(),
				ActiveMobs:  make([]*mob.Mob, 0),
			}

			// Convert string mob types to MobType
			for _, mobType := range props.MobTypes {
				sp.MobTypes = append(sp.MobTypes, mob.MobType(mobType))
			}

			ws.spawnPoints[coords] = sp
		}
	}
}

// UpdateSpawner updates the spawner state
func (ws *WorldSpawner) UpdateSpawner() {
	now := time.Now()
	for _, sp := range ws.spawnPoints {
		// Check if it's time to spawn new mobs
		if now.Sub(sp.LastSpawn).Seconds() >= float64(sp.RespawnTime) {
			// Remove dead mobs
			activeMobs := make([]*mob.Mob, 0)
			for _, mob := range sp.ActiveMobs {
				if mob.Health > 0 {
					activeMobs = append(activeMobs, mob)
				} else {
					delete(ws.mobs, mob.ID)
				}
			}
			sp.ActiveMobs = activeMobs

			// Spawn new mobs if needed
			if len(sp.ActiveMobs) < 3 { // Max 3 mobs per spawn point
				if mob, err := ws.SpawnMob(sp); err == nil {
					sp.ActiveMobs = append(sp.ActiveMobs, mob)
					ws.mobs[mob.ID] = mob
				}
			}

			sp.LastSpawn = now
		}
	}
}

// SpawnMob spawns a mob at the given spawn point
func (ws *WorldSpawner) SpawnMob(sp *SpawnPoint) (*mob.Mob, error) {
	// Select random mob type
	mobType := sp.MobTypes[rand.Intn(len(sp.MobTypes))]

	// Generate level within range
	level := sp.MinLevel
	if sp.MaxLevel > sp.MinLevel {
		level += rand.Intn(sp.MaxLevel - sp.MinLevel + 1)
	}

	// Generate random position within spawn radius
	angle := rand.Float64() * 2 * math.Pi
	distance := rand.Float64() * float64(sp.SpawnRadius)
	spawnX := sp.Location.X + int(math.Cos(angle)*distance)
	spawnY := sp.Location.Y + int(math.Sin(angle)*distance)

	// Create new mob
	newMob := mob.NewMob(mob.GenerateMobName(mobType), mobType, level)
	newMob.Abilities = mob.GetMobAbilities(mobType, level)
	newMob.LootTable = mob.GenerateLootTable(mobType, level)

	// Set mob location
	newMob.Location.X = spawnX
	newMob.Location.Y = spawnY

	return newMob, nil
}

// GetMobsInArea returns all mobs within a certain distance of coordinates
func (ws *WorldSpawner) GetMobsInArea(coords common.Coordinates, distance int) []*mob.Mob {
	var nearby []*mob.Mob
	for _, mob := range ws.mobs {
		mobLoc := common.Coordinates{X: mob.Location.X, Y: mob.Location.Y}
		if distanceBetween(coords, mobLoc) <= distance {
			nearby = append(nearby, mob)
		}
	}
	return nearby
}

// RemoveMob removes a mob from the spawner
func (ws *WorldSpawner) RemoveMob(mobID string) {
	if _, exists := ws.mobs[mobID]; !exists {
		return
	}

	// Find and remove from spawn point
	for _, sp := range ws.spawnPoints {
		for i, m := range sp.ActiveMobs {
			if m.ID == mobID {
				sp.ActiveMobs = append(sp.ActiveMobs[:i], sp.ActiveMobs[i+1:]...)
				break
			}
		}
	}

	delete(ws.mobs, mobID)
}

// GetAllMobs returns all currently active mobs
func (ws *WorldSpawner) GetAllMobs() []*mob.Mob {
	mobs := make([]*mob.Mob, 0, len(ws.mobs))
	for _, mob := range ws.mobs {
		mobs = append(mobs, mob)
	}
	return mobs
}
