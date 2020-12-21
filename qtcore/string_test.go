package qtcore

import "testing"

func TestQString1(t *testing.T) {
	s := NewQString5("foo123")
	if s.Length() != 6 {
		t.Fail()
	}
}
func BenchmarkQStringLenght(b *testing.B) {
	s := NewQString5("foo123")
	for i := 0; i < b.N; i++ {
		s.Length()
	}
}
