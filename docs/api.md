# API Documentation

## Base URL
All API endpoints are prefixed with `/api`

## Authentication
All requests require a valid JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

## Character Management

### Create Character
```http
POST /character/create
```

Request:
```json
{
    "name": "string",
    "class": "string",
    "attributes": {
        "strength": 10,
        "intelligence": 10,
        "dexterity": 10,
        "vitality": 10
    }
}
```

Response:
```json
{
    "id": "string",
    "name": "string",
    "class": "string",
    "level": 1,
    "experience": 0,
    "attributes": {
        "strength": 10,
        "intelligence": 10,
        "dexterity": 10,
        "vitality": 10
    },
    "equipment": {},
    "inventory": [],
    "money": {
        "copper": 0,
        "silver": 0,
        "gold": 0,
        "platinum": 0
    }
}
```

### Get Character
```http
GET /character/{id}
```

Response:
```json
{
    "id": "string",
    "name": "string",
    "class": "string",
    "level": 1,
    "experience": 0,
    "attributes": {
        "strength": 10,
        "intelligence": 10,
        "dexterity": 10,
        "vitality": 10
    },
    "equipment": {},
    "inventory": [],
    "money": {
        "copper": 0,
        "silver": 0,
        "gold": 0,
        "platinum": 0
    }
}
```

### Update Character
```http
PUT /character/{id}
```

Request:
```json
{
    "name": "string",
    "attributes": {
        "strength": 10,
        "intelligence": 10,
        "dexterity": 10,
        "vitality": 10
    }
}
```

Response:
```json
{
    "id": "string",
    "name": "string",
    "class": "string",
    "level": 1,
    "experience": 0,
    "attributes": {
        "strength": 10,
        "intelligence": 10,
        "dexterity": 10,
        "vitality": 10
    },
    "equipment": {},
    "inventory": [],
    "money": {
        "copper": 0,
        "silver": 0,
        "gold": 0,
        "platinum": 0
    }
}
```

## Combat System

### Start Battle
```http
POST /combat/start
```

Request:
```json
{
    "type": "string",
    "participants": [
        {
            "id": "string",
            "type": "string"
        }
    ],
    "terrain": "string",
    "weather": "string"
}
```

Response:
```json
{
    "id": "string",
    "type": "string",
    "state": "string",
    "currentTurn": "string",
    "participants": [
        {
            "id": "string",
            "type": "string",
            "health": 100,
            "steamPower": 100,
            "status": []
        }
    ],
    "terrain": "string",
    "weather": "string"
}
```

### Execute Action
```http
POST /combat/{id}/action
```

Request:
```json
{
    "type": "string",
    "target": "string",
    "ability": "string"
}
```

Response:
```json
{
    "id": "string",
    "type": "string",
    "state": "string",
    "currentTurn": "string",
    "participants": [
        {
            "id": "string",
            "type": "string",
            "health": 100,
            "steamPower": 100,
            "status": []
        }
    ],
    "terrain": "string",
    "weather": "string",
    "lastAction": {
        "type": "string",
        "source": "string",
        "target": "string",
        "ability": "string",
        "damage": 0,
        "effects": []
    }
}
```

## World System

### Get Location
```http
GET /world/location/{id}
```

Response:
```json
{
    "id": "string",
    "name": "string",
    "type": "string",
    "terrain": "string",
    "weather": "string",
    "resources": [],
    "npcs": [],
    "mobs": []
}
```

### Get Region
```http
GET /world/region/{id}
```

Response:
```json
{
    "id": "string",
    "name": "string",
    "level": 1,
    "locations": [
        {
            "id": "string",
            "name": "string",
            "type": "string"
        }
    ]
}
```

## Chat System

### Send Message
```http
POST /chat/message
```

Request:
```json
{
    "channel": "string",
    "content": "string"
}
```

Response:
```json
{
    "id": "string",
    "channel": "string",
    "sender": "string",
    "content": "string",
    "timestamp": "string"
}
```

### Get Channel History
```http
GET /chat/channel/{id}/history
```

Response:
```json
{
    "messages": [
        {
            "id": "string",
            "channel": "string",
            "sender": "string",
            "content": "string",
            "timestamp": "string"
        }
    ]
}
```

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
    "error": "string",
    "message": "string"
}
```

### 401 Unauthorized
```json
{
    "error": "Unauthorized",
    "message": "Invalid or missing authentication token"
}
```

### 403 Forbidden
```json
{
    "error": "Forbidden",
    "message": "Insufficient permissions"
}
```

### 404 Not Found
```json
{
    "error": "Not Found",
    "message": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
    "error": "Internal Server Error",
    "message": "An unexpected error occurred"
} 