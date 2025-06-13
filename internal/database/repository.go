package database

import (
	"fmt"
	"time"

	"roleplay/internal/character"
	"roleplay/internal/combat"
	"roleplay/internal/mob"
)

const (
	// Key prefixes
	characterPrefix = "character:"
	mobPrefix       = "mob:"
	combatPrefix    = "combat:"
	spawnPrefix     = "spawn:"
)

// Repository handles data persistence for the game
type Repository struct {
	db *RedisDB
}

// NewRepository creates a new repository instance
func NewRepository(db *RedisDB) *Repository {
	return &Repository{db: db}
}

// SaveCharacter saves a character to Redis
func (r *Repository) SaveCharacter(char *character.Character) error {
	key := characterPrefix + char.ID
	return r.db.Set(key, char, 24*time.Hour)
}

// GetCharacter retrieves a character from Redis
func (r *Repository) GetCharacter(id string) (*character.Character, error) {
	var char character.Character
	key := characterPrefix + id
	if err := r.db.Get(key, &char); err != nil {
		return nil, fmt.Errorf("failed to get character: %v", err)
	}
	return &char, nil
}

// DeleteCharacter removes a character from Redis
func (r *Repository) DeleteCharacter(id string) error {
	key := characterPrefix + id
	return r.db.Delete(key)
}

// SaveMob saves a mob to Redis
func (r *Repository) SaveMob(m *mob.Mob) error {
	key := mobPrefix + m.ID
	return r.db.Set(key, m, 1*time.Hour)
}

// GetMob retrieves a mob from Redis
func (r *Repository) GetMob(id string) (*mob.Mob, error) {
	var m mob.Mob
	key := mobPrefix + id
	if err := r.db.Get(key, &m); err != nil {
		return nil, fmt.Errorf("failed to get mob: %v", err)
	}
	return &m, nil
}

// DeleteMob removes a mob from Redis
func (r *Repository) DeleteMob(id string) error {
	key := mobPrefix + id
	return r.db.Delete(key)
}

// SaveCombatState saves a combat state to Redis
func (r *Repository) SaveCombatState(state *combat.CombatState) error {
	key := combatPrefix + state.ID
	return r.db.Set(key, state, 1*time.Hour)
}

// GetCombatState retrieves a combat state from Redis
func (r *Repository) GetCombatState(id string) (*combat.CombatState, error) {
	var state combat.CombatState
	key := combatPrefix + id
	if err := r.db.Get(key, &state); err != nil {
		return nil, fmt.Errorf("failed to get combat state: %v", err)
	}
	return &state, nil
}

// DeleteCombatState removes a combat state from Redis
func (r *Repository) DeleteCombatState(id string) error {
	key := combatPrefix + id
	return r.db.Delete(key)
}

// SaveSpawnPoint saves a spawn point to Redis
func (r *Repository) SaveSpawnPoint(sp *character.SpawnPoint) error {
	key := spawnPrefix + sp.ID
	return r.db.Set(key, sp, 0) // No expiration for spawn points
}

// GetSpawnPoint retrieves a spawn point from Redis
func (r *Repository) GetSpawnPoint(id string) (*character.SpawnPoint, error) {
	var sp character.SpawnPoint
	key := spawnPrefix + id
	if err := r.db.Get(key, &sp); err != nil {
		return nil, fmt.Errorf("failed to get spawn point: %v", err)
	}
	return &sp, nil
}

// DeleteSpawnPoint removes a spawn point from Redis
func (r *Repository) DeleteSpawnPoint(id string) error {
	key := spawnPrefix + id
	return r.db.Delete(key)
}

// GetAllCharacters retrieves all characters from Redis
func (r *Repository) GetAllCharacters() ([]*character.Character, error) {
	keys, err := r.db.Keys(characterPrefix + "*")
	if err != nil {
		return nil, fmt.Errorf("failed to get character keys: %v", err)
	}

	characters := make([]*character.Character, 0, len(keys))
	for _, key := range keys {
		var char character.Character
		if err := r.db.Get(key, &char); err != nil {
			return nil, fmt.Errorf("failed to get character %s: %v", key, err)
		}
		characters = append(characters, &char)
	}

	return characters, nil
}

// GetAllMobs retrieves all mobs from Redis
func (r *Repository) GetAllMobs() ([]*mob.Mob, error) {
	keys, err := r.db.Keys(mobPrefix + "*")
	if err != nil {
		return nil, fmt.Errorf("failed to get mob keys: %v", err)
	}

	mobs := make([]*mob.Mob, 0, len(keys))
	for _, key := range keys {
		var m mob.Mob
		if err := r.db.Get(key, &m); err != nil {
			return nil, fmt.Errorf("failed to get mob %s: %v", key, err)
		}
		mobs = append(mobs, &m)
	}

	return mobs, nil
}

// GetAllSpawnPoints retrieves all spawn points from Redis
func (r *Repository) GetAllSpawnPoints() ([]*character.SpawnPoint, error) {
	keys, err := r.db.Keys(spawnPrefix + "*")
	if err != nil {
		return nil, fmt.Errorf("failed to get spawn point keys: %v", err)
	}

	spawnPoints := make([]*character.SpawnPoint, 0, len(keys))
	for _, key := range keys {
		var sp character.SpawnPoint
		if err := r.db.Get(key, &sp); err != nil {
			return nil, fmt.Errorf("failed to get spawn point %s: %v", key, err)
		}
		spawnPoints = append(spawnPoints, &sp)
	}

	return spawnPoints, nil
}
