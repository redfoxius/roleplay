package character

// Class represents a character class in the game
type Class string

const (
	Engineer        Class = "Engineer"
	Alchemist       Class = "Alchemist"
	Aeronaut        Class = "Aeronaut"
	ClockworkKnight Class = "ClockworkKnight"
	SteamMage       Class = "SteamMage"
)

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

// Character represents a player character
type Character struct {
	ID         string
	Name       string
	Class      Class
	Level      int
	Experience int
	Health     int
	MaxHealth  int
	Attributes Attributes
	Equipment  []Equipment
	Inventory  []Item
}

// Equipment represents items that can be equipped
type Equipment struct {
	ID           string
	Name         string
	Type         string
	Slot         string
	Attributes   map[string]int
	Requirements map[string]int
}

// Item represents items that can be carried
type Item struct {
	ID          string
	Name        string
	Type        string
	Description string
	Stackable   bool
	Quantity    int
}

// NewCharacter creates a new character with default attributes
func NewCharacter(name string, class Class) *Character {
	return &Character{
		Name:  name,
		Class: class,
		Level: 1,
		Attributes: Attributes{
			Strength:            Attribute{Name: "Strength", Value: 10},
			Dexterity:           Attribute{Name: "Dexterity", Value: 10},
			Constitution:        Attribute{Name: "Constitution", Value: 10},
			Intelligence:        Attribute{Name: "Intelligence", Value: 10},
			Wisdom:              Attribute{Name: "Wisdom", Value: 10},
			Charisma:            Attribute{Name: "Charisma", Value: 10},
			TechnicalAptitude:   Attribute{Name: "Technical Aptitude", Value: 10},
			SteamPower:          Attribute{Name: "Steam Power", Value: 10},
			MechanicalPrecision: Attribute{Name: "Mechanical Precision", Value: 10},
			ArcaneKnowledge:     Attribute{Name: "Arcane Knowledge", Value: 10},
		},
		Health:    100,
		MaxHealth: 100,
	}
}

// ApplyClassBonuses applies the initial class bonuses to attributes
func (c *Character) ApplyClassBonuses() {
	switch c.Class {
	case Engineer:
		c.Attributes.TechnicalAptitude.Value += 5
		c.Attributes.MechanicalPrecision.Value += 5
	case Alchemist:
		c.Attributes.Intelligence.Value += 5
		c.Attributes.ArcaneKnowledge.Value += 5
	case Aeronaut:
		c.Attributes.Dexterity.Value += 5
		c.Attributes.SteamPower.Value += 5
	case ClockworkKnight:
		c.Attributes.Strength.Value += 5
		c.Attributes.Constitution.Value += 5
	case SteamMage:
		c.Attributes.ArcaneKnowledge.Value += 5
		c.Attributes.SteamPower.Value += 5
	}
}
