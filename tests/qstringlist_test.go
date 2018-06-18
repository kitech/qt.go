package main

import (
	"log"
	"testing"

	"github.com/kitech/qt.go/qtcore"
)

func Test0(t *testing.T) {
	lst := qtcore.NewQStringList()
	log.Println(lst)
	lstx := qtcore.NewQStringListxFromPointer(lst.GetCthis())
	log.Println(lstx)

	log.Println(lstx.Count())
}
