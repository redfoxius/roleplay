# Environment Variables

This document describes the environment variables used in the Roleplay application.

## Root Configuration

These variables are used in the root `.env` file and affect the entire application:

```env
# Redis Configuration
REDIS_PASSWORD=change-in-production  # Redis password for authentication
REDIS_PORT=6379                     # Redis port number

# Service Ports
GAME_SERVER_PORT=8080               # Port for the game server
AUTH_SERVICE_PORT=8081              # Port for the authentication service
CHAT_SERVICE_PORT=8082              # Port for the chat service
CLIENT_PORT=3000                    # Port for the Flutter client

# Development Settings
DEBUG=true                          # Enable debug mode
NODE_ENV=development                # Node.js environment

# Security
JWT_SECRET=your-secret-key          # Secret key for JWT token generation
```

## Auth Service Configuration

These variables are used in the auth service's `.env` file:

```env
# Auth Service Configuration
PORT=8081                           # Service port
JWT_SECRET=your-secret-key          # Secret key for JWT token generation
REDIS_URL=redis:6379                # Redis connection URL

# Development Settings
DEBUG=true                          # Enable debug mode
CORS_ALLOWED_ORIGINS=http://localhost:3000  # Allowed CORS origins
```

## Chat Service Configuration

These variables are used in the chat service's `.env` file:

```env
# Chat Service Configuration
PORT=8082                           # Service port
REDIS_URL=redis:6379                # Redis connection URL

# Development Settings
DEBUG=true                          # Enable debug mode
CORS_ALLOWED_ORIGINS=http://localhost:3000  # Allowed CORS origins

# WebSocket Settings
WS_MAX_CONNECTIONS=1000             # Maximum number of WebSocket connections
WS_MESSAGE_SIZE_LIMIT=4096          # Maximum size of WebSocket messages in bytes
```

## Game Server Configuration

These variables are used in the game server's `.env` file:

```env
# Game Server Configuration
PORT=8080                           # Service port
REDIS_URL=redis:6379                # Redis connection URL

# Service URLs
AUTH_SERVICE_URL=http://auth-service:8081  # Authentication service URL
CHAT_SERVICE_URL=http://chat-service:8082  # Chat service URL

# Development Settings
DEBUG=true                          # Enable debug mode
CORS_ALLOWED_ORIGINS=http://localhost:3000  # Allowed CORS origins

# Game Settings
MAX_PLAYERS=100                     # Maximum number of players
WORLD_SIZE=1000                     # Size of the game world
TICK_RATE=60                        # Game tick rate (updates per second)
```

## Setting Up Environment Variables

1. Copy the example files to create your environment files:
   ```bash
   cp .env.example .env
   cp services/auth-service/.env.example services/auth-service/.env
   cp services/chat-service/.env.example services/chat-service/.env
   cp services/game-server/.env.example services/game-server/.env
   ```

2. Update the values in each `.env` file with your specific configuration.

3. For production:
   - Use strong, unique passwords
   - Set `DEBUG=false`
   - Set `NODE_ENV=production`
   - Use a strong `JWT_SECRET`
   - Configure proper `CORS_ALLOWED_ORIGINS`

## Docker Compose

The `docker-compose.yml` file uses these environment variables to configure the services. You can override any of these values by setting them in your `.env` file or by passing them directly to the `docker-compose` command.

Example:
```bash
# Override specific variables
REDIS_PASSWORD=my-secure-password docker-compose up

# Use a different environment file
docker-compose --env-file .env.production up
``` 