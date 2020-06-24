package main

import (
	"github.com/clinaresl/pgnparser/pgntools"
	"github.com/juanfgarcia/uci"
	"fmt"
	"log"

)

const lowerBound = 250 
const depth = 3

func main() {

	var lastWhiteScore, lastBlackScore, turn = 0,0,-1

	eng, err := uci.NewEngine("cmd/stockfish")
	if err != nil {
		log.Fatal(err)
	}

	// Set config variables for pgnparser
	var pgnfile string  = "examples/test.pgn"
	var showboard int 
	var query, sort string
	var verbose bool = false
	
	games := pgntools.GetGamesFromFile(pgnfile, showboard, query, sort, verbose)

	// For each game, get each position
	// and for each position get it's FEN string
	for _, game := range games.GetGames(){
		
		board := pgntools.InitPgnBoard()
 
	    for _, move := range game.GetMoves() {
			board.UpdateBoard(move, false)
			fen :=  board.GetFen()
			//fmt.Println(fen)
			eng.SetFEN(fen)
			resultOpts := uci.HighestDepthOnly
			results, _ := eng.GoDepth(depth, resultOpts)
			
			score := results.Results[0].Score

			if (turn==1){
				if lastBlackScore-score >= lowerBound {
					fmt.Println(fen)
				}
				lastBlackScore = score 
			}else{
				if lastWhiteScore-score >= lowerBound {
					fmt.Println(fen)
				}
				lastWhiteScore = score
			}

			turn = -turn
	    }

	}
}