package main

import (
	"container/list"
	"flag"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/thewizardplusplus/go-chess-cli/encoding/ascii"
	models "github.com/thewizardplusplus/go-chess-models"
	"github.com/thewizardplusplus/go-chess-models/boards"
	"github.com/thewizardplusplus/go-chess-models/common"
	"github.com/thewizardplusplus/go-chess-models/encoding/uci"
	"github.com/thewizardplusplus/go-chess-models/pieces"
)

type namedPieceStorageFactory struct {
	name    string
	factory uci.PieceStorageFactory
}

type namedPieceStorage struct {
	name    string
	storage common.PieceStorage
}

type state struct {
	namedPieceStorages []namedPieceStorage
	color              common.Color
	currentDeep        int
	maximalDeep        int
}

type moveGenerator interface {
	MovesForColor(storage common.PieceStorage, color common.Color) (
		[]common.Move,
		error,
	)
}

var namedPieceStorageFactories = []namedPieceStorageFactory{
	{
		name:    "MapBoard",
		factory: boards.NewMapBoard,
	},
	{
		name:    "SliceBoard",
		factory: boards.NewSliceBoard,
	},
	{
		name: "BitBoard",
		factory: func(
			size common.Size,
			pieceGroup []common.Piece,
		) common.PieceStorage {
			return boards.NewBitBoard(size, pieceGroup, pieces.NewPiece)
		},
	},
}

func main() {
	fen := flag.String("fen", "rnbqk/ppppp/5/PPPPP/RNBQK",
		"board in Forsyth-Edwards Notation (default: Gardner's minichess)")
	color := flag.String("color", "white",
		"color that moves first (allowed: black, white)")
	mode := flag.String("mode", "depth-first",
		"comparing mode (allowed: depth-first, breadth-first)")
	deep := flag.Int("deep", 5,
		"analysis deep (should be greater than or equal to zero)")
	flag.Parse()

	var namedPieceStorages []namedPieceStorage
	for _, namedPieceStorageFactory := range namedPieceStorageFactories {
		storageName, storageFactory :=
			namedPieceStorageFactory.name, namedPieceStorageFactory.factory
		storage, err := uci.DecodePieceStorage(*fen, pieces.NewPiece, storageFactory)
		if err != nil {
			const message = "unable to decode the board to the %q storage: %s"
			log.Fatalf(message, storageName, err)
		}

		namedPieceStorages = append(namedPieceStorages, namedPieceStorage{
			name:    storageName,
			storage: storage,
		})
	}

	parsedColor, err := ascii.DecodeColor(*color)
	if err != nil {
		log.Fatalf("unable to decode the color: %s", err)
	}

	if *deep < 0 {
		log.Fatal("incorrect analysis deep")
	}

	var generator models.MoveGenerator
	initialState := state{
		namedPieceStorages: namedPieceStorages,
		color:              parsedColor,
		currentDeep:        0,
		maximalDeep:        *deep,
	}
	switch *mode {
	case "depth-first":
		runDepthFirst(generator, initialState)
	case "breadth-first":
		runBreadthFirst(generator, initialState)
	default:
		log.Fatal("incorrect comparing mode")
	}
}

func runDepthFirst(generator moveGenerator, currentState state) {
	handleCurrentState(generator, currentState, func(nextState state) {
		runDepthFirst(generator, nextState)
	})
}

func runBreadthFirst(generator moveGenerator, initialState state) {
	states := list.New()
	states.PushBack(initialState)

	for states.Len() != 0 {
		currentState := states.Remove(states.Front()).(state)
		handleCurrentState(generator, currentState, func(nextState state) {
			states.PushBack(nextState)
		})
	}
}

func handleCurrentState(
	generator moveGenerator,
	currentState state,
	nextStateHandler func(nextState state),
) {
	log.Printf(
		"%d ply, %s color",
		currentState.currentDeep,
		ascii.EncodeColor(currentState.color),
	)
	for _, namedPieceStorage := range currentState.namedPieceStorages {
		log.Printf("  * piece storage kind: %s", namedPieceStorage.name)
		log.Printf("    FEN: %s", uci.EncodePieceStorage(namedPieceStorage.storage))
	}

	if currentState.currentDeep == currentState.maximalDeep {
		log.Print("  maximal analysis deep")
		return
	}

	moves, err := generateMoves(generator, currentState)
	if err != nil {
		log.Printf("  %s", err)
		return
	}

	for _, move := range moves {
		log.Printf("  apply move %s", uci.EncodeMove(move))

		var nextNamedPieceStorages []namedPieceStorage
		for _, namedPieceStorageInstance := range currentState.namedPieceStorages {
			nextNamedPieceStorages = append(nextNamedPieceStorages, namedPieceStorage{
				name:    namedPieceStorageInstance.name,
				storage: namedPieceStorageInstance.storage.ApplyMove(move),
			})
		}

		nextStateHandler(state{
			namedPieceStorages: nextNamedPieceStorages,
			color:              currentState.color.Negative(),
			currentDeep:        currentState.currentDeep + 1,
			maximalDeep:        currentState.maximalDeep,
		})
	}
}

func generateMoves(
	generator moveGenerator,
	currentState state,
) ([]common.Move, error) {
	var previousNamedPieceStorage namedPieceStorage
	var previousMoves []common.Move
	for _, namedPieceStorage := range currentState.namedPieceStorages {
		moves, err :=
			generator.MovesForColor(namedPieceStorage.storage, currentState.color)
		if err != nil {
			return nil, fmt.Errorf("unable to generate moves: %w", err)
		}

		sort.Slice(moves, func(i int, j int) bool {
			startA, startB := moves[i].Start, moves[j].Start
			if startA == startB {
				finishA, finishB := moves[i].Finish, moves[j].Finish
				return less(finishA, finishB)
			}

			return less(startA, startB)
		})

		if previousMoves != nil && !reflect.DeepEqual(moves, previousMoves) {
			return nil, fmt.Errorf(
				"incorrect generated moves:\n"+
					"  piece storage kind: %s\n"+
					"  expected FEN: %s\n"+
					"  actual FEN: %s\n"+
					"  expected moves: %+v\n"+
					"  actual moves: %+v",
				namedPieceStorage.name,
				uci.EncodePieceStorage(previousNamedPieceStorage.storage),
				uci.EncodePieceStorage(namedPieceStorage.storage),
				encodeMoves(previousMoves),
				encodeMoves(moves),
			)
		}

		previousNamedPieceStorage = namedPieceStorage
		previousMoves = moves
	}

	return previousMoves, nil
}

func less(positionOne common.Position, positionTwo common.Position) bool {
	if positionOne.File == positionTwo.File {
		return positionOne.Rank < positionTwo.Rank
	}

	return positionOne.File < positionTwo.File
}

func encodeMoves(moves []common.Move) string {
	var encodedMoves []string
	for _, move := range moves {
		encodedMoves = append(encodedMoves, uci.EncodeMove(move))
	}

	return strings.Join(encodedMoves, ", ")
}
