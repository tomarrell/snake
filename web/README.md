# Web Adapter
The web adapater provides a TCP interface to the engine using websockets.

A single game can be created per connection. Attempts to create more than one game per connection will simply return the currently existing game.

The server will push `state` messages to the client when it computes a state change. This will be correlated to the tick rate of the game.

## API

Endpoint: `/ws`
Method: `GET`
Response `Upgrade to websocket connection`

## Messages


- `new`: Create a singleton game for connection and return the ID or return the ID of the existing game. This ID is a local identifier, and only used for confirmation.
```bash
snd > { "type": "new" }
rec > { "type": "id", "data": [int] }
```

- `input`: Send an input command to the game running on the current engine.
```bash
type direction = "left" | "right" | "up" | "down"

snd > { "type": "input", "direction": [direction] }
```

- `state`: Receive the game state from the server.
```bash
type Part = {
  x: int
  y: int
}

type Snake = {
  parts: []Part
}

type Fruit {
  value: int
  x:     int
  y:     int
}

type GameState = {
  width:  int
  height: int
  snake:  Snake
  fruit:  []Fruit
  score:  int
}

rcv > { "type": "state", "data": [GameState] }
```

## License
Licensed under MIT or GPLv3.0, whichever you prefer.
