package uci

import (
	"testing"
)

func TestGo(T *testing.T){
	eng, _ := NewEngine("../cmd/stockfish")

	_ = eng.Go(1,4)
	eng.Close()
}

//func BenchmarkFib(b *testing.B) {
//	eng, _ := NewEngine("../cmd/stockfish")
//	for n := 0; n < b.N; n++ {
//		eng.Go(3)
//	}
//}

