package character

import (
	"time"
)

// AIState represents the current state of a mob's AI
type AIState string

const (
	Idle       AIState = "idle"
	Aggressive AIState = "aggressive"
	Defensive  AIState = "defensive"
	Fleeing    AIState = "fleeing"
	Patrolling AIState = "patrolling"
)

// AIBehavior represents the AI behavior of a mob
type AIBehavior struct {
	State              AIState
	Target             *Character
	AggroRange         int
	FleeThreshold      float32 // Percentage of health to trigger flee
	PatrolPath         []string
	CurrentPatrolIndex int
	LastActionTime     time.Time
	Cooldowns          map[string]time.Time
}

// NewAIBehavior creates a new AI behavior
func NewAIBehavior(mobType MobType) *AIBehavior {
	behavior := &AIBehavior{
		State:          Idle,
		AggroRange:     10,
		FleeThreshold:  0.2, // 20% health
		LastActionTime: time.Now(),
		Cooldowns:      make(map[string]time.Time),
	}

	// Set type-specific behavior
	switch mobType {
	case Mechanical:
		behavior.AggroRange = 15
		behavior.FleeThreshold = 0.1
	case Biological:
		behavior.AggroRange = 8
		behavior.FleeThreshold = 0.3
	case Hybrid:
		behavior.AggroRange = 12
		behavior.FleeThreshold = 0.25
	case Elemental:
		behavior.AggroRange = 20
		behavior.FleeThreshold = 0.15
	case Construct:
		behavior.AggroRange = 5
		behavior.FleeThreshold = 0.05
	}

	return behavior
}

// UpdateAI updates the mob's AI state and behavior
func (ai *AIBehavior) UpdateAI(mob *Mob, nearbyPlayers []*Character) {
	// Check if we should flee
	if float32(mob.Health)/float32(mob.MaxHealth) <= ai.FleeThreshold {
		ai.State = Fleeing
		return
	}

	// Find nearest player in aggro range
	var nearestPlayer *Character
	var minDistance int = ai.AggroRange + 1

	for _, player := range nearbyPlayers {
		// TODO: Implement actual distance calculation
		distance := 5 // Placeholder
		if distance < minDistance {
			minDistance = distance
			nearestPlayer = player
		}
	}

	// Update state based on nearest player
	if nearestPlayer != nil {
		ai.Target = nearestPlayer
		ai.State = Aggressive
	} else {
		ai.Target = nil
		if ai.State == Aggressive {
			ai.State = Patrolling
		}
	}
}

// ChooseAction selects the next action for the mob to take
func (ai *AIBehavior) ChooseAction(mob *Mob) *Ability {
	if ai.State == Fleeing {
		return nil // Mob will try to escape
	}

	if ai.State == Aggressive && ai.Target != nil {
		// Check cooldowns and choose best ability
		var bestAbility *Ability
		var bestScore float32

		for i := range mob.Abilities {
			ability := &mob.Abilities[i]

			// Skip if on cooldown
			if lastUse, exists := ai.Cooldowns[ability.Name]; exists {
				if time.Since(lastUse).Seconds() < float64(ability.Cooldown) {
					continue
				}
			}

			// Calculate ability score based on current situation
			score := ai.calculateAbilityScore(ability, mob)
			if score > bestScore {
				bestScore = score
				bestAbility = ability
			}
		}

		if bestAbility != nil {
			ai.Cooldowns[bestAbility.Name] = time.Now()
			return bestAbility
		}
	}

	return nil
}

// calculateAbilityScore calculates how good an ability is for the current situation
func (ai *AIBehavior) calculateAbilityScore(ability *Ability, mob *Mob) float32 {
	var score float32

	// Base score on ability type
	switch ability.Type {
	case "damage":
		score = float32(ability.Damage)
	case "healing":
		// Prioritize healing when health is low
		healthPercent := float32(mob.Health) / float32(mob.MaxHealth)
		score = float32(ability.Healing) * (1.0 - healthPercent)
	case "buff":
		score = 50 // Base score for buffs
	case "debuff":
		score = 40 // Base score for debuffs
	}

	// Adjust score based on steam power cost
	steamPercent := float32(mob.SteamPower) / float32(mob.MaxSteamPower)
	if steamPercent < 0.3 && ability.SteamCost > 20 {
		score *= 0.5 // Reduce score for expensive abilities when low on steam
	}

	return score
}

// SetPatrolPath sets the patrol path for the mob
func (ai *AIBehavior) SetPatrolPath(path []string) {
	ai.PatrolPath = path
	ai.CurrentPatrolIndex = 0
	ai.State = Patrolling
}

// GetNextPatrolPoint returns the next point in the patrol path
func (ai *AIBehavior) GetNextPatrolPoint() string {
	if len(ai.PatrolPath) == 0 {
		return ""
	}

	point := ai.PatrolPath[ai.CurrentPatrolIndex]
	ai.CurrentPatrolIndex = (ai.CurrentPatrolIndex + 1) % len(ai.PatrolPath)
	return point
}
