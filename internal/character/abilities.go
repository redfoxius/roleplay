package character

// Ability represents a mob's special ability
type Ability struct {
	Name        string
	Description string
	Damage      int
	Healing     int
	SteamCost   int
	Cooldown    int
	CurrentCD   int
	Type        string // "damage", "healing", "buff", "debuff"
	Effect      string // Additional effect description
}

// GetMobAbilities returns abilities based on mob type and level
func GetMobAbilities(mobType MobType, level int) []Ability {
	var abilities []Ability

	switch mobType {
	case Mechanical:
		abilities = append(abilities, []Ability{
			{
				Name:        "Steam Blast",
				Description: "Releases a burst of superheated steam",
				Damage:      10 + (level * 2),
				SteamCost:   15,
				Cooldown:    3,
				Type:        "damage",
				Effect:      "Chance to burn target",
			},
			{
				Name:        "Clockwork Repair",
				Description: "Repairs mechanical components",
				Healing:     20 + (level * 3),
				SteamCost:   20,
				Cooldown:    4,
				Type:        "healing",
				Effect:      "Restores steam power",
			},
			{
				Name:        "Gear Shift",
				Description: "Increases mechanical precision",
				SteamCost:   10,
				Cooldown:    5,
				Type:        "buff",
				Effect:      "Increases damage for 3 turns",
			},
		}...)

	case Biological:
		abilities = append(abilities, []Ability{
			{
				Name:        "Toxic Bite",
				Description: "A poisonous attack",
				Damage:      15 + (level * 2),
				SteamCost:   10,
				Cooldown:    2,
				Type:        "damage",
				Effect:      "Poisons target",
			},
			{
				Name:        "Steam Regeneration",
				Description: "Heals using steam energy",
				Healing:     25 + (level * 3),
				SteamCost:   25,
				Cooldown:    4,
				Type:        "healing",
				Effect:      "Gradual healing over time",
			},
			{
				Name:        "Chemical Rage",
				Description: "Increases strength temporarily",
				SteamCost:   15,
				Cooldown:    6,
				Type:        "buff",
				Effect:      "Increases strength for 4 turns",
			},
		}...)

	case Hybrid:
		abilities = append(abilities, []Ability{
			{
				Name:        "Steam Blade",
				Description: "A cutting attack with steam",
				Damage:      20 + (level * 2),
				SteamCost:   20,
				Cooldown:    3,
				Type:        "damage",
				Effect:      "High critical chance",
			},
			{
				Name:        "Cybernetic Repair",
				Description: "Repairs both organic and mechanical parts",
				Healing:     30 + (level * 3),
				SteamCost:   30,
				Cooldown:    5,
				Type:        "healing",
				Effect:      "Restores both health and steam",
			},
			{
				Name:        "Steam Overcharge",
				Description: "Temporarily enhances all attributes",
				SteamCost:   40,
				Cooldown:    8,
				Type:        "buff",
				Effect:      "Increases all attributes for 3 turns",
			},
		}...)

	case Elemental:
		abilities = append(abilities, []Ability{
			{
				Name:        "Steam Explosion",
				Description: "Creates a powerful steam burst",
				Damage:      25 + (level * 3),
				SteamCost:   30,
				Cooldown:    4,
				Type:        "damage",
				Effect:      "Area of effect damage",
			},
			{
				Name:        "Elemental Shield",
				Description: "Creates a protective steam barrier",
				Healing:     15 + (level * 2),
				SteamCost:   25,
				Cooldown:    5,
				Type:        "healing",
				Effect:      "Reduces incoming damage",
			},
			{
				Name:        "Steam Surge",
				Description: "Enhances steam-based abilities",
				SteamCost:   35,
				Cooldown:    6,
				Type:        "buff",
				Effect:      "Increases steam power for 4 turns",
			},
		}...)

	case Construct:
		abilities = append(abilities, []Ability{
			{
				Name:        "Heavy Slam",
				Description: "A powerful mechanical attack",
				Damage:      30 + (level * 3),
				SteamCost:   25,
				Cooldown:    4,
				Type:        "damage",
				Effect:      "Stuns target",
			},
			{
				Name:        "Reinforced Structure",
				Description: "Strengthens mechanical components",
				Healing:     20 + (level * 2),
				SteamCost:   20,
				Cooldown:    3,
				Type:        "healing",
				Effect:      "Increases defense",
			},
			{
				Name:        "Overclock",
				Description: "Temporarily increases combat capabilities",
				SteamCost:   45,
				Cooldown:    7,
				Type:        "buff",
				Effect:      "Increases damage and defense for 3 turns",
			},
		}...)
	}

	return abilities
}

