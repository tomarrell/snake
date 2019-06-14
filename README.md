# Snake
A highly parallel, abstract, ~~completely overengineered~~ snake game engine.

Simulate multiple snake games in parallel as a large set of finite state machines. Abstracted over directional input returning state changes.

Each game has a unique board size and tick rate. Allowing it to support games across varied clients.

## TODO
- [x] Engine creation
- [x] Game creation
- [x] Game ending
- [x] Snake logic
- [x] Input handling
- [ ] State output
- [ ] Fruit

## Overview
The core of the project can be found within `./engine`, with the clients occupying the sibling directories.

## License
Licensed under MIT or GPLv3.0, whichever you prefer.
