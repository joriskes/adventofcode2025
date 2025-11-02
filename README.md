# Advent Of Code 2025

Advent of code challenges 2025, in GO.

## Installation

Copy `.env.example` to `.env` and set the environment variable called `AOC_SESSION` with the session cookie of
adventofcode.com

## Running

To run all days run: `go run main.go run` from the root.
To run a spefic day add that to the run command: `go run main.go run 1`.
You can also specify the day in the environment variable `AOC_DAY` and run: `go run main.go current` 

## Auto download / day creation script

To create / update a day run:
`go run main.go create <DAY_NUMBER>` from the root. It will create a new day as a directory, download the
AoC input and copy the template there