// GetCharacterAbilities returns abilities for a character based on class and level
func GetCharacterAbilities(class Class, level int) []Ability {
	var abilities []Ability

	switch class {
	case Engineer:
		abilities = append(abilities, Ability{
			Name:        "Steam Blast",
			Description: "A powerful blast of steam",
			Type:        "damage",
			Damage:      15 + level,
			SteamCost:   20,
			Cooldown:    2,
		})
		if level >= 5 {
			abilities = append(abilities, Ability{
				Name:        "Repair Bot",
				Description: "Deploy a repair bot that heals over time",
				Type:        "healing",
				Healing:     10 + level*2,
				SteamCost:   30,
				Cooldown:    4,
				Effect:      "heal_over_time",
			})
		}
		if level >= 10 {
			abilities = append(abilities, Ability{
				Name:        "Steam Shield",
				Description: "Create a protective steam barrier",
				Type:        "buff",
				SteamCost:   40,
				Cooldown:    6,
				Effect:      "damage_reduction",
			})
		}

	case Alchemist:
		abilities = append(abilities, Ability{
			Name:        "Acid Splash",
			Description: "Throw a vial of corrosive acid",
			Type:        "damage",
			Damage:      12 + level,
			SteamCost:   15,
			Cooldown:    2,
			Effect:      "armor_reduction",
		})
		if level >= 5 {
			abilities = append(abilities, Ability{
				Name:        "Healing Vapor",
				Description: "Release healing steam vapors",
				Type:        "healing",
				Healing:     15 + level*2,
				SteamCost:   25,
				Cooldown:    3,
			})
		}
		if level >= 10 {
			abilities = append(abilities, Ability{
				Name:        "Toxic Cloud",
				Description: "Create a cloud of toxic steam",
				Type:        "debuff",
				Damage:      5 + level,
				SteamCost:   35,
				Cooldown:    5,
				Effect:      "poison",
			})
		}

	case Aeronaut:
		abilities = append(abilities, Ability{
			Name:        "Steam Propulsion",
			Description: "Use steam to propel yourself at the enemy",
			Type:        "damage",
			Damage:      18 + level,
			SteamCost:   25,
			Cooldown:    3,
		})
		if level >= 5 {
			abilities = append(abilities, Ability{
				Name:        "Aerial Maneuver",
				Description: "Quickly dodge attacks",
				Type:        "buff",
				SteamCost:   20,
				Cooldown:    4,
				Effect:      "dodge_chance",
			})
		}
		if level >= 10 {
			abilities = append(abilities, Ability{
				Name:        "Steam Cyclone",
				Description: "Create a powerful steam tornado",
				Type:        "damage",
				Damage:      25 + level*2,
				SteamCost:   45,
				Cooldown:    6,
				Effect:      "knockback",
			})
		}

	case ClockworkKnight:
		abilities = append(abilities, Ability{
			Name:        "Steam-powered Strike",
			Description: "A powerful melee attack enhanced by steam",
			Type:        "damage",
			Damage:      20 + level,
			SteamCost:   30,
			Cooldown:    2,
		})
		if level >= 5 {
			abilities = append(abilities, Ability{
				Name:        "Clockwork Armor",
				Description: "Enhance your armor with steam power",
				Type:        "buff",
				SteamCost:   35,
				Cooldown:    5,
				Effect:      "armor_increase",
			})
		}
		if level >= 10 {
			abilities = append(abilities, Ability{
				Name:        "Steam-powered Charge",
				Description: "Charge at the enemy with steam-enhanced speed",
				Type:        "damage",
				Damage:      30 + level*2,
				SteamCost:   50,
				Cooldown:    6,
				Effect:      "stun",
			})
		}

	case SteamMage:
		abilities = append(abilities, Ability{
			Name:        "Steam Bolt",
			Description: "Launch a bolt of condensed steam",
			Type:        "damage",
			Damage:      15 + level,
			SteamCost:   20,
			Cooldown:    2,
		})
		if level >= 5 {
			abilities = append(abilities, Ability{
				Name:        "Steam Shield",
				Description: "Create a protective barrier of steam",
				Type:        "buff",
				SteamCost:   30,
				Cooldown:    4,
				Effect:      "damage_reduction",
			})
		}
		if level >= 10 {
			abilities = append(abilities, Ability{
				Name:        "Steam Nova",
				Description: "Release a burst of steam in all directions",
				Type:        "damage",
				Damage:      20 + level*2,
				SteamCost:   45,
				Cooldown:    6,
				Effect:      "area_damage",
			})
		}
	}

	return abilities
}
