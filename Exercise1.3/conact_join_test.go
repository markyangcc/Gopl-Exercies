// concat/join string benchmark
// name the .go file with suffix _test to run go test command, according to go standard
//
// run with: go test -bench=.

package concat_join_test

import (
	"strings"
	"testing"
)

var Args = []string{"-p", "-a", "-c", "-x", "-v", "-f", "-o"}

func concatargs(args []string) {
	var s, sep string
	for _, item := range args {
		s += sep + item
		sep = " "
	}
}

func joinargs(args []string) {
	var _ string
	_ = strings.Join(args, " ")
}

func BenchmarkConcatargs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatargs(Args)
	}
}

func BenchmarkJoinargs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinargs(Args)
	}
}
