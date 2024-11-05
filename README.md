# Pokédex CLI

Bring the Pokémon world to your terminal! **Pokédex CLI** is a command-line game where you can explore, catch, and interact with Pokémon, all from your console.

## 🌟 Description

Pokédex CLI is a terminal-based application that emulates the experience of being a Pokémon Trainer. Travel through different locations, encounter wild Pokémon, attempt to catch them, and build your own Pokédex. The game connects to the real-time [PokéAPI](https://pokeapi.co/) to fetch accurate data about Pokémon and locations.

## 🤔 Why?

As a fan of Pokémon and minimalist interfaces, I wanted a way to enjoy the Pokémon universe without the need for graphics or complex setups. Existing games are great, but they don't offer the lightweight, fast-paced experience that a terminal can provide. So, I built Pokédex CLI to bring the excitement of catching and learning about Pokémon directly to the command line.

## 🚀 Quick Start

### Prerequisites

- Go 1.16 or higher installed on your system.

### Installation

Install Pokédex CLI using the Go toolchain:

```bash
go install github.com/srinivassivaratri/pokedexcli@latest
```

### Run the Game

```bash
pokedexcli
```

## 📖 Usage

Once you start the game, you'll see the prompt:

```
Pokédex >
```

Available commands:

- `help`: Display all available commands.
- `map`: Show the next set of locations you can explore.
- `mapb`: Show the previous set of locations.
- `explore <location_name>`: Explore a specific location to find Pokémon.
- `catch <pokemon_name>`: Attempt to catch a Pokémon you've encountered.
- `inspect <pokemon_name>`: View detailed information about a caught Pokémon.
- `pokedex`: Display all the Pokémon you've caught so far.
- `exit`: Exit the game.

### Examples

Explore a location:

```bash
Pokédex > explore viridian-forest
```

Catch a Pokémon:

```bash
Pokédex > catch pikachu
```

Inspect a caught Pokémon:

```bash
Pokédex > inspect pikachu
```

## 🎮 Features

- **Real-time Data**: Fetches live data from the PokéAPI for up-to-date Pokémon information.
- **Caching System**: Implements a custom caching mechanism to optimize performance.
- **Graceful Error Handling**: Handles incorrect inputs and unexpected situations smoothly.
- **Modular Design**: Easy to extend with new features and commands.

## 🛠️ For Developers

### Tech Stack

- **Language**: Go
- **API**: PokéAPI
- **Caching**: Custom caching system in `internal/pokecache`

### Project Structure

- **main.go**: Entry point of the application.
- **repl.go**: Core game loop handling user input.
- **Commands**: Each command is defined in its own file (e.g., `command_catch.go`, `command_explore.go`).
- **internal/pokeapi**: Handles API interactions with the PokéAPI.
- **internal/pokecache**: Implements caching for API responses.

## 🤝 Contributing

### Clone the Repository

```bash
git clone https://github.com/srinivassivaratri/pokedexcli.git
cd pokedexcli
```

### Build the Project

```bash
go build
```

### Run the Project

```bash
./pokedexcli
```

### Run Tests

```bash
go test ./...
```

### Submit a Pull Request

- Fork the repository.
- Create a new branch for your feature or bug fix.
- Commit your changes and push to your fork.
- Open a pull request against the `main` branch.


## 🙏 Acknowledgments

- [PokéAPI](https://pokeapi.co/) for providing the Pokémon data.
- Inspired by the classic Pokémon games and the open-source community.