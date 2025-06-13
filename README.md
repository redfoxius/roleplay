# Steam-Powered RPG

A steampunk-themed role-playing game with a focus on steam-powered combat and exploration.

## Features

### World Map
The game features a procedurally generated world map with various regions and locations:

#### Regions
- **Steam City**: A bustling metropolis powered by steam technology (Level 1)
- **Mountain Range**: A treacherous mountain range with valuable resources (Level 3)
- **Ancient Forest**: A mysterious forest with ancient technology (Level 2)
- **Steam Desert**: A scorching desert with steam-powered sandstorms (Level 4)
- **Steam Swamp**: A misty swamp with steam-powered research facilities (Level 3)
- **Steam Plains**: Vast plains dotted with steam-powered windmills (Level 1)

#### Location Types
- **Towns**: Safe havens with shops, inns, and workshops
- **Dungeons**: Dangerous areas with traps, treasure rooms, and boss encounters
- **Resource Nodes**: Rich sources of valuable materials and steam technology
- **Ruins**: Ancient remains of steam-powered civilizations
- **Normal Locations**: Typical areas with standard encounters

#### Terrain Types
- **Forest**: Dense forests with steam-powered trees and mechanical wildlife
- **Mountain**: Towering mountains with steam vents and mechanical caves
- **Plains**: Vast plains with steam-powered windmills and mechanical herds
- **Desert**: Scorching deserts with steam-powered sandstorms
- **Swamp**: Misty swamps with steam-powered research facilities
- **Steam City**: Bustling cities powered by steam technology

Each terrain type affects:
- Movement cost
- Steam power regeneration
- Available resources
- Enemy types
- Visibility

### Character Classes
- **Engineer**: Masters of mechanical technology and steam power
- **Alchemist**: Experts in chemical reactions and steam-based potions
- **Aeronaut**: Skilled pilots of steam-powered airships
- **Clockwork Knight**: Warriors enhanced by steam-powered armor
- **Steam Mage**: Wielders of arcane steam magic

### Combat System
The game features a turn-based combat system with:
- Steam power management
- Terrain and weather effects
- Status effects
- Team-based combat
- Multiple battle types (PvP, PvE, Raid)

#### Battle Types
- **PvP**: Player vs Player combat
- **PvE**: Player vs Environment combat
- **Raid**: Multiple players vs powerful enemies

#### Status Effects
- **Steam Burn**: Damage over time from steam-based attacks
- **Mechanical Enhancement**: Temporary boost to mechanical abilities
- **Chemical Reaction**: Random effects from chemical attacks
- **Arcane Binding**: Movement restriction from arcane attacks

#### Terrain Effects
- **Steam-rich**: Enhanced steam-based abilities
- **Mechanical**: Bonus for mechanical abilities
- **Toxic**: Bonus for chemical abilities
- **Forest**: Bonus for nature-based abilities
- **Mountain**: Bonus for defensive abilities
- **Plains**: Balanced combat environment
- **Desert**: Penalty to steam power regeneration
- **Swamp**: Penalty to movement speed
- **Steam City**: Bonus to all steam-based abilities

#### Weather Effects
- **Steam Fog**: Reduced visibility and ranged attack accuracy
- **Acid Rain**: Damage to steam-powered equipment
- **Strong Wind**: Affects steam power regeneration
- **Clear**: Optimal combat conditions

### World Systems
- Dynamic weather and time of day
- Resource gathering and crafting
- Steam power management
- Character progression
- Quest system

### Money Management
- Multiple currency types (Copper, Silver, Gold, Platinum)
- Trading system
- Item economy

## Docker Setup
The application can be run using Docker Compose, which sets up the game server, Redis database, and Flutter client.

### Docker Compose
The `docker-compose.yml` file defines the services:
- **Server**: Built from the current directory using `Dockerfile.server`, exposes port 8080.
- **Redis**: Uses the official Redis image, exposes port 6379.
- **Client**: Built from the `roleplay_client` directory using `Dockerfile.client`, exposes port 3000.

### Dockerfiles
- **Server**: Uses a multi-stage build with `golang:1.21` for building the Go binary and `alpine:latest` for running it.
- **Client**: Uses a multi-stage build with `ubuntu:latest` for building the Flutter app and `nginx:alpine` for serving it.

### Running the Application
To run the application, ensure Docker and Docker Compose are installed, then execute:
```bash
docker-compose up
```

## Installation

1. Clone the repository
2. Navigate to the project directory
3. Run `docker-compose up` to start the server, database, and client

## Game Mechanics

### Steam Power
Steam power is the primary resource in the game:
- Regenerates over time
- Affected by terrain and weather
- Used for abilities and movement
- Can be enhanced by equipment

### Terrain Effects
Different terrains provide various effects:
- Movement cost
- Steam power bonus/penalty
- Visibility range
- Resource availability
- Enemy types

### Weather Effects
Weather conditions affect gameplay:
- Steam fog reduces visibility
- Acid rain damages steam-powered equipment
- Strong winds affect steam power regeneration
- Clear weather provides optimal conditions

### Status Effects
Various status effects can be applied:
- Steam burn (damage over time)
- Mechanical enhancement (attribute boost)
- Chemical reaction (random effects)
- Arcane binding (movement restriction)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.