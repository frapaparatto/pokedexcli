# pokedexcli

A command-line Pokédex: a REPL that talks to [PokeAPI](https://pokeapi.co/)
to explore locations, encounter Pokémon, try to catch them, and inspect the
ones you've caught.

## Learning Context

The goal was learning the language: `struct`s and methods, pointer
vs. value receivers, `interface`s (or, in one place, the lack of one — see
[Known Limitations](#known-limitations)), goroutines and `sync.Mutex`,
package layout, and table-driven tests. It is not written to the same
production-readiness bar as a service meant to run unattended — see
[Known Limitations](#known-limitations) below for what that means
concretely, persistence chief among it.

## Overview

Run the binary and you get a `Pokedex >` prompt. From there you can page
through in-game locations, explore a location to see which Pokémon appear
there, throw a Pokéball at one, and inspect or list the ones you've caught.
State (which Pokémon you've caught, which page of locations you're on)
lives only in memory for the lifetime of the process.

### Commands

| Command | Usage | Description |
|---|---|---|
| `help` | `help` | List all commands |
| `map` | `map` | Show the next 20 location areas |
| `mapb` | `mapb` | Show the previous 20 location areas |
| `explore` | `explore <location-area>` | List Pokémon encountered in a location area |
| `catch` | `catch <pokemon-name>` | Attempt to catch a Pokémon; success is chance-based |
| `inspect` | `inspect <pokemon-name>` | Show height, weight, stats, and types for a caught Pokémon |
| `pokedex` | `pokedex` | List the names of every Pokémon caught so far |
| `exit` | `exit` | Quit |

## Architecture

The codebase is split into one `internal` package per concern, plus a thin
`main.go` that wires them together.

- **`internal/pokeapi`**: the PokeAPI HTTP client (`Client`) and the
  response shapes it decodes into (`RespLocations`, `RespLocation`,
  `RespPokemon`). Each method (`ListLocations`, `ListPokemon`, `GetPokemon`)
  checks the cache first, falls back to an HTTP GET on a miss, and writes
  the raw response body back into the cache before returning.
- **`internal/pokecache`**: a small in-memory, thread-safe TTL cache
  (`map[string]cacheEntry` behind a `sync.Mutex`). A background goroutine
  started in `NewCache` sweeps expired entries on a ticker at the
  configured interval; there's no size cap or eviction beyond age.
- **`internal/pokedex`**: the caught-Pokémon collection itself — a
  `map[string]pokeapi.RespPokemon` keyed by name, wrapped in a couple of
  methods (`Add`, and construction via `NewPokedex`). Re-catching a
  Pokémon overwrites the earlier entry.
- **`internal/repl`**: the REPL loop (`Start`, in `repl.go`) and one file
  per command (`command_map.go`, `command_explore.go`, `command_catch.go`,
  `command_inspect.go`, `command_pokedex.go`, `command_exit.go`,
  `command_help.go`). `commands.go` holds the `Config` struct (shared
  REPL state: the PokeAPI client, the Pokedex, and pagination cursors) and
  the command registry, a `map[string]cliCommand` built in `getCommands`.
- **`main.go`**: the composition root. Builds a `pokeapi.Client` with a
  5-second HTTP timeout and a 5-minute cache interval, wraps it in a
  `repl.Config`, and starts the REPL.

Each REPL command has the same signature,
`func(conf *Config, options ...string) error`; `repl.Start` looks up the
first word of the input line in the command map and calls it with the rest
as arguments. An unrecognized command prints `Unknown command`; a returned
error is printed as `Error: <message>` and the loop continues either way.

### Design notes

- **Pagination**: `map`/`mapb` don't track a page number. They follow the
  `next`/`previous` URLs PokeAPI returns in each response, stored on
  `Config` as `nextLocationsURL`/`prevLocationsURL`. `mapb` on the first
  page returns an error (`"you're on the first page"`) rather than calling
  the API with a nil URL.
- **Catch chance**: implemented in `command_catch.go` as
  `100 / (100 + baseExperience)`, clamped to `[0.05, 0.90]`, so every
  Pokémon is catchable and none is a sure thing. The constant (`100`) and
  the clamp bounds are hardcoded, not configurable.
- **Caching is response-level, not request-aware**: the cache key is the
  request URL, so paginated location pages and individual location/Pokémon
  lookups are each cached independently, keyed by their own URL.

## Building and Running

### Prerequisites

* Go 1.26.5 (see `go.mod`)

There are no third-party dependencies — the module uses only the Go
standard library, so there's no `go.sum` to fetch.

### Run it

```bash
go run .
```

Or build a binary first:

```bash
go build -o pokedexcli .
./pokedexcli
```

There's also a `Makefile` that chains `test` → `fmt` → `vet` → `build` (requires `make`):

```bash
make
```

### Example session

```
Pokedex > map
canalave-city-area
eterna-city-area
...
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - magikarp
Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!
Pokedex > inspect tentacool
Name: tentacool
Height: 9
Weight: 455
Stats:
  -hp: 40
  -attack: 40
  ...
Types:
  - water
  - poison
Pokedex > pokedex
Pokedex:
 - tentacool
```

### Run the test suite

```bash
go test ./...
```

## Testing

Tests are plain `go test`, table-driven where the boot.dev course material
suggested it (`repl_test.go`'s `TestCleanInput`, `pokecache_test.go`).

- `internal/pokecache/pokecache_test.go`: add/get round-trip, and that the
  reap loop actually expires an entry after its interval elapses.
- `internal/pokedex/pokedex_test.go`: a new Pokedex starts empty, `Add`
  inserts, and re-adding the same name overwrites rather than duplicating.
- `internal/repl/repl_test.go`: `cleanInput`'s whitespace/case handling.
- `internal/repl/command_inspect_test.go`,
  `internal/repl/command_pokedex_test.go`: the inspect and pokedex
  commands against a `Config` built directly in the test, no PokeAPI
  client involved.

`internal/pokeapi` has no tests. `Client` is a concrete struct with no
interface behind it, so there is no seam to substitute a fake HTTP
response the way `command_inspect_test.go` substitutes a `Config` — see
[Known Limitations](#known-limitations).

## Known Limitations

This was a project for learning Go, and it stops short of several things
a longer-lived tool would need:

- **No persistence.** The Pokedex and pagination cursors live in
  `Config`, in memory, for the life of the process. Exiting the REPL
  (or a crash) loses every caught Pokémon; there's no save file, no
  database, nothing written to disk.
- **`pokeapi.Client` has no interface.** It's a concrete struct, so
  `internal/pokeapi` has no tests, and nothing above it (`repl`) can run
  against a fake client either — command tests that touch the network
  path (`catch`, `explore`, `map`) aren't covered, only the ones that
  don't call the API.
- **No configuration.** The base URL, HTTP timeout (5s), and cache TTL
  (5m) are hardcoded (`baseURL` in `pokeapi.go`, the literals passed to
  `pokeapi.NewClient` in `main.go`). None of it is overridable via flags,
  env vars, or a config file.
- **No retry/backoff or rate-limit handling** for PokeAPI calls, and no
  distinction between "PokeAPI is down" and "this location/Pokémon
  doesn't exist" beyond the 404 check already in `pokeapi`.
- **No structured logging.** Everything is `fmt.Print*` straight to
  stdout; there's no log level, no request correlation, nothing suited
  to running unattended.
- **The cache has no size bound**, only age-based eviction — a long
  session that explores many locations grows the map without limit until
  the reap loop catches up.
