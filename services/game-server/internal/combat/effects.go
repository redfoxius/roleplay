package combat

// Effect represents a status effect in combat
type Effect struct {
	Name        string
	Description string
	Type        string // buff, debuff, dot, hot
	Duration    int    // in turns
	Value       int    // effect value (damage, healing, stat modifier)
	Remaining   int    // remaining duration
}

// EffectManager manages active effects on a character or mob
type EffectManager struct {
	Effects map[string]*Effect
}

// NewEffectManager creates a new effect manager
func NewEffectManager() *EffectManager {
	return &EffectManager{
		Effects: make(map[string]*Effect),
	}
}

// AddEffect adds a new effect
func (em *EffectManager) AddEffect(effect *Effect) {
	em.Effects[effect.Name] = effect
}

// RemoveEffect removes an effect
func (em *EffectManager) RemoveEffect(name string) {
	delete(em.Effects, name)
}

// UpdateEffects updates all effects and returns any expired effects
func (em *EffectManager) UpdateEffects() []string {
	var expired []string
	for name, effect := range em.Effects {
		effect.Remaining--
		if effect.Remaining <= 0 {
			expired = append(expired, name)
			delete(em.Effects, name)
		}
	}
	return expired
}

// GetEffectValue returns the total value of effects of a specific type
func (em *EffectManager) GetEffectValue(effectType string) int {
	total := 0
	for _, effect := range em.Effects {
		if effect.Type == effectType {
			total += effect.Value
		}
	}
	return total
}

// Common effects
var (
	PoisonEffect = &Effect{
		Name:        "Poison",
		Description: "Takes damage over time",
		Type:        "dot",
		Duration:    3,
		Value:       5,
		Remaining:   3,
	}

	BurningEffect = &Effect{
		Name:        "Burning",
		Description: "Takes fire damage over time",
		Type:        "dot",
		Duration:    2,
		Value:       8,
		Remaining:   2,
	}

	HealingOverTimeEffect = &Effect{
		Name:        "Healing",
		Description: "Heals over time",
		Type:        "hot",
		Duration:    3,
		Value:       10,
		Remaining:   3,
	}

	StrengthBuffEffect = &Effect{
		Name:        "Strength Buff",
		Description: "Increased strength",
		Type:        "buff",
		Duration:    3,
		Value:       5,
		Remaining:   3,
	}

	WeaknessEffect = &Effect{
		Name:        "Weakness",
		Description: "Reduced strength",
		Type:        "debuff",
		Duration:    2,
		Value:       -3,
		Remaining:   2,
	}

	SteamPoweredEffect = &Effect{
		Name:        "Steam Powered",
		Description: "Increased steam power regeneration",
		Type:        "buff",
		Duration:    3,
		Value:       2,
		Remaining:   3,
	}

	SteamExhaustionEffect = &Effect{
		Name:        "Steam Exhaustion",
		Description: "Reduced steam power regeneration",
		Type:        "debuff",
		Duration:    2,
		Value:       -2,
		Remaining:   2,
	}
)
