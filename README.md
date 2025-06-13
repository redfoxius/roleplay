# Roleplay Game

A steampunk-themed roleplaying game with a focus on character development, combat, and world exploration.

## Features

### Character System
- Multiple character classes: Engineer, Alchemist, Aeronaut, Clockwork Knight, and Steam Mage
- Attribute-based progression system
- Steam power management
- Equipment and inventory system
- Money system with multiple denominations (Copper, Silver, Gold, Platinum)

### Combat System
- Unified battle system for both PvP and PvE combat
- Turn-based combat mechanics with initiative-based order
- Ability-based combat actions with cooldowns
- Steam power consumption for abilities
- Experience and leveling system
- Money drops from defeated enemies
- Status effects system with duration-based effects
- Terrain and weather effects on combat
- Team-based combat support
- Area of effect abilities
- Combat logging system
- Battle rewards distribution

### Battle Types
- PvP (Player vs Player)
- PvE (Player vs Environment)
- Raid (Multiple players vs powerful enemies)

### Status Effects
- Duration-based effects
- Class-specific effects:
  - Engineer: Steam Burn
  - Alchemist: Poison
  - Steam Mage: Steam Weakness
- Effects can modify:
  - Health
  - Steam Power
  - Attributes
  - Combat effectiveness

### Terrain Effects
- Steam-rich: Enhanced steam-based abilities
- Mechanical: Bonus for mechanical abilities
- Toxic: Bonus for chemical abilities
- Affects damage and healing calculations

### Weather Effects
- Steam-fog: Penalty to ranged abilities
- Acid-rain: General combat penalty
- Affects ability effectiveness

### World System
- Dynamic world map with different regions
- Terrain-based movement costs
- Steam power regeneration based on terrain
- Mob spawning system with different types:
  - Mechanical (Steam-powered constructs)
  - Biological (Mutated creatures)
  - Hybrid (Steam-cyborgs)
  - Elemental (Steam and metal elementals)
  - Construct (Massive mechanical beings)

### Money System
- Four denominations: Copper, Silver, Gold, and Platinum
- Automatic conversion between denominations
- Money drops from defeated mobs based on:
  - Mob level
  - Mob type (different multipliers)
  - Random variation
- Money transfer between characters
- Starting money: 100 copper for new characters

## API Documentation

The game provides a RESTful API for all interactions. See [API Documentation](docs/api.md) for detailed endpoint information.

### Key Endpoints
- Character Management (`/api/character/*`)
- Combat System (`/api/combat/*`, `/api/mob-combat/*`)
- World Interaction (`/api/world/*`)
- Money Management (`/api/character/{id}/money`, `/api/character/transfer-money`)

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Redis server

### Installation
1. Clone the repository
```bash
git clone https://github.com/yourusername/roleplay.git
cd roleplay
```

2. Install dependencies
```bash
go mod download
```

3. Start Redis server
```bash
redis-server
```

4. Run the game server
```bash
go run cmd/server/main.go
```

## Game Mechanics

### Character Creation
- Choose from 5 unique classes
- Each class has specific attribute bonuses
- Starting equipment based on class
- 100 copper starting money

### Combat
- Turn-based system with initiative
- Steam power management
- Ability cooldowns
- Money and experience rewards

### World Exploration
- Terrain affects movement and steam power
- Different regions with unique properties
- Dynamic mob spawning
- Money drops scale with mob difficulty

### Money System
- Automatic denomination conversion
- Safe money transfer between characters
- Mob-specific money drop tables
- Money can be used for:
  - Purchasing items
  - Trading with other players
  - Upgrading equipment
  - Learning new abilities

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.