# Game Server Documentation

## Overview
The game server is the core service that handles game mechanics, combat, character management, and world interactions.

## Architecture

### Core Components

#### Character System
- Character creation and management
- Class-based abilities and stats
- Equipment and inventory management
- Experience and leveling system
- Steam power mechanics

#### Combat System
- Turn-based combat with steam power mechanics
- Multiple battle types (PvP, PvE, Raid)
- Terrain and weather effects
- Status effects and buffs/debuffs
- Team-based combat

#### World System
- Procedurally generated world map
- Dynamic weather and time of day
- Resource gathering and crafting
- Quest system
- NPC and mob interactions

### Data Models

#### Character
```go
type Character struct {
    ID                string
    Name              string
    Class             Class
    Level             int
    Experience        int
    Health            int
    MaxHealth         int
    SteamPower        int
    MaxSteamPower     int
    MovementPoints    int
    MaxMovementPoints int
    Position          common.Coordinates
    Inventory         []common.Item
    Equipment         map[string]common.Item
    Abilities         []common.Ability
    Stats             common.Stats
}
```

#### Mob
```go
type Mob struct {
    ID            string
    Name          string
    Type          string
    Level         int
    Health        int
    MaxHealth     int
    Attributes    common.Attributes
    SteamPower    int
    MaxSteamPower int
    Experience    int
    Abilities     []common.Ability
    LootTable     common.LootTable
    Location      common.Coordinates
    MoneyDrop     common.Currency
    Position      common.Coordinates
    Stats         common.Stats
}
```

#### Battle
```go
type Battle struct {
    ID            string
    Type          BattleType
    State         string
    Participants  []*Participant
    TurnOrder     []*Participant
    CurrentTurn   int
    Round         int
    CombatLog     []CombatLogEntry
    Terrain       string
    Weather       string
    Teams         map[string][]string
    StatusEffects map[string][]StatusEffect
    ActiveIndex   int
}
```

## API Endpoints

### Character Management
- `POST /api/character/create` - Create a new character
- `GET /api/character/{id}` - Get character details
- `PUT /api/character/{id}` - Update character
- `POST /api/character/{id}/equip` - Equip an item
- `POST /api/character/{id}/unequip` - Unequip an item

### Combat System
- `POST /api/combat/start` - Start a new battle
- `POST /api/combat/{id}/action` - Execute a combat action
- `GET /api/combat/{id}` - Get battle status
- `POST /api/combat/{id}/surrender` - Surrender from battle

### World System
- `GET /api/world/location/{id}` - Get location details
- `GET /api/world/region/{id}` - Get region details
- `POST /api/world/move` - Move character to new location
- `GET /api/world/nearby/{x}/{y}/{distance}` - Get nearby locations

## Game Mechanics

### Steam Power
- Regenerates over time
- Used for abilities and movement
- Affected by terrain and weather
- Can be enhanced by equipment

### Combat
- Turn-based system
- Initiative based on Dexterity and Steam Power
- Multiple ability types (mechanical, chemical, arcane)
- Area effects and status effects
- Team-based combat support

### Status Effects
- Steam Burn (damage over time)
- Poison (damage over time)
- Steam Weakness (reduced steam power)
- Mechanical Enhancement (attribute boost)
- Chemical Reaction (random effects)

### Terrain Effects
- Steam-rich (enhanced steam abilities)
- Mechanical (bonus for mechanical abilities)
- Toxic (bonus for chemical abilities)
- Forest (bonus for nature abilities)
- Mountain (bonus for defensive abilities)

### Weather Effects
- Steam Fog (reduced visibility)
- Acid Rain (damage to steam equipment)
- Strong Wind (affects steam regeneration)
- Clear (optimal conditions)

## Development

### Prerequisites
- Go 1.21 or later
- Redis server
- Docker (optional)

### Setup
1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
```bash
export REDIS_URL=localhost:6379
export SERVER_PORT=8080
```

4. Run the server:
```bash
go run cmd/server/main.go
```

### Testing
Run the test suite:
```bash
go test ./...
```

### Building
Build the binary:
```bash
go build -o game-server cmd/server/main.go
```

### Docker
Build and run with Docker:
```bash
docker build -t game-server .
docker run -p 8080:8080 game-server
``` 