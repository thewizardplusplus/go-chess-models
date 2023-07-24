# Change Log

## [v1.9.0](https://github.com/thewizardplusplus/go-chess-models/tree/v1.9.0) (2023-07-24)

Adding the `boards.BitBoard` structure; adding the `go-chess-moves` and `go-chess-comparator` tools.

- new features:
  - the `boards.BitBoard` structure:
    - add the `boards.BitBoard` structure;
    - add the `boards.BitBoard` structure to the `go-chess-perft` tool;
  - add the `go.mod` file;
  - tools:
    - improve the `go-chess-perft` tool:
      - improve the documentation;
      - improve the options;
      - validate the `deep` option;
    - add the `go-chess-moves` tool;
    - add the `go-chess-comparator` tool;
- refactoring:
  - fix the code style with the `golangci-lint` tool;
- unit testing:
  - improve the tests of the `ApplyMove()` methods;
- examples:
  - move the tests of the `boards` package to the right place.

## [v1.8.1](https://github.com/thewizardplusplus/go-chess-models/tree/v1.8.1) (2022-04-07)

Adding the `boards.WrapBasePieceStorage` function that universally implements the `Pieces()` and `CheckMove()` methods of the `common.PieceStorage` interface if necessary and refactoring.

- new features:
  - the `boards.WrapBasePieceStorage` function that universally implements the `Pieces()` and `CheckMove()` methods of the `common.PieceStorage` interface if necessary;
- refactoring:
  - separate out the following packages:
    - separate out the `common` package:
      - move the `chessmodels.Color` type to the `common` package;
      - move the `chessmodels.Kind` type to the `common` package;
      - move the `chessmodels.Position` structure to the `common` package;
      - move the `chessmodels.Move` structure to the `common` package;
      - move the `chessmodels.CheckMove()` function to the `common` package;
      - move the `chessmodels.Size` structure to the `common` package;
      - move the `chessmodels.Piece` interface to the `common` package;
      - move the `chessmodels.PieceStorage` interface to the `common` package;
      - move the `chessmodels.Pieces()` function to the `common` package;
    - separate out the `boards` package:
      - move the `chessmodels.BaseBoard` structure to the `boards` package;
      - move the `chessmodels.MapBoard` structure to the `boards` package;
      - move the `chessmodels.SliceBoard` structure to the `boards` package;
    - add aliases to the `chessmodels` package for the moved entities:
      - mark them as deprecated;
  - add the `common.BasePieceStorage` interface;
  - add the `common.PieceGroupGetter` interface;
  - add the `common.MoveChecker` interface;
- unit testing:
  - use the `chessmodels.MockPieceStorage` structure in the tests of the `chessmodels.CheckMove()` function.

## [v1.8](https://github.com/thewizardplusplus/go-chess-models/tree/v1.8) (2022-02-05)

Adding the `chessmodels.SliceBoard` structure; optimization.

- new features:
  - the `chessmodels.BaseBoard` structure that stores a board size;
  - the `chessmodels.SliceBoard` structure:
    - rename the `chessmodels.Board` structure to `chessmodels.MapBoard`;
    - add the `chessmodels.SliceBoard` structure;
    - add the `chessmodels.SliceBoard` structure to the `go-chess-perft` tool;
    - make the `chessmodels.Board` type redirect to the `chessmodels.SliceBoard` structure;
  - transform the `Pieces()` methods of the `chessmodels.MapBoard` and `chessmodels.SliceBoard` structures to an independent function;
  - the `chessmodels.Size` structure:
    - add the `PositionIndex()` method;
    - add the `PositionCount()` method;
    - add the `IteratePositions()` method;
- refactoring:
  - optimize memory allocations;
- unit testing:
  - slightly improve the tests of the `chessmodels.MapBoard` structure:
    - of the `Piece()` method;
    - of the `ApplyMove()` method;
  - add the benchmarks based on the `chessmodels.Perft()` function:
    - for the `chessmodels.MapBoard` structure;
    - for the `chessmodels.SliceBoard` structure.

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
