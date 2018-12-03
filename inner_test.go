package log

import "testing"

type llvl struct{}

func (l *llvl) Print(...interface{}) {}

type llog struct {
	l   *llvl
	lvl LogLevel
}

func (l *llog) print() {
	if l.lvl == Off {
		l.l.Print()
	}
}

type llogx struct {
	l Simple
}

func (l *llogx) print() {
	l.l.Print()
}

func BenchmarkIfLevel(t *testing.B) {
	l := &llog{l: &llvl{}}
	for i := 0; i < t.N; i++ {
		l.print()
	}
}

func BenchmarkInterfaceCast(t *testing.B) {
	l := &llogx{l: &llvl{}}
	for i := 0; i < t.N; i++ {
		l.print()
	}
}
