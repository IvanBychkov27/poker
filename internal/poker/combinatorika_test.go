package poker

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func Test_combinatorika(t *testing.T) {
	p := &Poker{logger: zap.NewNop()}
	res := p.combinatorika(uint64(7), uint64(52))
	assert.Equal(t, uint64(133784560), res)
}

func Benchmark_combinatorika(b *testing.B) {
	p := &Poker{logger: zap.NewNop()}
	for i := 0; i < b.N; i++ {
		p.combinatorika(uint64(7), uint64(52))
	}
}

func Test_factorial(t *testing.T) {
	p := &Poker{logger: zap.NewNop()}
	res := p.factorial(uint64(10))
	assert.Equal(t, uint64(3628800), res)
}

func Test_factorial_01(t *testing.T) {
	p := &Poker{logger: zap.NewNop()}
	res := p.factorial_01(uint64(10))
	assert.Equal(t, uint64(3628800), res)
}

func Test_factorial_02(t *testing.T) {
	p := &Poker{logger: zap.NewNop()}
	res := p.factorial_02(uint64(10))
	assert.Equal(t, uint64(3628800), res)
}

//-------   Benchmark ------------------------
func Benchmark_factorial_00(b *testing.B) {
	p := &Poker{logger: zap.NewNop()}
	for i := 0; i < b.N; i++ {
		p.factorial(uint64(7))
	}
}

func Benchmark_factorial_01(b *testing.B) {
	p := &Poker{logger: zap.NewNop()}
	for i := 0; i < b.N; i++ {
		p.factorial_01(uint64(7))
	}
}

func Benchmark_factorial_02(b *testing.B) {
	p := &Poker{logger: zap.NewNop()}
	for i := 0; i < b.N; i++ {
		p.factorial_02(uint64(7))
	}
}
