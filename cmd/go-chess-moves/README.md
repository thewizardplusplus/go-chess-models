# go-chess-moves

The utility for generating all possible chess moves.

## Features

- generating all possible moves:
  - parameters:
    - representing the board:
      - as an associative array of pieces with their positions as keys;
      - as a plain array of pieces with exact correspondence array indices to piece positions;
      - as a set of integers corresponding to a particular combination of piece color and type, and where each bit corresponds to a particular piece position (so-called a [bitboard](https://en.wikipedia.org/wiki/Bitboard));
    - position;
    - color that moves first.

## Installation

```
$ go install github.com/thewizardplusplus/go-chess-models/cmd/go-chess-moves@latest
```

## Usage

```
$ go-chess-moves -h | -help | --help
$ go-chess-moves [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-storage {map|slice|bits}` &mdash; piece storage kind (default: `slice`);
- `-fen STRING` &mdash; board in [Forsyth-Edwards Notation](https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation) (default: `rnbqk/ppppp/5/PPPPP/RNBQK`, i.e., [Gardner's minichess](https://en.wikipedia.org/wiki/Minichess#5%C3%975_chess));
- `-color {black|white}` &mdash; color that moves first (default: `white`).
