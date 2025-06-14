package combat

import (
	"fmt"
	"sort"

	"github.com/redfoxius/roleplay/services/game-server/internal/character"
	"github.com/redfoxius/roleplay/services/game-server/internal/common"

	"github.com/google/uuid"
)

// BattleType represents the type of battle
type BattleType string

const (
	BattleTypePvP  BattleType = "pvp"
	BattleTypePvE  BattleType = "pve"
	BattleTypeRaid BattleType = "raid"
)

// Battle represents a battle instance
type Battle struct {
	ID            string
	Type          BattleType
	State         string // "active", "completed", "cancelled"
	Participants  []*Participant
	TurnOrder     []*Participant
	CurrentTurn   int
	Round         int
	CombatLog     []CombatLogEntry
	Terrain       string
	Weather       string
	Teams         map[string][]string       // Team ID -> Participant IDs
	StatusEffects map[string][]StatusEffect // Participant ID -> Status Effects
	ActiveIndex   int                       // Index of the active participant in TurnOrder
}

// Participant represents a battle participant (player or mob)
type Participant struct {
	ID            string
	Name          string
	Type          string // "player" or "mob"
	Class         character.Class
	Team          string // Team ID, empty for free-for-all
	Health        int
	MaxHealth     int
	SteamPower    int
	MaxSteamPower int
	Attributes    common.Attributes
	Abilities     []common.Ability
	IsActive      bool
	Experience    int
	Money         common.Currency
}

// NewBattle creates a new battle instance
func NewBattle(battleType BattleType) *Battle {
	return &Battle{
		ID:            uuid.New().String(),
		Type:          battleType,
		State:         "active",
		Participants:  make([]*Participant, 0),
		TurnOrder:     make([]*Participant, 0),
		CurrentTurn:   0,
		Round:         1,
		CombatLog:     make([]CombatLogEntry, 0),
		Terrain:       "neutral",
		Weather:       "clear",
		Teams:         make(map[string][]string),
		StatusEffects: make(map[string][]StatusEffect),
		ActiveIndex:   0,
	}
}

// AddPlayer adds a player to the battle
func (b *Battle) AddPlayer(player *character.Character) {
	participant := &Participant{
		ID:            player.ID,
		Name:          player.Name,
		Type:          "player",
		Class:         player.Class,
		Health:        player.Health,
		MaxHealth:     player.MaxHealth,
		SteamPower:    player.SteamPower,
		MaxSteamPower: player.MaxSteamPower,
		Attributes: common.Attributes{
			Strength:     common.Attribute{Name: "Strength", Value: player.Stats.Strength},
			Dexterity:    common.Attribute{Name: "Dexterity", Value: player.Stats.Dexterity},
			Intelligence: common.Attribute{Name: "Intelligence", Value: player.Stats.Intelligence},
			Constitution: common.Attribute{Name: "Constitution", Value: player.Stats.Vitality},
			SteamPower:   common.Attribute{Name: "SteamPower", Value: player.Stats.SteamPower},
		},
		Abilities:  player.GetAbilities(),
		IsActive:   true,
		Experience: 0,
		Money:      common.NewCurrency(0, 0, 0, 0), // Mobs don't have money until defeated
	}

	b.Participants = append(b.Participants, participant)
	b.TurnOrder = append(b.TurnOrder, participant)
	b.sortTurnOrder()
}

