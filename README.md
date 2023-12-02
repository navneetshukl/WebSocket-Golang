# WebSocket Chat Server

This project is a simple chat server built with Go. It uses the Gorilla WebSocket package and the Jet template engine.

## Features

- **WebSocket Connection**: The server upgrades HTTP connections to WebSocket connections, allowing real-time, two-way communication between the server and the client.
- **Message Broadcasting**: Messages sent by a client are broadcasted to all connected clients.
- **User Management**: The server keeps track of connected users and broadcasts a list of users whenever a new user joins or an existing user leaves.

## Code Overview

The code is organized into several functions, each responsible for a specific task:

- `Home`: Renders the home page.
- `WsEndPoint`: Upgrades the HTTP connection to a WebSocket connection.
- `ListenForWs`: Listens for incoming messages from a WebSocket connection.
- `ListenToWsChannel`: Listens to the WebSocket channel for incoming payloads and handles them based on their action type.
- `getUserList`: Returns a list of connected users.
- `broadcastToAll`: Broadcasts a message to all connected clients.
- `renderPage`: Renders a specified page using the Jet template engine.

## Usage

To run the server, simply execute the main Go file. The server will start and listen for incoming connections.

Please note that this is a basic implementation of a WebSocket server and is not intended for production use. It's a great starting point if you're learning about WebSockets or building a real-time application with Go.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source, under the terms of the [MIT license].
