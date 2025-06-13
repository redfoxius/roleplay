# Roleplay Game API Documentation

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

### Start Combat
```http
POST /api/combat/start
```

Initiates a new combat between multiple players.

**Request Body:**
```json
{
    "participants": ["string"] // Array of character IDs
}
```

**Response:**
```json
{
    "game_id": "string",
    "state": {
        "id": "string",
        "active_character": "string",
        "turn_order": ["string"],
        "current_turn": "integer",
        "round": "integer"
    },
    "error": "string" // Optional error message
}
```

### Start Mob Combat
```http
POST /api/mob-combat/start
```

Starts combat between a character and a mob.

**Request Body:**
```json
{
    "character_id": "string",
    "mob_id": "string"
}
```

**Response:**
```json
{
    "combat": {
        "character": {
            "id": "string",
            "name": "string",
            "health": "integer",
            "steam_power": "integer",
            "money": {
                "copper": "integer",
                "silver": "integer",
                "gold": "integer",
                "platinum": "integer"
            }
        },
        "mob": {
            "id": "string",
            "name": "string",
            "type": "string",
            "health": "integer",
            "steam_power": "integer",
            "money_drop": {
                "copper": "integer",
                "silver": "integer",
                "gold": "integer",
                "platinum": "integer"
            }
        },
        "turn": "integer"
    },
    "error": "string" // Optional error message
}
```

### Mob Combat Action
```http
POST /api/mob-combat/action
```

Performs a combat action between a character and mob.

**Request Body:**
```json
{
    "character_id": "string",
    "mob_id": "string",
    "ability": {
        "name": "string",
        "description": "string",
        "type": "string", // One of: "damage", "healing", "buff", "debuff"
        "damage": "integer",
        "steam_cost": "integer",
        "cooldown": "integer"
    }
}
```

**Response:**
```json
{
    "character_damage": "integer",
    "mob_damage": "integer",
    "money_gained": {
        "copper": "integer",
        "silver": "integer",
        "gold": "integer",
        "platinum": "integer"
    },
    "result": "string", // One of: "victory", "defeat", "ongoing"
    "error": "string" // Optional error message
}
```

## World Interaction

### Get Location
```http
GET /api/world/location?x={x}&y={y}
```

Gets information about a location in the world.

**Response:**
```json
{
    "location": {
        "name": "string",
        "description": "string",
        "region": "string",
        "terrain": "string",
        "coordinates": {
            "x": "integer",
            "y": "integer"
        }
    }
}
```

### Get Nearby Locations
```http
GET /api/world/nearby?x={x}&y={y}&distance={distance}
```

Gets all locations within a certain distance of the given coordinates.

**Response:**
```json
{
    "locations": [
        {
            "name": "string",
            "description": "string",
            "region": "string",
            "terrain": "string",
            "coordinates": {
                "x": "integer",
                "y": "integer"
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
    "error": "string" // Optional error message
}
```

### Get Nearby Mobs
```http
GET /api/world/nearby-mobs?x={x}&y={y}&distance={distance}
```

Gets all mobs within a certain distance of the given coordinates.

**Response:**
```json
{
    "mobs": [
        {
            "id": "string",
            "name": "string",
            "type": "string",
            "level": "integer",
            "health": "integer",
            "steam_power": "integer",
            "money_drop": {
                "copper": "integer",
                "silver": "integer",
                "gold": "integer",
                "platinum": "integer"
            },
            "location": {
                "x": "integer",
                "y": "integer"
            }
        }
    ]
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
``` 