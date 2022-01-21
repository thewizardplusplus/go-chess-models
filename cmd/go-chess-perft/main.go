package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/thewizardplusplus/go-chess-cli/encoding/ascii"
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	fen := flag.String("fen", "rnbqk/ppppp/5/PPPPP/RNBQK",
		"board in Forsyth–Edwards Notation (default: Gardner's minichess)")
	color := flag.String("color", "white",
		"color that moves first (allowed: black, white)")
	deep := flag.Int("deep", 5, "analysis deep")
	flag.Parse()

	storage, err := uci.DecodePieceStorage(*fen, pieces.NewPiece, models.NewBoard)
	if err != nil {
		log.Fatalf("unable to decode the board: %s", err)
	}

	parsedColor, err := ascii.DecodeColor(*color)
	if err != nil {
		log.Fatalf("unable to decode the color: %s", err)
	}

	moveCount := models.Perft(storage, parsedColor, *deep, nil)
	fmt.Printf("%d moves\n", moveCount)
}
