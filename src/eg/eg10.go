package main

// import "qt5"
import "fmt"
import "reflect"
import "qtrt"

type QString struct{}
type QChar struct{}
type QByteArray struct{}

func testResolve(args ...interface{}) {
	var vtys = make(map[int32]map[int32]reflect.Type)
	if false {
		fmt.Println(vtys)
	}
	vtys[0] = make(map[int32]reflect.Type)
	vtys[1] = make(map[int32]reflect.Type)
	vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
	vtys[2] = make(map[int32]reflect.Type)
	vtys[2][0] = qtrt.Int32Ty(false)     // "int"
	vtys[2][1] = reflect.TypeOf(QChar{}) // "QChar"
	vtys[3] = make(map[int32]reflect.Type)
	vtys[3][0] = qtrt.ByteTy(true) // "const char *"
	vtys[4] = make(map[int32]reflect.Type)
	vtys[4][0] = reflect.TypeOf(QChar{}) // "const QChar *"
	vtys[4][1] = qtrt.Int32Ty(false)     // "int"
	vtys[5] = make(map[int32]reflect.Type)
	vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
	vtys[6] = make(map[int32]reflect.Type)
	vtys[6][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

	var matched_index = qtrt.SymbolResolve(args, vtys)
	if true {
		fmt.Println(matched_index)
	}

}
func main() {
	vt0 := "vioug"
	if false {
		fmt.Println(vt0)
	}
	vt1 := &QChar{}
	if false {
		fmt.Println(vt1)
	}
	vt2 := QString{}
	testResolve(&vt2)

	// s := qt5.NewQString("vioug")
	// fmt.Println(321, s, 123)
	// len := s.Length()
	// lty := reflect.TypeOf(len)
	// fmt.Println(len, lty.Kind().String(), len.(int32))
}
