package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/thewizardplusplus/go-chess-cli/encoding/ascii"
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

func main() {
	storageKind := flag.String("storage", "slice",
		"piece storage kind (allowed: map, slice; default: slice)")
	fen := flag.String("fen", "rnbqk/ppppp/5/PPPPP/RNBQK",
		"board in Forsyth–Edwards Notation (default: Gardner's minichess)")
	color := flag.String("color", "white",
		"color that moves first (allowed: black, white)")
	deep := flag.Int("deep", 5, "analysis deep")
	cpuProfile := flag.String("cpuProfile", "", "file for CPU profile writing")
	memoryProfile := flag.String("memoryProfile", "",
		"file for memory profile writing")
	flag.Parse()

	if *cpuProfile != "" {
		cpuProfileFile, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatalf("unable to create CPU profile: %s", err)
		}
		defer func() {
			if err := cpuProfileFile.Close(); err != nil {
				log.Fatalf("unable to close the CPU profile: %s", err)
			}
		}()

		if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
			log.Fatalf("unable to start CPU profiling: %s", err)
		}
		defer pprof.StopCPUProfile()
	}

	var pieceStorageFactory uci.PieceStorageFactory
	switch *storageKind {
	case "map":
		pieceStorageFactory = models.NewMapBoard
	case "slice":
		pieceStorageFactory = models.NewSliceBoard
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
	moveCount := models.Perft(generator, storage, parsedColor, *deep, nil)
	var unitEnding string
	if moveCount != 1 {
		unitEnding = "s"
	}
	fmt.Printf("%d move%s\n", moveCount, unitEnding)

	if *memoryProfile != "" {
		memoryProfileFile, err := os.Create(*memoryProfile)
		if err != nil {
			log.Fatalf("unable to create memory profile: %s", err)
		}
		defer func() {
			if err := memoryProfileFile.Close(); err != nil {
				log.Fatalf("unable to close the memory profile: %s", err)
			}
		}()

		runtime.GC() // get up-to-date statistics

		if err := pprof.WriteHeapProfile(memoryProfileFile); err != nil {
			log.Fatalf("unable to write memory profile: %s", err)
		}
	}
}
