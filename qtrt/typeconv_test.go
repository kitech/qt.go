package qtrt

import "testing"

func TestSizeof(t *testing.T) {
	StringSliceToCCharPP(nil)
}

func TestF2i(t *testing.T) {

}

func TestToCharPP(t *testing.T) {
	lst := []string{"foo", "bar", "baz"}
	pp := StringSliceToCCharPP(lst)
	lst2 := CCharPPToStringSlice(pp)
	if len(lst2) != 3 {
		t.Fail()
	}
	for i := 0; i < len(lst); i++ {
		if lst2[i] != lst[i] {
			t.Fail()
		}
	}
}
