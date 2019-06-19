# Snake
A highly parallel, abstract, ~~completely overengineered~~ snake game engine.

Simulate multiple snake games in parallel as a large set of finite state machines. Abstracted over directional input returning state changes.

Each game has a unique board size and tick rate. Allowing it to support games across varied clients using custom built adapters for the given interface.

Demo of the term-snake adpater:
![Term-snake](./images/term_snake_demo.gif)

## Installation
You can build your own adapter by importing `github.com/tomarrell/snake` and then using the exported methods there to build your own adapter on top of the simulation.

Alternatively, you can use one of the pre-built adapters. Currently the supported clients are:
- Terminal
- Web [In Progress]

In order to play around with the terminal adapter, run the following.
```bash
> go get github.com/tomarrell/snake/term-snake
> term-snake
```
Just make sure you have your `$GOPATH` setup and `$PATH` pointing to `$GOPATH/bin`.

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
