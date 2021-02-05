package main

import (
	"fmt"
	"github.com/clinaresl/pgnparser/pgntools"
	"github.com/juanfgarcia/ChessTraining/lichess"
	"github.com/juanfgarcia/ChessTraining/uci"
	"log"
)

const lowerBound = 300
const depth = 3

func main() {
	eng, err := uci.NewEngine("cmd/stockfish")
	if err != nil {
		log.Fatal(err)
	}
    
    games := Lichess.GetGames("juanfgcas")
	// For each game, get each position
	// and for each position get it's FEN string
	for _, game := range games.GetGames() {
		turn := 1
		board := pgntools.InitPgnBoard()
		for i, move := range game.GetMoves() {
            if i == len(game.GetMoves())-1 {
                break
            }
			board.UpdateBoard(move, false)
			turn *= -1 
			fen := board.GetFen()
			eng.SetFEN(fen)
			results, _ := eng.Go(2)
			p0 := results.Score
			results, _ = eng.Go(4)
			p1 := results.Score
			if p1-p0 > lowerBound {
				fmt.Println(p1, p0)
				fmt.Println(fen)
				fmt.Println(results.BestMove)
			}
		}
	}
}
