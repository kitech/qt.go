package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qtouchdevice.h
// dst-file: /src/gui/qtouchdevice.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QTouchDevice)=8
type QTouchDevice struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTouchDevice) setName(args ...interface{}) () {
  // setName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTouchDevice7setNameERK7QString
  default:
    qtrt.ErrorResolve("QTouchDevice", "setName", args)
 }

}


func (this *QTouchDevice) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTouchDevice4nameEv
  default:
    qtrt.ErrorResolve("QTouchDevice", "name", args)
 }

}


func (this *QTouchDevice) setMaximumTouchPoints(args ...interface{}) () {
  // setMaximumTouchPoints(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTouchDevice21setMaximumTouchPointsEi
  default:
    qtrt.ErrorResolve("QTouchDevice", "setMaximumTouchPoints", args)
 }

}


func (this *QTouchDevice) devices_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTouchDevice", "devices", args)
 }

}


func NewQTouchDevice(args ...interface{})() {
}


func (this *QTouchDevice) FreeQTouchDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTouchDevice", "~QTouchDevice", args)
 }

}


func (this *QTouchDevice) maximumTouchPoints(args ...interface{}) () {
  // maximumTouchPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QTouchDevice18maximumTouchPointsEv
  default:
    qtrt.ErrorResolve("QTouchDevice", "maximumTouchPoints", args)
 }

}

// <= body block end

