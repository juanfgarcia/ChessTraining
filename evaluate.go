package main

import (
	"os"
	"log"
	"github.com/notnil/chess"
	"fmt"
	"github.com/freeeve/uci"
)

func main() {
	file, err := os.Open("examples/test.pgn")
	if err != nil {
		log.Fatal(err)
	}
	pgn, err := chess.PGN(file)
	if err != nil {
		log.Fatal(err)
	}
	game := chess.NewGame(pgn)
	
	eng, err := uci.NewEngine("cmd/stockfish")
	if err != nil {
		log.Fatal(err)
	}
	
	for _, position := range game.Positions(){
		fmt.Println(position)
		eng.SetFEN(position.String())
		// set some result filter options
		resultOpts := uci.HighestDepthOnly | uci.IncludeUpperbounds | uci.IncludeLowerbounds
		results, _ := eng.GoDepth(12, resultOpts)

		fmt.Printf("%d, %b\n" ,results.Results[0].Score, results.Results[0].Mate)
	}




	
}