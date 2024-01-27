# Project Description
This Pokedex CLI tool is for listing information about locations in pokemon and what pokemon can be encountered there using https://pokeapi.co/.

## Available Commands:
- exit - Exits the CLI
- map - Lists a page of 20 location areas in the Pokemon world. Subsequent calls to map will return the next page of data
- mapb - Lists the previous page of 20 location areas in the Pokemon world. Subsequent calls to mapb will return the previous page of data
- explore <location-area> - Lists the pokemon encountered in the provided <location-area>
- catch <pokemon> - Attempts to catch and store the provided <pokemon>
- inspect <pokemon> - Displays data about the provided <pokemon> if they have been caught and added to the pokedex
- pokedex - Lists the names of the pokemon that have been caught
- help - Displays a help message

## Running the project
Compile with `go build` to generate executable 'pokedexcli'
Run with `./pokedexcli`

## Other Features
### Request Caching
Default request cache TTL policy is 60 seconds, this can be configured with `-CacheTTL <time in seconds>` commandline argument
Example Usage: `./pokedexcli -CacheTTL 120`
Note: all values <= 0 will default to 60 seconds

### Command History
Supports command history with up/down arrow keys.
