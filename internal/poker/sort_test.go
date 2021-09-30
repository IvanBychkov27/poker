package poker

import (
	"go.uber.org/zap"
	"testing"
)

func BenchmarkPoker_sortComb(b *testing.B) {
	p := &Poker{
		logger: zap.NewNop(),
	}

	comb := p.deckCardsFull()

	for i := 0; i < b.N; i++ {
		p.sortComb(comb)
	}
}

func BenchmarkPoker_sortComb2(b *testing.B) {
	p := &Poker{
		logger: zap.NewNop(),
	}

	comb := p.deckCardsFull()

	for i := 0; i < b.N; i++ {
		p.sortComb2(comb)
	}
}
