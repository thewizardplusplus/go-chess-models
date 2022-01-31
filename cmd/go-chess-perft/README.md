# go-chess-perft

The utility for counting all possible chess moves (based on the [perft](https://www.chessprogramming.org/Perft) function).

## Features

- counting all possible moves:
  - parameters:
    - position;
    - color that moves first;
    - analysis deep;
- profiling:
  - targets:
    - CPU usage;
    - memory usage;
  - storing the results to a file.

## Installation

```
$ go get github.com/thewizardplusplus/go-chess-models/cmd/go-chess-perft
```

## Usage

```
$ go-chess-perft -h | -help | --help
$ go-chess-perft [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-storage {map|slice}` &mdash; piece storage kind (default: `slice`);
- `-fen STRING` &mdash; board in [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation) (default: `rnbqk/ppppp/5/PPPPP/RNBQK`, i.e., [Gardner's minichess](https://en.wikipedia.org/wiki/Minichess#5%C3%975_chess));
- `-color {black|white}` &mdash; color that moves first (default: `white`);
- `-deep INTEGER` &mdash; analysis deep (default: `5`);
- `-cpuProfile STRING` &mdash; file for CPU profile writing;
- `-memoryProfile STRING` &mdash; file for memory profile writing.
