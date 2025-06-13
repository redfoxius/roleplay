package character

import "fmt"

// EquipmentSlot represents a slot where equipment can be worn
type EquipmentSlot string

const (
	Head      EquipmentSlot = "head"
	Chest     EquipmentSlot = "chest"
	Hands     EquipmentSlot = "hands"
	Legs      EquipmentSlot = "legs"
	Feet      EquipmentSlot = "feet"
	Weapon    EquipmentSlot = "weapon"
	OffHand   EquipmentSlot = "off_hand"
	Accessory EquipmentSlot = "accessory"
)

// Equipment represents a piece of equipment
type Equipment struct {
	ID            string
	Name          string
	Slot          EquipmentSlot
	Level         int
	Rarity        string
	Stats         map[string]int
	SteamPower    int
	Durability    int
	MaxDurability int
	Description   string
}

// EquipmentSet represents all equipped items
type EquipmentSet struct {
	Slots map[EquipmentSlot]*Equipment
}

// NewEquipmentSet creates a new empty equipment set
func NewEquipmentSet() *EquipmentSet {
	return &EquipmentSet{
		Slots: make(map[EquipmentSlot]*Equipment),
	}
}

// EquipItem equips an item in the appropriate slot
func (es *EquipmentSet) EquipItem(item *Equipment) (*Equipment, error) {
	// Check if slot is valid
	if item.Slot == "" {
		return nil, fmt.Errorf("invalid equipment slot")
	}

	// Store currently equipped item
	oldItem := es.Slots[item.Slot]

	// Equip new item
	es.Slots[item.Slot] = item

	return oldItem, nil
}

// UnequipItem removes an item from a slot
func (es *EquipmentSet) UnequipItem(slot EquipmentSlot) (*Equipment, error) {
	item, exists := es.Slots[slot]
	if !exists {
		return nil, fmt.Errorf("no item equipped in slot %s", slot)
	}

	delete(es.Slots, slot)
	return item, nil
}

// GetTotalStats returns the sum of all stats from equipped items
func (es *EquipmentSet) GetTotalStats() map[string]int {
	total := make(map[string]int)
	for _, item := range es.Slots {
		for stat, value := range item.Stats {
			total[stat] += value
		}
	}
	return total
}

// GetTotalSteamPower returns the total steam power bonus from equipment
func (es *EquipmentSet) GetTotalSteamPower() int {
	total := 0
	for _, item := range es.Slots {
		total += item.SteamPower
	}
	return total
}

// Common equipment types
var (
	// Head equipment
	SteamGoggles = &Equipment{
		ID:     "steam_goggles",
		Name:   "Steam Goggles",
		Slot:   Head,
		Level:  1,
		Rarity: "common",
		Stats: map[string]int{
			"TechnicalAptitude": 2,
		},
		SteamPower:    5,
		Durability:    100,
		MaxDurability: 100,
		Description:   "Basic steam-powered goggles",
	}

	// Chest equipment
	SteamVest = &Equipment{
		ID:     "steam_vest",
		Name:   "Steam Vest",
		Slot:   Chest,
		Level:  1,
		Rarity: "common",
		Stats: map[string]int{
			"Constitution": 3,
		},
		SteamPower:    10,
		Durability:    100,
		MaxDurability: 100,
		Description:   "A vest with steam-powered enhancements",
	}

	// Weapon equipment
	SteamPistol = &Equipment{
		ID:     "steam_pistol",
		Name:   "Steam Pistol",
		Slot:   Weapon,
		Level:  1,
		Rarity: "common",
		Stats: map[string]int{
			"MechanicalPrecision": 2,
			"SteamPower":          5,
		},
		SteamPower:    15,
		Durability:    100,
		MaxDurability: 100,
		Description:   "A basic steam-powered pistol",
	}
)
