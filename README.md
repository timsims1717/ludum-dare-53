# Overtime at the Tetromino Factory

This is an [entry](https://ldj.am/$343105) for the [Ludum Dare 53](https://ldjam.com/) game jam. The theme: **Delivery**

The gaming industry has made some cuts recently, and in order to play the latest exciting game, you must put in some work.

Hundreds of individual blocks are being delivered to the Tetromino factory, and you must construct the Tetrominos and deliver them to the playing field.

Use the mouse to construct tetrominos on the factory floor. Only shapes with four blocks can be put on the conveyor belt. If you make one too big, you'll have to throw it in the trash.

This game is singleplayer, but can also easily be played by two players, one controlling the factory, and one playing tetris.

The controls for the factory are:

* **Mouse Left Click** to pick up and drop blocks on the grid and conveyor
* Blocks are automatically combined with other blocks of the same color they connect to on the grid.

The controls for the block falling game are:

* **A** to move the block left
* **D** to move the block right
* **S** to send the block down
* **W** to rotate the block
* **Escape** to pause
* **F5** to toggle fullscreen

![You're playing the game!](Screenshot2.png)

![You've assembled 5 Tetrominos!](Screenshot1.png)

See if you can find all the sticky notes left by Management.

Made by Ben and Tim Sims

## Links:

* [Github Page](https://github.com/timsims1717/ludum-dare-53)
* [itch.io Page](https://thetimsims.itch.io/overtime-at-the-tetromino-factory)

## Resources

* [Pixel Library](https://github.com/faiface/pixel)
* [Beep Library](https://github.com/faiface/beep)
* [Aseprite](https://aseprite.itch.io/)

## Post Jam Change Log

* Added Instruction Sticky Note to beginning of the game
* Achievement language update
* Bug fixes
  * Moving a piece down can occasionally cause a crash
  * The first block after a reset can't be controlled
  * Switched Level and Score to be correct
  * Moved Sticky notes around to remove overlap