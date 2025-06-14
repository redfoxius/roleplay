package combat

import (
	"roleplay/internal/character"
	"roleplay/internal/common"
	"roleplay/internal/mob"
)

// MobCombat handles combat between characters and mobs
type MobCombat struct {
	Character *character.Character
	Mob       *mob.Mob
	Turn      int
}

// NewMobCombat creates a new mob combat instance
func NewMobCombat(char *character.Character, m *mob.Mob) *MobCombat {
	return &MobCombat{
		Character: char,
		Mob:       m,
		Turn:      1,
	}
}

// CharacterAttack handles a character's attack on a mob
func (mc *MobCombat) CharacterAttack(ability *common.Ability) (int, error) {
	if mc.Character.SteamPower < ability.SteamCost {
		return 0, ErrInsufficientSteamPower
	}

	// Calculate damage based on character attributes and ability
	damage := ability.Damage
	damage += mc.Character.Attributes.Strength.Value / 2

	// Apply damage to mob
	mc.Mob.Health -= damage
	mc.Character.SteamPower -= ability.SteamCost

	// Check if mob is defeated
	if mc.Mob.Health <= 0 {
		mc.Character.AddExperience(mc.Mob.Experience)
		mc.Character.AddMoney(mc.Mob.MoneyDrop)
		// TODO: Handle loot drops
		return damage, nil
	}

	return damage, nil
}

// MobAttack handles a mob's attack on a character
func (mc *MobCombat) MobAttack() (int, error) {
	// Get mob's AI behavior
	ai := mob.NewAIBehavior(mob.MobType(mc.Mob.Type))
	ai.UpdateAI(mc.Mob, []*character.Character{mc.Character})

	// Choose and execute ability
	ability := ai.ChooseAction(mc.Mob)
	if ability == nil {
		return 0, nil
	}

	// Calculate damage based on mob attributes and ability
	damage := ability.Damage
	damage += mc.Mob.Attributes.Strength.Value / 2

	// Apply damage to character
	mc.Character.Health -= damage
	mc.Mob.SteamPower -= ability.SteamCost

	return damage, nil
}

// IsCombatOver checks if the combat is finished
func (mc *MobCombat) IsCombatOver() bool {
	return mc.Character.Health <= 0 || mc.Mob.Health <= 0
}

// GetCombatResult returns the result of the combat
func (mc *MobCombat) GetCombatResult() string {
	if mc.Character.Health <= 0 {
		return "defeat"
	}
	if mc.Mob.Health <= 0 {
		return "victory"
	}
	return "ongoing"
}
