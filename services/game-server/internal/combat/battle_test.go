package combat

import (
	"math"
	"testing"

	"github.com/redfoxius/roleplay/services/game-server/internal/character"
	"github.com/redfoxius/roleplay/services/game-server/internal/common"
)

func TestNewBattle(t *testing.T) {
	battle := NewBattle("PvP")

	// Test battle creation
	if battle.Type != "PvP" {
		t.Errorf("Expected battle type 'PvP', got '%s'", battle.Type)
	}
	if battle.State != "active" {
		t.Errorf("Expected battle state 'active', got '%s'", battle.State)
	}
	if len(battle.Participants) != 0 {
		t.Errorf("Expected 0 participants, got %d", len(battle.Participants))
	}
}

func TestAddPlayer(t *testing.T) {
	battle := NewBattle("PvP")
	char := &character.Character{
		ID:         "player1",
		Name:       "Player 1",
		Health:     100,
		SteamPower: 50,
	}

	// Test adding a player to the battle
	battle.AddPlayer(char)
	if len(battle.Participants) != 1 {
		t.Errorf("Expected 1 participant, got %d", len(battle.Participants))
	}
	if battle.Participants[0].ID != "player1" {
		t.Errorf("Expected participant ID 'player1', got '%s'", battle.Participants[0].ID)
	}
}

func TestAddMob(t *testing.T) {
	battle := NewBattle("PvE")
	mob := &character.Mob{
		ID:         "mob1",
		Name:       "Mob 1",
		Health:     50,
		SteamPower: 30,
	}

	// Test adding a mob to the battle
	battle.AddMob(mob)
	if len(battle.Participants) != 1 {
		t.Errorf("Expected 1 participant, got %d", len(battle.Participants))
	}
	if battle.Participants[0].ID != "mob1" {
		t.Errorf("Expected participant ID 'mob1', got '%s'", battle.Participants[0].ID)
	}
}

func TestExecuteBattleAction(t *testing.T) {
	battle := NewBattle("PvP")
	char := &character.Character{
		ID:         "player1",
		Name:       "Player 1",
		Health:     100,
		SteamPower: 50,
	}
	battle.AddPlayer(char)

	ability := &common.Ability{
		Name:        "Test Attack",
		Description: "A test attack",
		Type:        "damage",
		Damage:      20,
		SteamCost:   10,
		Cooldown:    0,
	}

	// Test executing an action
	damage, err := battle.ExecuteAction(ability, "player1")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if damage != 20 {
		t.Errorf("Expected damage 20, got %d", damage)
	}
	if char.Health != 80 {
		t.Errorf("Expected character health 80, got %d", char.Health)
	}
}