// AddMob adds a mob to the battle
func (b *Battle) AddMob(mob *character.Mob) {
	participant := &Participant{
		ID:            mob.ID,
		Name:          mob.Name,
		Type:          "mob",
		Health:        mob.Health,
		MaxHealth:     mob.MaxHealth,
		SteamPower:    mob.SteamPower,
		MaxSteamPower: mob.MaxSteamPower,
		Attributes: common.Attributes{
			Strength:     common.Attribute{Value: mob.Attributes.Strength.Value},
			Dexterity:    common.Attribute{Value: mob.Attributes.Dexterity.Value},
			Intelligence: common.Attribute{Value: mob.Attributes.Intelligence.Value},
			Constitution: common.Attribute{Value: mob.Attributes.Constitution.Value},
			SteamPower:   common.Attribute{Value: mob.Attributes.SteamPower.Value},
		},
		Abilities:  convertAbilities(mob.Abilities),
		IsActive:   true,
		Experience: mob.Experience,
		Money:      common.NewCurrency(0, 0, 0, 0), // Mobs don't have money until defeated
	}

	b.Participants = append(b.Participants, participant)
	b.TurnOrder = append(b.TurnOrder, participant)
	b.sortTurnOrder()
}

// convertAbilities converts character.Ability to common.Ability
func convertAbilities(abilities []character.Ability) []common.Ability {
	result := make([]common.Ability, len(abilities))
	for i, a := range abilities {
		result[i] = common.Ability{
			Name:        a.Name,
			Description: a.Description,
			Type:        a.Type,
			Damage:      a.Damage,
			Healing:     a.Healing,
			SteamCost:   a.SteamCost,
			Range:       1, // Default range
			Area:        0, // Default single target
			Cooldown:    a.Cooldown,
		}
	}
	return result
}

// AddToTeam adds a participant to a team
func (b *Battle) AddToTeam(participantID, teamID string) {
	participant := b.getParticipant(participantID)
	if participant != nil {
		participant.Team = teamID
		b.Teams[teamID] = append(b.Teams[teamID], participantID)
	}
}

// ExecuteAction executes a combat action
func (b *Battle) ExecuteAction(ability *common.Ability, targetID string) (int, error) {
	if b.State != "active" {
		return 0, fmt.Errorf("battle is not active")
	}

	attacker := b.TurnOrder[b.CurrentTurn]
	if !attacker.IsActive {
		return 0, fmt.Errorf("attacker is not active")
	}

	target := b.getParticipant(targetID)
	if target == nil || !target.IsActive {
		return 0, ErrInvalidTarget
	}

	// Check if action is valid
	if !b.isValidAction(ability, attacker, target) {
		return 0, ErrInvalidAction
	}

	// Calculate terrain and weather effects
	terrainBonus := b.calculateTerrainBonus(ability)
	weatherPenalty := b.calculateWeatherPenalty(ability)

	// Apply action effects
	var damage, healing int
	if ability.Damage > 0 {
		damage = b.calculateDamage(ability.Damage, attacker, target)
		damage = int(float64(damage) * terrainBonus * weatherPenalty)
		target.Health -= damage
	}

	if ability.Healing > 0 {
		healing = b.calculateHealing(ability.Healing, attacker)
		healing = int(float64(healing) * terrainBonus * weatherPenalty)
		target.Health = min(target.Health+healing, target.MaxHealth)
	}

	// Apply area effect if applicable
	if ability.Area > 0 {
		b.applyAreaEffect(ability, attacker, target, damage, healing)
	}

	// Apply status effects
	b.applyStatusEffects(attacker, target, ability)

	// Log the action
	b.logAction(ability, attacker, target, damage, healing)

	// Check if target is defeated
	if target.Health <= 0 {
		target.IsActive = false
		b.checkBattleCompletion()
	}

	// Move to next turn
	b.nextTurn()

	return damage, nil
}

// applyAreaEffect applies an area effect to all valid targets
func (b *Battle) applyAreaEffect(ability *common.Ability, attacker, centerTarget *Participant, baseDamage, baseHealing int) {
	for _, target := range b.Participants {
		if !target.IsActive || target == centerTarget {
			continue
		}

		// Check if target is in range
		if !b.isInRange(attacker, target, ability.Area) {
			continue
		}

		// Apply reduced damage/healing to area targets
		if ability.Damage > 0 {
			damage := baseDamage / 2
			target.Health -= damage
		}

		if ability.Healing > 0 {
			healing := baseHealing / 2
			target.Health = min(target.Health+healing, target.MaxHealth)
		}
	}
}

