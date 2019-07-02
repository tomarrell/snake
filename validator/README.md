# Validation Adapter
The validation adapter provides an API to allow for the client to be run simultaneously on the client, while sending information back to the server at specific times for validation.

This allows the server to remain completely stateless across each game, while also preventing the possibility of the client falsly reporting a score, or any other game parameters in order to gain an advantage.

## API
The API is implemented as a series of REST endpoints listed below.

- `/new`: generates a new game with randomized fruit location.
```golang
type Part struct {
  x int
  y int
}

type Snake struct {
  velX int
  velY int
  parts []Part
}

>> JSON Request Payload:
{
  width: int, // the number of grid squares the arena is wide
  height: int, // the number of grid squares the arena is high
  snake: Snake, // the starting position of the snake
}

<< JSON Response Payload:
{
  gameId: string,
  width: int,
  height: int,
  score: int,
  fruit: []engine.Fruit, // two new generated piece of fruit
	snake: Snake, // the starting position of the snake
  signature: string,
}
```

The client receives the gameID, the current score, and the array of fruit. It also receives a signature that it can validate from the server. It then records the velX and velY of the snake at each tick of the game until it reaches the next fruit. At which point, it calls the endpoint below.

- `validate`: validates given a set of ticks, that the snake reached the location of a fruit.
```golang
>> JSON Request Payload
{
  gameId: string,
  width: int,
  height: int,
  score: int, // the previously signed score
  snake: Snake, // the previously signed position of the snake
  fruit: []engine.Fruit, // the previously signed position of the fruit
  signature: string, // the most recent signature, corresponding to gameID, width, height score, snake and fruit
  ticks: [ // in the order that they occurred since the last fruit was eaten
    { velX: int, velY: int },
    { velX: int, velY: int },
    ...
  ],
}

<< JSON Response Payload
{
  gameId: string,
  width: int,
  height: int,
  score: int,
  fruit: []Fruit, // fruit contains a new piece of fruit, replacing the one that was eaten
  snake: Snake, // the validated position of the snake
  signature: string, // a new signature for the validated state
}
```
