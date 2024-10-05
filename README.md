# Pokedexcli

## What is this?

Pokedexcli is a fun command-line game that lets you explore the world of Pokemon right from your computer's terminal. It's like having a Pokedex (a digital Pokemon encyclopedia) at your fingertips!

## What can I do with it?

1. **Explore Pokemon locations**: You can move around different areas where Pokemon live.
2. **Catch Pokemon**: Try your luck at catching Pokemon you find.
3. **View your Pokedex**: See all the Pokemon you've caught.
4. **Inspect Pokemon**: Get more details about the Pokemon you've caught.

## How does it work?

When you start the game, you'll see a prompt that looks like this:

Pokedex >

You can type different commands here to play the game. Here are the main commands:

- `help`: Shows you all the available commands.
- `map`: Shows you the next set of locations you can explore.
- `mapb`: Shows you the previous set of locations.
- `explore <location name>`: Lets you explore a specific location to find Pokemon.
- `catch <pokemon name>`: Tries to catch a Pokemon you've found.
- `inspect <pokemon name>`: Gives you more information about a Pokemon you've caught.
- `pokedex`: Shows you all the Pokemon you've caught.
- `exit`: Closes the game.

## Cool features

1. **It remembers things**: The game uses a clever system (called caching) to remember information it's already looked up. This makes the game faster!

2. **It's connected**: The game connects to a big Pokemon database (called PokeAPI) to get real information about Pokemon.

3. **It's smart**: The game can handle different situations, like if you type a command wrong or try to catch a Pokemon you haven't found yet.

## For the techies

If you're interested in the code behind the game:

- It's written in a programming language called Go.
- The main game loop is in the `repl.go` file (lines 19-48).
- The commands are defined in separate files like `command_catch.go`, `command_explore.go`, etc.
- It uses a custom caching system (in the `internal/pokecache` folder) to store Pokemon data.
- The connection to the Pokemon API is handled in the `internal/pokeapi` folder.

## Want to play?

To start playing, you'll need to have Go installed on your computer. Then, you can run the game using the command line. Have fun exploring the world of Pokemon!