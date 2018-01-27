package qtsymbols

import "github.com/emirpasic/gods/maps/hashbidimap"

var Qtsymbolsbi = hashbidimap.New()

func init() {
	for sym, sign := range Qtsymbols {
		Qtsymbolsbi.Put(sym, sign)
	}
}
