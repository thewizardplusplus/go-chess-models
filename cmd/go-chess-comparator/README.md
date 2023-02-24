# go-chess-comparator

The utility for comparing the generation of all possible chess moves by different board representations.

## Features

- comparing the generation of all possible chess moves by different board representations:
  - representing the board:
    - as an associative array of pieces with their positions as keys;
    - as a plain array of pieces with exact correspondence array indices to piece positions;
    - as a set of integers corresponding to a particular combination of piece color and type, and where each bit corresponds to a particular piece position (so-called a [bitboard](https://en.wikipedia.org/wiki/Bitboard));
  - parameters:
    - position;
    - color that moves first;
    - comparing mode:
      - depth-first;
      - breadth-first;
    - analysis deep.

## Installation

```
$ go install github.com/thewizardplusplus/go-chess-models/cmd/go-chess-comparator@latest
```

## Usage

```
$ go-chess-comparator -h | -help | --help
$ go-chess-comparator [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-fen STRING` &mdash; board in [Forsyth-Edwards Notation](https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation) (default: `rnbqk/ppppp/5/PPPPP/RNBQK`, i.e., [Gardner's minichess](https://en.wikipedia.org/wiki/Minichess#5%C3%975_chess));
- `-color {black|white}` &mdash; color that moves first (default: `white`);
- `-mode {depth-first|breadth-first}` &mdash; comparing mode (default: `depth-first`);
- `-deep INTEGER` &mdash; analysis deep (should be greater than or equal to zero; default: `5`).
