# Authentication Service Documentation

## Overview
The authentication service handles user authentication, authorization, and session management for the game server.

## Architecture

### Core Components

#### User Management
- User registration and login
- Password hashing and validation
- Session management
- Role-based access control

#### Security
- JWT token generation and validation
- Password encryption
- Rate limiting
- IP blocking

### Data Models

#### User
```go
type User struct {
    ID        string
    Username  string
    Email     string
    Password  string
    Role      string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

#### Session
```go
type Session struct {
    ID        string
    UserID    string
    Token     string
    ExpiresAt time.Time
    CreatedAt time.Time
}
```

## API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login user
- `POST /api/auth/logout` - Logout user
- `POST /api/auth/refresh` - Refresh authentication token

### User Management
- `GET /api/users/{id}` - Get user details
- `PUT /api/users/{id}` - Update user details
- `DELETE /api/users/{id}` - Delete user
- `GET /api/users/{id}/sessions` - Get user sessions

### Session Management
- `GET /api/sessions/{id}` - Get session details
- `DELETE /api/sessions/{id}` - Invalidate session
- `GET /api/sessions/active` - Get active sessions

## Security

### Authentication Flow
1. User submits credentials
2. Service validates credentials
3. Service generates JWT token
4. Token is returned to client
5. Client includes token in subsequent requests

### Token Structure
```json
{
    "header": {
        "alg": "HS256",
        "typ": "JWT"
    },
    "payload": {
        "sub": "user_id",
        "role": "user_role",
        "exp": "expiration_time",
        "iat": "issued_at"
    }
}
```

### Password Security
- Passwords are hashed using bcrypt
- Salt is automatically generated
- Minimum password requirements:
  - 8 characters minimum
  - At least one uppercase letter
  - At least one lowercase letter
  - At least one number
  - At least one special character

## Development

### Prerequisites
- Go 1.21 or later
- PostgreSQL
- Redis (for session management)
- Docker (optional)

### Setup
1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=auth_service
export REDIS_URL=localhost:6379
export JWT_SECRET=your-secret-key
export SERVER_PORT=8081
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
go build -o auth-service main.go
```

### Docker
Build and run with Docker:
```bash
docker build -t auth-service .
docker run -p 8081:8081 auth-service
```

## Integration

### Game Server Integration
The game server should:
1. Include JWT token in Authorization header
2. Handle 401 Unauthorized responses
3. Refresh tokens when expired
4. Implement proper error handling

### Client Integration
The client should:
1. Store JWT token securely
2. Include token in all API requests
3. Handle token expiration
4. Implement proper error handling
5. Provide login/logout functionality 