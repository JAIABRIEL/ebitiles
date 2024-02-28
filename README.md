# ebitiles

**This is under heavy development and there WILL be many breaking changes soon.**

This library allows the creation of tile maps of size $`2^n`$.
At the moment, it simply supports storing and drawing tiles at a given position and layer.

![example_animated](https://github.com/JAIABRIEL/ebitiles/assets/31685376/d828405b-ab0f-4a2a-b8da-9cd4797d42cb)

## Current Features

* Place tiles (`*ebiten.Image`) at position and layer.
* Tiles draw their layers on each other and will buffer the result for fast performance.
* Just draw chunks of given level (chunk at level 3 is 8x8 tiles) without drawing the entire map.


## Planned Features

### Example for Tiled maps

I already made it working with Tiled maps on my machine. 
It's easy but I don't intend to put Tiled support directly in this lib.
However, you'll find it in `examples/` soon :)

### Chunk loader

Load chunks based on a position. 
For example draw the chunk level 3 (8x8) on which the player stands and all nearby 8x8 chunks.
This shouldn't be that hard to implement, since all dependencies for this are already implemented.

### Animated tiles

Simply animated tiles based on my [animation lib](https://github.com/JAIABRIEL/gonimator).
It's a bit tricky with different layers and buffers but I intend to make it as performance friendly as possible.

### Pseudo 3D

Draw tiles that are `higher` and `lower` than entities and switch this based on the entities y-position.
I already made a dirty prototype for this.

### Improved buffers

At the moment, just tiles are buffered. There will me a major performance boost when entire
chunks are buffered.

