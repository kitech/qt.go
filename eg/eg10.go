package main

import (
	"log"
	"reflect"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

// 这和qtcore.QString并不冲突
type QString struct{}
type QChar struct{}
type QByteArray struct{}

func testResolve(args ...interface{}) int {
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		log.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)
	vtys[1] = make(map[int32]reflect.Type)
	vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
	vtys[2] = make(map[int32]reflect.Type)
	vtys[2][0] = qtrt.Int32Ty(false)            // "int"
	vtys[2][1] = reflect.TypeOf(qtcore.QChar{}) // "QChar"
	vtys[3] = make(map[int32]reflect.Type)
	vtys[3][0] = qtrt.ByteTy(true) // "const char *"
	vtys[4] = make(map[int32]reflect.Type)
	vtys[4][0] = reflect.TypeOf(qtcore.QChar{}) // "const QChar *"
	vtys[4][1] = qtrt.Int32Ty(false)            // "int"
	vtys[5] = make(map[int32]reflect.Type)
	vtys[5][0] = reflect.TypeOf(qtcore.QChar{}) // "QChar"
	vtys[6] = make(map[int32]reflect.Type)
	vtys[6][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if true {
		log.Println(matched_index, vtys[1])
	}
	return int(matched_index)
}
func main() {
	var midx int
	vt0 := "vioug"
	if false {
		log.Println(vt0)
	}
	vt1 := &qtcore.QChar{}
	if false {
		log.Println(vt1)
	}
	vt2 := qtcore.QString{}

	midx = testResolve(vt0)
	if midx != 3 {
		log.Println("resolution failed")
	} else {
		log.Println("OK", midx, []interface{}{vt0})
	}

	midx = testResolve(&vt2)
	if midx != 1 {
		log.Println("resolution failed")
	} else {
		log.Println("OK", midx, []interface{}{&vt2})
	}

	midx = testResolve(vt1)
	if midx != 5 {
		log.Println("resolution failed")
	} else {
		log.Println("OK", midx, []interface{}{vt1})
	}

	s := qtcore.NewQString("vioug")
	log.Println(321, s, 123)
	len := s.Length()
	lty := reflect.TypeOf(len)
	log.Println(len, lty.Kind().String(), len.(int32))
	if int(len.(int32)) == 5 {
	} else {
		log.Println("Length method call failed", len)
	}
}
