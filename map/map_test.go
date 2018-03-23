package main

import "testing"

func TestTring(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"assdas", 3},
		{"abcdss", 5},
		{"一二三一", 3},
	}

	for _, tt := range tests {
		if actual := getString(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkString(b *testing.B) {
	s := "assdas"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) %d", len(s))
	b.ResetTimer()
	ans := 3
	for i := 0; i < b.N; i++ {
		actual := getString(s)
		if actual != ans {
			b.Errorf("got %d for input %s expected %d", actual, s, ans)
		}
	}
}
