# go-chess-models

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-chess-models?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-chess-models)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-chess-models)](https://goreportcard.com/report/github.com/thewizardplusplus/go-chess-models)
[![Build Status](https://app.travis-ci.com/thewizardplusplus/go-chess-models.svg?branch=master)](https://app.travis-ci.com/thewizardplusplus/go-chess-models)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-chess-models/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-chess-models)

The library that implements checking and generating of chess moves.

_**Disclaimer:** this library was written directly on an Android smartphone with the AnGoIde IDE._

## Features

- representing the board:
  - as an associative array of pieces with their positions as keys;
  - as a plain array of pieces with exact correspondence array indices to piece positions;
  - as a set of integers corresponding to a particular combination of piece color and type, and where each bit corresponds to a particular piece position (so-called a [bitboard](https://en.wikipedia.org/wiki/Bitboard));
- immutable applicating moves to the board via copying the latter;
- checkings of moves:
  - universal;
  - individual for all types of pieces;
- generating moves via filtering from all possible ones;
- move restrictions (abandoned moves):
  - pawn double-move;
  - en passant capture;
  - promotion;
  - castling;
- [perft](https://www.chessprogramming.org/Perft) function;
- using an abstraction of a piece;
- [Forsyth-Edwards Notation](https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation):
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
    - of a board;
- utilities:
  - utility for counting all possible moves (based on the [perft](https://www.chessprogramming.org/Perft) function);
  - utility for generating all possible chess moves;
  - utility for comparing the generation of all possible chess moves by different board representations.

## Installation

```
$ go get github.com/thewizardplusplus/go-chess-models
```

## Examples

`boards.MapBoard.CheckMove()`:

```go
package main

import (
	"fmt"

	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	board := boards.NewMapBoard(common.Size{Width: 5, Height: 5}, []common.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewBishop(common.White, common.Position{File: 3, Rank: 3}),
	})

	moveOne := common.Move{
		Start:  common.Position{File: 2, Rank: 2},
		Finish: common.Position{File: 3, Rank: 3},
	}
	fmt.Printf("%+v: %v\n", moveOne, board.CheckMove(moveOne))

	moveTwo := common.Move{
		Start:  common.Position{File: 3, Rank: 3},
		Finish: common.Position{File: 2, Rank: 2},
	}
	fmt.Printf("%+v: %v\n", moveTwo, board.CheckMove(moveTwo))

	// Output:
	// {Start:{File:2 Rank:2} Finish:{File:3 Rank:3}}: illegal move
	// {Start:{File:3 Rank:3} Finish:{File:2 Rank:2}}: <nil>
}
```

`boards.MapBoard.ApplyMove()`:

```go
package main

import (
	"fmt"
	"sort"

	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type ByPosition []common.Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i int, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(i int, j int) bool {
	a, b := group[i].Position(), group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func main() {
	board := boards.NewMapBoard(common.Size{Width: 5, Height: 5}, []common.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewBishop(common.White, common.Position{File: 3, Rank: 3}),
	})
	pieces := board.Pieces()
	sort.Sort(ByPosition(pieces))
	fmt.Printf("%+v\n", pieces)

	updatedBoard := board.ApplyMove(common.Move{
		Start:  common.Position{File: 3, Rank: 3},
		Finish: common.Position{File: 2, Rank: 2},
	})
	updatedPieces := updatedBoard.Pieces()
	sort.Sort(ByPosition(updatedPieces))
	fmt.Printf("%+v\n", updatedPieces)

	// Output:
	// [{Base:{kind:2 color:0 position:{File:2 Rank:2}}} {Base:{kind:3 color:1 position:{File:3 Rank:3}}}]
	// [{Base:{kind:3 color:1 position:{File:2 Rank:2}}}]
}
```

`chessmodels.MoveGenerator.MovesForColor()`:

```go
package main

import (
	"fmt"
	"sort"

	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	board := boards.NewMapBoard(common.Size{Width: 5, Height: 5}, []common.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewKnight(common.White, common.Position{File: 3, Rank: 3}),
		pieces.NewPawn(common.White, common.Position{File: 4, Rank: 3}),
	})

	var generator models.MoveGenerator
	moves, _ := generator.MovesForColor(board, common.White)

	// sorting only by the final point will be sufficient for the reproducibility
	// of this example
	sort.Slice(moves, func(i int, j int) bool {
		a, b := moves[i].Finish, moves[j].Finish
		if a.File == b.File {
			return a.Rank < b.Rank
		}

		return a.File < b.File
	})

	for _, move := range moves {
		fmt.Printf("%+v\n", move)
	}

	// Output:
	// {Start:{File:3 Rank:3} Finish:{File:1 Rank:2}}
	// {Start:{File:3 Rank:3} Finish:{File:1 Rank:4}}
	// {Start:{File:3 Rank:3} Finish:{File:2 Rank:1}}
	// {Start:{File:3 Rank:3} Finish:{File:4 Rank:1}}
	// {Start:{File:4 Rank:3} Finish:{File:4 Rank:4}}
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

	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
)

func main() {
	move := uci.EncodeMove(common.Move{
		Start:  common.Position{File: 3, Rank: 3},
		Finish: common.Position{File: 2, Rank: 2},
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

	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type ByPosition []common.Piece

func (group ByPosition) Len() int {
	return len(group)
}

func (group ByPosition) Swap(i int, j int) {
	group[i], group[j] = group[j], group[i]
}

func (group ByPosition) Less(i int, j int) bool {
	a, b := group[i].Position(), group[j].Position()
	if a.File == b.File {
		return a.Rank < b.Rank
	}

	return a.File < b.File
}

func main() {
	const fen = "8/8/8/8/3B4/2r5/8/8"
	storage, _ := uci.DecodePieceStorage(fen, pieces.NewPiece, boards.NewMapBoard)
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

	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	board := boards.NewMapBoard(common.Size{Width: 5, Height: 5}, []common.Piece{
		pieces.NewRook(common.Black, common.Position{File: 2, Rank: 2}),
		pieces.NewBishop(common.White, common.Position{File: 3, Rank: 3}),
	})
	fen := uci.EncodePieceStorage(board)
	fmt.Printf("%v\n", fen)

	// Output: 5/3B1/2r2/5/5
}
```

## Benchmarks

The `chessmodels.Perft()` function using the `boards.MapBoard` structure:

```
BenchmarkPerft/MapBoard/initial/1Ply-8         	     486	   2239365 ns/op	  507723 B/op	   13638 allocs/op
BenchmarkPerft/MapBoard/initial/2Ply-8         	      38	  27071904 ns/op	 6253672 B/op	  165159 allocs/op
BenchmarkPerft/MapBoard/initial/3Ply-8         	       3	 394039297 ns/op	91006645 B/op	 2399815 allocs/op
BenchmarkPerft/MapBoard/kiwipete/1Ply-8        	     150	   7844416 ns/op	 1959120 B/op	   47637 allocs/op
BenchmarkPerft/MapBoard/kiwipete/2Ply-8        	       4	 311539966 ns/op	80169034 B/op	 1892731 allocs/op
```

The `chessmodels.Perft()` function used the `boards.SliceBoard` structure:

```
BenchmarkPerft/SliceBoard/initial/1Ply-8       	     810	   1419572 ns/op	   684016 B/op	   13646 allocs/op
BenchmarkPerft/SliceBoard/initial/2Ply-8       	      68	  17883575 ns/op	  8331384 B/op	  165074 allocs/op
BenchmarkPerft/SliceBoard/initial/3Ply-8       	       4	 253594389 ns/op	121325578 B/op	 2399248 allocs/op
BenchmarkPerft/SliceBoard/kiwipete/1Ply-8      	     220	   5340912 ns/op	  2583824 B/op	   47672 allocs/op
BenchmarkPerft/SliceBoard/kiwipete/2Ply-8      	       5	 208109435 ns/op	103011680 B/op	 1895999 allocs/op
```

The `chessmodels.Perft()` function used the `boards.BitBoard` structure:

```
BenchmarkPerft/BitBoard/initial/1Ply-8         	     284	   4036298 ns/op	   958048 B/op	   35888 allocs/op
BenchmarkPerft/BitBoard/initial/2Ply-8         	      20	  54032966 ns/op	 11627104 B/op	  433592 allocs/op
BenchmarkPerft/BitBoard/initial/3Ply-8         	       2	 708741154 ns/op	169201536 B/op	 6300499 allocs/op
BenchmarkPerft/BitBoard/kiwipete/1Ply-8        	      74	  14372268 ns/op	  3546305 B/op	  124803 allocs/op
BenchmarkPerft/BitBoard/kiwipete/2Ply-8        	       2	 637371490 ns/op	138270896 B/op	 4866744 allocs/op
```

## Utilities

- [go-chess-perft](cmd/go-chess-perft) &mdash; utility for counting all possible moves (based on the [perft](https://www.chessprogramming.org/Perft) function)
- [go-chess-moves](cmd/go-chess-moves) &mdash; utility for generating all possible chess moves
- [go-chess-comparator](cmd/go-chess-comparator) &mdash; utility for comparing the generation of all possible chess moves by different board representations

## License

The MIT License (MIT)

Copyright &copy; 2019, 2022-2023 thewizardplusplus
