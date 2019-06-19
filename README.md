# Snake
A highly parallel, abstract, ~~completely overengineered~~ snake game engine.

Simulate multiple snake games in parallel as a large set of finite state machines. Abstracted over directional input returning state changes.

Each game has a unique board size and tick rate. Allowing it to support games across varied clients using custom built adapters for the given interface.

## TODO

#### Core
- [x] Engine creation
- [x] Game creation
- [x] Game ending
- [x] Snake logic
- [x] Input handling
- [x] State output
- [x] Fruit
- [ ] Snake collisions with itself

#### Adapters
- [x] Basic terminal interface
- [ ] Web server
- [ ] Leaderboard


## Overview
The core of the project can be found within `./engine`, with the adapters occupying the sibling directories. This models allows for arbitrary clients to interface with the engine and have it support multiple mediums.

## License
Licensed under MIT or GPLv3.0, whichever you prefer.
