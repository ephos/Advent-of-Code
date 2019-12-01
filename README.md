# Advent of Code

## How I Get the Puzzles

Using `aocdl` utility to download puzzle inputs.  Requires Go if grabbing with `go get`.  Pre-compiled binary will also work if you download/compile and put in path.

```bash
go get github.com/GreenLightning/advent-of-code-downloader/aocdl

aocdl -session-cookie 0000000000000000000000000000000000000000000000 -year 2015 -day 1
```

## Note on the Code

### Go

Some of this is so I can learn a new language, _namely Go_.

For the Go puzzles you can compile and run the resulting binary or just run the following...

```bash
go run ./main.go ./part_1.go ./part_2.go
```

### PowerShell

**NOTE:** I run Linux primarily, so if you clone this down on Windows just know I am working entirely with PowerShell _(Not Windows PowerShell)_ and the paths may look funny.

In the case of the PowerShell this is very much quick and dirty solutions.  Things like tests, error checking, parameterization, and modularizing are minimal to non-existent.  Quick and dirty solves the puzzle, channeling some inner Battle Faction for this 😁!