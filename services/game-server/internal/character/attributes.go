package character

import "github.com/redfoxius/roleplay/services/game-server/internal/common"

// Attribute represents a character's base attribute
type Attribute struct {
	Name  string
	Value int
}

// Attributes represents all character attributes
type Attributes struct {
	Strength            Attribute
	Dexterity           Attribute
	Constitution        Attribute
	Intelligence        Attribute
	Wisdom              Attribute
	Charisma            Attribute
	TechnicalAptitude   Attribute
	SteamPower          Attribute
	MechanicalPrecision Attribute
	ArcaneKnowledge     Attribute
}

// getBaseAttributes returns the base attributes for a given class
func getBaseAttributes(class Class) common.Attributes {
	base := common.Attributes{
		Strength:            common.Attribute{Name: "Strength", Value: 10},
		Dexterity:           common.Attribute{Name: "Dexterity", Value: 10},
		Constitution:        common.Attribute{Name: "Constitution", Value: 10},
		Intelligence:        common.Attribute{Name: "Intelligence", Value: 10},
		Wisdom:              common.Attribute{Name: "Wisdom", Value: 10},
		Charisma:            common.Attribute{Name: "Charisma", Value: 10},
		TechnicalAptitude:   common.Attribute{Name: "Technical Aptitude", Value: 10},
		SteamPower:          common.Attribute{Name: "Steam Power", Value: 10},
		MechanicalPrecision: common.Attribute{Name: "Mechanical Precision", Value: 10},
		ArcaneKnowledge:     common.Attribute{Name: "Arcane Knowledge", Value: 10},
	}

	// Add class-specific bonuses
	switch class {
	case Engineer:
		base.TechnicalAptitude.Value += 3
		base.MechanicalPrecision.Value += 2
		base.Intelligence.Value += 1
	case Alchemist:
		base.Intelligence.Value += 3
		base.ArcaneKnowledge.Value += 2
		base.Wisdom.Value += 1
	case Aeronaut:
		base.Dexterity.Value += 3
		base.SteamPower.Value += 2
		base.TechnicalAptitude.Value += 1
	case ClockworkKnight:
		base.Strength.Value += 3
		base.Constitution.Value += 2
		base.MechanicalPrecision.Value += 1
	case SteamMage:
		base.ArcaneKnowledge.Value += 3
		base.SteamPower.Value += 2
		base.Intelligence.Value += 1
	}

	return base
}
