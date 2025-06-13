package combat

import (
	"time"
)

// LogEntry represents a single entry in the combat log
type LogEntry struct {
	Timestamp time.Time
	Type      string // attack, ability, effect, death, loot
	Source    string
	Target    string
	Value     int
	Message   string
}

// CombatLog manages the combat log
type CombatLog struct {
	Entries []LogEntry
}

// NewCombatLog creates a new combat log
func NewCombatLog() *CombatLog {
	return &CombatLog{
		Entries: make([]LogEntry, 0),
	}
}

// AddEntry adds a new entry to the combat log
func (cl *CombatLog) AddEntry(entryType, source, target string, value int, message string) {
	cl.Entries = append(cl.Entries, LogEntry{
		Timestamp: time.Now(),
		Type:      entryType,
		Source:    source,
		Target:    target,
		Value:     value,
		Message:   message,
	})
}

// GetRecentEntries returns the most recent entries
func (cl *CombatLog) GetRecentEntries(count int) []LogEntry {
	if count <= 0 || count > len(cl.Entries) {
		count = len(cl.Entries)
	}
	start := len(cl.Entries) - count
	return cl.Entries[start:]
}

// GetEntriesByType returns all entries of a specific type
func (cl *CombatLog) GetEntriesByType(entryType string) []LogEntry {
	var filtered []LogEntry
	for _, entry := range cl.Entries {
		if entry.Type == entryType {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// GetEntriesBySource returns all entries from a specific source
func (cl *CombatLog) GetEntriesBySource(source string) []LogEntry {
	var filtered []LogEntry
	for _, entry := range cl.Entries {
		if entry.Source == source {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// GetEntriesByTarget returns all entries targeting a specific entity
func (cl *CombatLog) GetEntriesByTarget(target string) []LogEntry {
	var filtered []LogEntry
	for _, entry := range cl.Entries {
		if entry.Target == target {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// Clear clears the combat log
func (cl *CombatLog) Clear() {
	cl.Entries = make([]LogEntry, 0)
}

// Helper functions for common log entries
func (cl *CombatLog) LogAttack(source, target string, damage int) {
	cl.AddEntry("attack", source, target, damage,
		formatMessage(source, target, "attacks", damage))
}

func (cl *CombatLog) LogAbility(source, target, abilityName string, value int) {
	cl.AddEntry("ability", source, target, value,
		formatMessage(source, target, "uses "+abilityName, value))
}

func (cl *CombatLog) LogEffect(source, target, effectName string, value int) {
	cl.AddEntry("effect", source, target, value,
		formatMessage(source, target, "is affected by "+effectName, value))
}

func (cl *CombatLog) LogDeath(entity string) {
	cl.AddEntry("death", entity, "", 0,
		entity+" has been defeated")
}

func (cl *CombatLog) LogLoot(source string, itemName string, quantity int) {
	cl.AddEntry("loot", source, "", quantity,
		formatMessage(source, "", "receives", quantity)+" "+itemName)
}

func formatMessage(source, target, action string, value int) string {
	if target == "" {
		return source + " " + action
	}
	if value > 0 {
		return source + " " + action + " " + target + " for " + string(value)
	}
	return source + " " + action + " " + target
}
