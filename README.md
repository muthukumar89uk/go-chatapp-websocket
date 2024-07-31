# Go Chat App Using WebSocket

This is a simple chat application built using Go and WebSockets. The application allows multiple users to send messages to each other in real-time.

## Features

- Real-time messaging using WebSockets
- Broadcast messages to all connected users

## Project Structure
The project is organized into several packages, each responsible for specific functionalities:
- `handlers`  : Contains the HTTP request handlers for different API endpoints.
- `config`    : Contains basic configuration for the Database.
- `models`    : Defines the data models used in the application.
- `drivers`   : Contains functions for establish a connection to database.
- `helpers`   : Custom package that contains all the constants.

## Getting Started

### Installation

1. Clone the repository Or Download:

   ```
   git clone https://github.com/muthukumar89uk/go-chatapp-websocket.git
Click here to directly [download it](https://github.com/muthukumar89uk/go-chatapp-websocket/zipball/master).

### Install dependencies:

      go mod tidy

### Run the Application
 1. Run the Server:
   ```
      go run .
   ```
 2. Create the users using the `http://localhost:8080/user` URL.
 3. Websocket server will start on `ws://localhost:8080/ws` URL.