// isInRange checks if two participants are within range of each other
func (b *Battle) isInRange(attacker, target *Participant, rangeValue int) bool {
	// For now, we'll consider all participants in range
	// In a real implementation, this would check actual coordinates
	return true
}

// checkBattleCompletion checks if the battle is complete
func (b *Battle) checkBattleCompletion() {
	activeTeams := make(map[string]bool)
	activeParticipants := 0

	for _, p := range b.Participants {
		if p.IsActive {
			activeParticipants++
			if p.Team != "" {
				activeTeams[p.Team] = true
			}
		}
	}

	// Battle is complete if:
	// 1. Only one participant remains (free-for-all)
	// 2. Only one team remains (team battle)
	// 3. No participants remain (draw)
	if activeParticipants <= 1 || len(activeTeams) <= 1 {
		b.State = "completed"
		b.distributeRewards()
	}
}

// distributeRewards distributes rewards to the winning team/participant
func (b *Battle) distributeRewards() {
	var winners []*Participant
	var winningTeam string

	// Find winners
	if len(b.Teams) > 0 {
		// Team battle
		for teamID, participantIDs := range b.Teams {
			teamActive := false
			for _, id := range participantIDs {
				if p := b.getParticipant(id); p != nil && p.IsActive {
					teamActive = true
					break
				}
			}
			if teamActive {
				winningTeam = teamID
				break
			}
		}

		// Add all active participants from winning team
		for _, id := range b.Teams[winningTeam] {
			if p := b.getParticipant(id); p != nil && p.IsActive {
				winners = append(winners, p)
			}
		}
	} else {
		// Free-for-all
		for _, p := range b.Participants {
			if p.IsActive {
				winners = append(winners, p)
			}
		}
	}

	// Distribute rewards
	rewardPerWinner := 200 // Base experience reward
	for _, winner := range winners {
		winner.Experience += rewardPerWinner
	}
}

// getParticipant returns a participant by ID
func (b *Battle) getParticipant(id string) *Participant {
	for _, p := range b.Participants {
		if p.ID == id {
			return p
		}
	}
	return nil
}

// sortTurnOrder sorts participants by initiative
func (b *Battle) sortTurnOrder() {
	sort.Slice(b.TurnOrder, func(i, j int) bool {
		initiativeI := b.TurnOrder[i].Attributes.Dexterity.Value + b.TurnOrder[i].Attributes.SteamPower.Value
		initiativeJ := b.TurnOrder[j].Attributes.Dexterity.Value + b.TurnOrder[j].Attributes.SteamPower.Value
		return initiativeI > initiativeJ
	})
}

// nextTurn advances to the next turn
func (b *Battle) nextTurn() {
	b.CurrentTurn++
	if b.CurrentTurn >= len(b.TurnOrder) {
		b.CurrentTurn = 0
		b.Round++
		b.UpdateStatusEffects()
	}
}

// isValidAction checks if an action is valid
func (b *Battle) isValidAction(ability *common.Ability, attacker, target *Participant) bool {
	if !target.IsActive {
		return false
	}

	if attacker.SteamPower < ability.SteamCost {
		return false
	}

	if ability.Range > 0 {
		// Replace with a default value or remove the check, since location is not tracked
		// distance := distanceBetween(attacker.Location, target.Location)
		// if distance > ability.Range {
		// 	return false
		// }
	}

	return true
}

// calculateDamage calculates damage for an action
func (b *Battle) calculateDamage(damage int, attacker, target *Participant) int {
	// Add attribute bonuses
	if attacker.Type == "player" {
		damage += attacker.Attributes.Strength.Value / 2
	} else {
		damage += attacker.Attributes.Strength.Value / 2
	}

	// Apply defense reduction
	defense := target.Attributes.Constitution.Value / 2
	damage = max(1, damage-defense)

	return damage
}

