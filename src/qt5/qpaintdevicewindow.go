package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpaintdevicewindow.h
// dst-file: /src/gui/qpaintdevicewindow.go
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
// class sizeof(QPaintDeviceWindow)=1
type QPaintDeviceWindow struct {
  /*qbase*/ QWindow;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPaintDeviceWindow) update(args ...interface{}) () {
  // update(const class QRegion &)
  // update()
  // update(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPaintDeviceWindow6updateERK7QRegion
  case 1:
    // invoke: _ZN18QPaintDeviceWindow6updateEv
  case 2:
    // invoke: _ZN18QPaintDeviceWindow6updateERK5QRect
  default:
    qtrt.ErrorResolve("QPaintDeviceWindow", "update", args)
  }

}


func NewQPaintDeviceWindow(args ...interface{}) QPaintDeviceWindow {
  return QPaintDeviceWindow{}
}


func (this *QPaintDeviceWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPaintDeviceWindow10metaObjectEv
  default:
    qtrt.ErrorResolve("QPaintDeviceWindow", "metaObject", args)
  }

}

// <= body block end

