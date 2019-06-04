# Snake
Snake API powering my personal website's leaderboard.


## The problem...
Currently, the game is played entirely client side. This means that the client is in charge of keeping score. If we were to simply relay this value back to the server, it would open up the game to false scores.

Therefore, we must assume we cannot trust the client to keep track of its own state.

We have a few solutions to this problem.

1. We trust the client
   - Pros
     - Simple to implement
     - Game appears very responsive
   - Cons
     - Abusable, cannot trust high scores
2. We simulate the game on the server
   - Pros
     - Prevent simple cheating methods
     - Leaderboard is reliable
   - Cons
     - Latency is increased as state must be computed on server
     - More challenging to implement
      
