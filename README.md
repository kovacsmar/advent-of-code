# advent-of-code

**Solving Advent of Code Puzzles**

If you do not know what Advent of Code is, check out <https://adventofcode.com/>

I share my solutions here, but not the puzzles or inputs. Inputs are ignored by git.

To access the puzzles and your own input, visit <https://adventofcode.com/>

If you clone this repository, you need to add your own `puzzle.txt` and `test.txt` under each day's `input` directory, or the code will break.

I do not own or take credit for the Advent of Code website, puzzles or related property.

Solutions are written in Go.

## Generator Script

The generator script sets up a basic structure for the given year's puzzles.

It is intentionally minimal as puzzle inputs vary.

It supports the 12-day format introduced in 2025. Earlier years are not supported.

```bash
./generate 2026
```

It generates a directory for each of the 12 days with the following layout under the `2026` directory:

```text
2026/
├── Day-1
│   ├── go.mod
│   ├── input
│   │   ├── puzzle.txt
│   │   └── test.txt
│   ├── main.go
│   ├── part-1
│   │   ├── solution.go
│   │   └── solution_test.go
│   ├── part-2
│   │   ├── solution.go
│   │   └── solution_test.go
│   └── reader
│       └── reader.go
└── Day-2 ... Day-12
```

## Solutions

- [2026](2026)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
