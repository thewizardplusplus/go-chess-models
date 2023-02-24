package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	"github.com/thewizardplusplus/go-chess-cli/encoding/ascii"
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	storageKind := flag.String("storage", "slice",
		"piece storage kind (allowed: map, slice, bits)")
	fen := flag.String("fen", "rnbqk/ppppp/5/PPPPP/RNBQK",
		"board in Forsyth-Edwards Notation (default: Gardner's minichess)")
	color := flag.String("color", "white",
		"color that moves first (allowed: black, white)")
	flag.Parse()

	var pieceStorageFactory uci.PieceStorageFactory
	switch *storageKind {
	case "map":
		pieceStorageFactory = boards.NewMapBoard
	case "slice":
		pieceStorageFactory = boards.NewSliceBoard
	case "bits":
		pieceStorageFactory = func(
			size common.Size,
			pieceGroup []common.Piece,
		) common.PieceStorage {
			return boards.NewBitBoard(size, pieceGroup, pieces.NewPiece)
		}
	default:
		log.Fatal("incorrect piece storage kind")
	}

	storage, err :=
		uci.DecodePieceStorage(*fen, pieces.NewPiece, pieceStorageFactory)
	if err != nil {
		log.Fatalf("unable to decode the board: %s", err)
	}

	parsedColor, err := ascii.DecodeColor(*color)
	if err != nil {
		log.Fatalf("unable to decode the color: %s", err)
	}

	var generator models.MoveGenerator
	moves, err := generator.MovesForColor(storage, parsedColor)
	if err != nil {
		log.Fatalf("unable to generate moves: %s", err)
	}

	sort.Slice(moves, func(i int, j int) bool {
		startA, startB := moves[i].Start, moves[j].Start
		if startA == startB {
			finishA, finishB := moves[i].Finish, moves[j].Finish
			return less(finishA, finishB)
		}

		return less(startA, startB)
	})

	unitEnding, linkingVerb := "", "is"
	if len(moves) != 1 {
		unitEnding, linkingVerb = "s", "are"
	}
	fmt.Printf("%d move%s %s generated:\n", len(moves), unitEnding, linkingVerb)

	for _, move := range moves {
		nextStorage := storage.ApplyMove(move)
		fmt.Printf(
			"* %s -> %s\n",
			uci.EncodeMove(move),
			uci.EncodePieceStorage(nextStorage),
		)
	}
}

func less(positionOne common.Position, positionTwo common.Position) bool {
	if positionOne.File == positionTwo.File {
		return positionOne.Rank < positionTwo.Rank
	}

	return positionOne.File < positionTwo.File
}
