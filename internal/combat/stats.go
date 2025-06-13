package combat

import (
	"time"
)

// CombatStats tracks statistics for a character or mob
type CombatStats struct {
	TotalDamageDealt     int
	TotalDamageTaken     int
	TotalHealingDone     int
	TotalHealingReceived int
	TotalSteamPowerUsed  int
	AbilitiesUsed        map[string]int
	CriticalHits         int
	Misses               int
	Kills                int
	Deaths               int
	LongestCombat        time.Duration
	TotalCombatTime      time.Duration
	CombatCount          int
	Victories            int
	Defeats              int
	LootCollected        map[string]int
	ExperienceGained     int
}

// NewCombatStats creates a new combat statistics tracker
func NewCombatStats() *CombatStats {
	return &CombatStats{
		AbilitiesUsed: make(map[string]int),
		LootCollected: make(map[string]int),
	}
}

// RecordDamageDealt records damage dealt
func (cs *CombatStats) RecordDamageDealt(damage int) {
	cs.TotalDamageDealt += damage
}

// RecordDamageTaken records damage taken
func (cs *CombatStats) RecordDamageTaken(damage int) {
	cs.TotalDamageTaken += damage
}

// RecordHealingDone records healing done
func (cs *CombatStats) RecordHealingDone(healing int) {
	cs.TotalHealingDone += healing
}

// RecordHealingReceived records healing received
func (cs *CombatStats) RecordHealingReceived(healing int) {
	cs.TotalHealingReceived += healing
}

// RecordSteamPowerUsed records steam power usage
func (cs *CombatStats) RecordSteamPowerUsed(amount int) {
	cs.TotalSteamPowerUsed += amount
}

// RecordAbilityUse records the use of an ability
func (cs *CombatStats) RecordAbilityUse(abilityName string) {
	cs.AbilitiesUsed[abilityName]++
}

// RecordCriticalHit records a critical hit
func (cs *CombatStats) RecordCriticalHit() {
	cs.CriticalHits++
}

// RecordMiss records a missed attack
func (cs *CombatStats) RecordMiss() {
	cs.Misses++
}

// RecordKill records a kill
func (cs *CombatStats) RecordKill() {
	cs.Kills++
}

// RecordDeath records a death
func (cs *CombatStats) RecordDeath() {
	cs.Deaths++
}

// RecordCombatDuration records the duration of a combat
func (cs *CombatStats) RecordCombatDuration(duration time.Duration) {
	cs.TotalCombatTime += duration
	cs.CombatCount++
	if duration > cs.LongestCombat {
		cs.LongestCombat = duration
	}
}

// RecordVictory records a victory
func (cs *CombatStats) RecordVictory() {
	cs.Victories++
}

// RecordDefeat records a defeat
func (cs *CombatStats) RecordDefeat() {
	cs.Defeats++
}

// RecordLoot records collected loot
func (cs *CombatStats) RecordLoot(itemName string, quantity int) {
	cs.LootCollected[itemName] += quantity
}

// RecordExperience records gained experience
func (cs *CombatStats) RecordExperience(amount int) {
	cs.ExperienceGained += amount
}

// GetAverageDamageDealt returns the average damage dealt per combat
func (cs *CombatStats) GetAverageDamageDealt() float64 {
	if cs.CombatCount == 0 {
		return 0
	}
	return float64(cs.TotalDamageDealt) / float64(cs.CombatCount)
}

// GetAverageDamageTaken returns the average damage taken per combat
func (cs *CombatStats) GetAverageDamageTaken() float64 {
	if cs.CombatCount == 0 {
		return 0
	}
	return float64(cs.TotalDamageTaken) / float64(cs.CombatCount)
}

// GetAverageCombatDuration returns the average combat duration
func (cs *CombatStats) GetAverageCombatDuration() time.Duration {
	if cs.CombatCount == 0 {
		return 0
	}
	return cs.TotalCombatTime / time.Duration(cs.CombatCount)
}

// GetWinRate returns the win rate as a percentage
func (cs *CombatStats) GetWinRate() float64 {
	total := cs.Victories + cs.Defeats
	if total == 0 {
		return 0
	}
	return float64(cs.Victories) / float64(total) * 100
}

// GetMostUsedAbility returns the most frequently used ability
func (cs *CombatStats) GetMostUsedAbility() (string, int) {
	var mostUsed string
	var maxUses int
	for ability, uses := range cs.AbilitiesUsed {
		if uses > maxUses {
			mostUsed = ability
			maxUses = uses
		}
	}
	return mostUsed, maxUses
}

// GetMostCollectedLoot returns the most frequently collected loot
func (cs *CombatStats) GetMostCollectedLoot() (string, int) {
	var mostCollected string
	var maxQuantity int
	for item, quantity := range cs.LootCollected {
		if quantity > maxQuantity {
			mostCollected = item
			maxQuantity = quantity
		}
	}
	return mostCollected, maxQuantity
}
