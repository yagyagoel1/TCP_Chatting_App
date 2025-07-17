# TCP Chatting App

A simple TCP-based chat server written in Go. This server allows multiple clients to connect over TCP and exchange messages either via broadcast or by sending messages directly to a specific client.

---

## Features

* **Broadcast messaging**: Clients can broadcast messages to all connected clients.
* **Direct messaging**: Clients can send messages to a specific connected client by their assigned ID.
* **Random client IDs**: Each client is assigned a unique random ID upon connection.
* **Concurrent handling**: Uses Go routines to handle multiple clients concurrently.

---

## Requirements

* Go 1.22 or later

---

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yagyagoel1/TCP_Chatting_App.git
   cd TCP_Chatting_App
   ```

2. **Build the server**

   ```bash
   go build -o tcp_chat_server ./cmd
   ```

---

## Usage

1. **Run the server**

   ```bash
   ./tcp_chat_server
   ```

   By default, the server listens on `localhost:8080`. You should see:

   ```
   Server is listening on port 8080
   ```

2. **Connect clients**

   You can use any TCP client (e.g., `netcat`, `telnet`) to connect:

   ```bash
   nc localhost 8080
   ```

3. **Communication protocol**

   The server expects a simple line-based protocol:

   * **Broadcast**

     1. Send the command `broadcast` followed by a newline (`\n`).
     2. Send your message followed by a newline.

     ```text
     broadcast
     Hello everyone!
     ```

   * **Direct message**

     1. Send the command `send` followed by a newline.
     2. Send the recipient's client ID followed by a newline.
     3. Send your message followed by a newline.

     ```text
     send
     123456
     Hey, how are you?
     ```

4. **Example session**

   ```bash
   # Client A connects
   $ nc localhost 8080
   # (Assigned ID: 345678)
   broadcast
   Hi there!

   # Client B connects
   $ nc localhost 8080
   # (Assigned ID: 123456)
   send
   345678
   Hello back!
   ```

---

## Project Structure

```
TCP_Chatting_App/
├── cmd/
│   └── main.go        # Entry point for the server
└── go.mod             # Module definition
```

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contact

Created by [yagyagoel1](https://github.com/yagyagoel1). Feel free to open an issue for questions or feature requests.
