# g048
Making a command-line version of 2048 in Go as an exercise and learning experience.

Features:
 - "Swiping" using WASD keys (followed by Enter/Return)
 - Same collapse mechanic as the smartphone game, with multiple combines possible in one swipe
 - Random tiles appearing with 10% probability of being the square of the normal value, just like the smartphone game
 - Execution arguments for game settings modification: g048 [size] [base] [pieces]
    - size: The height/width of the gameboard (default is 4)
    - base: The number being exponentated with each swipe (default is base 2)
    - pieces: The number of pieces the game can start with (default is 2)
 
To-Dos:
 - Autosaving to text file in program directory
