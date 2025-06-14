package character

// CombatStats tracks combat-related statistics
type CombatStats struct {
	TotalDamageDealt     int
	TotalDamageTaken     int
	TotalHealingDone     int
	TotalHealingReceived int
	CriticalHits         int
	CriticalMisses       int
	AbilitiesUsed        map[string]int
	CombatCount          int
	Victories            int
	Defeats              int
	TotalCombatTime      int64 // in milliseconds
	LootCollected        map[string]int
}

// NewCombatStats creates a new CombatStats instance
func NewCombatStats() *CombatStats {
	return &CombatStats{
		AbilitiesUsed: make(map[string]int),
		LootCollected: make(map[string]int),
	}
}

// RecordDamageDealt records damage dealt
func (s *CombatStats) RecordDamageDealt(amount int) {
	s.TotalDamageDealt += amount
}

// RecordDamageTaken records damage taken
func (s *CombatStats) RecordDamageTaken(amount int) {
	s.TotalDamageTaken += amount
}

// RecordHealingDone records healing done
func (s *CombatStats) RecordHealingDone(amount int) {
	s.TotalHealingDone += amount
}

// RecordHealingReceived records healing received
func (s *CombatStats) RecordHealingReceived(amount int) {
	s.TotalHealingReceived += amount
}

// RecordCriticalHit records a critical hit
func (s *CombatStats) RecordCriticalHit() {
	s.CriticalHits++
}

// RecordCriticalMiss records a critical miss
func (s *CombatStats) RecordCriticalMiss() {
	s.CriticalMisses++
}

// RecordAbilityUse records the use of an ability
func (s *CombatStats) RecordAbilityUse(abilityName string) {
	s.AbilitiesUsed[abilityName]++
}

// RecordCombatEnd records the end of a combat
func (s *CombatStats) RecordCombatEnd(victory bool, duration int64) {
	s.CombatCount++
	if victory {
		s.Victories++
	} else {
		s.Defeats++
	}
	s.TotalCombatTime += duration
}

// RecordLoot records collected loot
func (s *CombatStats) RecordLoot(itemName string) {
	s.LootCollected[itemName]++
}

// GetAverageDamageDealt returns the average damage dealt per combat
func (s *CombatStats) GetAverageDamageDealt() float64 {
	if s.CombatCount == 0 {
		return 0
	}
	return float64(s.TotalDamageDealt) / float64(s.CombatCount)
}

// GetAverageDamageTaken returns the average damage taken per combat
func (s *CombatStats) GetAverageDamageTaken() float64 {
	if s.CombatCount == 0 {
		return 0
	}
	return float64(s.TotalDamageTaken) / float64(s.CombatCount)
}

// GetWinRate returns the win rate as a percentage
func (s *CombatStats) GetWinRate() float64 {
	if s.CombatCount == 0 {
		return 0
	}
	return float64(s.Victories) / float64(s.CombatCount) * 100
}

// GetMostUsedAbility returns the most frequently used ability
func (s *CombatStats) GetMostUsedAbility() (string, int) {
	var maxAbility string
	var maxUses int
	for ability, uses := range s.AbilitiesUsed {
		if uses > maxUses {
			maxAbility = ability
			maxUses = uses
		}
	}
	return maxAbility, maxUses
}

// GetMostCollectedLoot returns the most frequently collected loot
func (s *CombatStats) GetMostCollectedLoot() (string, int) {
	var maxLoot string
	var maxCount int
	for loot, count := range s.LootCollected {
		if count > maxCount {
			maxLoot = loot
			maxCount = count
		}
	}
	return maxLoot, maxCount
}
