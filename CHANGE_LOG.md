# Change Log

## [v1.7.1](https://github.com/thewizardplusplus/go-chess-models/tree/v1.7.1) (2022-01-25)

Moving the `chessmodels.perft()` function from the tests to the main code; adding the `go-chess-perft` tool and the example for the `chessmodels.MoveGenerator.MovesForColor()` method; bug fixing.

- new features:
  - transform the `chessmodels.Board.CheckMove()` method to an independent function;
  - the `chessmodels.Perft()` function:
    - move the `chessmodels.perft()` function from the tests to the main code;
    - pass a handler to all levels in the `chessmodels.Perft()` function;
    - pass a move generator to the `chessmodels.Perft()` function via an interface;
  - add the `go-chess-perft` tool;
- bug fixing:
  - fix rune conversion from an integer to a string in the `uci.EncodePosition()` function;
  - fix incorrect calculation of the last file in the `uci.EncodePieceStorage()` function;
- refactoring:
  - simplify the utility functions in the `pieces` package;
  - merge the `chessmodels.pieceGroup` type with the `chessmodels.Board` structure;
- unit testing:
  - complete the tests:
    - of the `pieces.NewPiece()` function;
    - of the `uci.DecodePiece()` function;
    - of the `uci.EncodePiece()` function;
  - add the long-playing tests to the Travis CI configuration;
- examples:
  - remove the examples:
    - remove the examples with a mock piece;
    - remove the redundant examples;
  - add the example for the `chessmodels.MoveGenerator.MovesForColor()` method.

## [v1.7](https://github.com/thewizardplusplus/go-chess-models/tree/v1.7) (2019-11-19)

Improving repository decor.

- fixing the code style;
- adding usage examples;
- improving:
  - repository decor;
  - CI configuration.

## [v1.6](https://github.com/thewizardplusplus/go-chess-models/tree/v1.6) (2019-11-01)

Working with [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation) for positions and moves, extending moves checkings and refactoring.

- [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation):
  - parsing:
    - of a position;
    - of a move;
  - serialization:
    - of a position;
    - of a move;
- extending universal moves checkings;
- refactoring.

## [v1.5](https://github.com/thewizardplusplus/go-chess-models/tree/v1.5) (2019-08-19)

Extracting working with [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation) to the separate package and refactoring.

- extracting working with [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation) to the separate package;
- refactoring.

## [v1.4](https://github.com/thewizardplusplus/go-chess-models/tree/v1.4) (2019-07-13)

Serialization models to [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation) and refactoring.

- serialization to [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation):
  - of a piece kind;
  - of a piece color;
  - of a board;
- refactoring.

## [v1.3](https://github.com/thewizardplusplus/go-chess-models/tree/v1.3) (2019-06-28)

Parsing models from [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation) and improving moves checking.

- parsing from [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation):
  - of a piece kind;
  - of a piece color;
  - of a board;
- checking for king capture (including during moves generating);
- fixing the bug with checking bishop moves.

## [v1.2](https://github.com/thewizardplusplus/go-chess-models/tree/v1.2) (2019-05-30)

Implementing individual checkings of moves for all types of pieces.

- implementing individual checkings of moves for all types of pieces.

## [v1.2-alpha](https://github.com/thewizardplusplus/go-chess-models/tree/v1.2-alpha) (2019-05-29)

Implementing real types of pieces, but without individual checkings of moves.

- implementing real types of pieces, but without individual checkings of moves.

## [v1.1](https://github.com/thewizardplusplus/go-chess-models/tree/v1.1) (2019-05-29)

Adding the base for pieces implementations.

- adding the base for pieces implementations.

## [v1.0](https://github.com/thewizardplusplus/go-chess-models/tree/v1.0) (2019-05-28)

Major version.
