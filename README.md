# Steam-Powered RPG

A steampunk-themed role-playing game with a focus on steam-powered combat and exploration.

## Project Structure

```
.
├── services/
│   ├── game-server/        # Go-based game server
│   ├── auth-service/       # Authentication service
│   └── chat-service/       # Real-time chat service
├── roleplay_client/        # Flutter-based game client
├── docs/                   # Documentation
└── docker-compose.yml      # Docker configuration
```

## Features

### Character System
- Multiple character classes with unique abilities
- Steam power management
- Equipment and inventory system
- Experience and leveling
- Money management with multiple currencies

### Combat System
- Turn-based combat with steam power mechanics
- Multiple battle types (PvP, PvE, Raid)
- Terrain and weather effects
- Status effects and buffs/debuffs
- Team-based combat

### World System
- Procedurally generated world map
- Dynamic weather and time of day
- Resource gathering and crafting
- Quest system
- NPC and mob interactions

### Multiplayer Features
- Real-time chat system
- Player trading
- Team formation
- Guild system
- Global and local chat channels

## Getting Started

### Prerequisites
- Docker and Docker Compose
- Go 1.21 or later
- Flutter SDK (for client development)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/roleplay.git
cd roleplay
```

2. Start the services using Docker Compose:
```bash
docker-compose up
```

The following services will be available:
- Game Server: http://localhost:8080
- Auth Service: http://localhost:8081
- Chat Service: http://localhost:8082
- Game Client: http://localhost:3000

### Development Setup

1. Game Server:
```bash
cd services/game-server
go mod download
go run cmd/server/main.go
```

2. Auth Service:
```bash
cd services/auth-service
go mod download
go run main.go
```

3. Chat Service:
```bash
cd services/chat-service
go mod download
go run main.go
```

4. Game Client:
```bash
cd roleplay_client
flutter pub get
flutter run
```

## API Documentation

Detailed API documentation can be found in the [docs/api.md](docs/api.md) file.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.