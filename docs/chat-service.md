# Chat Service Documentation

## Overview
The chat service provides real-time communication between players, including global chat, private messages, and team chat.

## Architecture

### Core Components

#### Chat System
- Real-time messaging using WebSocket
- Multiple chat channels
- Private messaging
- Team chat
- Global announcements

#### Message Management
- Message persistence
- Message history
- Message filtering
- Rate limiting

### Data Models

#### Message
```go
type Message struct {
    ID        string
    Channel   string
    Sender    string
    Content   string
    Type      string
    Timestamp time.Time
}
```

#### Channel
```go
type Channel struct {
    ID          string
    Name        string
    Type        string
    Members     []string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## API Endpoints

### WebSocket
- `ws://localhost:8082/ws` - WebSocket connection endpoint

### REST API
- `POST /api/messages` - Send a message
- `GET /api/channels/{id}/messages` - Get channel history
- `POST /api/channels` - Create a new channel
- `GET /api/channels` - List available channels
- `POST /api/channels/{id}/join` - Join a channel
- `POST /api/channels/{id}/leave` - Leave a channel

## WebSocket Protocol

### Connection
1. Client connects to WebSocket endpoint
2. Server authenticates connection using JWT
3. Server sends connection confirmation
4. Client can start sending/receiving messages

### Message Types

#### Client to Server
```json
{
    "type": "message",
    "channel": "channel_id",
    "content": "message_content"
}
```

#### Server to Client
```json
{
    "type": "message",
    "channel": "channel_id",
    "sender": "user_id",
    "content": "message_content",
    "timestamp": "2024-01-01T12:00:00Z"
}
```

### Channel Types
- Global: All players can access
- Team: Only team members can access
- Private: Only invited members can access
- System: System announcements

## Features

### Message Filtering
- Profanity filter
- Spam detection
- Rate limiting
- Content moderation

### Channel Management
- Channel creation
- Member management
- Channel moderation
- Channel settings

### Message History
- Configurable retention period
- Message search
- Message deletion
- Message editing

## Development

### Prerequisites
- Go 1.21 or later
- Redis (for pub/sub)
- MongoDB (for message storage)
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
export MONGODB_URI=mongodb://localhost:27017
export MONGODB_DB=chat_service
export SERVER_PORT=8082
export JWT_SECRET=your-secret-key
```

4. Run the server:
```bash
go run main.go
```

### Testing
Run the test suite:
```bash
go test ./...
```

### Building
Build the binary:
```bash
go build -o chat-service main.go
```

### Docker
Build and run with Docker:
```bash
docker build -t chat-service .
docker run -p 8082:8082 chat-service
```

## Integration

### Game Server Integration
The game server should:
1. Forward player events to chat service
2. Handle chat-related game events
3. Implement proper error handling

### Client Integration
The client should:
1. Establish WebSocket connection
2. Handle connection errors
3. Implement reconnection logic
4. Display messages in appropriate channels
5. Handle different message types
6. Implement proper error handling

## Security

### Authentication
- JWT-based authentication
- Token validation
- Session management

### Rate Limiting
- Message rate limiting
- Connection rate limiting
- Channel join rate limiting

### Content Security
- Message encryption
- Content filtering
- Spam prevention
- Moderation tools 