func TestBattleEnd(t *testing.T) {
	battle := NewBattle(BattleTypePvP)

	// Create two players
	player1 := &character.Character{
		ID:            "player1",
		Name:          "Player 1",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(player1)

	player2 := &character.Character{
		ID:            "player2",
		Name:          "Player 2",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(player2)

	// Create ability
	ability := &common.Ability{
		Name:        "Test Attack",
		Description: "A test attack",
		Type:        "damage",
		Damage:      150, // Enough to defeat player 2
		SteamCost:   10,
		Range:       2,
	}

	// Execute action to defeat player 2
	_, err := battle.ExecuteAction(ability, player2.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check battle state
	if battle.State != "completed" {
		t.Errorf("Expected battle state 'completed', got %s", battle.State)
	}

	if !battle.Participants[1].IsActive {
		t.Error("Expected player 2 to be inactive")
	}

	// Check rewards
	if player1.Experience != 200 { // PvP victory experience
		t.Errorf("Expected player 1 experience 200, got %d", player1.Experience)
	}
}

func TestStatusEffects(t *testing.T) {
	battle := NewBattle(BattleTypePvP)
	battle.Terrain = "mechanical" // Engineer bonus terrain

	// Create engineer
	engineer := &character.Character{
		ID:            "engineer",
		Name:          "Engineer",
		Class:         character.Engineer,
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(engineer)

	// Create target
	target := &character.Character{
		ID:            "target",
		Name:          "Target",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(target)

	// Create mechanical ability
	ability := &common.Ability{
		Name:        "Steam Blast",
		Description: "A mechanical attack",
		Type:        "mechanical",
		Damage:      20,
		SteamCost:   15,
		Range:       2,
	}

	// Execute action
	_, err := battle.ExecuteAction(ability, target.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check status effects
	effects := battle.StatusEffects[target.ID]
	if len(effects) != 1 {
		t.Errorf("Expected 1 status effect, got %d", len(effects))
	}

	if effects[0].Name != "Steam Burn" {
		t.Errorf("Expected status effect 'Steam Burn', got %s", effects[0].Name)
	}

	// Update effects
	battle.UpdateStatusEffects()

	// Check effect duration
	effects = battle.StatusEffects[target.ID]
	if len(effects) != 1 {
		t.Errorf("Expected 1 status effect after update, got %d", len(effects))
	}

	if effects[0].Duration != 1 {
		t.Errorf("Expected effect duration 1, got %d", effects[0].Duration)
	}
}

func TestTerrainAndWeatherEffects(t *testing.T) {
	battle := NewBattle(BattleTypePvP)
	battle.Terrain = "steam-rich"
	battle.Weather = "steam-fog"

	// Create ranged attacker
	attacker := &character.Character{
		ID:            "attacker",
		Name:          "Attacker",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(attacker)

	// Create target
	target := &character.Character{
		ID:            "target",
		Name:          "Target",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(target)

	// Create ranged ability
	ability := &common.Ability{
		Name:        "Ranged Attack",
		Description: "A ranged attack",
		Type:        "ranged",
		Damage:      20,
		SteamCost:   15,
		Range:       3,
	}

	// Execute action
	effect, err := battle.ExecuteAction(ability, target.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check terrain and weather effects
	// Base damage: 20
	// Terrain bonus (steam-rich): 1.2x
	// Weather penalty (steam-fog): 0.8x
	expectedDamage := int(math.Round(float64(20) * 1.2 * 0.8))
	if effect != expectedDamage {
		t.Errorf("Expected damage %d, got %d", expectedDamage, effect)
	}
}

func TestInvalidActions(t *testing.T) {
	battle := NewBattle(BattleTypePvP)

	// Create attacker
	attacker := &character.Character{
		ID:            "attacker",
		Name:          "Attacker",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    10, // Low steam power
		MaxSteamPower: 50,
	}
	battle.AddPlayer(attacker)

	// Create target
	target := &character.Character{
		ID:            "target",
		Name:          "Target",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(target)

	// Test insufficient steam power
	ability := &common.Ability{
		Name:        "Test Attack",
		Description: "A test attack",
		Type:        "damage",
		Damage:      20,
		SteamCost:   20, // More than available
		Range:       2,
	}

	_, err := battle.ExecuteAction(ability, target.ID)
	if err != ErrInsufficientSteamPower {
		t.Errorf("Expected ErrInsufficientSteamPower, got %v", err)
	}

	// Test out of range
	ability.SteamCost = 5
	ability.Range = 1
	attacker.SteamPower = 50
	attacker.Position = common.Coordinates{X: 10, Y: 10}
	target.Position = common.Coordinates{X: 20, Y: 20}

	_, err = battle.ExecuteAction(ability, target.ID)
	if err != ErrInvalidTarget {
		t.Errorf("Expected ErrInvalidTarget, got %v", err)
	}

	// Test invalid target
	_, err = battle.ExecuteAction(ability, "nonexistent_id")
	if err != ErrInvalidTarget {
		t.Errorf("Expected ErrInvalidTarget, got %v", err)
	}

	// Test dead target
	target.Health = 0
	_, err = battle.ExecuteAction(ability, target.ID)
	if err != ErrInvalidTarget {
		t.Errorf("Expected ErrInvalidTarget, got %v", err)
	}
}

func TestHealingActions(t *testing.T) {
	battle := NewBattle(BattleTypePvP)
	battle.Terrain = "steam-rich" // Healing bonus terrain

	// Create healer
	healer := &character.Character{
		ID:            "healer",
		Name:          "Healer",
		Class:         character.Alchemist,
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
		Stats: common.Stats{
			Intelligence: 15,
		},
	}
	battle.AddPlayer(healer)

	// Create wounded target
	target := &character.Character{
		ID:            "target",
		Name:          "Target",
		Health:        50, // Wounded
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(target)

	// Create healing ability
	ability := &common.Ability{
		Name:        "Healing Vapor",
		Description: "A healing ability",
		Type:        "healing",
		Healing:     30,
		SteamCost:   20,
		Range:       2,
	}

	// Execute healing
	effect, err := battle.ExecuteAction(ability, target.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check healing calculation
	// Base healing: 30
	// Intelligence bonus: 15/5 = 3
	// Terrain bonus (steam-rich): 1.2x
	expectedHealing := int(math.Round(float64(33) * 1.2))
	if effect != expectedHealing {
		t.Errorf("Expected healing %d, got %d", expectedHealing, effect)
	}

	// Check target health
	expectedHealth := 50 + expectedHealing
	if expectedHealth > target.MaxHealth {
		expectedHealth = target.MaxHealth
	}
	if battle.Participants[1].Health != expectedHealth {
		t.Errorf("Expected target health %d, got %d", expectedHealth, battle.Participants[1].Health)
	}
}

func TestAreaEffectActions(t *testing.T) {
	battle := NewBattle(BattleTypePvP)

	// Create attacker
	attacker := &character.Character{
		ID:            "attacker",
		Name:          "Attacker",
		Class:         character.SteamMage,
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(attacker)

	// Create multiple targets
	target1 := &character.Character{
		ID:            "target1",
		Name:          "Target 1",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(target1)

	target2 := &character.Character{
		ID:            "target2",
		Name:          "Target 2",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
	}
	battle.AddPlayer(target2)

	// Create area effect ability
	ability := &common.Ability{
		Name:        "Steam Explosion",
		Description: "An area effect attack",
		Type:        "arcane",
		Damage:      15,
		SteamCost:   25,
		Range:       3,
		Area:        2,
	}

	// Execute area effect
	_, err := battle.ExecuteAction(ability, target1.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if both targets took damage
	if battle.Participants[1].Health == 100 || battle.Participants[2].Health == 100 {
		t.Error("Expected both targets to take damage from area effect")
	}

	// Check steam power consumption
	expectedSteam := 50 - 25
	if battle.Participants[0].SteamPower != expectedSteam {
		t.Errorf("Expected attacker steam power %d, got %d", expectedSteam, battle.Participants[0].SteamPower)
	}
}

func TestBattleTurnOrder(t *testing.T) {
	battle := NewBattle(BattleTypePvP)

	// Create participants with different initiative values
	fastChar := &character.Character{
		ID:            "fast",
		Name:          "Fast Character",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
		Stats: common.Stats{
			Dexterity:  20,
			SteamPower: 10,
		},
	}
	battle.AddPlayer(fastChar)

	slowChar := &character.Character{
		ID:            "slow",
		Name:          "Slow Character",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
		Stats: common.Stats{
			Dexterity:  10,
			SteamPower: 5,
		},
	}
	battle.AddPlayer(slowChar)

	// Verify turn order
	if battle.Participants[0].ID != "fast" {
		t.Error("Expected fast character to be first in turn order")
	}
	if battle.Participants[1].ID != "slow" {
		t.Error("Expected slow character to be second in turn order")
	}

	// Test turn progression
	battle.nextTurn()
	if battle.CurrentTurn != 1 {
		t.Errorf("Expected current turn 1, got %d", battle.CurrentTurn)
	}

	battle.nextTurn()
	if battle.CurrentTurn != 0 {
		t.Errorf("Expected current turn 0, got %d", battle.CurrentTurn)
	}
	if battle.Round != 2 {
		t.Errorf("Expected round 2, got %d", battle.Round)
	}
}

func TestBattleRewards(t *testing.T) {
	battle := NewBattle(BattleTypePvE)

	// Create player
	player := &character.Character{
		ID:            "player",
		Name:          "Player",
		Health:        100,
		MaxHealth:     100,
		SteamPower:    50,
		MaxSteamPower: 50,
		Experience:    0,
	}
	battle.AddPlayer(player)

	// Create mob
	m := &character.Mob{
		ID:            "mob",
		Name:          "Test Mob",
		Type:          string(character.Mechanical),
		Level:         5,
		Health:        50,
		MaxHealth:     50,
		SteamPower:    30,
		MaxSteamPower: 30,
		Attributes: character.Attributes{
			Dexterity:  character.Attribute{Name: "Dexterity", Value: 10},
			SteamPower: character.Attribute{Name: "Steam Power", Value: 5},
		},
	}
	battle.AddMob(m)

	// Create ability to defeat mob
	ability := &common.Ability{
		Name:        "Test Attack",
		Description: "A test attack",
		Type:        "damage",
		Damage:      100,
		SteamCost:   20,
		Range:       2,
	}

	// Execute action to defeat mob
	_, err := battle.ExecuteAction(ability, m.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check rewards
	if player.Experience != 500 {
		t.Errorf("Expected player experience 500, got %d", player.Experience)
	}

	expectedMoney := common.NewCurrency(100, 0, 0, 0)
	if player.GetMoney().Copper != expectedMoney.Copper ||
		player.GetMoney().Silver != expectedMoney.Silver ||
		player.GetMoney().Gold != expectedMoney.Gold ||
		player.GetMoney().Platinum != expectedMoney.Platinum {
		t.Errorf("Expected player money %v, got %v", expectedMoney, player.GetMoney())
	}
}

func TestMultiPlayerBattle(t *testing.T) {
	battle := NewBattle(BattleTypePvP)

	// Create multiple players with different attributes
	players := []*character.Character{
		{
			ID:            "tank",
			Name:          "Tank",
			Class:         character.ClockworkKnight,
			Health:        150,
			MaxHealth:     150,
			SteamPower:    50,
			MaxSteamPower: 50,
			Stats: common.Stats{
				Strength: 15,
				Vitality: 15,
			},
		},
		{
			ID:            "healer",
			Name:          "Healer",
			Class:         character.Alchemist,
			Health:        100,
			MaxHealth:     100,
			SteamPower:    60,
			MaxSteamPower: 60,
			Stats: common.Stats{
				Intelligence: 15,
			},
		},
		{
			ID:            "dps",
			Name:          "DPS",
			Class:         character.SteamMage,
			Health:        80,
			MaxHealth:     80,
			SteamPower:    70,
			MaxSteamPower: 70,
			Stats: common.Stats{
				Intelligence: 15,
			},
		},
	}

	// Add all players to battle
	for _, player := range players {
		battle.AddPlayer(player)
	}

	// Verify all players are added
	if len(battle.Participants) != 3 {
		t.Errorf("Expected 3 participants, got %d", len(battle.Participants))
	}

	// Test turn order based on initiative
	// DPS should go first (highest steam power)
	// Healer second
	// Tank last
	if battle.Participants[0].ID != "dps" {
		t.Error("Expected DPS to be first in turn order")
	}
	if battle.Participants[1].ID != "healer" {
		t.Error("Expected Healer to be second in turn order")
	}
	if battle.Participants[2].ID != "tank" {
		t.Error("Expected Tank to be last in turn order")
	}

	// Test multi-target combat
	// DPS attacks Tank
	dpsAbility := &common.Ability{
		Name:        "Steam Bolt",
		Description: "A powerful steam attack",
		Type:        "arcane",
		Damage:      25,
		SteamCost:   20,
		Range:       3,
	}

	_, err := battle.ExecuteAction(dpsAbility, "tank")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Healer heals Tank
	healAbility := &common.Ability{
		Name:        "Healing Vapor",
		Description: "A healing ability",
		Type:        "healing",
		Healing:     30,
		SteamCost:   25,
		Range:       2,
	}

	_, err = battle.ExecuteAction(healAbility, "tank")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Tank attacks DPS
	tankAbility := &common.Ability{
		Name:        "Steam Hammer",
		Description: "A powerful melee attack",
		Type:        "melee",
		Damage:      20,
		SteamCost:   15,
		Range:       1,
	}

	_, err = battle.ExecuteAction(tankAbility, "dps")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify health changes
	tank := battle.Participants[2]
	dps := battle.Participants[0]

	// Tank should have taken damage and been healed
	expectedTankHealth := 150 - 25 + 30
	if tank.Health != expectedTankHealth {
		t.Errorf("Expected tank health %d, got %d", expectedTankHealth, tank.Health)
	}

	// DPS should have taken damage from tank
	expectedDPSHealth := 80 - 20
	if dps.Health != expectedDPSHealth {
		t.Errorf("Expected DPS health %d, got %d", expectedDPSHealth, dps.Health)
	}

	// Test battle completion with multiple players
	// DPS attacks Tank with lethal damage
	lethalAbility := &common.Ability{
		Name:        "Steam Explosion",
		Description: "A powerful area attack",
		Type:        "arcane",
		Damage:      200,
		SteamCost:   50,
		Range:       3,
		Area:        2,
	}

	_, err = battle.ExecuteAction(lethalAbility, "tank")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify battle state
	if battle.State != "completed" {
		t.Error("Expected battle to be completed")
	}

	if tank.IsActive {
		t.Error("Expected tank to be inactive")
	}

	// Verify rewards
	if dps.Experience != 200 { // PvP victory experience
		t.Errorf("Expected DPS experience 200, got %d", dps.Experience)
	}
}

func TestTeamBattle(t *testing.T) {
	battle := NewBattle(BattleTypePvP)

	// Create two teams
	team1 := []*character.Character{
		{
			ID:            "team1_tank",
			Name:          "Team 1 Tank",
			Class:         character.ClockworkKnight,
			Health:        150,
			MaxHealth:     150,
			SteamPower:    50,
			MaxSteamPower: 50,
		},
		{
			ID:            "team1_healer",
			Name:          "Team 1 Healer",
			Class:         character.Alchemist,
			Health:        100,
			MaxHealth:     100,
			SteamPower:    60,
			MaxSteamPower: 60,
		},
	}

	team2 := []*character.Character{
		{
			ID:            "team2_tank",
			Name:          "Team 2 Tank",
			Class:         character.ClockworkKnight,
			Health:        150,
			MaxHealth:     150,
			SteamPower:    50,
			MaxSteamPower: 50,
		},
		{
			ID:            "team2_healer",
			Name:          "Team 2 Healer",
			Class:         character.Alchemist,
			Health:        100,
			MaxHealth:     100,
			SteamPower:    60,
			MaxSteamPower: 60,
		},
	}

	// Add all players to battle
	for _, player := range append(team1, team2...) {
		battle.AddPlayer(player)
	}

	// Test team-based combat
	// Team 1 attacks Team 2
	attackAbility := &common.Ability{
		Name:        "Team Attack",
		Description: "A coordinated attack",
		Type:        "damage",
		Damage:      30,
		SteamCost:   25,
		Range:       3,
		Area:        2,
	}

	// Team 1 Tank attacks Team 2 Tank
	_, err := battle.ExecuteAction(attackAbility, "team2_tank")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Team 1 Healer heals Team 1 Tank
	healAbility := &common.Ability{
		Name:        "Team Heal",
		Description: "A healing ability",
		Type:        "healing",
		Healing:     40,
		SteamCost:   30,
		Range:       2,
	}

	_, err = battle.ExecuteAction(healAbility, "team1_tank")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Team 2 counter-attacks
	// Team 2 Tank attacks Team 1 Healer
	_, err = battle.ExecuteAction(attackAbility, "team1_healer")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Team 2 Healer heals Team 2 Tank
	_, err = battle.ExecuteAction(healAbility, "team2_tank")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify team health states
	team1Tank := battle.Participants[0]
	team1Healer := battle.Participants[1]
	team2Tank := battle.Participants[2]
	team2Healer := battle.Participants[3]

	// Team 1 Tank should be at full health (healed)
	if team1Tank.Health != 150 {
		t.Errorf("Expected Team 1 Tank health 150, got %d", team1Tank.Health)
	}

	// Team 1 Healer should have taken damage
	if team1Healer.Health != 70 { // 100 - 30
		t.Errorf("Expected Team 1 Healer health 70, got %d", team1Healer.Health)
	}

	// Team 2 Tank should be healed
	if team2Tank.Health != 150 {
		t.Errorf("Expected Team 2 Tank health 150, got %d", team2Tank.Health)
	}

	// Test team elimination
	// Team 1 eliminates Team 2
	lethalAbility := &common.Ability{
		Name:        "Team Finisher",
		Description: "A finishing move",
		Type:        "damage",
		Damage:      200,
		SteamCost:   50,
		Range:       3,
		Area:        2,
	}

	// Team 1 Tank eliminates Team 2 Tank
	_, err = battle.ExecuteAction(lethalAbility, "team2_tank")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Team 1 Healer eliminates Team 2 Healer
	_, err = battle.ExecuteAction(lethalAbility, "team2_healer")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify battle completion
	if battle.State != "completed" {
		t.Error("Expected battle to be completed")
	}

	if team2Tank.IsActive || team2Healer.IsActive {
		t.Error("Expected Team 2 to be eliminated")
	}

	// Verify team rewards
	if team1Tank.Experience != 200 || team1Healer.Experience != 200 {
		t.Error("Expected Team 1 to receive victory experience")
	}
}
