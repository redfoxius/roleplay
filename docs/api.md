# API Documentation

## Overview
This document outlines the API endpoints for the Roleplay game server.

## Docker Setup
The application can be run using Docker Compose, which sets up the game server, Redis database, and Flutter client.

### Docker Compose
The `docker-compose.yml` file defines the services:
- **Server**: Built from the current directory using `Dockerfile.server`, exposes port 8080.
- **Redis**: Uses the official Redis image, exposes port 6379.
- **Client**: Built from the `roleplay_client` directory using `Dockerfile.client`, exposes port 3000.

### Running the Application
To run the application, ensure Docker and Docker Compose are installed, then execute:
```bash
docker-compose up
```

## API Endpoints

## Character Management

### Create Character
```http
POST /api/character/create
```

Creates a new character with the specified name and class.

**Request Body:**
```json
{
    "name": "string",
    "class": "string" // One of: "Engineer", "Alchemist", "Aeronaut", "ClockworkKnight", "SteamMage"
}
```

**Response:**
```json
{
    "character": {
        "id": "string",
        "name": "string",
        "class": "string",
        "level": "integer",
        "health": "integer",
        "steam_power": "integer",
        "money": {
            "copper": "integer",
            "silver": "integer",
            "gold": "integer",
            "platinum": "integer"
        },
        "attributes": {
            "strength": { "name": "string", "value": "integer" },
            "dexterity": { "name": "string", "value": "integer" },
            "constitution": { "name": "string", "value": "integer" },
            "intelligence": { "name": "string", "value": "integer" },
            "wisdom": { "name": "string", "value": "integer" },
            "charisma": { "name": "string", "value": "integer" },
            "technical_aptitude": { "name": "string", "value": "integer" },
            "steam_power": { "name": "string", "value": "integer" },
            "mechanical_precision": { "name": "string", "value": "integer" },
            "arcane_knowledge": { "name": "string", "value": "integer" }
        }
    },
    "error": "string" // Optional error message
}
```

### Get Character Money
```http
GET /api/character/{id}/money
```

Gets the current money balance for a character.

**Response:**
```json
{
    "money": {
        "copper": "integer",
        "silver": "integer",
        "gold": "integer",
        "platinum": "integer"
    }
}
```

### Transfer Money
```http
POST /api/character/transfer-money
```

Transfers money between characters.

**Request Body:**
```json
{
    "from_character_id": "string",
    "to_character_id": "string",
    "amount": {
        "copper": "integer",
        "silver": "integer",
        "gold": "integer",
        "platinum": "integer"
    }
}
```

**Response:**
```json
{
    "success": "boolean",
    "error": "string" // Optional error message
}
```

## Combat

### Start Battle
```http
POST /api/battle/start
```

Initiates a new battle with the specified type and participants.

**Request Body:**
```json
{
    "type": "string", // One of: "pvp", "pve", "raid"
    "participants": ["string"], // Array of character/mob IDs
    "teams": { // Optional, for team battles
        "team1": ["string"],
        "team2": ["string"]
    },
    "terrain": "string", // One of: "neutral", "steam-rich", "mechanical", "toxic", "forest", "mountain", "plains", "desert", "swamp", "steam-city"
    "weather": "string"  // One of: "clear", "steam-fog", "acid-rain", "strong-wind"
}
```

**Response:**
```json
{
    "battle_id": "string",
    "battle": {
        "id": "string",
        "type": "string",
        "state": "string", // "active", "completed", "cancelled"
        "participants": [
            {
                "id": "string",
                "name": "string",
                "type": "string", // "player" or "mob"
                "class": "string",
                "team": "string",
                "health": "integer",
                "steam_power": "integer",
                "attributes": {
                    "strength": { "name": "string", "value": "integer" },
                    "dexterity": { "name": "string", "value": "integer" },
                    "constitution": { "name": "string", "value": "integer" },
                    "intelligence": { "name": "string", "value": "integer" },
                    "wisdom": { "name": "string", "value": "integer" },
                    "charisma": { "name": "string", "value": "integer" },
                    "technical_aptitude": { "name": "string", "value": "integer" },
                    "steam_power": { "name": "string", "value": "integer" },
                    "mechanical_precision": { "name": "string", "value": "integer" },
                    "arcane_knowledge": { "name": "string", "value": "integer" }
                },
                "abilities": [
                    {
                        "name": "string",
                        "description": "string",
                        "type": "string",
                        "damage": "integer",
                        "healing": "integer",
                        "steam_cost": "integer",
                        "range": "integer",
                        "area": "integer",
                        "cooldown": "integer"
                    }
                ],
                "is_active": "boolean",
                "experience": "integer",
                "money": {
                    "copper": "integer",
                    "silver": "integer",
                    "gold": "integer",
                    "platinum": "integer"
                }
            }
        ],
        "turn_order": ["string"],
        "current_turn": "integer",
        "round": "integer",
        "terrain": "string",
        "weather": "string",
        "teams": {
            "team1": ["string"],
            "team2": ["string"]
        },
        "status_effects": {
            "participant_id": [
                {
                    "name": "string",
                    "duration": "integer",
                    "value": "integer",
                    "description": "string"
                }
            ]
        },
        "active_index": "integer" // Index of the active participant in turn order
    },
    "error": "string" // Optional error message
}
```

### Execute Battle Action
```http
POST /api/battle/{battle_id}/action
```

Executes a combat action in the battle.

