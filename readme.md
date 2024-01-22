# Project Description
This Pokedex CLI tool is for listing information about locations in pokemon and what pokemon can be encountered there using https://pokeapi.co/.

Available Commands:
    exit - Exits the CLI
    map - Lists a page of 20 location areas in the Pokemon world. Subsequent calls to map will return the next page of data
    mapb - Lists the previous page of 20 location areas in the Pokemon world. Subsequent calls to mapb will return the previous page of data
    explore <location-area> - Lists the pokemon encountered in the provided <location-area>
    help - Displays a help message

# Running the project
Compile with `go build` to generate executable 'pokedexcli'
Run with `./pokedexcli`