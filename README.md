# go-chess-models

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-chess-models?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-chess-models)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-chess-models)](https://goreportcard.com/report/github.com/thewizardplusplus/go-chess-models)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-chess-models.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-chess-models)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-chess-models/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-chess-models)

The library that implements checking and generating of chess moves.

_**Disclaimer:** this library was written directly on an Android smartphone with the AnGoIde IDE._

## Features

- representing the board as an associative array of pieces with their positions as keys;
- immutable applicating moves to the board via copying the latter;
- checkings of moves:
  - universal;
  - individual for all types of pieces;
- generating moves via filtering from all possible ones;
- using an abstraction of a piece;
- [Forsyth–Edwards Notation](https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation):
  - parsing:
    - of a position;
    - of a move;
    - of a piece kind;
    - of a piece color;
    - of a board;
  - serialization:
    - of a position;
    - of a move;
    - of a piece kind;
    - of a piece color;
    - of a board.

## Installation

```
$ go get github.com/thewizardplusplus/go-chess-models
```

## Examples

`chessmodels.Board.CheckMove()`:

```go
package main

import (
	"fmt"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewRook(models.Black, models.Position{File: 2, Rank: 2}),
		pieces.NewBishop(models.White, models.Position{File: 3, Rank: 3}),
	})

	moveOne := models.Move{
		Start:  models.Position{File: 2, Rank: 2},
		Finish: models.Position{File: 3, Rank: 3},
	}
	fmt.Printf("%+v: %v\n", moveOne, board.CheckMove(moveOne))

	moveTwo := models.Move{
		Start:  models.Position{File: 3, Rank: 3},
		Finish: models.Position{File: 2, Rank: 2},
	}
	fmt.Printf("%+v: %v\n", moveTwo, board.CheckMove(moveTwo))

	// Output:
	// {Start:{File:2 Rank:2} Finish:{File:3 Rank:3}}: illegal move
	// {Start:{File:3 Rank:3} Finish:{File:2 Rank:2}}: <nil>
}
```

`chessmodels.Board.ApplyMove()`:

```go
package main

import (
	"fmt"
	"sort"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type ByPosition []models.Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(
	i, j int,
) bool {
	a := group[i].Position()
	b := group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func main() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewRook(models.Black, models.Position{File: 2, Rank: 2}),
		pieces.NewBishop(models.White, models.Position{File: 3, Rank: 3}),
	})
	pieces := board.Pieces()
	sort.Sort(ByPosition(pieces))
	fmt.Printf("%+v\n", pieces)

	updatedBoard := board.ApplyMove(models.Move{
		Start:  models.Position{File: 3, Rank: 3},
		Finish: models.Position{File: 2, Rank: 2},
	})
	updatedPieces := updatedBoard.Pieces()
	sort.Sort(ByPosition(updatedPieces))
	fmt.Printf("%+v\n", updatedPieces)

	// Output:
	// [{Base:{kind:2 color:0 position:{File:2 Rank:2}}} {Base:{kind:3 color:1 position:{File:3 Rank:3}}}]
	// [{Base:{kind:3 color:1 position:{File:2 Rank:2}}}]
}
```

`uci.DecodeMove()`:

```go
package main

import (
	"fmt"

	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func main() {
	move, _ := uci.DecodeMove("d4c3")
	fmt.Printf("%+v\n", move)

	// Output: {Start:{File:3 Rank:3} Finish:{File:2 Rank:2}}
}
```

`uci.EncodeMove()`:

```go
package main

import (
	"fmt"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func main() {
	move := uci.EncodeMove(models.Move{
		Start:  models.Position{File: 3, Rank: 3},
		Finish: models.Position{File: 2, Rank: 2},
	})
	fmt.Printf("%v\n", move)

	// Output: d4c3
}
```

`uci.DecodePieceStorage()`:

```go
package main

import (
	"fmt"
	"sort"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type ByPosition []models.Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(
	i, j int,
) bool {
	a := group[i].Position()
	b := group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func main() {
	const fen = "8/8/8/8/3B4/2r5/8/8"
	storage, _ := uci.DecodePieceStorage(fen, pieces.NewPiece, models.NewBoard)
	pieces := storage.Pieces()
	sort.Sort(ByPosition(pieces))
	fmt.Printf("%+v\n", pieces)

	// Output: [{Base:{kind:2 color:0 position:{File:2 Rank:2}}} {Base:{kind:3 color:1 position:{File:3 Rank:3}}}]
}
```

`uci.EncodePieceStorage()`:

```go
package main

import (
	"fmt"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	board := models.NewBoard(models.Size{Width: 5, Height: 5}, []models.Piece{
		pieces.NewRook(models.Black, models.Position{File: 2, Rank: 2}),
		pieces.NewBishop(models.White, models.Position{File: 3, Rank: 3}),
	})
	fen := uci.EncodePieceStorage(board)
	fmt.Printf("%v\n", fen)

	// Output: 5/3B1/2r2/5/5
}
```

## Utilities

- [go-chess-perft](cmd/go-chess-perft) &mdash; utility for counting all possible moves (based on the [perft](https://www.chessprogramming.org/Perft) function)

## License

The MIT License (MIT)

Copyright &copy; 2019 thewizardplusplus
