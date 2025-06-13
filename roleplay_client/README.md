# Roleplay Game Flutter Client

## Overview
The Flutter client for the Roleplay game provides a user interface for interacting with the game server. It is designed to be run using Docker, alongside the game server and Redis database.

## Features
- **Character Creation**: Create and manage your character.
- **World Map**: Explore the dynamic world map with various regions and locations.
- **Combat System**: Engage in turn-based combat with unique abilities and effects.
- **Resource Management**: Manage your character's resources and inventory.

## Installation
1. Ensure Docker and Docker Compose are installed on your system.
2. Clone the repository and navigate to the `roleplay_client` directory.

## Running the Application
To run the Flutter client using Docker, execute the following command from the project root:
```bash
docker-compose up
```
This will start the game server, Redis database, and Flutter client. The client will be accessible at `http://localhost:3000`.

## Development
- **Flutter**: The client is built using Flutter, a UI toolkit for building natively compiled applications.
- **Docker**: The client is containerized using Docker, ensuring consistent deployment across different environments.

## Contributing
Contributions are welcome! Please read the contributing guidelines before submitting a pull request.

## License
This project is licensed under the MIT License.
