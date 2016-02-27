package main

/*
#cgo CFLAGS: -std=c11
#include <stdint.h>
*/
import "C"
import "qtcore"
import "fmt"
import "reflect"
import "time"
import "qtrt"

func testResolve(args ...interface{}) {
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
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
		fmt.Println(matched_index)
	}

}

func testgc() {
	s := qtcore.NewQString("vioug")
	qtcore.GcfreeQString(s)
}

func main() {
	testResolve("vioug")

	s := qtcore.NewQString("vioug")
	fmt.Println(321, s, 123)
	len := s.Length()
	lty := reflect.TypeOf(len)

	s1 := s.Append("abc")
	fmt.Println(len, lty.Kind().String(), len.(int32), s1, s)
	fmt.Println(s1.(*qtcore.QString).Length(), s1.(*qtcore.QString).Qclsinst, s.Qclsinst)
	s.Free()
	func() {
		defer func() {
			// fatal error can not recover, only panic can recover
			r := recover()
			fmt.Println(r)
		}()

		// len = s.Length()
		// fmt.Println(len)
	}()

	testgc()
	time.Sleep(1 * time.Second)
}