**Request Body:**
```json
{
    "ability": {
        "name": "string",
        "description": "string",
        "type": "string", // One of: "mechanical", "chemical", "arcane", "ranged", "melee"
        "damage": "integer",
        "healing": "integer",
        "steam_cost": "integer",
        "range": "integer",
        "area": "integer",
        "cooldown": "integer"
    },
    "target_id": "string"
}
```

**Response:**
```json
{
    "damage": "integer",
    "healing": "integer",
    "status_effects": [
        {
            "name": "string",
            "duration": "integer",
            "value": "integer",
            "description": "string"
        }
    ],
    "combat_log": {
        "round": "integer",
        "turn": "integer",
        "character": "string",
        "action": "string",
        "target": "string",
        "damage": "integer",
        "healing": "integer",
        "effects": ["string"]
    },
    "battle_state": "string", // "active", "completed", "cancelled"
    "error": "string" // Optional error message
}
```

### Get Battle Status
```http
GET /api/battle/{battle_id}
```

Gets the current status of a battle.

**Response:**
```json
{
    "battle": {
        "id": "string",
        "type": "string",
        "state": "string",
        "participants": ["string"],
        "turn_order": ["string"],
        "current_turn": "integer",
        "round": "integer",
        "terrain": "string",
        "weather": "string",
        "teams": {
            "team1": ["string"],
            "team2": ["string"]
        },
        "status_effects": {
            "participant_id": [
                {
                    "name": "string",
                    "duration": "integer",
                    "value": "integer",
                    "description": "string"
                }
            ]
        },
        "combat_log": [
            {
                "round": "integer",
                "turn": "integer",
                "character": "string",
                "action": "string",
                "target": "string",
                "damage": "integer",
                "healing": "integer",
                "effects": ["string"]
            }
        ]
    }
}
```

## World Map

### Get World Map
```http
GET /api/world/map
```

Returns the current state of the world map.

**Response:**
```json
{
    "width": "integer",
    "height": "integer",
    "regions": [
        {
            "name": "string",
            "level": "integer",
            "terrain": "string",
            "bounds": {
                "min_x": "integer",
                "min_y": "integer",
                "max_x": "integer",
                "max_y": "integer"
            }
        }
    ],
    "locations": [
        {
            "id": "string",
            "name": "string",
            "type": "string", // One of: "normal", "dungeon", "town", "resource", "ruin"
            "description": "string",
            "terrain": "string",
            "region": "string",
            "coordinates": {
                "x": "integer",
                "y": "integer"
            },
            "properties": {
                "difficulty": "integer",
                "steam_power_bonus": "integer",
                "resources": ["string"],
                "mob_types": ["string"],
                "special_features": ["string"]
            }
        }
    ]
}
```

### Get Location
```http
GET /api/world/location/{x}/{y}
```

Returns information about a specific location on the map.

**Response:**
```json
{
    "id": "string",
    "name": "string",
    "type": "string",
    "description": "string",
    "terrain": "string",
    "region": "string",
    "coordinates": {
        "x": "integer",
        "y": "integer"
    },
    "properties": {
        "difficulty": "integer",
        "steam_power_bonus": "integer",
        "resources": ["string"],
        "mob_types": ["string"],
        "special_features": ["string"]
    }
}
```

### Get Nearby Locations
```http
GET /api/world/nearby/{x}/{y}/{distance}
```

Returns all locations within the specified distance of the given coordinates.

**Response:**
```json
{
    "locations": [
        {
            "id": "string",
            "name": "string",
            "type": "string",
            "description": "string",
            "terrain": "string",
            "region": "string",
            "coordinates": {
                "x": "integer",
                "y": "integer"
            },
            "properties": {
                "difficulty": "integer",
                "steam_power_bonus": "integer",
                "resources": ["string"],
                "mob_types": ["string"],
                "special_features": ["string"]
            }
        }
    ]
}
```

### Get Nearby Locations by Type
```http
GET /api/world/nearby/{type}/{x}/{y}/{distance}
```

Returns all locations of a specific type within the specified distance.

**Parameters:**
- `type`: One of: "dungeon", "town", "resource", "ruin"

**Response:**
```json
{
    "locations": [
        {
            "id": "string",
            "name": "string",
            "type": "string",
            "description": "string",
            "terrain": "string",
            "region": "string",
            "coordinates": {
                "x": "integer",
                "y": "integer"
            },
            "properties": {
                "difficulty": "integer",
                "steam_power_bonus": "integer",
                "resources": ["string"],
                "mob_types": ["string"],
                "special_features": ["string"]
            }
        }
    ]
}
```

### Move Character
```http
POST /api/world/move
```

Moves a character to a new location.

**Request Body:**
```json
{
    "character_id": "string",
    "x": "integer",
    "y": "integer"
}
```

**Response:**
```json
{
    "success": "boolean",
    "new_location": {
        "id": "string",
        "name": "string",
        "type": "string",
        "description": "string",
        "terrain": "string",
        "region": "string",
        "coordinates": {
            "x": "integer",
            "y": "integer"
        },
        "properties": {
            "difficulty": "integer",
            "steam_power_bonus": "integer",
            "resources": ["string"],
            "mob_types": ["string"],
            "special_features": ["string"]
        }
    },
    "steam_power_cost": "integer",
    "steam_power_bonus": "integer",
    "error": "string" // Optional error message
}
```

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
    "error": "string" // Description of the error
}
```

### 404 Not Found
```json
{
    "error": "string" // Description of the error
}
```

### 500 Internal Server Error
```json
{
    "error": "string" // Description of the error
} 