// calculateHealing calculates healing for an action
func (b *Battle) calculateHealing(healing int, attacker *Participant) int {
	// Add attribute bonuses
	if attacker.Type == "player" {
		healing += attacker.Attributes.Wisdom.Value / 2
	} else {
		healing += attacker.Attributes.Wisdom.Value / 2
	}

	return healing
}

// calculateTerrainBonus calculates bonus damage/healing based on terrain
func (b *Battle) calculateTerrainBonus(ability *common.Ability) float64 {
	bonus := 1.0
	switch b.Terrain {
	case "steam-rich":
		if ability.SteamCost > 0 {
			bonus = 1.2
		}
	case "mechanical":
		if b.TurnOrder[b.CurrentTurn].Type == "player" {
			// Check class through abilities
			for _, ab := range b.TurnOrder[b.CurrentTurn].Abilities {
				if ab.Type == "mechanical" {
					bonus = 1.15
					break
				}
			}
		}
	case "toxic":
		if b.TurnOrder[b.CurrentTurn].Type == "player" {
			// Check class through abilities
			for _, ab := range b.TurnOrder[b.CurrentTurn].Abilities {
				if ab.Type == "chemical" {
					bonus = 1.15
					break
				}
			}
		}
	}
	return bonus
}

// calculateWeatherPenalty calculates penalties based on weather
func (b *Battle) calculateWeatherPenalty(ability *common.Ability) float64 {
	penalty := 1.0
	switch b.Weather {
	case "steam-fog":
		if ability.Type == "ranged" {
			penalty = 0.8
		}
	case "acid-rain":
		penalty = 0.9
	}
	return penalty
}

// applyStatusEffects applies status effects to participants
func (b *Battle) applyStatusEffects(attacker, target *Participant, ability *common.Ability) {
	if attacker.Type != "player" {
		return
	}

	// Apply effects based on character class and ability type
	switch attacker.Class {
	case character.Engineer:
		if ability.Type == "mechanical" {
			b.addStatusEffect(target.ID, StatusEffect{
				Name:        "Steam Burn",
				Duration:    2,
				Value:       5,
				Description: "Takes 5 damage per turn",
			})
		}
	case character.Alchemist:
		if ability.Type == "chemical" {
			b.addStatusEffect(target.ID, StatusEffect{
				Name:        "Poison",
				Duration:    3,
				Value:       3,
				Description: "Takes 3 damage per turn",
			})
		}
	case character.SteamMage:
		if ability.Type == "arcane" {
			b.addStatusEffect(target.ID, StatusEffect{
				Name:        "Steam Weakness",
				Duration:    2,
				Value:       -10,
				Description: "Steam power reduced by 10",
			})
		}
	}
}

// addStatusEffect adds a status effect to a participant
func (b *Battle) addStatusEffect(participantID string, effect StatusEffect) {
	b.StatusEffects[participantID] = append(b.StatusEffects[participantID], effect)
}

// logAction adds an entry to the combat log
func (b *Battle) logAction(ability *common.Ability, attacker, target *Participant, damage, healing int) {
	entry := CombatLogEntry{
		Round:     b.Round,
		Turn:      b.CurrentTurn,
		Character: attacker.Name,
		Action:    ability.Name,
		Target:    target.Name,
		Damage:    damage,
		Healing:   healing,
		Effects:   make([]string, 0),
	}

	// Add status effects to log
	for _, effect := range b.StatusEffects[target.ID] {
		entry.Effects = append(entry.Effects, effect.Name)
	}

	b.CombatLog = append(b.CombatLog, entry)
}

// UpdateStatusEffects updates all status effects
func (b *Battle) UpdateStatusEffects() {
	for participantID, effects := range b.StatusEffects {
		var remainingEffects []StatusEffect
		for _, effect := range effects {
			effect.Duration--
			if effect.Duration > 0 {
				remainingEffects = append(remainingEffects, effect)
			}
		}
		b.StatusEffects[participantID] = remainingEffects
	}
}
