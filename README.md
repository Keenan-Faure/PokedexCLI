# PokemonCLI

A Pokemon game in a command-line REPL using the [PokéAPI](https://pokeapi.co)

This project takes the following into consideration:

-   Parsing JSON in Go
-   Making HTTP requests in Go
-   Build a CLI tool that makes interacting with a back-end server easier
-   Go development and tooling
-   Caching and how to use it to improve performance
-   Internal go modules

## How to install

Download the code from Github repository :)

## How to run

Once you navigated to the folder containing the (extracted) repository in your command line enter:

```
go build && ./pokemon
```

## Running tests

To run the tests:

```
go test ./internal/{{module-name}}
```

where `module-name` is the name of the folder or module eg. `fetch`

## Commands

Supported commands that are used in the CLI app can be found by entering `help` in the cli app.

## Pokemon types

Pokemon types have associated emoji's. This can be changed in `./internal/fetch/fetch.go` under the varaible `PokeTypes`

## PokeFarm

A pokemon farm has been added to evolve pokemon. It is worth noting that pokemon can only evolve inside while inside the farm.

Pokemon obtain EXP at a rate of `5` per second by default, this can be manually changed in `./internal/pokefarm/pokefarm.go`

```
const baseExp = 5          //exp received every second
```

Pokemon Evolution base exp for the first and second form can also be found and adjusted in `./internal/pokefarm/pokefarm.go`

```
const expFirstForm = 1050  //exp required to evolve to 1st form
const expSecondFrom = 2500 //exp required to evolve to 2nd form
```

Thank you for reading 👹
