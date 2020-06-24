package main

import (
	"github.com/clinaresl/pgnparser/pgntools"
	"fmt"
)

func main() {
	
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
			fmt.Println(board.GetFen())
	    }

	}
}